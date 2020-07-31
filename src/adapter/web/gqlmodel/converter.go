package gqlmodel

import "github.com/sky0621/wht/application/domain"

func ToTextContentForCreate(inputs []TextContentInput) []domain.TextContentForCreate {
	var contents []domain.TextContentForCreate
	for _, in := range inputs {
		contents = append(contents,
			domain.TextContentForCreate{
				Name: in.Name,
				Text: in.Text,
			},
		)
	}
	return contents
}

func FromTextContent(inputs []*domain.TextContent) []*TextContent {
	var results []*TextContent
	for _, in := range inputs {
		results = append(results, &TextContent{
			ID:   in.ID,
			Name: in.Name,
			Text: in.Text,
		})
	}
	return results
}
