package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
	"time"
)

func (u *Apartomat) CreateProject(
	ctx context.Context,
	workspaceID int,
	name string,
	startAt, endAt *time.Time,
) (*store.Project, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.IntEq(workspaceID)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "workspace (id=%d)", workspaceID)
	}

	var (
		workspace = workspaces[0]
	)

	if !u.CanCreateProject(ctx, UserFromCtx(ctx), workspace) {
		return nil, errors.Wrapf(ErrForbidden, "can't create project in workspace (id=%d)", workspace.ID)
	}

	project := &store.Project{
		Name:        name,
		WorkspaceID: workspaceID,
		Status:      store.ProjectStatusNew,
		StartAt:     startAt,
		EndAt:       endAt,
	}

	return u.Projects.Save(ctx, project)
}

func (u *Apartomat) CanCreateProject(ctx context.Context, subj *UserCtx, obj *store.Workspace) bool {
	return true
}
