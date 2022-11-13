package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

func (u *Apartomat) GetUserProfile(ctx context.Context, id string) (*store.User, error) {
	users, err := u.Users.List(ctx, store.UserStoreQuery{ID: expr.StrEq(id)})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user (id=%s): %w", id, ErrNotFound)
	}

	var (
		user = users[0]
	)

	if ok, err := u.CanGetUserProfile(ctx, UserFromCtx(ctx), user); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get user profile (id=%s): %w", user.ID, ErrForbidden)
	}

	return user, nil
}

func (u *Apartomat) CanGetUserProfile(ctx context.Context, subj *UserCtx, obj *store.User) (bool, error) {
	if subj == nil {
		return false, nil
	}

	return subj.ID == obj.ID, nil
}
