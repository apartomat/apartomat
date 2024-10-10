package crm

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	. "github.com/apartomat/apartomat/internal/store/users"
)

func (u *CRM) GetUserProfile(ctx context.Context, id string) (*User, error) {
	user, err := u.Users.Get(ctx, IDIn(id))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanGetUserProfile(ctx, auth.UserFromCtx(ctx), user); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get user profile (id=%s): %w", user.ID, ErrForbidden)
	}

	return user, nil
}
