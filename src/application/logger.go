package application

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
	WithRequestContext(ctx context.Context) *zerolog.Logger
}

type logger struct {
	zLog zerolog.Logger
}

func NewLogger() Logger {
	l := zerolog.New(os.Stdout)
	return &logger{zLog: l}
}

func (l *logger) WithRequestContext(ctx context.Context) *zerolog.Logger {
	zLgr := l.zLog.With().Logger()
	return &zLgr
}
