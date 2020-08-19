package web

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/wht/lib"

	"github.com/sky0621/wht/adapter/storage"

	"github.com/sky0621/wht/adapter/web/gqlmodel"
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
	// FIXME:
	if err := r.gcsClient.ExecUploadObject(ctx, storage.ImageContentsBucket, wht.Image.Filename, wht.Image.File); err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return &gqlmodel.MutationResponse{ID: lib.FromInt64ToPStr(id)}, nil
}

// ------------------------------------------------------------------
// Query
// ------------------------------------------------------------------

func (r *queryResolver) FindWht(ctx context.Context) ([]gqlmodel.Wht, error) {
	logger := lib.RequestCtxLogger(ctx)
	logger.Info().Msg("FindWht___START")

	records, err := r.wht.ReadWhts(ctx)
	if err != nil {
		logger.Err(err).Msgf("%+v", err)
		return nil, err
	}
	var results []gqlmodel.Wht
	for _, record := range records {
		results = append(results, gqlmodel.Wht{
			ID:         gqlmodel.WhtID(record.ID),
			RecordDate: record.RecordDate,
			Path:       record.Path, // FIXME: 署名付きURL生成
		},
		)
	}

	return results, nil
}
