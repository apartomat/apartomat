package postgres

import "context"

type Context string

const (
	queryContext Context = "query"
)

func WithQueryContext(ctx context.Context, val string) context.Context {
	return context.WithValue(ctx, queryContext, val)
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
