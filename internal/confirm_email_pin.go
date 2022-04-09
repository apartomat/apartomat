package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type ConfirmEmailPINToken interface {
	Email() string
	PIN() string
}

type ConfirmEmailPINTokenIssuer interface {
	Issue(email, pin string) (string, error)
}

type ConfirmEmailPINTokenVerifier interface {
	Verify(str, pin string) (ConfirmEmailPINToken, error)
}

func (u *Apartomat) CheckConfirmEmailPINToken(ctx context.Context, str, pin string) (string, error) {
	confirmToken, err := u.ConfirmEmailPINTokenVerifier.Verify(str, pin)
	if err != nil {
		return "", err
	}

	if confirmToken.PIN() != pin {
		return "", ErrForbidden
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
