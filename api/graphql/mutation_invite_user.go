package graphql

import (
	"context"
	"errors"
	"math"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
	"go.uber.org/zap"
)

func (r *mutationResolver) InviteUser(
	ctx context.Context,
	workspaceID string,
	email string,
	role WorkspaceUserRole,
) (InviteUserToWorkspaceResult, error) {
	res, expiration, err := r.useCases.InviteUserToWorkspace(ctx, workspaceID, email, workspace_users.WorkspaceUserRole(role))
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		if errors.Is(err, apartomat.ErrAlreadyExists) {
			return AlreadyInWorkspace{Message: "user already in workspace"}, nil
		}

		r.logger.Error("can't invite user", zap.Error(err))

		return serverError()
	}

	return InviteSent{To: res, TokenExpiration: int(math.Round(expiration.Seconds()))}, nil
}
