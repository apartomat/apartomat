package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"math/rand"
)

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
		user      *store.User
		workspace *store.Workspace
	)

	users, err := u.Users.List(ctx, store.UserStoreQuery{Email: expr.StrEq(email)})
	if err != nil {
		return "", "", err
	}

	if len(users) == 0 {
		id, err := NewNanoID()
		if err != nil {
			return "", "", err
		}

		user = &store.User{
			ID:       id,
			Email:    email,
			IsActive: true,
		}

		user, err = u.Users.Save(ctx, user)
		if err != nil {
			return "", "", err
		}
	} else {
		user = users[0]
	}

	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{UserID: expr.StrEq(user.ID)})
	if err != nil {
		return "", "", err
	}

	if len(workspaces) == 0 {
		id, err := NewNanoID()
		if err != nil {
			return "", "", err
		}

		workspace = &store.Workspace{
			ID:       id,
			Name:     workspaceName,
			IsActive: true,
			UserID:   user.ID,
		}

		workspace, err = u.Workspaces.Save(ctx, workspace)
		if err != nil {
			return "", "", err
		}

		wid, err := NewNanoID()
		if err != nil {
			return "", "", err
		}

		wu := &store.WorkspaceUser{
			ID:          wid,
			WorkspaceID: workspace.ID,
			UserID:      user.ID,
			Role:        store.WorkspaceUserRoleAdmin,
		}

		_, err = u.WorkspaceUsers.Save(ctx, wu)
		if err != nil {
			return "", "", err
		}
	}

	return email, token, nil
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
