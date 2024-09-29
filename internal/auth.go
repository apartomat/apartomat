package apartomat

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"

	"github.com/apartomat/apartomat/internal/auth"
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

	user, err := u.Users.Get(ctx, EmailIn(email))
	if err != nil {
		return "", err
	}

	return u.AuthTokenIssuer.Issue(user.ID)
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

	user, err := u.Users.Get(ctx, EmailIn(email))
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return "", err
	}

	if user == nil {
		id, err := GenerateNanoID()
		if err != nil {
			return "", err
		}

		user = NewUser(id, email, "", true, false)

		if err := u.Users.Save(ctx, user); err != nil {
			return "", err
		}
	}

	workspace, err := u.Workspaces.Get(ctx, workspaces.UserIDIn(user.ID))
	if err != nil && !errors.Is(err, workspaces.ErrWorkspaceNotFound) {
		return "", err
	}

	if workspace == nil {
		id, err := GenerateNanoID()
		if err != nil {
			return "", err
		}

		workspace = workspaces.NewWorkspace(id, workspaceName, true, user.ID)

		if err := u.Workspaces.Save(ctx, workspace); err != nil {
			return "", err
		}

		wid, err := GenerateNanoID()
		if err != nil {
			return "", err
		}

		wu := workspace_users.NewWorkspaceUser(wid, workspace_users.WorkspaceUserRoleAdmin, workspace.ID, user.ID)

		if err := u.WorkspaceUsers.Save(ctx, wu); err != nil {
			return "", err
		}
	}

	if user.DefaultWorkspaceID == nil {
		user.DefaultWorkspaceID = &workspace.ID

		if err := u.Users.Save(ctx, user); err != nil {
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

	slog.DebugContext(ctx, "pin", slog.String("pin", pin))

	if u.Params.SendPinByEmail {
		slog.DebugContext(ctx, "send pin to", slog.String("email", email), slog.String("pin", pin))

		err = u.Mailer.Send(u.MailFactory.MailPIN(email, pin))
		if err != nil {
			return "", "", fmt.Errorf("sent error: %w", err)
		}
	}

	user, err := u.Users.Get(ctx, EmailIn(email))
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return "", "", err
	}

	if user == nil {
		id, err := GenerateNanoID()
		if err != nil {
			return "", "", err
		}

		user = NewUser(id, email, "", true, false)

		if err := u.Users.Save(ctx, user); err != nil {
			return "", "", err
		}
	}

	workspace, err := u.Workspaces.Get(ctx, workspaces.UserIDIn(user.ID))
	if err != nil && !errors.Is(err, workspaces.ErrWorkspaceNotFound) {
		return "", "", err
	}

	if workspace == nil {
		id, err := GenerateNanoID()
		if err != nil {
			return "", "", err
		}

		workspace = workspaces.NewWorkspace(id, workspaceName, true, user.ID)

		if err := u.Workspaces.Save(ctx, workspace); err != nil {
			return "", "", err
		}

		wid, err := GenerateNanoID()
		if err != nil {
			return "", "", err
		}

		wu := workspace_users.NewWorkspaceUser(wid, workspace_users.WorkspaceUserRoleAdmin, workspace.ID, user.ID)

		if err := u.WorkspaceUsers.Save(ctx, wu); err != nil {
			return "", "", err
		}
	}

	if user.DefaultWorkspaceID == nil {
		user.DefaultWorkspaceID = &workspace.ID

		if err := u.Users.Save(ctx, user); err != nil {
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

	user, err := u.Users.Get(ctx, EmailIn(confirmToken.Email()))
	if err != nil {
		return "", err
	}

	return u.AuthTokenIssuer.Issue(user.ID)
}

func (u *Apartomat) AcceptInviteToWorkspace(ctx context.Context, str string) (string, error) {
	confirmToken, err := u.InviteTokenVerifier.Verify(str)
	if err != nil {
		return "", err
	}

	workspace, err := u.Workspaces.Get(ctx, workspaces.IDIn(confirmToken.WorkspaceID()))
	if err != nil {
		return "", err
	}

	user, err := u.Users.Get(ctx, EmailIn(confirmToken.Email()))
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return "", err
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
			return "", err
		}

		if wus != nil {
			return "", fmt.Errorf("user is in workspace (id=%s) already: %w", confirmToken.WorkspaceID(), ErrAlreadyExists)
		}

	} else {
		id, err := GenerateNanoID()
		if err != nil {
			return "", err
		}

		user = NewUser(id, confirmToken.Email(), "", true, true)
	}

	{
		id, err := GenerateNanoID()
		if err != nil {
			return "", err
		}

		wuser := workspace_users.NewWorkspaceUser(
			id,
			workspace_users.WorkspaceUserRole(confirmToken.Role()),
			workspace.ID,
			user.ID,
		)

		if user.DefaultWorkspaceID == nil {
			user.DefaultWorkspaceID = &workspace.ID
		}

		if err := u.Users.Save(ctx, user); err != nil {
			return "", err
		}

		if err := u.WorkspaceUsers.Save(ctx, wuser); err != nil {
			return "", err
		}
	}

	return u.AuthTokenIssuer.Issue(user.ID)
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
