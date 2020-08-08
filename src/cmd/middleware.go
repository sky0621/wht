package main

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/sky0621/wht/lib"
)

func requestCtxLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), lib.LoggerKey, log.Logger.With().Logger())))
		}
		return http.HandlerFunc(fn)
	}
}
