package graphql

import "context"

var (
	Version = "Development"
)

type queryResolver struct {
	*rootResolver
}

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return Version, nil
}

func (r *rootResolver) Query() QueryResolver { return &queryResolver{r} }
