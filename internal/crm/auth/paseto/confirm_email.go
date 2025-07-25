package paseto

import (
	"crypto/ed25519"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/o1egl/paseto"
	"time"
)

const (
	confirmEmailPurpose = "email-confirm"
)

type ConfirmEmailToken struct {
	paseto.JSONToken
}

func (token ConfirmEmailToken) Email() string {
	return token.Subject
}

func NewConfirmEmailToken(email string) ConfirmEmailToken {
	token := ConfirmEmailToken{
		JSONToken: paseto.JSONToken{
			Subject:    email,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(10 * time.Minute),
		},
	}

	token.Set(claimKeyPurpose, confirmEmailPurpose)

	return token
}

func (token ConfirmEmailToken) Validate(validators ...paseto.Validator) error {
	if len(validators) == 0 {
		validators = append(validators, paseto.ValidAt(time.Now()), hasPurpose(confirmEmailPurpose))
	}

	return token.JSONToken.Validate(validators...)
}

type ConfirmEmailTokenIssuerVerifier struct {
	privateKey ed25519.PrivateKey
}

func NewConfirmEmailTokenIssuerVerifier(key ed25519.PrivateKey) *ConfirmEmailTokenIssuerVerifier {
	return &ConfirmEmailTokenIssuerVerifier{key}
}

func (p *ConfirmEmailTokenIssuerVerifier) Issue(email string) (string, error) {
	token := NewConfirmEmailToken(email)
	str, err := paseto.NewV2().Sign(p.privateKey, token, "")

	if err != nil {
		return "", fmt.Errorf("can't sign: %w", err)
	}

	return str, nil
}

func (p *ConfirmEmailTokenIssuerVerifier) Verify(str string) (auth.EmailConfirmToken, error) {
	var (
		token  ConfirmEmailToken
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
