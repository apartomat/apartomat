package auth

type AuthToken interface {
	UserID() string
}

type AuthTokenIssuer interface {
	Issue(id string) (string, error)
}

type AuthTokenVerifier interface {
	Verify(str string) (AuthToken, error)
}
