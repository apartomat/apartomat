package token

import (
	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
)

var (
	ErrTokenVerificationError = errors.New("incorrect token format")
	ErrTokenValidationError   = errors.New("token expired or not valid")
)

const (
	ClaimKeyPurpose = "pur"
	ClaimKeyPIN     = "pin"
)

var (
	hasPurpose = func(purpose string) paseto.Validator {
		return hasClaim(ClaimKeyPurpose, purpose)
	}
)

func hasClaim(key, value string) paseto.Validator {
	return func(token *paseto.JSONToken) error {
		if token.Get(key) != value {
			return errors.Wrapf(paseto.ErrTokenValidationError, "incorrect token claim %s", value)
		}

		return nil
	}
}

func hasPIN(pin string) paseto.Validator {
	return hasClaim(ClaimKeyPIN, pin)
}
