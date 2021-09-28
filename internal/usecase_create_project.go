package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type CreateProject struct {
	workspaces store.WorkspaceStore
	projects   store.ProjectStore
	acl        *Acl
}

func NewCreateProject(
	workspaces store.WorkspaceStore,
	projects store.ProjectStore,
	acl *Acl,
) *CreateProject {
	return &CreateProject{workspaces, projects, acl}
}

func (u *CreateProject) Do(ctx context.Context, workspaceID int, name string) (*store.Project, error) {
	workspaces, err := u.workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.IntEq(workspaceID)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace (id=%d)", workspaceID)
	}

	var (
		workspace = workspaces[0]
	)

	if !u.acl.CanCreateProject(ctx, UserFromCtx(ctx), workspace) {
		return nil, errors.Wrapf(ErrForbidden, "can't create project in workspace (id=%d)", workspace.ID)
	}

	project := &store.Project{
		Name:        name,
		IsActive:    true,
		WorkspaceID: workspaceID,
	}

	return u.projects.Save(ctx, project)
}
