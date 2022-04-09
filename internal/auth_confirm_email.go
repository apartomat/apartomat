package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type EmailConfirmToken interface {
	Email() string
}

type EmailConfirmTokenIssuer interface {
	Issue(email string) (string, error)
}

type EmailConfirmTokenVerifier interface {
	Verify(str string) (EmailConfirmToken, error)
}

func (u *Apartomat) ConfirmEmailByToken(ctx context.Context, str string) (string, error) {
	confirmToken, err := u.ConfirmTokenByEmailVerifier.Verify(str)
	if err != nil {
		return "", err
	}

	email := confirmToken.Email()

	users, err := u.Users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.Wrapf(ErrNotFound, "user %s", str)
	}

	return u.AuthTokenIssuer.Issue(users[0].ID)
}
