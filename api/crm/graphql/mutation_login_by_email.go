package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

const (
	sendLink = false
)

func (r *mutationResolver) LoginByEmail(ctx context.Context, email string, workspace string) (LoginByEmailResult, error) {
	switch sendLink {
	case true:
		e, err := r.useCases.LoginByEmail(ctx, email, workspace)
		if err != nil {
			if errors.Is(err, crm.ErrInvalidEmail) {
				return InvalidEmail{Message: err.Error()}, nil
			}

			slog.ErrorContext(ctx, "can't send token by an email", slog.Any("err", err))

			return serverError()
		}

		return LinkSentByEmail{Email: e}, nil
	default:
		e, token, err := r.useCases.LoginEmailPIN(ctx, email, workspace)
		if err != nil {
			if errors.Is(err, crm.ErrInvalidEmail) {
				return InvalidEmail{Message: err.Error()}, nil
			}

			slog.ErrorContext(ctx, "can't send PIN by an email", slog.Any("err", err))

			return serverError()
		}

		return PinSentByEmail{Email: e, Token: token}, nil
	}
}
