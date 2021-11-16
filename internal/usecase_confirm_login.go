package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

func (u *Apartomat) ConfirmLogin(ctx context.Context, str string) (string, error) {
	confirmToken, _, err := u.ConfirmTokenByEmailVerifier.Verify(str)
	if err != nil {
		return "", err
	}

	email := confirmToken.Subject

	if !u.CanConfirmLogin(ctx, nil, email) {
		return "", errors.Wrapf(ErrForbidden, "can't confirm login")
	}

	users, err := u.Users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.Wrapf(ErrNotFound, "user %s", str)
	}

	return u.AuthTokenIssuer.Issue(users[0].ID, confirmToken.Subject)
}

func (u *Apartomat) CanConfirmLogin(ctx context.Context, subj *UserCtx, obj string) bool {
	return true
}
