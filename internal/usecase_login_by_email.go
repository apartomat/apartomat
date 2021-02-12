package apartomat

import (
	"context"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrSendError    = errors.New("can't send email")
)

// Use case: anonymous inputs own email address to get confirmation token
type LoginByEmail struct {
	users          store.UserStore
	workspaces     store.WorkspaceStore
	workspaceUsers store.WorkspaceUserStore
	issuer         EmailConfirmTokenIssuer
	mailer         MailSender
}

func NewLoginByEmail(
	users store.UserStore,
	workspaces store.WorkspaceStore,
	workspaceUsers store.WorkspaceUserStore,
	issuer EmailConfirmTokenIssuer,
	mailer MailSender,
) *LoginByEmail {
	return &LoginByEmail{users, workspaces, workspaceUsers, issuer, mailer}
}

func (lbe *LoginByEmail) Do(ctx context.Context, email string, workspaceName string) (string, error) {
	if err := validation.Validate(email, is.EmailFormat); err != nil {
		return "", ErrInvalidEmail
	}

	token, err := lbe.issuer.Issue(email)
	if err != nil {
		return "", err
	}

	err = lbe.mailer.Send(NewMailAuth("no-reply@zaibatsu.ru", email, token))
	if err != nil {
		return "", fmt.Errorf("sent error: %w", ErrSendError)
	}

	var (
		user      *store.User
		workspace *store.Workspace
	)

	users, err := lbe.users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		user = &store.User{
			Email:    email,
			IsActive: true,
		}

		user, err = lbe.users.Save(ctx, user)
		if err != nil {
			return "", err
		}
	} else {
		user = users[0]
	}

	workspaces, err := lbe.workspaces.List(ctx, store.WorkspaceStoreQuery{UserID: expr.IntEq(user.ID)})
	if err != nil {
		return "", err
	}

	if len(workspaces) == 0 {
		workspace = &store.Workspace{
			Name:     workspaceName,
			IsActive: true,
			UserID:   user.ID,
		}

		workspace, err = lbe.workspaces.Save(ctx, workspace)
		if err != nil {
			return "", err
		}

		wu := &store.WorkspaceUser{
			WorkspaceID: workspace.ID,
			UserID:      user.ID,
			Role:        store.WorkspaceUserRoleAdmin,
		}

		_, err = lbe.workspaceUsers.Save(ctx, wu)
		if err != nil {
			return "", err
		}
	}

	return email, nil
}
