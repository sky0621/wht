package lib

import (
	"context"

	"github.com/rs/zerolog"
)

const LoggerKey = "request-ctx-logger"

func RequestCtxLogger(ctx context.Context) zerolog.Logger {
	if l, ok := ctx.Value(LoggerKey).(zerolog.Logger); ok {
		return l
	}
	return zerolog.Nop()
}
