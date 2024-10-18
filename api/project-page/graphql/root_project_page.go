package graphql

import (
	"context"
)

func (r *rootResolver) ProjectPage() ProjectPageResolver {
	return &projectPageResolver{r}
}

type projectPageResolver struct {
	*rootResolver
}

func (r *projectPageResolver) Visualizations(ctx context.Context, obj *ProjectPage) (*ProjectPageVisualizations, error) {
	return &ProjectPageVisualizations{}, nil
}
