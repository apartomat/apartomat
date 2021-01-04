package graphql

import (
	"context"
	"errors"
	"github.com/ztsu/apartomat/internal"
	"log"
)

type mutationResolver struct {
	*rootResolver
}

func (r *mutationResolver) LoginByEmail(ctx context.Context, email string, workspace string) (LoginByEmailResult, error) {
	e, err := r.useCases.LoginByEmail.Do(ctx, email, workspace)
	if err != nil {
		if errors.Is(err, apartomat.ErrInvalidEmail) {
			return InvalidEmail{Message: err.Error()}, nil
		}

		log.Printf("can't send email with token: %s", err)

		return ServerError{Message: "can't send auth token via email"}, nil
	}

	return CheckEmail{Email: e}, nil
}

func (r *mutationResolver) ConfirmLogin(ctx context.Context, token string) (ConfirmLoginResult, error) {
	str, err := r.useCases.ConfirmLogin.Do(token)
	if err != nil {
		if errors.Is(err, apartomat.ErrTokenValidationError) {
			return InvalidToken{Message: "token expired or not valid"}, nil
		}

		log.Printf("can't verify token: %s", err)

		return ServerError{Message: "can't verify token"}, nil
	}

	return LoginConfirmed{Token: str}, nil
}
