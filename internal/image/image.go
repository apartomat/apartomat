package image

import (
	"context"
	"io"
)

type Uploader interface {
	Upload(ctx context.Context, reader io.Reader, size int64, path, contentType string) (string, error)
}

type Resizer interface {
	Resize(ctx context.Context, path string, options ResizeOptions) (io.Reader, error)
}

type ResizeOptions struct {
	Width  uint
	Height uint
	Fit    uint
}
