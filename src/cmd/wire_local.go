//+build wireinject

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/wht/adapter/store"
	"github.com/sky0621/wht/adapter/web"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

// ローカルマシン上で動かす際の固有設定
func buildLocal(ctx context.Context, cfg config) (*app, func(), error) {
	wire.Build(
		setupLocalRDB,
		appSet,
		web.NewResolver,
		setupRouter,
		setupLocalCloudStorageClient,
	)
	return nil, nil, nil
}

func setupLocalRDB(cfg config) (boil.ContextExecutor, func(), error) {
	log.Debug().Msg("setupLocalRDB___START")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPass)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to time.LoadLocation: %w", err)
	}
	boil.SetLocation(loc)

	log.Debug().Msg("setupLocalRDB___END")
	return db, func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Err(err).Send()
			}
		}
	}, nil
}

func setupLocalCloudStorageClient(ctx context.Context, cfg config) (store.CloudStorageClient, error) {
	log.Debug().Msg("setupLocalCloudStorageClient___START")

	// FIXME: モックを返却
	bucketNameMap := map[store.BucketPurpose]string{
		store.ImageContentsBucket: cfg.ImageContentsBucket,
	}
	return store.NewCloudStorageClient(ctx, bucketNameMap, "")
}
