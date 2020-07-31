package web

import (
	"context"
	"net/http"
	"time"

	"github.com/sky0621/wht/adapter/web/gqlmodel"
)

const loadersKey = "dataLoaders"

type DataLoaders struct {
	textContentLoader  *gqlmodel.TextContentLoader
	imageContentLoader *gqlmodel.ImageContentLoader
	voiceContentLoader *gqlmodel.VoiceContentLoader
	movieContentLoader *gqlmodel.MovieContentLoader
}

func DataLoaderMiddleware(resolver *Resolver, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &DataLoaders{
			textContentLoader: gqlmodel.NewTextContentLoader(gqlmodel.TextContentLoaderConfig{
				MaxBatch: 100,
				Wait:     1 * time.Millisecond,
				Fetch: func(whtIDs []int64) ([][]*gqlmodel.TextContent, []error) {
					contentsByID, err := resolver.wht.ReadTextContents(r.Context(), whtIDs)
					if err != nil {
						return nil, replicateError(err, len(whtIDs))
					}
					results := make([][]*gqlmodel.TextContent, len(whtIDs))
					for i, whtID := range whtIDs {
						results[i] = gqlmodel.FromTextContent(contentsByID[whtID])
					}
					return results, nil
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

func replicateError(err error, count int) []error {
	errors := make([]error, count)
	errors[0] = err
	for bp := 1; bp < count; bp *= 2 {
		copy(errors[bp:], errors[:bp])
	}
	return errors
}
