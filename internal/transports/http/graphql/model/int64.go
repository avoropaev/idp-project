// Package model describers marshal/unmarshal methods for custom type graphql
package model

import (
	"fmt"
	"io"

	gqlgen "github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
)

// MarshalInt64 is a marshal method of int64 type
func MarshalInt64(i int64) gqlgen.Marshaler {
	return gqlgen.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, fmt.Sprintf("%d", i))
	})
}

// UnmarshalInt64 is an unmarshal method of int64 type
func UnmarshalInt64(v interface{}) (int64, error) {
	if v, ok := v.(int64); ok {
		return v, nil
	}

	return int64(0), errors.Wrapf(ErrUnmarshal, "%T is not an int64", v)
}
