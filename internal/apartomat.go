package apartomat

import (
	"github.com/apartomat/apartomat/internal/image"
	"github.com/apartomat/apartomat/internal/mail"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/apartomat/apartomat/internal/token"
	"github.com/pkg/errors"
)

type Apartomat struct {
	AuthTokenIssuer   token.AuthTokenIssuer
	AuthTokenVerifier token.AuthTokenVerifier

	ConfirmTokenByEmailIssuer   token.EmailConfirmTokenIssuer
	ConfirmTokenByEmailVerifier token.EmailConfirmTokenVerifier

	Mailer   mail.Sender
	Uploader image.Uploader

	Contacts       contacts.Store
	Projects       store.ProjectStore
	ProjectFiles   store.ProjectFileStore
	Users          store.UserStore
	Workspaces     store.WorkspaceStore
	WorkspaceUsers store.WorkspaceUserStore
}

var (
	ErrForbidden     = errors.New("forbidden")
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)
