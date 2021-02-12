package apartomat

import (
	"context"
	"github.com/pkg/errors"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

type GetUserProfile struct {
	users store.UserStore
}

func NewGetUserProfile(
	users store.UserStore,
) *GetUserProfile {
	return &GetUserProfile{users}
}

func (gup *GetUserProfile) Do(ctx context.Context, email string) (*store.User, error) {
	users, err := gup.users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "user %s", email)
	}

	return users[0], nil
}
