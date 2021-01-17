package apartomat

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ztsu/apartomat/internal/pkg/expr"
	"github.com/ztsu/apartomat/internal/store"
)

type ConfirmLogin struct {
	verifier EmailConfirmTokenVerifier
	issuer   AuthTokenIssuer
	users    store.UserStore
	acl      *Acl
}

func NewConfirmLogin(verifier EmailConfirmTokenVerifier, issuer AuthTokenIssuer, users store.UserStore, acl *Acl) *ConfirmLogin {
	return &ConfirmLogin{verifier: verifier, issuer: issuer, users: users, acl: acl}
}

func (uc *ConfirmLogin) Do(ctx context.Context, str string) (string, error) {
	confirmToken, _, err := uc.verifier.Verify(str)
	if err != nil {
		return "", err
	}

	email := confirmToken.Subject

	if !uc.acl.CanConfirmLogin(ctx, nil, email) {
		return "", errors.Wrapf(ErrForbidden, "can't confirm login")
	}

	users, err := uc.users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.Wrapf(ErrNotFound, "user %s", str)
	}

	return uc.issuer.Issue(users[0].ID, confirmToken.Subject)
}
