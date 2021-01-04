package apartomat

import (
	"crypto/ed25519"
	"fmt"
	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
	"time"
)

type AuthTokenIssuer interface {
	Issue(email string) (string, error)
}

type AuthTokenVerifier interface {
	Verify(str string) (*AuthToken, string, error)
}

type AuthToken struct {
	paseto.JSONToken
}

const (
	authPurpose = "auth"
)

func NewAuthToken(login string) AuthToken {
	token := AuthToken{
		JSONToken: paseto.JSONToken{
			Subject:    login,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(1 * 365 * 24 * time.Hour),
		},
	}

	token.Set("purpose", authPurpose)

	return token
}

func (token AuthToken) Validate(validators ...paseto.Validator) error {
	if len(validators) == 0 {
		validators = append(validators, paseto.ValidAt(time.Now()), hasPurpose(authPurpose))
	}

	return token.JSONToken.Validate(validators...)
}

type pasetoAuthTokenIssuerVerifier struct {
	privateKey ed25519.PrivateKey
}

func NewPasetoAuthTokenIssuerVerifier(key ed25519.PrivateKey) *pasetoAuthTokenIssuerVerifier {
	return &pasetoAuthTokenIssuerVerifier{key}
}

func (p *pasetoAuthTokenIssuerVerifier) Issue(email string) (string, error) {
	token := NewAuthToken(email)
	str, err := paseto.NewV2().Sign(p.privateKey, token, "")

	if err != nil {
		return "", fmt.Errorf("can't sign: %s", err)
	}

	return str, nil
}

func (p *pasetoAuthTokenIssuerVerifier) Verify(str string) (*AuthToken, string, error) {
	var (
		token  AuthToken
		footer string
	)

	err := paseto.NewV2().Verify(str, p.privateKey.Public(), &token, &footer)
	if err != nil {
		return nil, "", errors.Wrapf(ErrTokenVerificationError, "%s", err)
	}

	err = token.Validate()
	if err != nil {
		return nil, "", errors.Wrapf(ErrTokenValidationError, "%s", err)
	}

	return &token, footer, nil
}
