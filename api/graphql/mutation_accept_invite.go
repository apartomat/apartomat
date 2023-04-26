package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/auth/paseto"
	"go.uber.org/zap"
)

func (r *mutationResolver) AcceptInvite(ctx context.Context, str string) (AcceptInviteResult, error) {
	str, err := r.useCases.AcceptInviteToWorkspace(ctx, str)
	if err != nil {
		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		if errors.Is(err, apartomat.ErrAlreadyExists) {
			return AlreadyInWorkspace{Message: "user already in workspace"}, nil
		}

		r.logger.Error("can't accept invite", zap.Error(err))

		return serverError()
	}

	return InviteAccepted{Token: str}, nil
}
