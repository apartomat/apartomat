package apartomat

import (
	"context"
	"github.com/pkg/errors"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

type GetWorkspaceProjects struct {
	workspaces store.WorkspaceStore
	projects   store.ProjectStore
	acl        *Acl
}

func NewGetWorkspaceProjects(
	workspaces store.WorkspaceStore,
	projects store.ProjectStore,
	acl *Acl,
) *GetWorkspaceProjects {
	return &GetWorkspaceProjects{workspaces, projects, acl}
}

func (u *GetWorkspaceProjects) Do(ctx context.Context, id, limit, offset int) ([]*store.Project, error) {
	workspaces, err := u.workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.IntEq(id)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace %d", id)
	}

	workspace := workspaces[0]

	if !u.acl.CanGetWorkspaceProjects(ctx, UserFromCtx(ctx), workspace) {
		return nil, errors.Wrapf(ErrForbidden, "can't get workspace %d projects", workspace.ID)
	}

	p, err := u.projects.List(ctx, store.ProjectStoreQuery{WorkspaceID: expr.IntEq(id)})
	if err != nil {
		return nil, err
	}

	return p, nil
}
