package setup

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

type ExitCode int

const (
	normalEnd   = 0
	abnormalEnd = -1
)

func ExecMain() ExitCode {
	cfg := newConfig()
	log.Println(cfg)

	ctx := context.Background()
	var app app
	var err error
	if cfg.IsLocal() {
		app, err = buildLocal(ctx, cfg)
	} else {
		app, err = build(ctx, cfg)
	}
	if err != nil {
		log.Printf("%+v", err)
		return abnormalEnd
	}
	defer app.shutdown()

	go func() {
		q := make(chan os.Signal)
		signal.Notify(q, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-q
		app.shutdown()
		os.Exit(abnormalEnd)
	}()

	log.Printf("wht: listening on port %s", cfg.WebPort)
	if err := http.ListenAndServe(":"+cfg.WebPort, app.r); err != nil {
		log.Printf("%+v", err)
		return abnormalEnd
	}
	return normalEnd
}

type app struct {
	db *sqlx.DB
	r  *chi.Mux
}

func newApp(db *sqlx.DB, r *chi.Mux) app {
	return app{
		db: db,
		r:  r,
	}
}

func (a app) shutdown() {
	if a.db == nil {
		return
	}
	if err := a.db.Close(); err != nil {
		log.Printf("%+v", err)
	}
}
