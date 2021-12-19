package apartomat

import (
	"context"
	"errors"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

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
		user      *store.User
		workspace *store.Workspace
	)

	users, err := u.Users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		user = &store.User{
			Email:    email,
			IsActive: true,
		}

		user, err = u.Users.Save(ctx, user)
		if err != nil {
			return "", err
		}
	} else {
		user = users[0]
	}

	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{UserID: expr.IntEq(user.ID)})
	if err != nil {
		return "", err
	}

	if len(workspaces) == 0 {
		workspace = &store.Workspace{
			Name:     workspaceName,
			IsActive: true,
			UserID:   user.ID,
		}

		workspace, err = u.Workspaces.Save(ctx, workspace)
		if err != nil {
			return "", err
		}

		wu := &store.WorkspaceUser{
			WorkspaceID: workspace.ID,
			UserID:      user.ID,
			Role:        store.WorkspaceUserRoleAdmin,
		}

		_, err = u.WorkspaceUsers.Save(ctx, wu)
		if err != nil {
			return "", err
		}
	}

	return email, nil
}
