package paseto

import (
	"crypto/ed25519"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/o1egl/paseto"
	"time"
)

const (
	authPurpose = "auth"
)

type AuthToken struct {
	paseto.JSONToken
}

func (token AuthToken) UserID() string {
	return token.Subject
}

func NewAuthToken(id string) AuthToken {
	token := AuthToken{
		JSONToken: paseto.JSONToken{
			Subject:    id,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(1 * 365 * 24 * time.Hour),
		},
	}

	token.Set(claimKeyPurpose, authPurpose)

	return token
}

func (token AuthToken) Validate(validators ...paseto.Validator) error {
	if len(validators) == 0 {
		validators = append(validators, paseto.ValidAt(time.Now()), hasPurpose(authPurpose))
	}

	return token.JSONToken.Validate(validators...)
}

type AuthTokenIssuerVerifier struct {
	privateKey ed25519.PrivateKey
}

func NewAuthTokenIssuerVerifier(key ed25519.PrivateKey) *AuthTokenIssuerVerifier {
	return &AuthTokenIssuerVerifier{key}
}

func (p *AuthTokenIssuerVerifier) Issue(id string) (string, error) {
	token := NewAuthToken(id)
	str, err := paseto.NewV2().Sign(p.privateKey, token, "")

	if err != nil {
		return "", fmt.Errorf("can't sign: %s", err)
	}

	return str, nil
}

func (p *AuthTokenIssuerVerifier) Verify(str string) (auth.AuthToken, error) {
	var (
		token  AuthToken
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
