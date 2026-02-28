package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Uploader struct {
	client     *s3.Client
	region     string
	bucketName string
}

func NewS3ImageUploader(client *s3.Client, region, bucketName string) *Uploader {
	return &Uploader{client, region, bucketName}
}

func NewS3ImageUploaderWithCred(
	ctx context.Context,
	accessKeyID, secretAccessKey, region, bucketName string,
) (*Uploader, error) {
	var (
		cred aws.CredentialsProviderFunc = func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     accessKeyID,
				SecretAccessKey: secretAccessKey,
			}, nil
		}
	)

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(cred),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return NewS3ImageUploader(client, region, bucketName), nil
}

func (u *Uploader) Upload(
	ctx context.Context,
	reader io.Reader,
	_ int64,
	path,
	contentType string,
) (string, error) {
	inp := &s3.PutObjectInput{
		Bucket:      aws.String(u.bucketName),
		Body:        reader,
		Key:         aws.String(path),
		ContentType: aws.String(contentType),
	}

	_, err := u.client.PutObject(ctx, inp)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", u.bucketName, path)

	return url, nil
}
