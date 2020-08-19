package domain

import "time"

// 今日こと
type Wht struct {
	ID         int64     // ID
	RecordDate time.Time // 記録日
	Path       string    // 画像パス
}

// TODO: validator.v9
type WhtForCreate struct {
	RecordDate time.Time // 記録日
	Path       string    // 画像パス
}
