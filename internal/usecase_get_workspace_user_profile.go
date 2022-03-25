package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

func (u *Apartomat) GetWorkspaceUserProfile(ctx context.Context, workspaceID, userID string) (*store.User, error) {
	if !u.CanGetWorkspaceUserProfile(
		ctx,
		UserFromCtx(ctx),
		struct{ WorkspaceID, UserID string }{workspaceID, userID},
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

func (u *Apartomat) CanGetWorkspaceUserProfile(ctx context.Context, subj *UserCtx, obj struct{ WorkspaceID, UserID string }) bool {
	// todo check subj has access to workspace
	return true
}
