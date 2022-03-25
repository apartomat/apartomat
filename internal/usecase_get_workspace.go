package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

func (u *Apartomat) GetWorkspace(ctx context.Context, id string) (*store.Workspace, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.StrEq(id)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace %d", id)
	}

	workspace := workspaces[0]

	if ok, err := u.CanGetWorkspace(ctx, UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't get workspace (id=%d)", workspace.ID)
	}

	return workspaces[0], nil
}

func (u *Apartomat) CanGetWorkspace(ctx context.Context, subj *UserCtx, obj *store.Workspace) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.StrEq(obj.ID), UserID: expr.StrEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}
