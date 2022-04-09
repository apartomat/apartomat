package mail

import (
	"fmt"
)

type Factory struct {
	Hostname string
	From     string
}

func NewFactory(hostname, from string) *Factory {
	return &Factory{Hostname: hostname, From: from}
}

func (f *Factory) MailAuth(to, token string) *Mail {
	return &Mail{
		From:    f.From,
		To:      to,
		Subject: "Login to Apartomat...",
		Body: fmt.Sprintf(`
Hello!

Please, open %s/confirm?token=%s
`, f.Hostname, token),
	}
}

func (f *Factory) MailPIN(to, pin string) *Mail {
	return &Mail{
		From:    f.From,
		To:      to,
		Subject: "Login to Apartomat...",
		Body: fmt.Sprintf(`
Hello!

Your PIN is: %s
`, pin),
	}
}
