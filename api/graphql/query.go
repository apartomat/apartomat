package graphql

import "context"

type queryResolver struct {
	*rootResolver
}

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return "", nil // TODO
}

func (r *rootResolver) Query() QueryResolver { return &queryResolver{r} }
