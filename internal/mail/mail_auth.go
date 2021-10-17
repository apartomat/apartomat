package mail

import (
	"fmt"
)

func NewMailAuth(from, to, token string) *Mail {
	return &Mail{
		From:    from,
		To:      to,
		Subject: "Login to Apartomat...",
		Body: fmt.Sprintf(`
Hello!

Please, open http://localhost:3000/confirm?token=%s
`, token),
	}
}
