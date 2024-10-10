package paseto

import (
	"crypto/ed25519"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/o1egl/paseto"
	"time"
)

const (
	invitePurpose = "invite"

	claimKeyWorkspaceID   = "wid"
	claimKeyWorkspaceRole = "role"
)

type InviteToken struct {
	paseto.JSONToken
}

func (token InviteToken) Email() string {
	return token.Subject
}

func (token InviteToken) WorkspaceID() string {
	return token.Get(claimKeyWorkspaceID)
}

func (token InviteToken) Role() string {
	return token.Get(claimKeyWorkspaceRole)
}

func NewInviteToken(email, workspaceID, role string, tokenExpiration time.Duration) InviteToken {
	token := InviteToken{
		JSONToken: paseto.JSONToken{
			Subject:    email,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(tokenExpiration),
		},
	}

	token.Set(claimKeyPurpose, invitePurpose)

	token.Set(claimKeyWorkspaceID, workspaceID)
	token.Set(claimKeyWorkspaceRole, role)

	return token
}

func (token InviteToken) Validate(validators ...paseto.Validator) error {
	if len(validators) == 0 {
		validators = append(validators, paseto.ValidAt(time.Now()), hasPurpose(invitePurpose))
	}

	return token.JSONToken.Validate(validators...)
}

type inviteTokenIssuerVerifier struct {
	privateKey ed25519.PrivateKey
}

func NewInviteTokenIssuerVerifier(key ed25519.PrivateKey) *inviteTokenIssuerVerifier {
	return &inviteTokenIssuerVerifier{key}
}

func (p *inviteTokenIssuerVerifier) Issue(email, workspaceID, role string, tokenExpiration time.Duration) (string, error) {
	token := NewInviteToken(email, workspaceID, role, tokenExpiration)

	str, err := paseto.NewV2().Sign(p.privateKey, token, "")

	if err != nil {
		return "", fmt.Errorf("can't sign: %w", err)
	}

	return str, nil
}

func (p *inviteTokenIssuerVerifier) Verify(str string) (auth.InviteToken, error) {
	var (
		token  InviteToken
		footer string
	)

	err := paseto.NewV2().Verify(str, p.privateKey.Public(), &token, &footer)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, ErrTokenVerificationError)
	}

	err = token.Validate()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, ErrTokenValidationError)
	}

	return &token, nil
}
