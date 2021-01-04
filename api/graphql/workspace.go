package graphql

import (
	"context"
	"errors"
	apartomat "github.com/ztsu/apartomat/internal"
)

func (r *queryResolver) Workspace(ctx context.Context, id int) (WorkspaceResult, error) {
	ws, err := r.useCases.GetWorkspace.Do(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrWorkspaceNotFound) {
			return NotFound{}, nil
		}

		return ServerError{Message: "internal server error"}, nil
	}

	return workspaceToGraphQL(ws), nil
}
