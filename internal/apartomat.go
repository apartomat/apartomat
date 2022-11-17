package apartomat

import (
	"errors"
	"github.com/apartomat/apartomat/internal/image"
	"github.com/apartomat/apartomat/internal/mail"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/apartomat/apartomat/internal/store/visualizations"
	nanoid "github.com/matoous/go-nanoid/v2"
)

type Apartomat struct {
	AuthTokenIssuer   AuthTokenIssuer
	AuthTokenVerifier AuthTokenVerifier

	ConfirmTokenByEmailIssuer   EmailConfirmTokenIssuer
	ConfirmTokenByEmailVerifier EmailConfirmTokenVerifier

	ConfirmEmailPINTokenIssuer   ConfirmEmailPINTokenIssuer
	ConfirmEmailPINTokenVerifier ConfirmEmailPINTokenVerifier

	Mailer      mail.Sender
	MailFactory *mail.Factory

	Uploader image.Uploader

	Contacts       contacts.Store
	Houses         houses.Store
	Projects       projects.Store
	ProjectFiles   store.ProjectFileStore
	Rooms          rooms.Store
	Users          store.UserStore
	Visualizations visualizations.Store
	Workspaces     store.WorkspaceStore
	WorkspaceUsers store.WorkspaceUserStore
}

var (
	ErrForbidden     = errors.New("forbidden")
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

const nanoidAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewNanoID() (string, error) {
	return nanoid.Generate(nanoidAlphabet, 21)
}
