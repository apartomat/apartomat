package graphql

import (
	"context"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
	"log"
)

func (r *queryResolver) Workspace(ctx context.Context, id int) (WorkspaceResult, error) {
	ws, err := r.useCases.GetWorkspace(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't resolve workspace (id=%d): %s", id, err)

		return ServerError{}, nil
	}

	return workspaceToGraphQL(ws), nil
}

func workspaceToGraphQL(w *store.Workspace) *Workspace {
	return &Workspace{
		ID:   w.ID,
		Name: w.Name,
	}
}
