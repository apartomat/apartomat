package paseto

import (
	"crypto/ed25519"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/o1egl/paseto"
	"time"
)

const (
	confirmEmailPINPurpose = "confirm-email-pin"

	claimKeyPIN = "pin"
)

type ConfirmEmailPINToken struct {
	paseto.JSONToken
}

func (token ConfirmEmailPINToken) Email() string {
	return token.Subject
}

func (token ConfirmEmailPINToken) PIN() string {
	return token.Get(claimKeyPIN)
}

func NewConfirmEmailPINToken(email, pin string) ConfirmEmailPINToken {
	token := ConfirmEmailPINToken{
		JSONToken: paseto.JSONToken{
			Subject:    email,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(5 * time.Minute),
		},
	}

	token.Set(claimKeyPIN, pin)

	token.Set(claimKeyPurpose, confirmEmailPINPurpose)

	return token
}

func (token ConfirmEmailPINToken) Validate(validators ...paseto.Validator) error {
	if len(validators) == 0 {
		validators = append(validators, paseto.ValidAt(time.Now()), hasPurpose(confirmEmailPINPurpose))
	}

	return token.JSONToken.Validate(validators...)
}

type confirmEmailPINTokenIssuerVerifier struct {
	privateKey ed25519.PrivateKey
	secret     []byte
}

func NewConfirmEmailPINTokenIssuerVerifier(key ed25519.PrivateKey) *confirmEmailPINTokenIssuerVerifier {
	return &confirmEmailPINTokenIssuerVerifier{key, []byte("YELLOW SUBMARINE, BLACK WIZARDRY")}
}

func (p *confirmEmailPINTokenIssuerVerifier) Issue(email, pin string) (string, error) {
	token := NewConfirmEmailPINToken(email, pin)
	str, err := paseto.NewV2().Encrypt(p.secret, token, "")

	if err != nil {
		return "", fmt.Errorf("can't encrypt: %s", err)
	}

	return str, nil
}

func (p *confirmEmailPINTokenIssuerVerifier) Verify(str, pin string) (auth.ConfirmEmailPINToken, error) {
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

func hasPIN(pin string) paseto.Validator {
	return hasClaim(claimKeyPIN, pin)
}
