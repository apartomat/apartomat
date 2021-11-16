package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

func (u *Apartomat) GetUserProfile(ctx context.Context, email string) (*store.User, error) {
	users, err := u.Users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "user %s", email)
	}

	return users[0], nil
}
