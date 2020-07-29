package web

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sky0621/wht/adapter/web/gqlmodel"
)

const loadersKey = "dataLoaders"

type DataLoaders struct {
	contentLoader *gqlmodel.ContentLoader
}

func DataLoaderMiddleware(resolver *Resolver, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &DataLoaders{
			contentLoader: gqlmodel.NewContentLoader(gqlmodel.ContentLoaderConfig{
				MaxBatch: 100,
				Wait:     1 * time.Millisecond,
				Fetch: func(keys []int64) ([][]gqlmodel.Content, []error) {
					fmt.Println(r.Context())
					// FIXME:

					return nil, nil
				},
			}),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *DataLoaders {
	return ctx.Value(loadersKey).(*DataLoaders)
}
