package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"log"
)

func (r *queryResolver) Project(ctx context.Context, id int) (ProjectResult, error) {
	project, err := r.useCases.GetProject.Do(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		log.Printf("can't resolve project (id=%d): %s", id, err)

		return serverError()
	}

	return projectToGraphQL(project), nil
}

func projectToGraphQL(p *store.Project) *Project {
	if p == nil {
		return nil
	}

	return &Project{
		ID:    p.ID,
		Title: p.Name, // @todo
	}
}

func (r *rootResolver) Project() ProjectResolver { return &projectResolver{r} }

type projectResolver struct {
	*rootResolver
}

func (r *projectResolver) Files(ctx context.Context, obj *Project) (*ProjectFiles, error) {
	return &ProjectFiles{Project: &ID{ID: obj.ID}}, nil
}
