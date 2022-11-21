package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/dataloader"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/users"
	"github.com/apartomat/apartomat/internal/store/workspaces"
)

func (u *Apartomat) GetWorkspace(ctx context.Context, id string) (*workspaces.Workspace, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(id), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(ws) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", id, ErrNotFound)
	}

	workspace := ws[0]

	if ok, err := u.CanGetWorkspace(ctx, UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	return ws[0], nil
}

func (u *Apartomat) CanGetWorkspace(ctx context.Context, subj *UserCtx, obj *workspaces.Workspace) (bool, error) {
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

	return wu[0].UserID == subj.ID, nil
}

func (u *Apartomat) GetDefaultWorkspace(ctx context.Context, userID string) (*workspaces.Workspace, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.UserIDIn(userID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(ws) == 0 {
		return nil, fmt.Errorf("workspace of user (id=%s): %w", userID, ErrNotFound)
	}

	return ws[0], nil
}

func (u *Apartomat) GetWorkspaceProjects(
	ctx context.Context,
	workspaceID string,
	status []projects.Status,
	limit,
	offset int,
) ([]*projects.Project, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(workspaceID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(ws) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", workspaceID, ErrNotFound)
	}

	workspace := ws[0]

	if ok, err := u.CanGetWorkspaceProjects(ctx, UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s) projects: %w", workspace.ID, ErrForbidden)
	}

	p, err := u.Projects.List(ctx,
		projects.And(
			projects.WorkspaceIDIn(workspaceID),
			projects.StatusIn(status...),
		),
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *Apartomat) CanGetWorkspaceProjects(ctx context.Context, subj *UserCtx, obj *workspaces.Workspace) (bool, error) {
	return u.isWorkspaceUser(ctx, subj, obj)
}

func (u *Apartomat) GetWorkspaceUserProfile(ctx context.Context, workspaceID, userID string) (*users.User, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(workspaceID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(ws) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", workspaceID, ErrNotFound)
	}

	workspace := ws[0]

	if ok, err := u.CanGetWorkspaceUsers(ctx, UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s) users: %w", workspace.ID, ErrForbidden)
	}

	loader, err := dataloader.UserLoaderFromCtx(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get workspace user profile: %w", err)
	}

	user, err := loader.Load(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Apartomat) GetWorkspaceUsers(ctx context.Context, id string, limit, offset int) ([]*store.WorkspaceUser, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(id), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(ws) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", id, ErrNotFound)
	}

	workspace := ws[0]

	if ok, err := u.CanGetWorkspaceUsers(ctx, UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s) users: %w", workspace.ID, ErrForbidden)
	}

	wu, err := u.WorkspaceUsers.List(ctx, store.WorkspaceUserStoreQuery{WorkspaceID: expr.StrEq(id)})
	if err != nil {
		return nil, err
	}

	return wu, nil
}

func (u *Apartomat) CanGetWorkspaceUsers(ctx context.Context, subj *UserCtx, obj *workspaces.Workspace) (bool, error) {
	return u.isWorkspaceUser(ctx, subj, obj)
}
