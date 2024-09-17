package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

const (
	sendLink = false
)

func (r *mutationResolver) LoginByEmail(ctx context.Context, email string, workspace string) (LoginByEmailResult, error) {
	switch sendLink {
	case true:
		e, err := r.useCases.LoginByEmail(ctx, email, workspace)
		if err != nil {
			if errors.Is(err, apartomat.ErrInvalidEmail) {
				return InvalidEmail{Message: err.Error()}, nil
			}

			r.logger.Error("can't send token by an email", zap.Error(err))

			return serverError()
		}

		return LinkSentByEmail{Email: e}, nil
	default:
		e, token, err := r.useCases.LoginEmailPIN(ctx, email, workspace)
		if err != nil {
			if errors.Is(err, apartomat.ErrInvalidEmail) {
				return InvalidEmail{Message: err.Error()}, nil
			}

			r.logger.Error("can't send PIN by an email", zap.Error(err))

			return serverError()
		}

		return PinSentByEmail{Email: e, Token: token}, nil
	}
}
