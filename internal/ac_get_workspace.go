package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

func (acl *Acl) CanGetWorkspace(ctx context.Context, subj *UserCtx, obj *store.Workspace) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := acl.store.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.IntEq(obj.ID), UserID: expr.IntEq(subj.ID)},
	)
	if err != nil || len(wu) == 0 {
		return false, err
	}

	return wu[0].UserID == subj.ID, nil
}
