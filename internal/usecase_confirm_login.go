package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/token"
	"github.com/pkg/errors"
)

type ConfirmLogin struct {
	verifier token.EmailConfirmTokenVerifier
	issuer   token.AuthTokenIssuer
	users    store.UserStore
	acl      *Acl
}

func NewConfirmLogin(verifier token.EmailConfirmTokenVerifier, issuer token.AuthTokenIssuer, users store.UserStore, acl *Acl) *ConfirmLogin {
	return &ConfirmLogin{verifier: verifier, issuer: issuer, users: users, acl: acl}
}

func (u *ConfirmLogin) Do(ctx context.Context, str string) (string, error) {
	confirmToken, _, err := u.verifier.Verify(str)
	if err != nil {
		return "", err
	}

	email := confirmToken.Subject

	if !u.acl.CanConfirmLogin(ctx, nil, email) {
		return "", errors.Wrapf(ErrForbidden, "can't confirm login")
	}

	users, err := u.users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.Wrapf(ErrNotFound, "user %s", str)
	}

	return u.issuer.Issue(users[0].ID, confirmToken.Subject)
}
