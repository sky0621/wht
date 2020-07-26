package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/sky0621/wht/adapter/controller/gqlmodel"
)

func (r *mutationResolver) CreateWht(ctx context.Context, wht gqlmodel.WhtInput) (string, error) {
	// FIXME:
	return "9097d7a4-5a9b-4581-84e4-80b0db4282aa", nil
}

func (r *queryResolver) FindWht(ctx context.Context) ([]gqlmodel.Wht, error) {
	// FIXME:
	t := "いい一日4"
	return []gqlmodel.Wht{
		{ID: "001", RecordDate: "2020-07-26", Title: &t, Text: "今日は、いい一日だった。"},
	}, nil
}
