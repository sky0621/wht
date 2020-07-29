//+build wireinject

package setup

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wht/adapter/web"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

// ローカルマシン上で動かす際の固有設定
func buildLocal(ctx context.Context, cfg config) (app, error) {
	wire.Build(
		connectLocalDB,
		web.NewResolver,
		setupRouter,
		newApp,
	)
	return app{}, nil
}

func connectLocalDB(cfg config) (*sqlx.DB, error) {
	log.Println("connectLocalDB() start...")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPass)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, xerrors.Errorf("failed to time.LoadLocation: %w", err)
	}
	boil.SetLocation(loc)

	return db, nil
}
