package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	"log/slog"
)

func (r *queryResolver) Workspace(ctx context.Context, id string) (WorkspaceResult, error) {
	ws, err := r.useCases.GetWorkspace(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
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
