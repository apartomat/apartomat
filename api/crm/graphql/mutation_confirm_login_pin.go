package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm/auth/paseto"
)

func (r *mutationResolver) ConfirmLoginPin(ctx context.Context, t, pin string) (ConfirmLoginPinResult, error) {
	str, err := r.crm.CheckConfirmEmailPINToken(ctx, t, pin)
	if err != nil {
		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		slog.ErrorContext(ctx, "can't verify token", slog.Any("err", err))

		return serverError()
	}

	return LoginConfirmed{Token: str}, nil
}
