package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"time"
)

func (u *Apartomat) CreateProject(
	ctx context.Context,
	workspaceID string,
	name string,
	startAt, endAt *time.Time,
) (*store.Project, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.StrEq(workspaceID)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", workspaceID, ErrNotFound)
	}

	var (
		workspace = workspaces[0]
	)

	if ok, err := u.CanCreateProject(ctx, UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't create project in workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	id, err := NewNanoID()
	if err != nil {
		return nil, err
	}

	project := &store.Project{
		ID:          id,
		Name:        name,
		WorkspaceID: workspaceID,
		Status:      store.ProjectStatusNew,
		StartAt:     startAt,
		EndAt:       endAt,
	}

	return u.Projects.Save(ctx, project)
}

func (u *Apartomat) CanCreateProject(ctx context.Context, subj *UserCtx, obj *store.Workspace) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.StrEq(obj.ID), UserID: expr.StrEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return true, nil
}
