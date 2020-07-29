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
	// タイトル
	Title *string `json:"title"`
	// コンテンツリスト
	//Contents []Content `json:"contents"`
}

func (Wht) IsNode() {}
