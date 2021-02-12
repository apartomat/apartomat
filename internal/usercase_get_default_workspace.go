package apartomat

import (
	"context"
	"github.com/pkg/errors"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

type GetDefaultWorkspace struct {
	workspaces store.WorkspaceStore
}

func NewGetDefaultWorkspace(
	workspaces store.WorkspaceStore,
) *GetDefaultWorkspace {
	return &GetDefaultWorkspace{workspaces}
}

func (u *GetDefaultWorkspace) Do(ctx context.Context, userID int) (*store.Workspace, error) {
	workspaces, err := u.workspaces.List(ctx, store.WorkspaceStoreQuery{UserID: expr.IntEq(userID), Limit: 1})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace of user %d", userID)
	}

	return workspaces[0], nil
}
