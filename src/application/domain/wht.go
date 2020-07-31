package domain

import "time"

// 今日こと
type Wht struct {
	ID            int64           // ID
	RecordDate    time.Time       // 記録日
	Title         *string         // タイトル
	TextContents  []*TextContent  // テキストコンテンツ群
	ImageContents []*ImageContent // 画像コンテンツ群
	VoiceContents []*VoiceContent // 音声コンテンツ群
	MovieContents []*MovieContent // 動画コンテンツ群
}

// 「今日こと」検索条件
type WhtCondition struct {
	ID         *int64     // ID
	RecordDate *time.Time // 記録日
}

// TODO: validator.v9
type WhtForCreate struct {
	RecordDate time.Time // 記録日
	Title      *string   // タイトル
}
