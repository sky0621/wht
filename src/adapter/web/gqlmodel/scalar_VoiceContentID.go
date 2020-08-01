package gqlmodel

import (
	"io"
	"log"
	"strconv"

	"golang.org/x/xerrors"
)

const (
	typeNameVoiceContentID   = "VoiceContentID"
)

type VoiceContentID int64

// UnmarshalGQL GraphQL -> Domain
func (id *VoiceContentID) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return xerrors.New("not string")
	}
	typeName, dbUniqueID, err := decodeID(s)
	if err != nil {
		return xerrors.Errorf("failed to decodeID[%s] - typeName:%s: %w", s, typeName, err)
	}
	*id = VoiceContentID(dbUniqueID)
	return nil
}

// MarshalGQL Domain -> GraphQL
func (id VoiceContentID) MarshalGQL(w io.Writer) {
	_, err := io.WriteString(w, strconv.Quote(id.NodeID()))
	if err != nil {
		log.Print(err)
	}
}

func (id VoiceContentID) NodeID() string {
	return encodeID(typeNameVoiceContentID, id.DBUniqueID())
}

func (id VoiceContentID) DBUniqueID() int64 {
	return int64(id)
}

func (id VoiceContentID) DBUniqueIDPtr() *int64 {
	r := id.DBUniqueID()
	return &r
}
