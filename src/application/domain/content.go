package domain

// テキストコンテンツ
type TextContent struct {
	ID    int64
	WhtID int64
	// コンテンツ名
	Name *string
	// テキスト
	Text string
}

// 画像コンテンツ
type ImageContent struct {
	ID    int64
	WhtID int64
	// コンテンツ名
	Name *string
	// 画像パス
	Path string
}

// 音声コンテンツ
type VoiceContent struct {
	ID    int64
	WhtID int64
	// コンテンツ名
	Name *string
	// 音声パス
	Path string
}

// 動画コンテンツ
type MovieContent struct {
	ID    int64
	WhtID int64
	// コンテンツ名
	Name *string
	// 動画パス
	Path string
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

// コンテンツ検索条件
type ContentCondition struct {
	ID     *int64  // ID
	WhtID  *int64  // 「今日こと」ID
	WhtIDs []int64 // 「今日こと」ID群（IN）
}
