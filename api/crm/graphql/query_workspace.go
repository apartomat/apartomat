package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/workspaces"
)

func (r *queryResolver) Workspace(ctx context.Context, id string) (WorkspaceResult, error) {
	ws, err := r.crm.GetWorkspace(ctx, id)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't resolve workspace", slog.String("workspace", id), slog.Any("err", err))

		return serverError()
	}
	return workspaceToGraphQL(ws), nil
}

func workspaceToGraphQL(w *workspaces.Workspace) *Workspace {
	return &Workspace{
		ID:   w.ID,
		Name: w.Name,
	}
}
