package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm/auth/paseto"
)

func (r *mutationResolver) ConfirmLoginLink(ctx context.Context, str string) (ConfirmLoginLinkResult, error) {
	str, err := r.useCases.ConfirmEmailByToken(ctx, str)
	if err != nil {
		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		slog.ErrorContext(ctx, "can't verify token", slog.Any("err", err))

		return serverError()
	}

	return LoginConfirmed{Token: str}, nil
}
