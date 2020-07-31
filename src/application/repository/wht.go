package repository

import (
	"context"

	"github.com/sky0621/wht/application/domain"
)

type Wht interface {
	Create(ctx context.Context, in *domain.WhtForCreate) (int64, error)
	Read(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error)
	Upsert(ctx context.Context, in domain.Wht) (*domain.Wht, error)
}
