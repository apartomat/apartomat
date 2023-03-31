package auth

type EmailConfirmToken interface {
	Email() string
}

type EmailConfirmTokenIssuer interface {
	Issue(email string) (string, error)
}

type EmailConfirmTokenVerifier interface {
	Verify(str string) (EmailConfirmToken, error)
}
