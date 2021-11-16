package apartomat

import "github.com/apartomat/apartomat/internal/token"

type CheckAuthToken struct {
	verifier token.AuthTokenVerifier
}

func (u *Apartomat) CheckAuthToken(str string) (*token.AuthToken, error) {
	token, _, err := u.AuthTokenVerifier.Verify(str)

	return token, err
}
