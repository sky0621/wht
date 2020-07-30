package main

import (
	"log"

	"github.com/google/wire"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func init() {
	_ = appSet
}

var appSet = wire.NewSet(
	newApp,
)

type app struct {
	db *sqlx.DB
	r  *chi.Mux
}

func newApp(db *sqlx.DB, r *chi.Mux) *app {
	return &app{
		db: db,
		r:  r,
	}
}

func (a app) Shutdown() {
	if a.db == nil {
		return
	}
	if err := a.db.Close(); err != nil {
		log.Printf("%+v", err)
	}
}
