package graphql

import (
	"context"
	"errors"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *screenQueryResolver) Project(ctx context.Context, obj *ScreenQuery, id string) (*ProjectScreen, error) {
	return &ProjectScreen{
		Project: Project{ID: id},
	}, nil
}

func (r *rootResolver) ProjectScreen() ProjectScreenResolver {
	return &projectScreenResolver{r}
}

type projectScreenResolver struct {
	*rootResolver
}

func (r *projectScreenResolver) Project(ctx context.Context, obj *ProjectScreen) (ProjectResult, error) {
	if p, ok := obj.Project.(Project); ok {
		project, err := r.useCases.GetProject(ctx, p.ID)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return forbidden()
			}

			if errors.Is(err, apartomat.ErrNotFound) {
				return notFound()
			}

			log.Printf("can't resolve project (id=%s): %s", p.ID, err)

			return nil, fmt.Errorf("can't resolve project (id=%s): %w", p.ID, err)
		}

		return projectToGraphQL(project), nil
	}

	log.Printf("obj.Project is not a Project")

	return nil, errors.New("obj.Project is not a Project")
}

func (r *projectScreenResolver) Menu(ctx context.Context, obj *ProjectScreen) (MenuResult, error) {
	if p, ok := obj.Project.(Project); ok {
		return MenuItems{Items: []*MenuItem{
			{Title: "Файлы", URL: fmt.Sprintf("/p/%s/files", p.ID)},
			{Title: "Комплектация", URL: fmt.Sprintf("/p/%s/spec", p.ID)},
		}}, nil
	}

	log.Printf("obj.Project is not a Project")

	return serverError()
}

func (r *projectScreenResolver) Enums(ctx context.Context, obj *ProjectScreen) (*ProjectEnums, error) {
	return &ProjectEnums{
		Status: &ProjectStatusEnum{
			Items: []*ProjectStatusEnumItem{
				{ProjectStatusNew, "Новый"},
				{ProjectStatusInProgress, "В работе"},
				{ProjectStatusDone, "Завершен"},
				{ProjectStatusCanceled, "Отменен"},
			},
		},
	}, nil
}
