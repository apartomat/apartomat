package minio

import (
	"bytes"
	"context"
	"fmt"
	parent "github.com/apartomat/apartomat/internal/image"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

func (m *Uploader) Resize(ctx context.Context, path string, opts parent.ResizeOptions) (io.Reader, error) {
	cl, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("nsLMWG53xOq2ekLz", "h5OiDt9sDAGIB1YQPhg65xSHRQHrV5r4", ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("can't connect to minio: %w", err)
	}

	obj, err := cl.GetObject(ctx, m.bucketName, strings.TrimLeft(path, "/"), minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	img, f, err := image.Decode(obj)
	if err != nil {
		return nil, err
	}

	res := resize.Resize(opts.Width, opts.Height, img, resize.Lanczos3)

	var (
		buf bytes.Buffer
	)

	switch f {
	case "jpeg":
		err = jpeg.Encode(&buf, res, nil)
		if err != nil {
			return nil, err
		}
	case "png":
		err = png.Encode(&buf, res)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown format %s", f)
	}

	return &buf, nil
}
