package graphql

import (
	"context"
	"errors"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *queryResolver) Files(ctx context.Context, id string) (*FilesScreen, error) {
	return &FilesScreen{
		Project: Project{ID: id},
	}, nil
}

func (r *rootResolver) FilesScreen() FilesScreenResolver {
	return &filesScreenResolver{r}
}

type filesScreenResolver struct {
	*rootResolver
}

func (r *filesScreenResolver) Project(ctx context.Context, obj *FilesScreen) (ProjectResult, error) {
	if p, ok := obj.Project.(Project); ok {
		project, err := r.useCases.GetProject(ctx, p.ID)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return forbidden()
			}

			if errors.Is(err, apartomat.ErrNotFound) {
				return notFound()
			}

			log.Printf("can't resolve project (id=%d): %s", p.ID, err)

			return serverError()
		}

		return projectToGraphQL(project), nil
	}

	log.Printf("obj.Project is not a Project")

	return serverError()
}

func (r *filesScreenResolver) Menu(ctx context.Context, obj *FilesScreen) (MenuResult, error) {
	if p, ok := obj.Project.(Project); ok {
		return MenuItems{Items: []*MenuItem{
			{Title: "Файлы", URL: fmt.Sprintf("/p/%d/files", p.ID)},
			{Title: "Комплектация", URL: fmt.Sprintf("/p/%d/spec", p.ID)},
		}}, nil
	}

	log.Printf("obj.Project is not a Project")

	return serverError()
}
