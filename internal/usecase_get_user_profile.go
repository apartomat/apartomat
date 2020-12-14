package apartomat

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ztsu/apartomat/internal/pkg/expr"
	"github.com/ztsu/apartomat/internal/store"
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
		return nil, ErrUserNotFound
	}

	return users[0], nil
}

var (
	ErrUserNotFound = errors.New("user not found")
)
