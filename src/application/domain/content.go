package domain

// コンテンツ
type Content interface {
	IsContent()
}

type TextContent struct {
	ID   int64
	Name *string
	Text string
}

func (c *TextContent) IsContent() {}

// TODO: validator.v9
type TextContentForCreate struct {
	Name *string
	Text string
}

// TODO: validator.v9
type TextContentForUpdate struct {
	ID   int64
	Name *string
	Text string
}
