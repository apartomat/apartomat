package graphql

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/token"
	"log"
)

func (r *mutationResolver) ConfirmLogin(ctx context.Context, str string) (ConfirmLoginResult, error) {
	str, err := r.useCases.ConfirmLogin.Do(ctx, str)
	if err != nil {
		if errors.Is(err, token.ErrTokenValidationError) {
			return InvalidToken{Message: "token expired or not valid"}, nil
		}

		log.Printf("can't verify token: %s", err)

		return ServerError{Message: "can't verify token"}, nil
	}

	return LoginConfirmed{Token: str}, nil
}
