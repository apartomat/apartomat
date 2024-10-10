package auth

type ConfirmEmailPINToken interface {
	Email() string
	PIN() string
}

type ConfirmEmailPINTokenIssuer interface {
	Issue(email, pin string) (string, error)
}

type ConfirmEmailPINTokenVerifier interface {
	Verify(str, pin string) (ConfirmEmailPINToken, error)
}
