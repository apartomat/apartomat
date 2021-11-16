package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

func (u *Apartomat) GetWorkspaceUsers(ctx context.Context, id, limit, offset int) ([]*store.WorkspaceUser, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.IntEq(id)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace %d", id)
	}

	workspace := workspaces[0]

	if !u.CanGetWorkspaceUsers(ctx, UserFromCtx(ctx), workspace) {
		return nil, errors.Wrapf(ErrForbidden, "can't get workspace %d users", workspace.ID)
	}

	wu, err := u.WorkspaceUsers.List(ctx, store.WorkspaceUserStoreQuery{WorkspaceID: expr.IntEq(id)})
	if err != nil {
		return nil, err
	}

	return wu, nil
}

func (u *Apartomat) CanGetWorkspaceUsers(ctx context.Context, subj *UserCtx, obj *store.Workspace) bool {
	// todo check subj is workspace owner or admin
	return true
}
