package gqlmodel

// テキストコンテンツ
type TextContent struct {
	ID TextContentID `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// テキスト
	Text string `json:"text"`
}

// 画像コンテンツ
type ImageContent struct {
	ID ImageContentID `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// 画像パス
	Path string `json:"path"`
}

// 音声コンテンツ
type VoiceContent struct {
	ID VoiceContentID `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// 音声パス
	Path string `json:"path"`
}

// 動画コンテンツ
type MovieContent struct {
	ID MovieContentID `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// 動画パス
	Path string `json:"path"`
}
