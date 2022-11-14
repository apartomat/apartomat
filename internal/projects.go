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

func (u *Apartomat) ChangeProjectStatus(ctx context.Context, projectID string, status store.ProjectStatus) (*store.Project, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID), Limit: 1, Offset: 0})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanUpdateProject(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update project (id=%s): %w", project.ID, ErrForbidden)
	}

	if project.Status != status {
		project.Status = status
		return u.Projects.Save(ctx, project)
	}

	return project, nil
}

func (u *Apartomat) ChangeProjectDates(ctx context.Context, projectID string, startAt, endAt *time.Time) (*store.Project, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID), Limit: 1, Offset: 0})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanUpdateProject(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update project (id=%s): %w", project.ID, ErrForbidden)
	}

	project.StartAt = startAt
	project.EndAt = endAt

	return u.Projects.Save(ctx, project)
}

func (u *Apartomat) CanUpdateProject(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.StrEq(obj.WorkspaceID), UserID: expr.StrEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}

func (u *Apartomat) GetProject(ctx context.Context, id string) (*store.Project, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(id)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", id, ErrNotFound)
	}

	project := projects[0]

	if ok, err := u.CanGetProject(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s): %w", project.ID, ErrForbidden)
	}

	return project, nil
}

func (u *Apartomat) CanGetProject(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}
