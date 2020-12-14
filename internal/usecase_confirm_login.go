package apartomat

type ConfirmLogin struct {
	verifier EmailConfirmTokenVerifier
	issuer   AuthTokenIssuer
}

func NewConfirmLogin(verifier EmailConfirmTokenVerifier, issuer AuthTokenIssuer) *ConfirmLogin {
	return &ConfirmLogin{verifier: verifier, issuer: issuer}
}

func (uc *ConfirmLogin) Do(str string) (string, error) {
	confirmToken, _, err := uc.verifier.Verify(str)
	if err != nil {
		return "", err
	}

	return uc.issuer.Issue(confirmToken.Subject)
}
