package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
)

func (r *mutationResolver) CreateWht(ctx context.Context, wht WhtInput) (string, error) {
	// FIXME:
	return "", nil
}

func (r *queryResolver) FindWht(ctx context.Context) ([]*Wht, error) {
	// FIXME:
	t := "いい一日"
	return []*Wht{
		{ID: "001", RecordDate: "2020-07-26", Title: &t, Text: "今日は、いい一日だった。"},
	}, nil
}
