package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) ConfirmLogin(ctx context.Context, token string) (ConfirmLoginResult, error) {
	str, err := r.useCases.ConfirmLogin.Do(ctx, token)
	if err != nil {
		if errors.Is(err, apartomat.ErrTokenValidationError) {
			return InvalidToken{Message: "token expired or not valid"}, nil
		}

		log.Printf("can't verify token: %s", err)

		return ServerError{Message: "can't verify token"}, nil
	}

	return LoginConfirmed{Token: str}, nil
}
