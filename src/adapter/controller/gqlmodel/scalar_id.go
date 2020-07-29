package gqlmodel

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

const sep = ":"
const (
	typeNameWht = "wht"
)

func decodeID(id string) (string, int64, error) {
	b, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return "", 0, xerrors.Errorf("failed to DecodeString[%s]: %w", id, err)
	}
	items := strings.SplitN(string(b), sep, 2)
	dbUniqueID, err := strconv.ParseInt(items[1], 10, 64)
	if err != nil {
		return items[0], -1, xerrors.Errorf("failed to ParseInt[%s]: %w", items[1], err)
	}
	return items[0], dbUniqueID, nil
}

func encodeID(typeName string, dbUniqueID int64) string {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s%s%d", typeName, sep, dbUniqueID)))
}

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
	return encodeID(typeNameWht, id.DBUniqueID())
}

func (id WhtID) DBUniqueID() int64 {
	return int64(id)
}
