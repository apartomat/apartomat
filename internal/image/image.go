package image

import (
	"context"
	"io"
)

type ImageUploader interface {
	Upload(ctx context.Context, reader io.Reader, path, contentType string) (string, error)
}
