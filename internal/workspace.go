package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/dataloader"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/users"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	"time"
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

	if ok, err := u.CanGetWorkspace(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	return ws[0], nil
}

func (u *Apartomat) CanGetWorkspace(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(obj.ID),
			workspace_users.UserIDIn(subj.ID),
		),
		1,
		0,
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

	if ok, err := u.CanGetWorkspaceProjects(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
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

func (u *Apartomat) CanGetWorkspaceProjects(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
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

	if ok, err := u.CanGetWorkspaceUsers(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
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

func (u *Apartomat) GetWorkspaceUsers(ctx context.Context, id string, limit, offset int) ([]*workspace_users.WorkspaceUser, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(id), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(ws) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", id, ErrNotFound)
	}

	workspace := ws[0]

	if ok, err := u.CanGetWorkspaceUsers(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s) users: %w", workspace.ID, ErrForbidden)
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.WorkspaceIDIn(id),
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return wu, nil
}

func (u *Apartomat) CanGetWorkspaceUsers(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	return u.isWorkspaceUser(ctx, subj, obj)
}

func (u *Apartomat) InviteUserToWorkspace(
	ctx context.Context,
	workspaceID string,
	email string,
	role workspace_users.WorkspaceUserRole,
) (string, time.Duration, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(workspaceID), 1, 0)
	if err != nil {
		return "", 0, err
	}

	if len(ws) == 0 {
		return "", 0, fmt.Errorf("workspace (id=%s): %w", workspaceID, ErrNotFound)
	}

	var (
		workspace = ws[0]
	)

	if ok, err := u.CanInviteUsersToWorkspace(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
		return "", 0, err
	} else if !ok {
		return "", 0, fmt.Errorf("can't invite users to workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	us, err := u.Users.List(ctx, users.EmailIn(email), 1, 0)
	if err != nil {
		return "", 0, err
	}

	if len(us) != 0 {
		var (
			user = us[0]
		)

		wus, err := u.WorkspaceUsers.List(
			ctx,
			workspace_users.And(
				workspace_users.UserIDIn(user.ID),
				workspace_users.WorkspaceIDIn(workspace.ID),
			),
			1,
			0,
		)
		if err != nil {
			return "", 0, err
		}

		if len(wus) == 1 {
			return "", 0, fmt.Errorf("user (email=%s): %w", email, ErrAlreadyExists)
		}
	}

	const (
		tokenExpiration = 60 * time.Minute
	)

	token, err := u.InviteTokenIssuer.Issue(email, workspace.ID, string(role), tokenExpiration)
	if err != nil {
		return "", 0, err
	}

	err = u.Mailer.Send(u.MailFactory.MailInvite(email, token))
	if err != nil {
		return "", 0, fmt.Errorf("can't send email to %s: %w", email, err)
	}

	return email, tokenExpiration, err
}

func (u *Apartomat) CanInviteUsersToWorkspace(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	return u.isWorkspaceUserAndRoleIs(ctx, subj, obj, workspace_users.WorkspaceUserRoleAdmin)
}
