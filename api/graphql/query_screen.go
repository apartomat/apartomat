package graphql

import "context"

func (r *rootResolver) ScreenQuery() ScreenQueryResolver {
	return &screenQueryResolver{r}
}

type screenQueryResolver struct {
	*rootResolver
}

func (r *queryResolver) Screen(ctx context.Context) (*ScreenQuery, error) {
	return &ScreenQuery{}, nil
}
