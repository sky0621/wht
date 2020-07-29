package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wht/adapter/controller/gqlmodel"
	"github.com/sky0621/wht/adapter/gateway"
	"github.com/sky0621/wht/application"
	"github.com/sky0621/wht/application/domain"
	"github.com/sky0621/wht/application/util"
)

// ------------------------------------------------------------------
// Mutation
// ------------------------------------------------------------------

func (r *mutationResolver) CreateWht(ctx context.Context, wht gqlmodel.WhtInput) (*gqlmodel.MutationResponse, error) {
	res, err := gateway.Tx(ctx, r.db, func(ctx context.Context, txx *sqlx.Tx) (*gateway.TxResponse, error) {
		id, err := application.NewWht(gateway.NewWhtRepository(txx), gateway.NewContentRepository(txx)).
			CreateWht(ctx, domain.Wht{RecordDate: wht.RecordDate, Title: wht.Title})
		return &gateway.TxResponse{CreatedID: id}, err
	})
	if err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return &gqlmodel.MutationResponse{ID: util.FromInt64ToPStr(res.CreatedID)}, nil
}

func (r *mutationResolver) CreateTextContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.TextContentInput) (*gqlmodel.MutationResponse, error) {
	res, err := gateway.Tx(ctx, r.db, func(ctx context.Context, txx *sqlx.Tx) (*gateway.TxResponse, error) {
		var contents []domain.TextContent
		for _, in := range inputs {
			contents = append(contents, domain.NewTextContent(in.Name, in.Text))
		}
		err := application.NewWht(gateway.NewWhtRepository(txx), gateway.NewContentRepository(txx)).
			CreateTextContents(ctx, recordDate, contents)
		return &gateway.TxResponse{CreatedID: 0}, err // TODO: think returning id when batch create
	})
	if err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return &gqlmodel.MutationResponse{ID: util.FromInt64ToPStr(res.CreatedID)}, nil
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

// ------------------------------------------------------------------
// Query
// ------------------------------------------------------------------

func (r *queryResolver) FindWht(ctx context.Context, condition *gqlmodel.WhtConditionInput) ([]gqlmodel.Wht, error) {
	c := &domain.WhtCondition{}
	if condition != nil {
		id := condition.ID.DBUniqueID()
		c.ID = &id
	}

	records, err := application.NewWht(gateway.NewWhtRepository(r.db), gateway.NewContentRepository(r.db)).
		ReadWht(ctx, c)
	if err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}

	var results []gqlmodel.Wht
	for _, r := range records {
		if r.ID == nil || r.RecordDate == util.NilTime {
			err := errors.New("id or recordDate is nil")
			fmt.Printf("%#+v", err) // TODO: use custom logger
			return nil, err
		}
		results = append(results, gqlmodel.Wht{
			ID:         gqlmodel.WhtID(*r.ID),
			RecordDate: r.RecordDate,
			Title:      r.Title,
		})
	}
	return results, nil
}

func (r *whtResolver) Contents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.Content, error) {
	contents, err := For(ctx).contentLoader.Load(obj.ID.DBUniqueID())
	if err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return contents, nil
}

func (r *Resolver) Wht() WhtResolver { return &whtResolver{r} }

type whtResolver struct{ *Resolver }
