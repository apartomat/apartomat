package apartomat

import (
	"crypto/ed25519"
	"fmt"
	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type AuthTokenIssuer interface {
	Issue(id int, email string) (string, error)
}

type AuthTokenVerifier interface {
	Verify(str string) (*AuthToken, string, error)
}

type AuthToken struct {
	UserID int
	paseto.JSONToken
}

const (
	authPurpose    = "auth"
	userIdClaimKey = "userId"
)

func NewAuthToken(id int, email string) AuthToken {
	token := AuthToken{
		JSONToken: paseto.JSONToken{
			Subject:    email,
			IssuedAt:   time.Now(),
			Expiration: time.Now().Add(1 * 365 * 24 * time.Hour),
		},
	}

	token.Set("purpose", authPurpose)
	token.Set(userIdClaimKey, strconv.Itoa(id))

	return token
}

func (token AuthToken) Validate(validators ...paseto.Validator) error {
	if len(validators) == 0 {
		validators = append(validators, paseto.ValidAt(time.Now()), hasPurpose(authPurpose), hasID)
	}

	return token.JSONToken.Validate(validators...)
}

type pasetoAuthTokenIssuerVerifier struct {
	privateKey ed25519.PrivateKey
}

func NewPasetoAuthTokenIssuerVerifier(key ed25519.PrivateKey) *pasetoAuthTokenIssuerVerifier {
	return &pasetoAuthTokenIssuerVerifier{key}
}

func (p *pasetoAuthTokenIssuerVerifier) Issue(id int, email string) (string, error) {
	token := NewAuthToken(id, email)
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

	token.UserID, _ = strconv.Atoi(token.Get(userIdClaimKey))

	return &token, footer, nil
}

func hasID(token *paseto.JSONToken) error {
	_, err := strconv.Atoi(token.Get(userIdClaimKey))
	if err != nil {
		return errors.Wrapf(paseto.ErrTokenValidationError, "token has no %s", userIdClaimKey)
	}

	return nil
}
