package apartomat

type CheckAuthToken struct {
	verifier AuthTokenVerifier
}

func NewCheckAuthToken(verifier AuthTokenVerifier) *CheckAuthToken {
	return &CheckAuthToken{verifier}
}

func (u *CheckAuthToken) Do(str string) (*AuthToken, error) {
	token, _, err := u.verifier.Verify(str)

	return token, err
}
