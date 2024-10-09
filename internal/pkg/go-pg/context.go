package go_pg

import (
	"context"
	"strings"
)

type Context string

const (
	queryContext Context = "query"
)

func WithQueryContext(ctx context.Context, val string) context.Context {
	var (
		qc = QueryContext(ctx)

		nval = make([]string, 0, 2)
	)

	if qc != "Unknown" {
		nval = append(nval, qc)
	}

	nval = append(nval, val)

	return context.WithValue(ctx, queryContext, strings.Join(nval, ":"))
}

func QueryContext(ctx context.Context) string {
	var (
		val = "Unknown"
	)

	if v := ctx.Value(queryContext); v != nil {
		if str, ok := v.(string); ok {
			val = str
		}
	}

	return val
}
