package internal

import "strconv"

func PStr(v string) *string {
	return &v
}

func FromInt64ToPStr(v int64) *string {
	return PStr(strconv.Itoa(int(v)))
}
