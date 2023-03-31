package apartomat

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/store/projects"
	. "github.com/apartomat/apartomat/internal/store/users"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (u *Apartomat) CheckAuthToken(str string) (auth.AuthToken, error) {
	return u.AuthTokenVerifier.Verify(str)
}

func (u *Apartomat) ConfirmEmailByToken(ctx context.Context, str string) (string, error) {
	confirmToken, err := u.ConfirmTokenByEmailVerifier.Verify(str)
	if err != nil {
		return "", err
	}

	email := confirmToken.Email()

	users, err := u.Users.List(ctx, EmailIn(email), 1, 0)
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", fmt.Errorf("user (email=%s): %w", email, ErrNotFound)
	}

	return u.AuthTokenIssuer.Issue(users[0].ID)
}

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrSendError    = errors.New("can't send email")
)

func (u *Apartomat) LoginByEmail(ctx context.Context, email string, workspaceName string) (string, error) {
	if err := validation.Validate(email, is.EmailFormat); err != nil {
		return "", ErrInvalidEmail
	}

	token, err := u.ConfirmTokenByEmailIssuer.Issue(email)
	if err != nil {
		return "", err
	}

	err = u.Mailer.Send(u.MailFactory.MailAuth(email, token))
	if err != nil {
		return "", fmt.Errorf("sent error: %w", err)
	}

	var (
		user      *User
		workspace *workspaces.Workspace
	)

	users, err := u.Users.List(ctx, EmailIn(email), 1, 0)
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		id, err := NewNanoID()
		if err != nil {
			return "", err
		}

		user = NewUser(id, email, "", true, false)

		if err := u.Users.Save(ctx, user); err != nil {
			return "", err
		}
	} else {
		user = users[0]
	}

	ws, err := u.Workspaces.List(ctx, workspaces.UserIDIn(user.ID), 1, 0)
	if err != nil {
		return "", err
	}

	if len(ws) == 0 {
		id, err := NewNanoID()
		if err != nil {
			return "", err
		}

		workspace = workspaces.NewWorkspace(id, workspaceName, true, user.ID)

		if err := u.Workspaces.Save(ctx, workspace); err != nil {
			return "", err
		}

		wid, err := NewNanoID()
		if err != nil {
			return "", err
		}

		wu := workspace_users.NewWorkspaceUser(wid, workspace_users.WorkspaceUserRoleAdmin, workspace.ID, user.ID)

		if err := u.WorkspaceUsers.Save(ctx, wu); err != nil {
			return "", err
		}
	}

	return email, nil
}

func (u *Apartomat) LoginEmailPIN(ctx context.Context, email string, workspaceName string) (string, string, error) {
	if err := validation.Validate(email, is.EmailFormat); err != nil {
		return "", "", ErrInvalidEmail
	}

	pin := randn(6)

	token, err := u.ConfirmEmailPINTokenIssuer.Issue(email, pin)
	if err != nil {
		return "", "", err
	}

	err = u.Mailer.Send(u.MailFactory.MailPIN(email, pin))
	if err != nil {
		return "", "", fmt.Errorf("sent error: %w", err)
	}

	var (
		user      *User
		workspace *workspaces.Workspace
	)

	users, err := u.Users.List(ctx, EmailIn(email), 1, 0)
	if err != nil {
		return "", "", err
	}

	if len(users) == 0 {
		id, err := NewNanoID()
		if err != nil {
			return "", "", err
		}

		user = NewUser(id, email, "", true, false)

		if err := u.Users.Save(ctx, user); err != nil {
			return "", "", err
		}
	} else {
		user = users[0]
	}

	ws, err := u.Workspaces.List(ctx, workspaces.UserIDIn(user.ID), 1, 0)
	if err != nil {
		return "", "", err
	}

	if len(ws) == 0 {
		id, err := NewNanoID()
		if err != nil {
			return "", "", err
		}

		workspace = workspaces.NewWorkspace(id, workspaceName, true, user.ID)

		if err := u.Workspaces.Save(ctx, workspace); err != nil {
			return "", "", err
		}

		wid, err := NewNanoID()
		if err != nil {
			return "", "", err
		}

		wu := workspace_users.NewWorkspaceUser(wid, workspace_users.WorkspaceUserRoleAdmin, workspace.ID, user.ID)

		if err := u.WorkspaceUsers.Save(ctx, wu); err != nil {
			return "", "", err
		}
	}

	return email, token, nil
}

func (u *Apartomat) CheckConfirmEmailPINToken(ctx context.Context, str, pin string) (string, error) {
	confirmToken, err := u.ConfirmEmailPINTokenVerifier.Verify(str, pin)
	if err != nil {
		return "", err
	}

	if confirmToken.PIN() != pin {
		return "", ErrForbidden
	}

	email := confirmToken.Email()

	users, err := u.Users.List(ctx, EmailIn(email), 1, 0)
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", fmt.Errorf("user (email=%s): %w", email, ErrNotFound)
	}

	return u.AuthTokenIssuer.Issue(users[0].ID)
}

func (u *Apartomat) isWorkspaceUser(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
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

func (u *Apartomat) isProjectUser(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(obj.WorkspaceID),
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

func randn(n int) string {
	var (
		digits = []rune("1234567890")

		res = make([]rune, n)
	)

	for i := range res {
		res[i] = digits[rand.Intn(len(digits))]
	}

	return string(res)
}
