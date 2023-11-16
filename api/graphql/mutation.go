package graphql

import (
	"context"
)

type mutationResolver struct {
	*rootResolver
}

func (r *mutationResolver) Pass(ctx context.Context) (bool, error) {
	return true, nil
}

func (r *rootResolver) Mutation() MutationResolver { return &mutationResolver{r} }
