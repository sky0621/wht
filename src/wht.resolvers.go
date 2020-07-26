package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/sky0621/wht/model"
)

func (r *mutationResolver) CreateWht(ctx context.Context, wht model.WhtInput) (string, error) {
	// FIXME:
	return "", nil
}

func (r *queryResolver) FindWht(ctx context.Context) ([]*model.Wht, error) {
	// FIXME:
	t := "いい一日2"
	return []*model.Wht{
		{ID: "001", RecordDate: "2020-07-26", Title: &t, Text: "今日は、いい一日だった。"},
	}, nil
}
