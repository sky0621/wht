package lib

import (
	"strconv"

	"github.com/volatiletech/null/v8"
)

func PStr(v string) *string {
	return &v
}

func FromInt64ToPStr(v int64) *string {
	return PStr(strconv.Itoa(int(v)))
}

func FromNullStringToPStr(v null.String) *string {
	return &v.String
}
