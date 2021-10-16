package graphql

import (
	"context"
)

type mutationResolver struct {
	*rootResolver
}

func (r *mutationResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

func (r *rootResolver) Mutation() MutationResolver { return &mutationResolver{r} }
