package gqlmodel

import (
	"io"
	"log"
	"strconv"

	"golang.org/x/xerrors"
)

const (
	typeNameMovieContentID   = "MovieContentID"
)

type MovieContentID int64

// UnmarshalGQL GraphQL -> Domain
func (id *MovieContentID) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return xerrors.New("not string")
	}
	typeName, dbUniqueID, err := decodeID(s)
	if err != nil {
		return xerrors.Errorf("failed to decodeID[%s] - typeName:%s: %w", s, typeName, err)
	}
	*id = MovieContentID(dbUniqueID)
	return nil
}

// MarshalGQL Domain -> GraphQL
func (id MovieContentID) MarshalGQL(w io.Writer) {
	_, err := io.WriteString(w, strconv.Quote(id.NodeID()))
	if err != nil {
		log.Print(err)
	}
}

func (id MovieContentID) NodeID() string {
	return encodeID(typeNameMovieContentID, id.DBUniqueID())
}

func (id MovieContentID) DBUniqueID() int64 {
	return int64(id)
}

func (id MovieContentID) DBUniqueIDPtr() *int64 {
	r := id.DBUniqueID()
	return &r
}
