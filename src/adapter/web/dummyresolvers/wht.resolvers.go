package web

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/sky0621/wht/adapter/web"
	"github.com/sky0621/wht/adapter/web/gqlmodel"
)

func (r *mutationResolver) CreateWht(ctx context.Context, wht gqlmodel.WhtInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTextContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.TextContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateImageContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.ImageContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVoiceContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.VoiceContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMovieContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.MovieContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindWht(ctx context.Context, condition *gqlmodel.WhtConditionInput) ([]gqlmodel.Wht, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *whtResolver) Contents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

// Wht returns web.WhtResolver implementation.
func (r *Resolver) Wht() web.WhtResolver { return &whtResolver{r} }

type whtResolver struct{ *Resolver }
