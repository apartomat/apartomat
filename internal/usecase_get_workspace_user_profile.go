package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type GetWorkspaceUserProfile struct {
	acl *Acl
}

func NewGetWorkspaceUserProfile(
	acl *Acl,
) *GetWorkspaceUserProfile {
	return &GetWorkspaceUserProfile{acl}
}

func (u *GetWorkspaceUserProfile) Do(ctx context.Context, workspaceID, userID int) (*store.User, error) {
	if !u.acl.CanGetWorkspaceUserProfile(
		ctx,
		UserFromCtx(ctx),
		struct{ WorkspaceID, UserID int }{workspaceID, userID},
	) {
		return nil, errors.Wrapf(ErrForbidden, "can't get workspace %d users profile %d", workspaceID, userID)
	}

	loader, err := UserLoaderFromCtx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can't get workspace user profile")
	}

	user, err := loader.Load(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
