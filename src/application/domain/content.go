package domain

// テキストコンテンツ
type TextContent struct {
	ID string `json:"id"`
	ID int64
	// コンテンツ名
	Name *string `json:"name"`
	// テキスト
	Text string `json:"text"`
}

// 画像コンテンツ
type ImageContent struct {
	ID string `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// 画像パス
	Path string `json:"path"`
}

// 音声コンテンツ
type VoiceContent struct {
	ID string `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// 音声パス
	Path string `json:"path"`
}

// 動画コンテンツ
type MovieContent struct {
	ID string `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// 動画パス
	Path string `json:"path"`
}

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

// テキストコンテンツ検索条件
type TextContentCondition struct {
	ID     *int64  // ID
	WhtID  *int64  // 「今日こと」ID
	WhtIDs []int64 // 「今日こと」ID群（IN）
}
