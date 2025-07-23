package graphql

import (
	"context"
	"errors"
	"log/slog"
	"math"

	"github.com/apartomat/apartomat/internal/crm"

	"github.com/apartomat/apartomat/internal/store/workspaceusers"
)

func (r *mutationResolver) InviteUser(
	ctx context.Context,
	workspaceID string,
	email string,
	role WorkspaceUserRole,
) (InviteUserToWorkspaceResult, error) {
	res, expiration, err := r.crm.InviteUserToWorkspace(ctx, workspaceID, email, workspaceusers.WorkspaceUserRole(role))
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		if errors.Is(err, crm.ErrAlreadyExists) {
			return AlreadyInWorkspace{Message: "user already in workspace"}, nil
		}

		slog.ErrorContext(ctx, "can't invite user", slog.Any("err", err))

		return serverError()
	}

	return InviteSent{To: res, TokenExpiration: int(math.Round(expiration.Seconds()))}, nil
}
