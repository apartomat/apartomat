package token

import (
	"crypto/ed25519"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/o1egl/paseto"
	"time"
)

const (
	emailConfirmPurpose = "email-confirm"
)

type EmailConfirmToken struct {
	paseto.JSONToken
}

func (token EmailConfirmToken) Email() string {
	return token.Subject
}

func NewEmailConfirmToken(email string) EmailConfirmToken {
	token := EmailConfirmToken{
		JSONToken: paseto.JSONToken{
			Subject:    email,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(10 * time.Minute),
		},
	}

	token.Set(ClaimKeyPurpose, emailConfirmPurpose)

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
		return "", fmt.Errorf("can't sign: %w", err)
	}

	return str, nil
}

func (p *pasetoMailConfirmTokenIssuerVerifier) Verify(str string) (apartomat.EmailConfirmToken, error) {
	var (
		token  EmailConfirmToken
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
