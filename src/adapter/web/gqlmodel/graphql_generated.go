// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Node interface {
	IsNode()
}

type MutationResponse struct {
	ID *string `json:"id"`
}

type NoopInput struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type NoopPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
}

// 今日ことインプット
type WhtInput struct {
	// 記録日
	RecordDate time.Time `json:"recordDate"`
	// 画像
	Image graphql.Upload `json:"image"`
}
