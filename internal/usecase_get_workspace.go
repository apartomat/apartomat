package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type GetWorkspace struct {
	workspaces store.WorkspaceStore
	acl        *Acl
}

func NewGetWorkspace(
	workspaces store.WorkspaceStore,
	acl *Acl,
) *GetWorkspace {
	return &GetWorkspace{workspaces, acl}
}

func (u *GetWorkspace) Do(ctx context.Context, id int) (*store.Workspace, error) {
	workspaces, err := u.workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.IntEq(id)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace %d", id)
	}

	workspace := workspaces[0]

	if ok, err := u.acl.CanGetWorkspace(ctx, UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't get workspace (id=%d)", workspace.ID)
	}

	return workspaces[0], nil
}
