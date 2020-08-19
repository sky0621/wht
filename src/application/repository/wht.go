package repository

import (
	"context"

	"github.com/sky0621/wht/application/domain"
)

type Wht interface {
	Create(ctx context.Context, in *domain.WhtForCreate) (int64, error)
	Read(ctx context.Context) ([]*domain.Wht, error)
}
