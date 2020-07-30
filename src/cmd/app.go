package main

import (
	"github.com/sky0621/wht/adapter/rdb"

	"github.com/sky0621/wht/application/usecase"

	"github.com/google/wire"

	"github.com/go-chi/chi"
)

func init() {
	_ = appSet
}

var appSet = wire.NewSet(
	repositoryHandler,
	usecaseHandler,
	newApp,
)

type app struct {
	r *chi.Mux
}

func newApp(r *chi.Mux) *app {
	return &app{r: r}
}

var usecaseHandler = wire.NewSet(
	usecase.NewWht,
)

var repositoryHandler = wire.NewSet(
	rdb.NewWhtRepository,
	rdb.NewContentRepository,
)
