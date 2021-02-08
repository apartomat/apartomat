package apartomat

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ztsu/apartomat/internal/store"
	"github.com/ztsu/apartomat/internal/store/dataloader"
)

type GetWorkspaceUserProfile struct {
	users *dataloader.UserLoader
	acl   *Acl
}

func NewGetWorkspaceUserProfile(
	users *dataloader.UserLoader,
	acl *Acl,
) *GetWorkspaceUserProfile {
	return &GetWorkspaceUserProfile{users, acl}
}

func (u *GetWorkspaceUserProfile) Do(ctx context.Context, workspaceID, userID int) (*store.User, error) {
	if !u.acl.CanGetWorkspaceUserProfile(
		ctx,
		UserFromCtx(ctx),
		struct{ WorkspaceID, UserID int }{workspaceID, userID},
	) {
		return nil, errors.Wrapf(ErrForbidden, "can't get workspace %d users profile %d", workspaceID, userID)
	}

	user, err := u.users.Load(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
