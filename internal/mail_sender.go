package apartomat

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"net/textproto"
	"time"
)

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

type SmtpServerConfig struct {
	Addr     string
	User     string
	Password string
}

type mailSender struct {
	config SmtpServerConfig
}

type MailSender interface {
	Send(m *Mail) error
}

func NewMailSender(config SmtpServerConfig) MailSender {
	return &mailSender{config: config}
}

func (ms *mailSender) Send(m *Mail) error {
	header := textproto.MIMEHeader{}
	header.Set(textproto.CanonicalMIMEHeaderKey("from"), m.From)
	header.Set(textproto.CanonicalMIMEHeaderKey("to"), m.To)
	header.Set(textproto.CanonicalMIMEHeaderKey("mime-version"), "1.0")
	header.Set(textproto.CanonicalMIMEHeaderKey("content-type"), "text/html; charset=UTF-8")
	header.Set(textproto.CanonicalMIMEHeaderKey("date"), time.Now().Format(time.RFC1123Z))
	header.Set(textproto.CanonicalMIMEHeaderKey("subject"), m.Subject)

	var (
		buffer bytes.Buffer
	)

	for key, value := range header {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", key, value[0]))
	}

	buffer.WriteString(fmt.Sprintf("\r\n%s", m.Body))

	body := buffer.Bytes()

	host, _, _ := net.SplitHostPort(ms.config.Addr)
	auth := smtp.PlainAuth("", ms.config.User, ms.config.Password, host)

	conn, err := tls.Dial(
		"tcp",
		ms.config.Addr,
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		},
	)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	if err = client.Auth(auth); err != nil {
		return err
	}

	//

	defer client.Close()

	if err := client.Mail(m.From); err != nil {
		return err
	}

	if err := client.Rcpt(m.To); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	if _, err = w.Write(body); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return client.Quit()
}
