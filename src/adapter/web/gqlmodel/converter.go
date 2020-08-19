package gqlmodel

import "github.com/sky0621/wht/application/domain"

// ------------------------------------------------------------------
// gqlmodel -> domain
// ------------------------------------------------------------------

func ToWhtForCreate(in WhtInput) *domain.WhtForCreate {
	return &domain.WhtForCreate{
		RecordDate: in.RecordDate,
		Path:       in.Image.Filename,
	}
}
