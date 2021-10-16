package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

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
