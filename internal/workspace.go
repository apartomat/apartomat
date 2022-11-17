package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/dataloader"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/projects"
)

func (u *Apartomat) GetWorkspace(ctx context.Context, id string) (*store.Workspace, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.StrEq(id)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", id, ErrNotFound)
	}

	workspace := workspaces[0]

	if ok, err := u.CanGetWorkspace(ctx, UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	return workspaces[0], nil
}

func (u *Apartomat) CanGetWorkspace(ctx context.Context, subj *UserCtx, obj *store.Workspace) (bool, error) {
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

func (u *Apartomat) GetWorkspaceProjects(
	ctx context.Context,
	workspaceID string,
	status []projects.Status,
	limit,
	offset int,
) ([]*projects.Project, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.StrEq(workspaceID)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", workspaceID, ErrNotFound)
	}

	workspace := workspaces[0]

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

func (u *Apartomat) CanGetWorkspaceProjects(ctx context.Context, subj *UserCtx, obj *store.Workspace) (bool, error) {
	return u.isWorkspaceUser(ctx, subj, obj)
}

func (u *Apartomat) GetWorkspaceUserProfile(ctx context.Context, workspaceID, userID string) (*store.User, error) {
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.StrEq(workspaceID)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", workspaceID, ErrNotFound)
	}

	workspace := workspaces[0]

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
	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{ID: expr.StrEq(id)})
	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", id, ErrNotFound)
	}

	workspace := workspaces[0]

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

func (u *Apartomat) CanGetWorkspaceUsers(ctx context.Context, subj *UserCtx, obj *store.Workspace) (bool, error) {
	return u.isWorkspaceUser(ctx, subj, obj)
}
