package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input CreateProjectInput) (CreateProjectResult, error) {
	project, err := r.crm.CreateProject(ctx, input.WorkspaceID, input.Name, input.StartAt, input.EndAt)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(
			ctx,
			"can't create project in workspace",
			slog.String("workspace", input.WorkspaceID),
			slog.Any("err", err),
		)

		return serverError()
	}

	return ProjectCreated{Project: projectToGraphQL(project)}, nil
}
