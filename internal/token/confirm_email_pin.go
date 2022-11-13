package token

import (
	"crypto/ed25519"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/o1egl/paseto"
	"time"
)

const (
	confirmEmailPINPurpose = "confirm-email-pin"
)

type ConfirmEmailPINToken struct {
	paseto.JSONToken
}

func (token ConfirmEmailPINToken) Email() string {
	return token.Subject
}

func (token ConfirmEmailPINToken) PIN() string {
	return token.Get(ClaimKeyPIN)
}

func NewConfirmEmailPINToken(email, pin string) ConfirmEmailPINToken {
	token := ConfirmEmailPINToken{
		JSONToken: paseto.JSONToken{
			Subject:    email,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(5 * time.Minute),
		},
	}

	token.Set(ClaimKeyPIN, pin)

	token.Set(ClaimKeyPurpose, confirmEmailPINPurpose)

	return token
}

func (token ConfirmEmailPINToken) Validate(validators ...paseto.Validator) error {
	if len(validators) == 0 {
		validators = append(validators, paseto.ValidAt(time.Now()), hasPurpose(confirmEmailPINPurpose))
	}

	return token.JSONToken.Validate(validators...)
}

type pasetoConfirmEmailPINTokenIssuerVerifier struct {
	privateKey ed25519.PrivateKey
	secret     []byte
}

func NewPasetoConfirmEmailPINTokenIssuerVerifier(key ed25519.PrivateKey) *pasetoConfirmEmailPINTokenIssuerVerifier {
	return &pasetoConfirmEmailPINTokenIssuerVerifier{key, []byte("YELLOW SUBMARINE, BLACK WIZARDRY")}
}

func (p *pasetoConfirmEmailPINTokenIssuerVerifier) Issue(email, pin string) (string, error) {
	token := NewConfirmEmailPINToken(email, pin)
	str, err := paseto.NewV2().Encrypt(p.secret, token, "")

	if err != nil {
		return "", fmt.Errorf("can't encrypt: %s", err)
	}

	return str, nil
}

func (p *pasetoConfirmEmailPINTokenIssuerVerifier) Verify(str, pin string) (apartomat.ConfirmEmailPINToken, error) {
	var (
		token  ConfirmEmailPINToken
		footer string
	)

	err := paseto.NewV2().Decrypt(str, p.secret, &token, &footer)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, ErrTokenVerificationError)
	}

	err = token.Validate(hasPIN(pin))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err, ErrTokenValidationError)
	}

	return &token, nil
}
