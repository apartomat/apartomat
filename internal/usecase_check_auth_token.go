package apartomat

type CheckAuthToken struct {
	verifier AuthTokenVerifier
}

func NewCheckAuthToken(verifier AuthTokenVerifier) *CheckAuthToken {
	return &CheckAuthToken{verifier}
}

func (uc *CheckAuthToken) Do(str string) (*AuthToken, error) {
	token, _, err := uc.verifier.Verify(str)

	return token, err
}
