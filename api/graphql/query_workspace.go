package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	"log"
)

func (r *queryResolver) Workspace(ctx context.Context, id string) (WorkspaceResult, error) {
	ws, err := r.useCases.GetWorkspace(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't resolve workspace (id=%s): %s", id, err)

		return ServerError{}, nil
	}
	return workspaceToGraphQL(ws), nil
}

func workspaceToGraphQL(w *workspaces.Workspace) *Workspace {
	return &Workspace{
		ID:   w.ID,
		Name: w.Name,
	}
}
