// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

type Node interface {
	IsNode()
}

type NoopInput struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type NoopPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
}

// 「今日こと」
type Wht struct {
	// ID
	ID string `json:"id"`
	// 記録日
	RecordDate string `json:"recordDate"`
	// タイトル
	Title *string `json:"title"`
	// テキスト
	Text string `json:"text"`
}

func (Wht) IsNode() {}

type WhtInput struct {
	// 記録日
	RecordDate string `json:"recordDate"`
	// タイトル
	Title *string `json:"title"`
	// テキスト
	Text string `json:"text"`
}
