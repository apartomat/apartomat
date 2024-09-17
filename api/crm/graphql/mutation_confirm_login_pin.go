package graphql

import (
	"context"
	"errors"

	"github.com/apartomat/apartomat/internal/auth/paseto"
	"go.uber.org/zap"
)

func (r *mutationResolver) ConfirmLoginPin(ctx context.Context, t, pin string) (ConfirmLoginPinResult, error) {
	str, err := r.useCases.CheckConfirmEmailPINToken(ctx, t, pin)
	if err != nil {
		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		r.logger.Error("can't verify token", zap.Error(err))

		return serverError()
	}

	return LoginConfirmed{Token: str}, nil
}
