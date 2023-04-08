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
) (string, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(workspaceID), 1, 0)
	if err != nil {
		return "", err
	}

	if len(ws) == 0 {
		return "", fmt.Errorf("workspace (id=%s): %w", workspaceID, ErrNotFound)
	}

	var (
		workspace = ws[0]
	)

	if ok, err := u.CanInviteUsersToWorkspace(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
		return "", err
	} else if !ok {
		return "", fmt.Errorf("can't invite users to workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	us, err := u.Users.List(ctx, users.EmailIn(email), 1, 0)
	if err != nil {
		return "", err
	}

	if len(us) != 0 {
		var (
			user = us[0]
		)

		wus, err := u.WorkspaceUsers.List(ctx, workspace_users.UserIDIn(user.ID), 1, 0)
		if err != nil {
			return "", err
		}

		if len(wus) == 1 {
			return "", fmt.Errorf("user (email=%s): %w", email, ErrAlreadyExists)
		}
	}

	token, err := u.InviteTokenIssuer.Issue(email, workspace.ID, string(role))
	if err != nil {
		return "", err
	}

	err = u.Mailer.Send(u.MailFactory.MailAuth(email, token))
	if err != nil {
		return "", fmt.Errorf("can't send email to %s: %w", email, err)
	}

	return email, err
}

func (u *Apartomat) CanInviteUsersToWorkspace(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	return u.isWorkspaceUserAndRoleIs(ctx, subj, obj, workspace_users.WorkspaceUserRoleAdmin)
}

func (u *Apartomat) AcceptInviteToWorkspace(ctx context.Context, str string) (string, error) {
	confirmToken, err := u.InviteTokenVerifier.Verify(str)
	if err != nil {
		return "", err
	}

	var (
		user      *users.User
		workspace *workspaces.Workspace
	)

	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(confirmToken.WorkspaceID()), 1, 0)
	if err != nil {
		return "", err
	}

	if len(ws) == 0 {
		return "", fmt.Errorf("can't find workspace (id=%s): %w", confirmToken.WorkspaceID(), ErrNotFound)
	}

	workspace = ws[0]

	us, err := u.Users.List(ctx, users.EmailIn(confirmToken.Email()), 1, 0)
	if err != nil {
		return "", err
	}

	if len(us) != 0 {
		user = us[0]

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
			return "", err
		}

		if len(wus) != 0 {
			return "", fmt.Errorf("user is in workspace (id=%s) already: %w", confirmToken.WorkspaceID(), ErrAlreadyExists)
		}

	} else {
		id, err := NewNanoID()
		if err != nil {
			return "", err
		}

		user = users.NewUser(id, confirmToken.Email(), "", true, true)

		if err := u.Users.Save(ctx, user); err != nil {
			return "", err
		}
	}

	{
		id, err := NewNanoID()
		if err != nil {
			return "", err
		}

		wuser := workspace_users.NewWorkspaceUser(
			id,
			workspace_users.WorkspaceUserRole(confirmToken.Role()),
			workspace.ID,
			user.ID,
		)

		if err := u.WorkspaceUsers.Save(ctx, wuser); err != nil {
			return "", err
		}
	}

	return u.AuthTokenIssuer.Issue(us[0].ID)
}
