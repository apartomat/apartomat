package smtp

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	transport "net/smtp"
	"net/textproto"
	"strings"
	"time"

	"github.com/apartomat/apartomat/internal/crm/mail"
)

type Config struct {
	Addr     string
	User     string
	Password string
}

type mailSender struct {
	config Config
}

const smtpSendTimeout = 60 * time.Second

func NewMailSender(config Config) mail.Sender {
	return &mailSender{config: config}
}

func (ms *mailSender) Send(m *mail.Mail) error {
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

	host, port, err := net.SplitHostPort(ms.config.Addr)
	if err != nil {
		return fmt.Errorf("invalid smtp addr %q: %w", ms.config.Addr, err)
	}
	auth := transport.PlainAuth("", ms.config.User, ms.config.Password, host)

	if port == "465" {
		dialer := &net.Dialer{Timeout: smtpSendTimeout}
		conn, err := tls.DialWithDialer(
			dialer,
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
		_ = conn.SetDeadline(time.Now().Add(smtpSendTimeout))

		client, err := transport.NewClient(conn, host)
		if err != nil {
			_ = conn.Close()
			return err
		}

		defer client.Close()

		return sendWithClient(client, auth, m.From, m.To, body)
	}

	conn, err := (&net.Dialer{Timeout: smtpSendTimeout}).Dial("tcp", ms.config.Addr)
	if err != nil {
		return err
	}
	_ = conn.SetDeadline(time.Now().Add(smtpSendTimeout))

	client, err := transport.NewClient(conn, host)
	if err != nil {
		_ = conn.Close()
		return err
	}

	defer client.Close()

	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(&tls.Config{InsecureSkipVerify: true, ServerName: host}); err != nil {
			return err
		}
	} else if strings.HasSuffix(port, "587") {
		return fmt.Errorf("smtp server does not support STARTTLS on %s", ms.config.Addr)
	}

	return sendWithClient(client, auth, m.From, m.To, body)
}

func sendWithClient(client *transport.Client, auth transport.Auth, from, to string, body []byte) error {
	if err := client.Auth(auth); err != nil {
		return err
	}

	if err := client.Mail(from); err != nil {
		return err
	}

	if err := client.Rcpt(to); err != nil {
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
