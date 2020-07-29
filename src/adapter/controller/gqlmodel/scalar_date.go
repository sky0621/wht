package gqlmodel

import (
	"io"
	"log"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sky0621/wht/application/util"
	"golang.org/x/xerrors"
)

const dateLayout = "2006/01/02"

// UnmarshalDate GraphQL -> Domain
func UnmarshalDate(v interface{}) (time.Time, error) {
	s, ok := v.(string)
	if !ok {
		return time.Time{}, xerrors.New("not string")
	}
	t, err := time.ParseInLocation(dateLayout, s, util.JST)
	if err != nil {
		return time.Time{}, xerrors.Errorf("failed to ParseInLocation: %w", err)
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, util.JST), nil
}

// MarshalDate Domain -> GraphQL
func MarshalDate(v time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.Quote(v.Format(dateLayout)))
		if err != nil {
			log.Print(err)
		}
	})
}
