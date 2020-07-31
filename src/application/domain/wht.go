package domain

import "time"

// 今日こと
type Wht struct {
	ID         *int64    // ID
	RecordDate time.Time // 記録日
	Title      *string   // タイトル
	//Contents   []Content // コンテンツ群
}

// 「今日こと」検索条件
type WhtCondition struct {
	ID         *int64     // ID
	RecordDate *time.Time // 記録日
}
