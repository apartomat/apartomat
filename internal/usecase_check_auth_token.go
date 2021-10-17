package apartomat

import "github.com/apartomat/apartomat/internal/token"

type CheckAuthToken struct {
	verifier token.AuthTokenVerifier
}

func NewCheckAuthToken(verifier token.AuthTokenVerifier) *CheckAuthToken {
	return &CheckAuthToken{verifier}
}

func (u *CheckAuthToken) Do(str string) (*token.AuthToken, error) {
	token, _, err := u.verifier.Verify(str)

	return token, err
}
