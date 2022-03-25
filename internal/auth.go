package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/token"
	"github.com/pkg/errors"
)

func (u *Apartomat) ConfirmLogin(ctx context.Context, str string) (string, error) {
	confirmToken, _, err := u.ConfirmTokenByEmailVerifier.Verify(str)
	if err != nil {
		return "", err
	}

	email := confirmToken.Subject

	users, err := u.Users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.Wrapf(ErrNotFound, "user %s", str)
	}

	return u.AuthTokenIssuer.Issue(users[0].ID, confirmToken.Subject)
}

func (u *Apartomat) CheckAuthToken(str string) (*token.AuthToken, error) {
	authToken, _, err := u.AuthTokenVerifier.Verify(str)
	println("CheckAuthToken", authToken)
	if err != nil {
		println("=========", err.Error())
	}

	return authToken, err
}
