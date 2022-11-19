package apartomat

import (
	"context"
	"errors"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/projects"
	. "github.com/apartomat/apartomat/internal/store/users"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"math/rand"
)

const userCtxKey = "user"

type UserCtx struct {
	ID    string
	Email string
}

func WithUserCtx(ctx context.Context, userCtx *UserCtx) context.Context {
	return context.WithValue(ctx, userCtxKey, userCtx)
}

func UserFromCtx(ctx context.Context) *UserCtx {
	user, _ := ctx.Value(userCtxKey).(*UserCtx)
	return user
}

type AuthToken interface {
	UserID() string
}

type AuthTokenIssuer interface {
	Issue(id string) (string, error)
}

type AuthTokenVerifier interface {
	Verify(str string) (AuthToken, error)
}

func (u *Apartomat) CheckAuthToken(str string) (AuthToken, error) {
	return u.AuthTokenVerifier.Verify(str)
}

type EmailConfirmToken interface {
	Email() string
}

type EmailConfirmTokenIssuer interface {
	Issue(email string) (string, error)
}

type EmailConfirmTokenVerifier interface {
	Verify(str string) (EmailConfirmToken, error)
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
		workspace *store.Workspace
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

		user = New(id, email, "", true, false)

		if err := u.Users.Save(ctx, user); err != nil {
			return "", err
		}
	} else {
		user = users[0]
	}

	workspaces, err := u.Workspaces.List(ctx, store.WorkspaceStoreQuery{UserID: expr.StrEq(user.ID)})
	if err != nil {
		return "", err
	}

	if len(workspaces) == 0 {
		id, err := NewNanoID()
		if err != nil {
			return "", err
		}

		workspace = &store.Workspace{
			ID:       id,
			Name:     workspaceName,
			IsActive: true,
			UserID:   user.ID,
		}

		workspace, err = u.Workspaces.Save(ctx, workspace)
		if err != nil {
			return "", err
		}

		wid, err := NewNanoID()
		if err != nil {
			return "", err
		}

		wu := &store.WorkspaceUser{
			ID:          wid,
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

func (u *Apartomat) isWorkspaceUser(ctx context.Context, subj *UserCtx, obj *store.Workspace) (bool, error) {
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

func (u *Apartomat) isProjectUser(ctx context.Context, subj *UserCtx, obj *projects.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.StrEq(obj.WorkspaceID), UserID: expr.StrEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
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
		workspace *store.Workspace
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

		user = New(id, email, "", true, false)

		if err := u.Users.Save(ctx, user); err != nil {
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

type ConfirmEmailPINToken interface {
	Email() string
	PIN() string
}

type ConfirmEmailPINTokenIssuer interface {
	Issue(email, pin string) (string, error)
}

type ConfirmEmailPINTokenVerifier interface {
	Verify(str, pin string) (ConfirmEmailPINToken, error)
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
