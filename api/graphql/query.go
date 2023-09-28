package graphql

import "context"

var (
	version = "0"
)

type queryResolver struct {
	*rootResolver
}

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return version, nil
}

func (r *rootResolver) Query() QueryResolver { return &queryResolver{r} }
