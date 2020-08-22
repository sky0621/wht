// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/wht/adapter/rdb"
	"github.com/sky0621/wht/adapter/storage"
	"github.com/sky0621/wht/adapter/web"
	"github.com/sky0621/wht/application/usecase"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.opencensus.io/trace"
	"gocloud.dev/gcp"
	"gocloud.dev/server"
	"gocloud.dev/server/health"
	"gocloud.dev/server/sdserver"
	"golang.org/x/xerrors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

import (
	_ "github.com/lib/pq"
	_ "gocloud.dev/blob/gcsblob"
)

// Injectors from wire.go:

func build(ctx context.Context, cfg config) (*app, func(), error) {
	contextExecutor, cleanup, err := setupRDB(cfg)
	if err != nil {
		return nil, nil, err
	}
	wht := rdb.NewWhtRepository(contextExecutor)
	usecaseWht := usecase.NewWht(wht)
	cloudStorageClient, err := setupCloudStorageClient(ctx, cfg)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	resolver := web.NewResolver(usecaseWht, cloudStorageClient)
	server, err := setupServer(ctx, cfg, resolver)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mainApp := newApp(server)
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
	usecaseWht := usecase.NewWht(wht)
	cloudStorageClient, err := setupLocalCloudStorageClient(ctx, cfg)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	resolver := web.NewResolver(usecaseWht, cloudStorageClient)
	server, err := setupLocalServer(ctx, cfg, resolver)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mainApp := newApp(server)
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

type GlobalMonitoredResource struct {
	projectID string
}

func (g GlobalMonitoredResource) MonitoredResource() (string, map[string]string) {
	return "global", map[string]string{"project_id": g.projectID}
}

type customHealthCheck struct {
	mu      sync.RWMutex
	healthy bool
}

func (h *customHealthCheck) CheckHealth() error {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if !h.healthy {
		return errors.New("not ready yet")
	}
	return nil
}

func setupServer(ctx context.Context, cfg config, resolver *web.Resolver) (*server.Server, error) {
	log.Debug().Msg("setupServer___START")

	credentials, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to DefaultCredentials: %w", err)
	}

	tokenSource := gcp.CredentialsTokenSource(credentials)

	var projectID gcp.ProjectID
	{
		var err error
		projectID, err = gcp.DefaultProjectID(credentials)
		if err != nil {
			return nil, xerrors.Errorf("failed to DefaultProjectID: %w", err)
		}
	}

	var exporter trace.Exporter
	if cfg.Trace {
		fmt.Println("Exporting traces to Stackdriver")
		mr := GlobalMonitoredResource{projectID: string(projectID)}
		exporter, _, err = sdserver.NewExporter(projectID, tokenSource, mr)
		if err != nil {
			return nil, xerrors.Errorf("failed to NewExporter: %w", err)
		}
	}

	healthCheck := new(customHealthCheck)
	time.AfterFunc(10*time.Second, func() {
		healthCheck.mu.Lock()
		defer healthCheck.mu.Unlock()
		healthCheck.healthy = true
	})

	options := &server.Options{
		RequestLogger: sdserver.NewRequestLogger(),
		HealthChecks:  []health.Checker{healthCheck},
		TraceExporter: exporter,

		DefaultSamplingPolicy: trace.AlwaysSample(),
		Driver:                &server.DefaultDriver{},
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(requestCtxLogger())

	r.Use(middleware.Timeout(60 * time.Second))

	r.HandleFunc("/pg", playground.Handler("GraphQL playground", "/query"))

	r.Handle("/query", graphQlServer(resolver))

	var workDir string
	{
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			return nil, xerrors.Errorf("failed to Getwd: %w", err)
		}
	}
	log.Info().Msgf("workDir:%s", workDir)

	filesDir := http.Dir(filepath.Join(workDir, "frontend"))
	log.Info().Msgf("filesDir:%#+v", filesDir)

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("Host", r.Host).Str("RequestURI", r.RequestURI).Msg("REQUEST")

		ctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		log.Info().Msgf("pathPrefix:%s", pathPrefix)

		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})

	return server.New(r, options), nil
}

func graphQlServer(resolver *web.Resolver) *handler.Server {
	log.Debug().Msg("graphQlServer___START")

	cfg := web.Config{Resolvers: resolver}

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

func setupCloudStorageClient(ctx context.Context, cfg config) (storage.CloudStorageClient, error) {
	log.Debug().Msg("setupCloudStorageClient___START")

	bucketNameMap := map[storage.BucketPurpose]string{storage.ImageContentsBucket: cfg.ImageContentsBucket}
	return storage.NewCloudStorageClient(ctx, bucketNameMap)
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

func setupLocalServer(ctx context.Context, cfg config, resolver *web.Resolver) (*server.Server, error) {
	log.Debug().Msg("setupLocalServer___START")

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(requestCtxLogger())

	r.Use(cors.Handler(cors.Options{

		AllowedOrigins: []string{"*"},

		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.HandleFunc("/pg", playground.Handler("GraphQL playground", "/query"))

	r.Handle("/query", graphQlServer(resolver))

	var workDir string
	{
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			return nil, xerrors.Errorf("failed to Getwd: %w", err)
		}
	}
	log.Info().Msgf("workDir:%s", workDir)

	filesDir := http.Dir(filepath.Join(workDir, "dist"))
	log.Info().Msgf("filesDir:%#+v", filesDir)

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("Host", r.Host).Str("RequestURI", r.RequestURI).Interface("URL", r.URL).Msg("REQUEST")

		ctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		log.Info().Msgf("pathPrefix:%s", pathPrefix)

		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})

	return server.New(r, nil), nil
}

func setupLocalCloudStorageClient(ctx context.Context, cfg config) (storage.CloudStorageClient, error) {
	log.Debug().Msg("setupLocalCloudStorageClient___START")

	bucketNameMap := map[storage.BucketPurpose]string{storage.ImageContentsBucket: cfg.ImageContentsBucket}
	return storage.NewCloudStorageClient(ctx, bucketNameMap)
}
