package token

import (
	"crypto/ed25519"
	"fmt"
	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
	"time"
)

type EmailConfirmTokenIssuer interface {
	Issue(email string) (string, error)
}

type EmailConfirmTokenVerifier interface {
	Verify(str string) (*EmailConfirmToken, string, error)
}

const (
	emailConfirmPurpose = "email-confirm"
)

type EmailConfirmToken struct {
	paseto.JSONToken
}

func NewEmailConfirmToken(login string) EmailConfirmToken {
	token := EmailConfirmToken{
		JSONToken: paseto.JSONToken{
			Subject:    login,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(10 * time.Minute),
		},
	}

	token.Set("purpose", emailConfirmPurpose)

	return token
}

func (token EmailConfirmToken) Validate(validators ...paseto.Validator) error {
	if len(validators) == 0 {
		validators = append(validators, paseto.ValidAt(time.Now()), hasPurpose(emailConfirmPurpose))
	}

	return token.JSONToken.Validate(validators...)
}

type pasetoMailConfirmTokenIssuerVerifier struct {
	privateKey ed25519.PrivateKey
}

func NewPasetoMailConfirmTokenIssuerVerifier(key ed25519.PrivateKey) *pasetoMailConfirmTokenIssuerVerifier {
	return &pasetoMailConfirmTokenIssuerVerifier{key}
}

func (p *pasetoMailConfirmTokenIssuerVerifier) Issue(email string) (string, error) {
	token := NewEmailConfirmToken(email)
	str, err := paseto.NewV2().Sign(p.privateKey, token, "")

	if err != nil {
		return "", fmt.Errorf("can't sign: %s", err)
	}

	return str, nil
}

func (p *pasetoMailConfirmTokenIssuerVerifier) Verify(str string) (*EmailConfirmToken, string, error) {
	var (
		token  EmailConfirmToken
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
