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

func (r *content) CreateTextContents(ctx context.Context, whtID int64, inputs []*domain.TextContentForCreate) error {
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
			mods = append(mods, boiled.ContentTextWhere.ID.EQ(*condition.ID))
		}
		if condition.WhtID != nil {
			mods = append(mods, boiled.ContentTextWhere.WHTID.EQ(*condition.WhtID))
		} else if len(condition.WhtIDs) > 0 {
			mods = append(mods, boiled.ContentTextWhere.WHTID.IN(condition.WhtIDs))
		}
	}
	contents, err := boiled.ContentTexts(mods...).All(ctx, r.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to insert content_text[condition:%#+v]: %w", condition, err)
	}
	var results []*domain.TextContent
	for _, c := range contents {
		results = append(results, ToTextContent(c))
	}
	return results, nil
}
