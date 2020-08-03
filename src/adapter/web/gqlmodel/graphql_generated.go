// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Node interface {
	IsNode()
}

// 画像コンテンツインプット
type ImageContentInput struct {
	// コンテンツ名
	Name *string `json:"name"`
	// 画像
	Image graphql.Upload `json:"image"`
}

// 動画コンテンツインプット
type MovieContentInput struct {
	// コンテンツ名
	Name *string `json:"name"`
	// 動画
	Movie graphql.Upload `json:"movie"`
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

// テキストコンテンツインプット
type TextContentInput struct {
	// コンテンツ名
	Name *string `json:"name"`
	// テキスト
	Text string `json:"text"`
}

// 音声コンテンツインプット
type VoiceContentInput struct {
	// コンテンツ名
	Name *string `json:"name"`
	// 音声
	Voice graphql.Upload `json:"voice"`
}

// 「今日こと」検索条件
type WhtConditionInput struct {
	// ID
	ID *WhtID `json:"id"`
	// 記録日
	RecordDate *time.Time `json:"recordDate"`
	// タイトル
	Title *string `json:"title"`
	// 検索方法
	Compare *Compare `json:"compare"`
}

// 今日ことインプット
type WhtInput struct {
	// 記録日
	RecordDate time.Time `json:"recordDate"`
	// タイトル
	Title *string `json:"title"`
}

// 検索方法
type Compare string

const (
	// 完全一致
	CompareEqual Compare = "Equal"
	// 曖昧検索
	CompareLike Compare = "Like"
)

var AllCompare = []Compare{
	CompareEqual,
	CompareLike,
}

func (e Compare) IsValid() bool {
	switch e {
	case CompareEqual, CompareLike:
		return true
	}
	return false
}

func (e Compare) String() string {
	return string(e)
}

func (e *Compare) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Compare(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Compare", str)
	}
	return nil
}

func (e Compare) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RoleManager Role = "MANAGER"
	RoleEditor  Role = "EDITOR"
	RoleViewer  Role = "VIEWER"
	RoleGuest   Role = "GUEST"
)

var AllRole = []Role{
	RoleAdmin,
	RoleManager,
	RoleEditor,
	RoleViewer,
	RoleGuest,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleManager, RoleEditor, RoleViewer, RoleGuest:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}