package repository

import (
	"context"

	"github.com/sky0621/wht/application/domain"
)

type Wht interface {
	Create(ctx context.Context, in domain.Wht) (int64, error)
	Read(ctx context.Context, condition *domain.WhtCondition) ([]*domain.Wht, error)
	Upsert(ctx context.Context, in domain.Wht) (*domain.Wht, error)
}

type Content interface {
	CreateTextContents(ctx context.Context, whtID int64, inputs []domain.TextContentForCreate) error
	ReadByWhtID(ctx context.Context, whtID int64) ([]domain.Content, error)
}
