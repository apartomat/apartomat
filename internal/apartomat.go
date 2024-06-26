package apartomat

import (
	"errors"
	"go.uber.org/zap"

	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/image"
	"github.com/apartomat/apartomat/internal/mail"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projects"
	sites "github.com/apartomat/apartomat/internal/store/public_sites"
	"github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/apartomat/apartomat/internal/store/users"
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	nanoid "github.com/matoous/go-nanoid/v2"
)

type Params struct {
	SendPinByEmail bool
}

type Apartomat struct {
	Params Params

	AuthTokenIssuer   auth.AuthTokenIssuer
	AuthTokenVerifier auth.AuthTokenVerifier

	ConfirmTokenByEmailIssuer   auth.EmailConfirmTokenIssuer
	ConfirmTokenByEmailVerifier auth.EmailConfirmTokenVerifier

	ConfirmEmailPINTokenIssuer   auth.ConfirmEmailPINTokenIssuer
	ConfirmEmailPINTokenVerifier auth.ConfirmEmailPINTokenVerifier

	InviteTokenIssuer   auth.InviteTokenIssuer
	InviteTokenVerifier auth.InviteTokenVerifier

	Mailer      mail.Sender
	MailFactory *mail.Factory

	Uploader image.Uploader

	Acl *Acl

	Albums         albums.Store
	AlbumFiles     albumFiles.Store
	Contacts       contacts.Store
	Houses         houses.Store
	Projects       projects.Store
	PublicSites    sites.Store
	Files          files.Store
	Rooms          rooms.Store
	Users          users.Store
	Visualizations visualizations.Store
	Workspaces     workspaces.Store
	WorkspaceUsers workspace_users.Store

	Logger *zap.Logger
}

var (
	ErrForbidden     = errors.New("forbidden")
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

const nanoidAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateNanoID() (string, error) {
	return nanoid.Generate(nanoidAlphabet, 21)
}

func MustGenerateNanoID() string {
	return nanoid.MustGenerate(nanoidAlphabet, 21)
}
