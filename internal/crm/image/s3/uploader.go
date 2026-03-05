package s3

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Uploader struct {
	client     *s3.Client
	region     string
	bucketName string
	endpoint   string
}

func NewS3ImageUploader(client *s3.Client, region, bucketName, endpoint string) *Uploader {
	return &Uploader{client, region, bucketName, endpoint}
}

func NewS3ImageUploaderWithCred(
	ctx context.Context,
	accessKeyID, secretAccessKey, region, bucketName, endpoint string,
) (*Uploader, error) {
	var (
		cred aws.CredentialsProviderFunc = func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     accessKeyID,
				SecretAccessKey: secretAccessKey,
			}, nil
		}

		clientOpts []func(*s3.Options)
	)

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(cred),
	)
	if err != nil {
		return nil, err
	}

	if endpoint != "" {
		clientOpts = append(clientOpts, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(strings.TrimSuffix(endpoint, "/"))
			o.UsePathStyle = true
		})
	}

	client := s3.NewFromConfig(cfg, clientOpts...)

	return NewS3ImageUploader(client, region, bucketName, endpoint), nil
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

	var (
		baseURL = fmt.Sprintf("s3.%s.amazonaws.com", u.region)
	)

	if u.endpoint != "" {
		baseURL = u.endpoint
	}

	baseURL = strings.TrimSuffix(baseURL, "/")

	return fmt.Sprintf("%s/%s/%s", baseURL, u.bucketName, path), nil
}
