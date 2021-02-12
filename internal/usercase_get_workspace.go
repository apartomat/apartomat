package apartomat

import (
	"context"
	"github.com/pkg/errors"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
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

	workspace := workspaces[0]

	return workspace, nil
}
