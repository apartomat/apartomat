package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/apartomat/apartomat/internal/resizer"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"os"
	"strconv"
)

type Request events.APIGatewayProxyRequest

type Response events.APIGatewayProxyResponse

func handler(ctx context.Context, req Request) (Response, error) {
	cfg := resizer.S3Config{
		AccessKeyID:     os.Getenv("S3_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("S3_SECRET_ACCESS_KEY"),
		Region:          os.Getenv("S3_REGION"),
		BucketName:      os.Getenv("S3_BUCKET_NAME"),
	}

	r, err := resizer.Resize(ctx, path(req.PathParameters), params(req.QueryStringParameters), cfg)
	if err != nil {
		return Response{StatusCode: 500, Body: fmt.Sprintf("can't resize image: %s", err)}, nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return Response{StatusCode: 500, Body: fmt.Sprintf("can't resize image: %s", err)}, nil
	}

	body := base64.StdEncoding.EncodeToString(b)

	return Response{
		StatusCode: 200,
		Body:       body,
		Headers: map[string]string{
			"Content-Type":   "image/jpeg",
			"Content-Length": strconv.Itoa(len(body)),
		},
		IsBase64Encoded: true,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func params(vals map[string]string) resizer.Params {
	var (
		p resizer.Params
	)

	for k, v := range vals {
		switch k {
		case "w", "width":
			if w, err := strconv.Atoi(v); err == nil {
				p.Width = uint(w)
			}
		case "h", "height":
			if h, err := strconv.Atoi(v); err == nil {
				p.Height = uint(h)
			}
		}
	}

	return p
}

func path(vals map[string]string) string {
	for k, v := range vals {
		if k == "path" {
			return v
		}
	}

	return ""
}
