package graphql

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
	"log"

	apartomat "github.com/apartomat/apartomat/internal"
)

func (r *mutationResolver) InviteUser(
	ctx context.Context,
	workspaceID string,
	email string,
	role WorkspaceUserRole,
) (InviteUserToWorkspaceResult, error) {
	res, err := r.useCases.InviteUserToWorkspace(ctx, workspaceID, email, workspace_users.WorkspaceUserRole(role))
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		if errors.Is(err, apartomat.ErrAlreadyExists) {
			return AlreadyInWorkspace{Message: "user already in workspace"}, nil
		}

		log.Printf("can't invite user: %s", err)

		return ServerError{Message: "can't invite user"}, nil
	}

	return InviteSent{To: res}, nil
}
