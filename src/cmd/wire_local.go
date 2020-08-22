//+build wireinject

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/wht/adapter/storage"
	"github.com/sky0621/wht/adapter/web"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gocloud.dev/server"
	"golang.org/x/xerrors"
)

// ローカルマシン上で動かす際の固有設定
func buildLocal(ctx context.Context, cfg config) (*app, func(), error) {
	wire.Build(
		setupLocalRDB,
		appSet,
		web.NewResolver,
		setupLocalServer,
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

func setupLocalServer(ctx context.Context, cfg config, resolver *web.Resolver) (*server.Server, error) {
	log.Debug().Msg("setupLocalServer___START")

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(requestCtxLogger())

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// FIXME: 本番はNG
	r.HandleFunc("/pg", playground.Handler("GraphQL playground", "/query"))

	r.Handle("/query", graphQlServer(resolver))

	// TODO: ローカルでは、ひとまず、フロント 3000 port、バック 8080 port で起動させておく。
	//var workDir string
	//{
	//	var err error
	//	workDir, err = os.Getwd()
	//	if err != nil {
	//		return nil, xerrors.Errorf("failed to Getwd: %w", err)
	//	}
	//}
	//log.Info().Msgf("workDir:%s", workDir)
	//
	//filesDir := http.Dir(filepath.Join(workDir, "dist"))
	//log.Info().Msgf("filesDir:%#+v", filesDir)
	//
	//r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
	//	log.Info().Str("Host", r.Host).Str("RequestURI", r.RequestURI).Interface("URL", r.URL).Msg("REQUEST")
	//
	//	ctx := chi.RouteContext(r.Context())
	//	pathPrefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
	//	log.Info().Msgf("pathPrefix:%s", pathPrefix)
	//
	//	fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
	//	fs.ServeHTTP(w, r)
	//})

	return server.New(r, nil), nil
}

func setupLocalCloudStorageClient(ctx context.Context, cfg config) (storage.CloudStorageClient, error) {
	log.Debug().Msg("setupLocalCloudStorageClient___START")

	// FIXME: モックを返却 or fileblob ?
	bucketNameMap := map[storage.BucketPurpose]string{
		storage.ImageContentsBucket: cfg.ImageContentsBucket,
	}
	return storage.NewCloudStorageClient(ctx, bucketNameMap)
}
