package domain

import "time"

// 今日こと
type Wht struct {
	// ID
	ID *int64
	// 記録日
	RecordDate time.Time
	// タイトル
	Title *string
}

// 「今日こと」検索条件
type WhtCondition struct {
	// ID
	ID *int64
	// 記録日
	RecordDate *time.Time
}
