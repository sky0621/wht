package gqlmodel

import "github.com/sky0621/wht/application/domain"

// ------------------------------------------------------------------
// gqlmodel -> domain
// ------------------------------------------------------------------

func ToWhtForCreate(in WhtInput) *domain.WhtForCreate {
	return &domain.WhtForCreate{
		RecordDate: in.RecordDate,
		Title:      in.Title,
	}
}

func FromWht(in *domain.Wht) Wht {
	if in == nil {
		return Wht{}
	}
	return Wht{
		ID:         WhtID(in.ID),
		RecordDate: in.RecordDate,
		Title:      in.Title,
	}
}

func FromWhts(inputs []*domain.Wht) []Wht {
	var results []Wht
	for _, in := range inputs {
		results = append(results, FromWht(in))
	}
	return results
}

func ToTextContentForCreate(in TextContentInput) *domain.TextContentForCreate {
	return &domain.TextContentForCreate{
		Name: in.Name,
		Text: in.Text,
	}
}

func ToTextContentForCreateSlice(inputs []TextContentInput) []*domain.TextContentForCreate {
	var contents []*domain.TextContentForCreate
	for _, in := range inputs {
		contents = append(contents, ToTextContentForCreate(in))
	}
	return contents
}

func FromTextContent(in *domain.TextContent) TextContent {
	if in == nil {
		return TextContent{}
	}
	return TextContent{
		ID:   TextContentID(in.ID),
		Name: in.Name,
		Text: in.Text,
	}
}

func FromTextContents(inputs []*domain.TextContent) []TextContent {
	var results []TextContent
	for _, in := range inputs {
		results = append(results, FromTextContent(in))
	}
	return results
}
