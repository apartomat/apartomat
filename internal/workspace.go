package apartomat

import (
	"context"
	"errors"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/dataloaders"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/users"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	"time"
)

func (u *Apartomat) GetWorkspace(ctx context.Context, id string) (*workspaces.Workspace, error) {
	workspace, err := u.Workspaces.Get(ctx, workspaces.IDIn(id))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanGetWorkspace(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	return workspace, nil
}

func (u *Apartomat) GetDefaultWorkspace(ctx context.Context, userID string) (*workspaces.Workspace, error) {
	workspace, err := u.Workspaces.Get(ctx, workspaces.UserIDIn(userID))
	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (u *Apartomat) GetWorkspaceProjects(
	ctx context.Context,
	workspaceID string,
	status []projects.Status,
	limit,
	offset int,
) ([]*projects.Project, error) {
	if ok, err := u.Acl.CanGetWorkspaceProjectsOfWorkspaceID(ctx, auth.UserFromCtx(ctx), workspaceID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s) projects: %w", workspaceID, ErrForbidden)
	}

	res, err := u.Projects.List(ctx,
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

	return res, nil
}

func (u *Apartomat) GetWorkspaceUserProfileDl(ctx context.Context, workspaceID, userID string) (*users.User, error) {
	var (
		loaders = dataloaders.FromContext(ctx)
	)
	if loaders == nil {
		return nil, errors.New("can't get dataloaders from context")
	}

	if ok, err := u.Acl.CanGetWorkspaceUsersOfWorkspaceID(ctx, auth.UserFromCtx(ctx), workspaceID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s) users: %w", workspaceID, ErrForbidden)
	}

	user, err := loaders.Users.Load(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Apartomat) GetWorkspaceUsers(ctx context.Context, workspaceID string, limit, offset int) ([]*workspace_users.WorkspaceUser, error) {
	if ok, err := u.Acl.CanGetWorkspaceUsersOfWorkspaceID(ctx, auth.UserFromCtx(ctx), workspaceID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get workspace (id=%s) users: %w", workspaceID, ErrForbidden)
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.WorkspaceIDIn(workspaceID),
		workspace_users.SortDefault,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return wu, nil
}

func (u *Apartomat) InviteUserToWorkspace(
	ctx context.Context,
	workspaceID string,
	email string,
	role workspace_users.WorkspaceUserRole,
) (string, time.Duration, error) {
	workspace, err := u.Workspaces.Get(ctx, workspaces.IDIn(workspaceID))
	if err != nil {
		return "", 0, err
	}

	if ok, err := u.Acl.CanInviteUsersToWorkspace(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
		return "", 0, err
	} else if !ok {
		return "", 0, fmt.Errorf("can't invite users to workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	user, err := u.Users.Get(ctx, users.EmailIn(email))
	if err != nil && errors.Is(err, users.ErrUserNotFound) {
		return "", 0, err
	}

	if user != nil {
		wus, err := u.WorkspaceUsers.Get(
			ctx,
			workspace_users.And(
				workspace_users.UserIDIn(user.ID),
				workspace_users.WorkspaceIDIn(workspace.ID),
			),
		)
		if err != nil && !errors.Is(err, workspace_users.ErrWorkspaceUserNotFound) {
			return "", 0, err
		}

		if wus != nil {
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
