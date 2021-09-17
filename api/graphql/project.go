package graphql

import (
	"context"
	"github.com/apartomat/apartomat/internal/store"
)

func (r *rootResolver) Project() ProjectResolver { return &projectResolver{r} }

type projectResolver struct {
	*rootResolver
}

func (r *projectResolver) Files(ctx context.Context, obj *Project) (*ProjectFiles, error) {
	return &ProjectFiles{Project: &ID{ID: obj.ID}}, nil
}

func projectToGraphQL(p *store.Project) *Project {
	if p == nil {
		return nil
	}

	return &Project{
		ID:    p.ID,
		Title: p.Name,
	}
}
