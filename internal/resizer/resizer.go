package resizer

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
	"strings"
)

type Params struct {
	Width  uint
	Height uint
}

type S3Config struct {
	AccessKeyID, SecretAccessKey, Region, BucketName string
}

func Resize(ctx context.Context, path string, params Params, s3cfg S3Config) (io.Reader, error) {
	var (
		cred aws.CredentialsProviderFunc = func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     s3cfg.AccessKeyID,
				SecretAccessKey: s3cfg.SecretAccessKey,
			}, nil
		}
	)

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(s3cfg.Region),
		config.WithCredentialsProvider(cred),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	outp, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s3cfg.BucketName),
		Key:    aws.String(strings.TrimLeft(path, "/")),
	})
	if err != nil {
		return nil, err
	}

	img, f, err := image.Decode(outp.Body)
	if err != nil {
		return nil, err
	}

	res := resize.Resize(params.Width, params.Height, img, resize.Lanczos3)

	var (
		buf bytes.Buffer
	)

	switch f {
	case "jpeg":
		err = jpeg.Encode(&buf, res, nil)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown format %s", f)
	}

	return &buf, nil

}
