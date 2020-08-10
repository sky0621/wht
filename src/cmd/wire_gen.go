// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/wht/adapter/rdb"
	"github.com/sky0621/wht/adapter/store"
	"github.com/sky0621/wht/adapter/web"
	"github.com/sky0621/wht/adapter/web/gqlmodel"
	"github.com/sky0621/wht/application/usecase"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
	"time"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func build(ctx context.Context, cfg config) (*app, func(), error) {
	contextExecutor, cleanup, err := setupRDB(cfg)
	if err != nil {
		return nil, nil, err
	}
	wht := rdb.NewWhtRepository(contextExecutor)
	content := rdb.NewContentRepository(contextExecutor)
	usecaseWht := usecase.NewWht(wht, content)
	cloudStorageClient, err := setupCloudStorageClient(ctx, cfg)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	resolver := web.NewResolver(usecaseWht, cloudStorageClient)
	mux := setupRouter(cfg, resolver)
	mainApp := newApp(mux)
	return mainApp, func() {
		cleanup()
	}, nil
}

// Injectors from wire_local.go:

func buildLocal(ctx context.Context, cfg config) (*app, func(), error) {
	contextExecutor, cleanup, err := setupLocalRDB(cfg)
	if err != nil {
		return nil, nil, err
	}
	wht := rdb.NewWhtRepository(contextExecutor)
	content := rdb.NewContentRepository(contextExecutor)
	usecaseWht := usecase.NewWht(wht, content)
	cloudStorageClient, err := setupLocalCloudStorageClient(ctx, cfg)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	resolver := web.NewResolver(usecaseWht, cloudStorageClient)
	mux := setupRouter(cfg, resolver)
	mainApp := newApp(mux)
	return mainApp, func() {
		cleanup()
	}, nil
}

// wire.go:

func setupRDB(cfg config) (boil.ContextExecutor, func(), error) {
	log.Debug().Msg("setupRDB")

	dsn := fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName)
	log.Info().Msgf("DSN:%s", dsn)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, nil, xerrors.Errorf("failed to ping: %w", err)
	}
	boil.DebugMode = true

	return db, func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Err(err).Send()
			}
		}
	}, nil
}

func setupRouter(cfg config, resolver *web.Resolver) *chi.Mux {
	log.Debug().Msg("setupRouter___START")

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(requestCtxLogger())

	r.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))

	r.Handle("/query", web.DataLoaderMiddleware(resolver, graphQlServer(resolver)))

	return r
}

func graphQlServer(resolver *web.Resolver) *handler.Server {
	log.Debug().Msg("graphQlServer___START")

	cfg := web.Config{Resolvers: resolver}

	cfg.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role gqlmodel.Role) (interface{}, error) {

		return next(ctx)
	}
	srv := handler.New(web.NewExecutableSchema(cfg))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	var mb int64 = 1 << 20
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     128 * mb,
		MaxUploadSize: 100 * mb,
	})

	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		log.Err(err).Msgf("failed to graphQL service: %+v", err)
		return graphql.DefaultErrorPresenter(ctx, err)
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return xerrors.Errorf("panic occurred: %w", err)
	})

	return srv
}

func setupCloudStorageClient(ctx context.Context, cfg config) (store.CloudStorageClient, error) {
	log.Debug().Msg("setupCloudStorageClient___START")

	bucketNameMap := map[store.BucketPurpose]string{store.ImageContentsBucket: cfg.ImageContentsBucket}
	return store.NewCloudStorageClient(ctx, bucketNameMap)
}

// wire_local.go:

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

	bucketNameMap := map[store.BucketPurpose]string{store.ImageContentsBucket: cfg.ImageContentsBucket}
	return store.NewCloudStorageClient(ctx, bucketNameMap)
}
