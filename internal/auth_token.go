package apartomat

type AuthToken interface {
	UserID() string
}

type AuthTokenIssuer interface {
	Issue(id string) (string, error)
}

type AuthTokenVerifier interface {
	Verify(str string) (AuthToken, error)
}

func (u *Apartomat) CheckAuthToken(str string) (AuthToken, error) {
	return u.AuthTokenVerifier.Verify(str)
}
