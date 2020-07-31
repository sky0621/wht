package rdb

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/sky0621/wht/application/repository"

	"github.com/sky0621/wht/adapter/rdb/boiled"
	"github.com/sky0621/wht/application/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/xerrors"
)

func NewContentRepository(db boil.ContextExecutor) repository.Content {
	return &content{db: db}
}

type content struct {
	db boil.ContextExecutor
}

func (r *content) CreateTextContents(ctx context.Context, whtID int64, inputs []domain.TextContentForCreate) error {
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

func (r *content) ReadTextContents(ctx context.Context, condition *domain.TextContentCondition) ([]*domain.TextContent, error) {
	var mods []qm.QueryMod
	if condition != nil {
		if condition.ID != nil {
			mods = append(mods, boiled.WHTWhere.ID.EQ(*condition.ID))
		}
		if condition.RecordDate != nil {
			mods = append(mods, boiled.WHTWhere.RecordDate.EQ(*condition.RecordDate))
		}
	}
	contents, err := boiled.ContentTexts(mods...).All(ctx, r.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to insert content_text[whtID:%d][in:%#+v]: %w", whtID, in, err)
	}
	// FIXME:
	var results []domain.TextContent
	for _, txtCon := range []*domain.TextContent{
		{ID: 100, Text: "今日は、いい１日だったよ。"},
		{ID: 200, Text: "今日も、いい１日だったよ。"},
	} {
		results = append(results, txtCon)
	}
	return results, nil
}
