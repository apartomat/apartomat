package mail

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

type Sender interface {
	Send(m *Mail) error
}
