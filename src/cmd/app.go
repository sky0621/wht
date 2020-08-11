package main

import (
	"github.com/google/wire"
	"github.com/sky0621/wht/adapter/rdb"
	"github.com/sky0621/wht/application/usecase"
	"gocloud.dev/server"
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
	s *server.Server
}

func newApp(s *server.Server) *app {
	return &app{s: s}
}

var usecaseHandler = wire.NewSet(
	usecase.NewWht,
)

var repositoryHandler = wire.NewSet(
	rdb.NewWhtRepository,
	rdb.NewContentRepository,
)
