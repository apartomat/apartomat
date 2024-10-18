package graphql

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/project-page"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/public_sites"
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

		slog.Error("failed to resolve project-page site", err)

		return serverError()
	}

	return projectPageToGraphQL(ps), nil
}

func projectPageToGraphQL(site *public_sites.PublicSite) ProjectPage {
	return ProjectPage{
		ID: site.ID,
	}
}
