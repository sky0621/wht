package rdb

import (
	"context"

	"github.com/sky0621/wht/adapter/rdb/boiled"
	"github.com/sky0621/wht/application"
	"github.com/sky0621/wht/application/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

func NewContentRepository(db boil.ContextExecutor) application.ContentRepository {
	return &contentRepository{db: db}
}

type contentRepository struct {
	db boil.ContextExecutor
}

func (r *contentRepository) CreateTextContents(ctx context.Context, whtID int64, inputs []domain.TextContentForCreate) error {
	// TODO: バッチ形式を検討！
	for _, in := range inputs {
		mdl := boiled.ContentText{
			WHTID: whtID,
			Name:  null.StringFromPtr(in.Name),
			Text:  in.Text,
		}
		if err := mdl.Insert(ctx, r.db, boil.Infer()); err != nil {
			return xerrors.Errorf("failed to insert content_text[whtID:%d][in:%#+v]: %w", whtID, in, err)
		}
	}
	return nil
}

func (r *contentRepository) ReadByWhtID(ctx context.Context, whtID int64) ([]domain.Content, error) {
	// FIXME:
	var results []domain.Content
	for _, txtCon := range []*domain.TextContent{
		{ID: 100, Text: "今日は、いい１日だったよ。"},
		{ID: 200, Text: "今日も、いい１日だったよ。"},
	} {
		results = append(results, txtCon)
	}
	return results, nil
}
