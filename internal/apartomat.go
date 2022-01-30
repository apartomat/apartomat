package apartomat

import (
	"github.com/apartomat/apartomat/internal/image"
	"github.com/apartomat/apartomat/internal/mail"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/apartomat/apartomat/internal/token"
	"github.com/pkg/errors"
)

type Apartomat struct {
	AuthTokenIssuer   token.AuthTokenIssuer
	AuthTokenVerifier token.AuthTokenVerifier

	ConfirmTokenByEmailIssuer   token.EmailConfirmTokenIssuer
	ConfirmTokenByEmailVerifier token.EmailConfirmTokenVerifier

	Mailer      mail.Sender
	MailFactory *mail.Factory

	Uploader image.Uploader

	Contacts       contacts.Store
	Houses         houses.Store
	Projects       store.ProjectStore
	ProjectFiles   store.ProjectFileStore
	Rooms          rooms.Store
	Users          store.UserStore
	Workspaces     store.WorkspaceStore
	WorkspaceUsers store.WorkspaceUserStore
}

var (
	ErrForbidden     = errors.New("forbidden")
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)
