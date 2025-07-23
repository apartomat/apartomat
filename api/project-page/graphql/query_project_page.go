package graphql

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/project-page"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/projectpages"
	"log/slog"
)

func (r *queryResolver) ProjectPage(ctx context.Context, id string) (ProjectPageResult, error) {
	ps, err := r.projectPage.GetProjectPage(ctx, id)
	if err != nil {
		if errors.Is(err, project_page.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, store.ErrNotFound) {
			return notFound()
		}

		slog.Error("failed to resolve project page", err)

		return serverError()
	}

	return projectPageToGraphQL(ps), nil
}

func projectPageToGraphQL(p *projectpages.ProjectPage) ProjectPage {
	return ProjectPage{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
	}
}
