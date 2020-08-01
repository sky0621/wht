package gqlmodel

import (
	"io"
	"log"
	"strconv"

	"golang.org/x/xerrors"
)

const (
	typeNameTextContentID   = "TextContentID"
)

type TextContentID int64

// UnmarshalGQL GraphQL -> Domain
func (id *TextContentID) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return xerrors.New("not string")
	}
	typeName, dbUniqueID, err := decodeID(s)
	if err != nil {
		return xerrors.Errorf("failed to decodeID[%s] - typeName:%s: %w", s, typeName, err)
	}
	*id = TextContentID(dbUniqueID)
	return nil
}

// MarshalGQL Domain -> GraphQL
func (id TextContentID) MarshalGQL(w io.Writer) {
	_, err := io.WriteString(w, strconv.Quote(id.NodeID()))
	if err != nil {
		log.Print(err)
	}
}

func (id TextContentID) NodeID() string {
	return encodeID(typeNameTextContentID, id.DBUniqueID())
}

func (id TextContentID) DBUniqueID() int64 {
	return int64(id)
}

func (id TextContentID) DBUniqueIDPtr() *int64 {
	r := id.DBUniqueID()
	return &r
}
