package rdb

import (
	"github.com/sky0621/wht/adapter/rdb/boiled"
	"github.com/sky0621/wht/application/domain"
	"github.com/sky0621/wht/application/util"
)

// ------------------------------------------------------------------
// boiled -> domain
// ------------------------------------------------------------------

func ToTextContent(from *boiled.ContentText) *domain.TextContent {
	if from == nil {
		return nil
	}
	return &domain.TextContent{
		ID:    from.ID,
		WhtID: from.WHTID,
		Name:  util.FromNullStringToPStr(from.Name),
		Text:  from.Text,
	}
}
