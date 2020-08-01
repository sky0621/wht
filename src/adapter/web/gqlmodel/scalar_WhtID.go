package gqlmodel

import (
	"io"
	"log"
	"strconv"

	"golang.org/x/xerrors"
)

const (
	typeNameWhtID   = "WhtID"
)

type WhtID int64

// UnmarshalGQL GraphQL -> Domain
func (id *WhtID) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return xerrors.New("not string")
	}
	typeName, dbUniqueID, err := decodeID(s)
	if err != nil {
		return xerrors.Errorf("failed to decodeID[%s] - typeName:%s: %w", s, typeName, err)
	}
	*id = WhtID(dbUniqueID)
	return nil
}

// MarshalGQL Domain -> GraphQL
func (id WhtID) MarshalGQL(w io.Writer) {
	_, err := io.WriteString(w, strconv.Quote(id.NodeID()))
	if err != nil {
		log.Print(err)
	}
}

func (id WhtID) NodeID() string {
	return encodeID(typeNameWhtID, id.DBUniqueID())
}

func (id WhtID) DBUniqueID() int64 {
	return int64(id)
}

func (id WhtID) DBUniqueIDPtr() *int64 {
	r := id.DBUniqueID()
	return &r
}
