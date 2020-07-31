package repository

import (
	"context"

	"github.com/sky0621/wht/application/domain"
)

type Content interface {
	CreateTextContents(ctx context.Context, whtID int64, inputs []*domain.TextContentForCreate) error
	ReadTextContents(ctx context.Context, condition *domain.TextContentCondition) ([]*domain.TextContent, error)
}
