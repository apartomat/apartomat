package image

import (
	"context"
	"io"
)

type Uploader interface {
	Upload(ctx context.Context, reader io.Reader, path, contentType string) (string, error)
}
