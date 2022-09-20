package image

import (
	"context"
	"io"
)

type Uploader interface {
	Upload(ctx context.Context, reader io.Reader, size int64, path, contentType string) (string, error)
}
