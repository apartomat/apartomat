package main

import (
	"context"
	"github.com/apartomat/apartomat/internal/crm/image"
	"github.com/apartomat/apartomat/internal/crm/image/minio"
	"github.com/apartomat/apartomat/internal/resizer"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func params(values url.Values) resizer.Params {
	var (
		p resizer.Params
	)

	for k, v := range values {
		if len(v) == 0 {
			continue
		}

		switch k {
		case "w", "width":
			if w, err := strconv.Atoi(v[0]); err == nil {
				p.Width = uint(w)
			}
		case "h", "height":
			if h, err := strconv.Atoi(v[0]); err == nil {
				p.Height = uint(h)
			}
		}
	}

	return p
}

func opts(values url.Values) image.ResizeOptions {
	var (
		p image.ResizeOptions
	)

	for k, v := range values {
		if len(v) == 0 {
			continue
		}

		switch k {
		case "w", "width":
			if w, err := strconv.Atoi(v[0]); err == nil {
				p.Width = uint(w)
			}
		case "h", "height":
			if h, err := strconv.Atoi(v[0]); err == nil {
				p.Height = uint(h)
			}
		case "f", "fit":
			if f, err := strconv.Atoi(v[0]); err == nil {
				p.Fit = uint(f)
			}
		}
	}

	return p
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cfg := resizer.S3Config{
		AccessKeyID:     os.Getenv("S3_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("S3_SECRET_ACCESS_KEY"),
		Region:          os.Getenv("S3_REGION"),
		BucketName:      os.Getenv("S3_BUCKET_NAME"),
	}

	img, err := resizer.Resize(context.TODO(), r.URL.Path, params(r.URL.Query()), cfg)
	if err != nil {
		log.Printf("can't resize: %s", err)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	_, err = io.Copy(w, img)
	if err != nil {
		log.Printf("can't write: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func mhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := minio.NewUploader("apartomat")

	img, err := res.Resize(context.TODO(), r.URL.Path, opts(r.URL.Query()))
	if err != nil {
		log.Printf("can't resize: %s", err)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	_, err = io.Copy(w, img)
	if err != nil {
		log.Printf("can't write: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	addr := "localhost:8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/", mhandler)

	s := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit

		log.Print("stopping server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Fatalf("can't stop server: %s", err)
		}

		close(done)
	}()

	log.Printf("starting server at %s...", s.Addr)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("can't start server: %s", err)
		}
	}()

	<-done

	log.Print("buy")
}
