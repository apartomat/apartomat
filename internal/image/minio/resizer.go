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
	"net/url"
	"strconv"
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

	var (
		origPath = strings.TrimLeft(path, "/")
		resPath  = origPath
	)

	q := &url.Values{}

	if opts.Width != 0 {
		q.Add("w", strconv.Itoa(int(opts.Width)))
	}

	if opts.Height != 0 {
		q.Add("h", strconv.Itoa(int(opts.Height)))
	}

	if sq := q.Encode(); sq != "" {
		resPath += "?" + sq
	}

	robj, err := cl.GetObject(
		ctx,
		m.bucketName,
		resPath,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}

	stat, err := robj.Stat()
	if err != nil {
		if er := minio.ToErrorResponse(err); er.Code != "NoSuchKey" {
			return nil, err
		}

		if sq := q.Encode(); sq != "" {
			oobj, err := cl.GetObject(
				ctx,
				m.bucketName,
				origPath,
				minio.GetObjectOptions{},
			)
			if err != nil {
				return nil, err
			}

			img, f, err := image.Decode(oobj)
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

			var (
				resbuf bytes.Buffer
			)

			resbuf.Write(buf.Bytes())

			if _, err = m.Upload(ctx, &buf, int64(buf.Len()), resPath, stat.ContentType); err != nil {
				return nil, err
			}

			return &resbuf, nil
		}
	}

	return robj, nil
}
