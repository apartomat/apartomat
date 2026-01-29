package bookbinder

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	_ "image/gif"
	_ "image/png"
	"io"
	"log/slog"
	"net/http"

	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/jung-kurt/gofpdf"
)

type Orientation string

const (
	Landscape Orientation = "LANDSCAPE"
	Portrait  Orientation = "PORTRAIT"
)

func (o Orientation) String() string {
	switch o {
	case Landscape:
		return "L"
	case Portrait:
		return "P"
	default:
		return ""
	}
}

type Format string

const (
	A3 Format = "A3"
	A4 Format = "A4"
)

func (f Format) String() string {
	return string(f)
}

type PageSizeUnit string

func (s PageSizeUnit) String() string {
	return string(s)
}

const (
	mm PageSizeUnit = "mm"
)

type PageSize struct {
	Width, Height float64
	Units         PageSizeUnit
}

var (
	defaultSizes map[Format]map[Orientation]PageSize
)

func init() {
	defaultSizes = map[Format]map[Orientation]PageSize{
		A3: {
			Portrait: {
				Width:  297,
				Height: 420,
				Units:  mm,
			},
			Landscape: {
				Width:  420,
				Height: 297,
				Units:  mm,
			},
		},
		A4: {
			Portrait: {
				Width:  210,
				Height: 297,
				Units:  mm,
			},
			Landscape: {
				Width:  297,
				Height: 210,
				Units:  mm,
			},
		},
	}
}

type Binder struct {
	filesStore files.Store
}

func NewBinder(filesStore files.Store) *Binder {
	return &Binder{filesStore: filesStore}
}

func (b *Binder) Bind(
	ctx context.Context,
	orientation Orientation,
	format Format,
	pages []albums.AlbumPage,
) (io.Reader, error) {
	var (
		pageSize = defaultSizes[format][orientation]

		// consider using https://github.com/signintech/gopdf
		pdf = gofpdf.New(orientation.String(), pageSize.Units.String(), format.String(), "")

		res = &bytes.Buffer{}
	)

	pdf.AddUTF8FontFromBytes("Arsenal", "", arsenalRegularTTF)
	pdf.SetFont("Arsenal", "", 16)

	for _, page := range pages {
		switch t := page.(type) {
		case albums.AlbumPageSplitCover:
			if err := b.addSplitCoverPage(ctx, pdf, t, orientation, pageSize); err != nil {
				return nil, err
			}
		case albums.AlbumPageVisualization:
			if err := b.addVisualizationPage(ctx, pdf, t, orientation, pageSize); err != nil {
				return nil, err
			}
		default:
			slog.Warn("unknown page type")
			continue
		}
	}

	return res, pdf.Output(res)
}

func (b *Binder) downloadFile(ctx context.Context, id string) ([]byte, string, error) {
	file, err := b.filesStore.Get(ctx, files.IDIn(id))
	if err != nil {
		return nil, "", fmt.Errorf("can't get file %s: %w", id, err)
	}

	resp, err := http.Get(file.URL)
	if err != nil {
		return nil, "", fmt.Errorf("can't download file from %s: %w", file.URL, err)
	}

	defer resp.Body.Close()

	imgBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("can't read file body: %w", err)
	}

	return imgBytes, file.MimeType, nil
}

func typeByMime(mimeType string) string {
	var (
		imgType string
	)

	switch mimeType {
	case "image/png":
		imgType = "png"
	case "image/jpeg", "image/jpg":
		imgType = "jpg"
	case "image/gif":
		imgType = "gif"
	}

	return imgType
}

func pageUniqName(page albums.AlbumPage) (string, error) {
	pageBytes, err := json.Marshal(page)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("split-%x", md5.Sum(pageBytes)), nil
}
