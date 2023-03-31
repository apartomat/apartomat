package graphql

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/auth/paseto"
	"log"
)

func (r *mutationResolver) ConfirmLoginPin(ctx context.Context, t, pin string) (ConfirmLoginPinResult, error) {
	str, err := r.useCases.CheckConfirmEmailPINToken(ctx, t, pin)
	if err != nil {
		if errors.Is(err, paseto.ErrTokenValidationError) {
			return InvalidToken{Message: "token is expired or not valid"}, nil
		}

		log.Printf("can't verify token: %s", err)

		return ServerError{Message: "can't verify token"}, nil
	}

	return LoginConfirmed{Token: str}, nil
}
