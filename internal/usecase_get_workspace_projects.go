package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

func (u *Apartomat) GetWorkspaceProjects(
	ctx context.Context,
	workspaceID int,
	filter GetWorkspaceProjectsFilter,
	limit,
	offset int,
) ([]*store.Project, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.IntEq(workspaceID)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace %d", workspaceID)
	}

	workspace := workspaces[0]

	if !u.CanGetWorkspaceProjects(ctx, UserFromCtx(ctx), workspace) {
		return nil, errors.Wrapf(ErrForbidden, "can't get workspace %d projects", workspace.ID)
	}

	p, err := u.Projects.List(
		ctx,
		store.ProjectStoreQuery{
			WorkspaceID: expr.IntEq(workspaceID),
			Status:      filter.Status,
			Limit:       limit,
			Offset:      offset,
		},
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

type GetWorkspaceProjectsFilter struct {
	Status store.ProjectStatusExpr
}

func (u *Apartomat) CanGetWorkspaceProjects(ctx context.Context, subj *UserCtx, obj *store.Workspace) bool {
	// todo check subj is workspace owner or admin
	return true
}
