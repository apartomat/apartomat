package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type GetWorkspace struct {
	workspaces store.WorkspaceStore
}

func NewGetWorkspace(
	workspaces store.WorkspaceStore,
) *GetWorkspace {
	return &GetWorkspace{workspaces}
}

func (u *GetWorkspace) Do(ctx context.Context, id int) (*store.Workspace, error) {
	workspaces, err := u.workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.IntEq(id)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace %d", id)
	}

	return workspaces[0], nil
}
