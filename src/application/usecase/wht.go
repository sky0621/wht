package usecase

import (
	"context"

	"github.com/sky0621/wht/lib"

	"github.com/sky0621/wht/application/domain"
	"github.com/sky0621/wht/application/repository"
)

type Wht interface {
	CreateWht(ctx context.Context, in *domain.WhtForCreate) (int64, error)
	ReadWhts(ctx context.Context) ([]*domain.Wht, error)
}

func NewWht(whtRepo repository.Wht) Wht {
	return &wht{whtRepo: whtRepo}
}

type wht struct {
	whtRepo repository.Wht
}

/*
 * CreateWht 「今日こと」を作成し、作成した「今日こと」のユニークIDを返す。
 */
func (w wht) CreateWht(ctx context.Context, in *domain.WhtForCreate) (int64, error) {
	return w.whtRepo.Create(ctx, in)
}

/*
 * ReadWhts 全ての「今日こと」を返す。
 */
func (w wht) ReadWhts(ctx context.Context) ([]*domain.Wht, error) {
	logger := lib.RequestCtxLogger(ctx)
	logger.Info().Msg("ReadWhts___START")
	return w.whtRepo.Read(ctx)
}
