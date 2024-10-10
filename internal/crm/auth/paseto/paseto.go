package paseto

import (
	"errors"
	"fmt"
	"github.com/o1egl/paseto"
)

var (
	ErrTokenVerificationError = errors.New("incorrect auth format")
	ErrTokenValidationError   = errors.New("auth expired or not valid")
)

const (
	claimKeyPurpose = "pur"
)

func hasPurpose(purpose string) paseto.Validator {
	return hasClaim(claimKeyPurpose, purpose)
}

func hasClaim(key, value string) paseto.Validator {
	return func(token *paseto.JSONToken) error {
		if token.Get(key) != value {
			return fmt.Errorf("incorrect auth claim %s: %w", value, paseto.ErrTokenValidationError)
		}

		return nil
	}
}
