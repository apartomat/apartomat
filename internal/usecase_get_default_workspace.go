package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

func (u *Apartomat) GetDefaultWorkspace(ctx context.Context, userID string) (*store.Workspace, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{UserID: expr.StrEq(userID), Limit: 1})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, fmt.Errorf("workspace of user (id=%s): %w", userID, ErrNotFound)
	}

	return workspaces[0], nil
}
