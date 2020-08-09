package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
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
		log.Err(err).Msgf("failed to buildApp: %+v", err)
		return abnormalEnd
	}
	defer shutdownFunc()

	log.Info().Msg(cfg.String())

	go func() {
		q := make(chan os.Signal, 1)
		signal.Notify(q, os.Interrupt, syscall.SIGTERM)
		<-q
		shutdownFunc()
		os.Exit(abnormalEnd)
	}()

	log.Printf("wht: listening on port %s", cfg.WebPort)
	if err := http.ListenAndServe(":"+cfg.WebPort, app.r); err != nil {
		log.Err(err).Str("WebPort", cfg.WebPort).Msgf("failed to ListenAndServe: %+v", err)
		return abnormalEnd
	}
	return normalEnd
}

func buildApp(ctx context.Context, cfg config) (*app, func(), error) {
	if cfg.IsLocal() {
		globalLocalLogSetting()
		return buildLocal(ctx, cfg)
	}
	globalLogSetting()
	return build(ctx, cfg)
}
