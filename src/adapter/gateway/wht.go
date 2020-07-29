package gateway

import (
	"context"

	"github.com/sky0621/wht/adapter/gateway/sqlboilermodel"
	"github.com/sky0621/wht/service"
	"github.com/sky0621/wht/service/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/xerrors"
)

func NewWhtRepository(db boil.ContextExecutor) service.WhtRepository {
	return &whtRepository{db: db}
}

type whtRepository struct {
	db boil.ContextExecutor
}

func (r *whtRepository) Create(ctx context.Context, in domain.Wht) (int64, error) {
	mdl := &sqlboilermodel.WHT{
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
			mod = append(mod, sqlboilermodel.WHTWhere.ID.EQ(*condition.ID))
		}
		if condition.RecordDate != nil {
			mod = append(mod, sqlboilermodel.WHTWhere.RecordDate.EQ(*condition.RecordDate))
		}
	}
	records, err := sqlboilermodel.WHTS(mod...).All(ctx, r.db)
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
	mdl := &sqlboilermodel.WHT{
		ID:         *in.ID,
		RecordDate: in.RecordDate,
		Title:      null.StringFromPtr(in.Title),
	}
	if err := mdl.Upsert(ctx, r.db, true,
		[]string{sqlboilermodel.WHTColumns.RecordDate},
		boil.Whitelist(
			sqlboilermodel.WHTColumns.Title,
			sqlboilermodel.WHTColumns.UpdatedAt,
		),
		boil.Whitelist(
			sqlboilermodel.WHTColumns.RecordDate,
			sqlboilermodel.WHTColumns.Title,
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
