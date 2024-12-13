package minio

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Uploader struct {
	bucketName string
}

func NewUploader(bucketName string) *Uploader {
	return &Uploader{bucketName: bucketName}
}

func (m *Uploader) Upload(ctx context.Context, r io.Reader, size int64, path, contentType string) (string, error) {
	cl, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("nsLMWG53xOq2ekLz", "h5OiDt9sDAGIB1YQPhg65xSHRQHrV5r4", ""),
		Secure: false,
	})
	if err != nil {
		return "", fmt.Errorf("can't connect to minio: %w", err)
	}

	inf, err := cl.PutObject(ctx, m.bucketName, path, r, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("can't upload to minio: %w", err)
	}

	return fmt.Sprintf("http://localhost:9000/apartomat/%s", inf.Key), nil
}
