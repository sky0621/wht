package rdb

import (
	"context"

	"github.com/sky0621/wht/lib"

	"github.com/sky0621/wht/application/repository"

	"github.com/sky0621/wht/adapter/rdb/boiled"
	"github.com/sky0621/wht/application/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

func NewWhtRepository(db boil.ContextExecutor) repository.Wht {
	return &wht{db: db}
}

type wht struct {
	db boil.ContextExecutor
}

func (r *wht) Create(ctx context.Context, in *domain.WhtForCreate) (int64, error) {
	mdl := &boiled.WHT{
		RecordDate: in.RecordDate,
		Path:       in.Path,
	}
	if err := mdl.Insert(ctx, r.db, boil.Infer()); err != nil {
		return -1, xerrors.Errorf("failed to insert wht: %w", err)
	}
	return mdl.ID, nil
}

func (r *wht) Read(ctx context.Context) ([]*domain.Wht, error) {
	logger := lib.RequestCtxLogger(ctx)
	logger.Info().Msg("Read___START")

	records, err := boiled.WHTS().All(ctx, r.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to select wht: %w", err)
	}
	var results []*domain.Wht
	for _, r := range records {
		results = append(results, &domain.Wht{
			ID:         r.ID,
			RecordDate: r.RecordDate,
			Path:       r.Path,
		})
	}
	return results, nil
}
