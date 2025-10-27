package crm

import (
	"errors"
	"net/url"

	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/apartomat/apartomat/internal/crm/image"
	"github.com/apartomat/apartomat/internal/crm/mail"
	"github.com/apartomat/apartomat/internal/store/albumfiles"
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projectpages"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/apartomat/apartomat/internal/store/users"
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	"github.com/apartomat/apartomat/internal/store/workspaceusers"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/uptrace/bun"
)

type Params struct {
	SendPinByEmail     bool
	ProjectPageBaseURL url.URL
}

type CRM struct {
	DB     *bun.DB
	Params Params

	AuthTokenIssuer   auth.AuthTokenIssuer
	AuthTokenVerifier auth.AuthTokenVerifier

	ConfirmTokenByEmailIssuer   auth.EmailConfirmTokenIssuer
	ConfirmTokenByEmailVerifier auth.EmailConfirmTokenVerifier

	ConfirmEmailPINTokenIssuer   auth.ConfirmEmailPINTokenIssuer
	ConfirmEmailPINTokenVerifier auth.ConfirmEmailPINTokenVerifier

	InviteTokenIssuer   auth.InviteTokenIssuer
	InviteTokenVerifier auth.InviteTokenVerifier

	MailSender  mail.Sender
	MailFactory *mail.Factory

	Uploader image.Uploader

	Acl *Acl

	Albums         albums.Store
	AlbumFiles     albumfiles.Store
	Contacts       contacts.Store
	Houses         houses.Store
	Projects       projects.Store
	ProjectPages   projectpages.Store
	Files          files.Store
	Rooms          rooms.Store
	Users          users.Store
	Visualizations visualizations.Store
	Workspaces     workspaces.Store
	WorkspaceUsers workspaceusers.Store
}

func NewCRM(
	db *bun.DB,
	params Params,
	authTokenIssuer auth.AuthTokenIssuer,
	authTokenVerifier auth.AuthTokenVerifier,
	confirmTokenByEmailIssuer auth.EmailConfirmTokenIssuer,
	confirmTokenByEmailVerifier auth.EmailConfirmTokenVerifier,
	confirmEmailPinIssuer auth.ConfirmEmailPINTokenIssuer,
	confirmEmailPinVerifier auth.ConfirmEmailPINTokenVerifier,
	inviteTokenIssuer auth.InviteTokenIssuer,
	inviteTokenVerifier auth.InviteTokenVerifier,
	mailSender mail.Sender,
	mailFactory *mail.Factory,
	uploader image.Uploader,
	acl *Acl,
	albumsStore albums.Store,
	albumFilesStore albumfiles.Store,
	contactsStore contacts.Store,
	housesStore houses.Store,
	projectsStore projects.Store,
	projectPagesStore projectpages.Store,
	filesStore files.Store,
	roomsStore rooms.Store,
	usersStore users.Store,
	visualizationsStore visualizations.Store,
	workspacesStore workspaces.Store,
	workspaceUsersStore workspaceusers.Store,
) *CRM {
	return &CRM{
		DB:                           db,
		Params:                       params,
		AuthTokenIssuer:              authTokenIssuer,
		AuthTokenVerifier:            authTokenVerifier,
		ConfirmTokenByEmailIssuer:    confirmTokenByEmailIssuer,
		ConfirmTokenByEmailVerifier:  confirmTokenByEmailVerifier,
		ConfirmEmailPINTokenIssuer:   confirmEmailPinIssuer,
		ConfirmEmailPINTokenVerifier: confirmEmailPinVerifier,
		InviteTokenIssuer:            inviteTokenIssuer,
		InviteTokenVerifier:          inviteTokenVerifier,
		MailSender:                   mailSender,
		MailFactory:                  mailFactory,
		Uploader:                     uploader,
		Acl:                          acl,
		Albums:                       albumsStore,
		AlbumFiles:                   albumFilesStore,
		Contacts:                     contactsStore,
		Houses:                       housesStore,
		Projects:                     projectsStore,
		ProjectPages:                 projectPagesStore,
		Files:                        filesStore,
		Rooms:                        roomsStore,
		Users:                        usersStore,
		Visualizations:               visualizationsStore,
		Workspaces:                   workspacesStore,
		WorkspaceUsers:               workspaceUsersStore,
	}
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
