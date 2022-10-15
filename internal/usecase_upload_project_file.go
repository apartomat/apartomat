package apartomat

import (
	"io"
)

type Upload struct {
	Name     string
	MimeType string
	Data     io.Reader
	Size     int64
}
