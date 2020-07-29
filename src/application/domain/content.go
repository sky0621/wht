package domain

func NewTextContent(name *string, text string) TextContent {
	return TextContent{
		Name: name,
		Text: text,
	}
}

type TextContent struct {
	Name *string
	Text string
}
