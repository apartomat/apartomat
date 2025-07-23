package graphql

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/crm/auth/paseto"
	"log/slog"
)

func (r *mutationResolver) AcceptInvite(ctx context.Context, str string) (AcceptInviteResult, error) {
	str, err := r.crm.AcceptInviteToWorkspace(ctx, str)
	if err != nil {
		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		if errors.Is(err, crm.ErrAlreadyExists) {
			return AlreadyInWorkspace{Message: "user already in workspace"}, nil
		}

		slog.ErrorContext(ctx, "can't accept invite", slog.Any("err", err))

		return serverError()
	}

	return InviteAccepted{Token: str}, nil
}
