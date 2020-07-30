package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

const (
	normalEnd   = 0
	abnormalEnd = -1
)

func main() {
	os.Exit(execMain())
}

func execMain() int {
	cfg := newConfig()

	app, shutdownFunc, err := buildApp(context.Background(), cfg)
	if err != nil {
		log.Printf("%+v", err)
		return abnormalEnd
	}
	defer shutdownFunc()

	go func() {
		q := make(chan os.Signal, 1)
		signal.Notify(q, os.Interrupt, syscall.SIGTERM)
		<-q
		shutdownFunc()
		os.Exit(abnormalEnd)
	}()

	log.Printf("wht: listening on port %s", cfg.WebPort)
	if err := http.ListenAndServe(":"+cfg.WebPort, app.r); err != nil {
		log.Printf("%+v", err)
		return abnormalEnd
	}
	return normalEnd
}

func buildApp(ctx context.Context, cfg config) (*app, func(), error) {
	if cfg.IsLocal() {
		log.Println("on Local...")
		return buildLocal(ctx, cfg)
	}
	log.Println("on GCP...")
	return build(ctx, cfg)
}
