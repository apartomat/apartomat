package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type GetWorkspaceUsers struct {
	workspaces     store.WorkspaceStore
	workspaceUsers store.WorkspaceUserStore
	acl            *Acl
}

func NewGetWorkspaceUsers(
	workspaces store.WorkspaceStore,
	workspaceUsers store.WorkspaceUserStore,
	acl *Acl,
) *GetWorkspaceUsers {
	return &GetWorkspaceUsers{workspaces, workspaceUsers, acl}
}

func (u *GetWorkspaceUsers) Do(ctx context.Context, id, limit, offset int) ([]*store.WorkspaceUser, error) {
	workspaces, err := u.workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.IntEq(id)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace %d", id)
	}

	workspace := workspaces[0]

	if !u.acl.CanGetWorkspaceUsers(ctx, UserFromCtx(ctx), workspace) {
		return nil, errors.Wrapf(ErrForbidden, "can't get workspace %d users", workspace.ID)
	}

	wu, err := u.workspaceUsers.List(ctx, store.WorkspaceUserStoreQuery{WorkspaceID: expr.IntEq(id)})
	if err != nil {
		return nil, err
	}

	return wu, nil
}
