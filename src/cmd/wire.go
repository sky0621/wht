//+build wireinject

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sky0621/wht/adapter/store"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wht/adapter/web"
	"github.com/sky0621/wht/adapter/web/gqlmodel"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

// デプロイ先インフラで動かす際の設定
func build(ctx context.Context, cfg config) (*app, func(), error) {
	wire.Build(
		connectDB,
		appSet,
		web.NewResolver,
		setupRouter,
		setupCloudStorageClient,
	)
	return nil, nil, nil
}

func connectDB(cfg config) (boil.ContextExecutor, func(), error) {
	log.Println("connectDB() start...")

	dsn := fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName)
	log.Printf("DSN:%s", dsn)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, nil, xerrors.Errorf("failed to ping: %w", err)
	}

	// FIXME: 本番はNG?
	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to time.LoadLocation: %w", err)
	}
	boil.SetLocation(loc)

	return db, func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Printf("%+v", err)
			}
		}
	}, nil
}

func setupRouter(cfg config, resolver *web.Resolver) *chi.Mux {
	r := chi.NewRouter()

	// FIXME: 本番はNG
	r.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))

	r.Handle("/query", web.DataLoaderMiddleware(resolver, graphQlServer(resolver)))

	return r
}

func graphQlServer(resolver *web.Resolver) *handler.Server {
	cfg := web.Config{Resolvers: resolver}
	// FIXME: 認可実装
	cfg.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role gqlmodel.Role) (interface{}, error) {
		// or let it pass through
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
		log.Println(err)
		return graphql.DefaultErrorPresenter(ctx, err)
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Println(err)
		return fmt.Errorf("util server error! %v", err)
	})

	return srv
}

func setupCloudStorageClient(ctx context.Context, cfg config) (store.CloudStorageClient, error) {
	bucketNameMap := map[store.BucketPurpose]string{
		store.ImageContentsBucket: cfg.ImageContentsBucket,
	}
	return store.NewCloudStorageClient(ctx, bucketNameMap)
}
