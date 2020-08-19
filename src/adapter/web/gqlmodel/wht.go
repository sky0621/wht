package gqlmodel

import (
	"time"
)

// 今日こと
type Wht struct {
	// ID
	ID WhtID `json:"id"`
	// 記録日
	RecordDate time.Time `json:"recordDate"`
	// 画像パス
	Path string `json:"path"`
}

func (Wht) IsNode() {}
