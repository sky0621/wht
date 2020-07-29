package rdb

import (
	"context"

	"github.com/sky0621/wht/adapter/rdb/boiled"
	"github.com/sky0621/wht/application"
	"github.com/sky0621/wht/application/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/xerrors"
)

func NewWhtRepository(db boil.ContextExecutor) application.WhtRepository {
	return &whtRepository{db: db}
}

type whtRepository struct {
	db boil.ContextExecutor
}

func (r *whtRepository) Create(ctx context.Context, in domain.Wht) (int64, error) {
	mdl := &boiled.WHT{
		RecordDate: in.RecordDate,
		Title:      null.StringFromPtr(in.Title),
	}
	if err := mdl.Insert(ctx, r.db, boil.Infer()); err != nil {
		return -1, xerrors.Errorf("failed to insert wht: %w", err)
	}
	return mdl.ID, nil
}

func (r *whtRepository) Read(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error) {
	var mod []qm.QueryMod
	if condition != nil {
		if condition.ID != nil {
			mod = append(mod, boiled.WHTWhere.ID.EQ(*condition.ID))
		}
		if condition.RecordDate != nil {
			mod = append(mod, boiled.WHTWhere.RecordDate.EQ(*condition.RecordDate))
		}
	}
	records, err := boiled.WHTS(mod...).All(ctx, r.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to select wht: %w", err)
	}
	var results []*domain.Wht
	for _, r := range records {
		results = append(results, &domain.Wht{
			ID:         &r.ID,
			RecordDate: r.RecordDate,
			Title:      r.Title.Ptr(),
		})
	}
	return results, nil
}

func (r *whtRepository) Upsert(ctx context.Context, in domain.Wht) (*domain.Wht, error) {
	mdl := &boiled.WHT{
		ID:         *in.ID,
		RecordDate: in.RecordDate,
		Title:      null.StringFromPtr(in.Title),
	}
	if err := mdl.Upsert(ctx, r.db, true,
		[]string{boiled.WHTColumns.RecordDate},
		boil.Whitelist(
			boiled.WHTColumns.Title,
			boiled.WHTColumns.UpdatedAt,
		),
		boil.Whitelist(
			boiled.WHTColumns.RecordDate,
			boiled.WHTColumns.Title,
		),
	); err != nil {
		return nil, xerrors.Errorf("failed to upsert wht: %w", err)
	}
	return &domain.Wht{
		ID:         &mdl.ID,
		RecordDate: mdl.RecordDate,
		Title:      mdl.Title.Ptr(),
	}, nil
}
