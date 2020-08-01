package gqlmodel

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

const sep = ":"

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
