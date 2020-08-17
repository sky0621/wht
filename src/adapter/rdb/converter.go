package rdb

import (
	"github.com/sky0621/wht/adapter/rdb/boiled"
	"github.com/sky0621/wht/application/domain"
	"github.com/sky0621/wht/lib"
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
		Name:  lib.FromNullStringToPStr(from.Name),
		Text:  from.Text,
	}
}

func ToImageContent(from *boiled.ContentImage) *domain.ImageContent {
	if from == nil {
		return nil
	}
	return &domain.ImageContent{
		ID:    from.ID,
		WhtID: from.WHTID,
		Name:  lib.FromNullStringToPStr(from.Name),
		Path:  from.Path,
	}
}
