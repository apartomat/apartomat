package auth

import "time"

type InviteToken interface {
	Email() string
	WorkspaceID() string
	Role() string
}

type InviteTokenIssuer interface {
	Issue(email, workspaceID, role string, tokenExpiration time.Duration) (string, error)
}

type InviteTokenVerifier interface {
	Verify(str string) (InviteToken, error)
}
