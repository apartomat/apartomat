package graphql

import (
	"context"
	"errors"

	"github.com/apartomat/apartomat/internal/auth/paseto"
	"go.uber.org/zap"
)

func (r *mutationResolver) ConfirmLoginLink(ctx context.Context, str string) (ConfirmLoginLinkResult, error) {
	str, err := r.useCases.ConfirmEmailByToken(ctx, str)
	if err != nil {
		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		r.logger.Error("can't verify token", zap.Error(err))

		return serverError()
	}

	return LoginConfirmed{Token: str}, nil
}
