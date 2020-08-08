package web

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/sky0621/wht/lib"

	"github.com/sky0621/wht/adapter/store"

	"github.com/sky0621/wht/adapter/web/gqlmodel"
	"github.com/sky0621/wht/application/domain"
	"github.com/sky0621/wht/application/util"
)

// ------------------------------------------------------------------
// Mutation
// ------------------------------------------------------------------

func (r *mutationResolver) CreateWht(ctx context.Context, wht gqlmodel.WhtInput) (*gqlmodel.MutationResponse, error) {
	id, err := r.wht.CreateWht(ctx, gqlmodel.ToWhtForCreate(wht))
	if err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return &gqlmodel.MutationResponse{ID: util.FromInt64ToPStr(id)}, nil
}

func (r *mutationResolver) CreateTextContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.TextContentInput) (*gqlmodel.MutationResponse, error) {
	if err := r.wht.CreateTextContents(ctx, recordDate, gqlmodel.ToTextContentForCreateSlice(inputs)); err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return &gqlmodel.MutationResponse{}, nil
}

func (r *mutationResolver) CreateImageContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.ImageContentInput) (*gqlmodel.MutationResponse, error) {
	// FIXME:
	for _, in := range inputs {
		if err := r.gcsClient.ExecUploadObject(ctx, store.ImageContentsBucket, in.Image.Filename, in.Image.File); err != nil {
			fmt.Printf("%#+v", err) // TODO: use custom logger
			return nil, err
		}
	}
	return &gqlmodel.MutationResponse{
		ID: nil,
	}, nil
}

func (r *mutationResolver) CreateVoiceContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.VoiceContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMovieContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.MovieContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// ------------------------------------------------------------------
// Query
// ------------------------------------------------------------------

func (r *queryResolver) FindWht(ctx context.Context, c *gqlmodel.WhtConditionInput) ([]gqlmodel.Wht, error) {
	logger := lib.RequestCtxLogger(ctx)
	logger.Info().Msg("FindWht___START")

	condition := &domain.WhtCondition{}
	if c != nil {
		condition.ID = c.ID.DBUniqueIDPtr()
	}

	records, err := r.wht.ReadWhts(ctx, condition)
	if err != nil {
		logger.Err(err).Msgf("%+v", err)
		return nil, err
	}

	return gqlmodel.FromWhts(records), nil
}

func (r *whtResolver) TextContents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.TextContent, error) {
	contents, err := For(ctx).textContentLoader.Load(obj.ID.DBUniqueID())
	if err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return contents, nil
}

func (r *whtResolver) ImageContents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.ImageContent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *whtResolver) VoiceContents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.VoiceContent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *whtResolver) MovieContents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.MovieContent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Wht() WhtResolver { return &whtResolver{r} }

type whtResolver struct{ *Resolver }
