package bookbinder

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"

	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/disintegration/gift"
	"github.com/jung-kurt/gofpdf"
)

func (b *Binder) addVisualizationPage(
	ctx context.Context,
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageVisualization,
	orientation Orientation,
	pageSize PageSize,
) error {
	imgBytes, mimeType, err := b.downloadFile(ctx, page.FileID)
	if err != nil {
		return fmt.Errorf("cannot download visualization page image: %w", err)
	}

	name, err := pageUniqName(page)
	if err != nil {
		return err
	}

	pdf.AddPage()

	_, img, _, x, y, w, h := process(bytes.NewReader(imgBytes), orientation, pageSize)

	opt := gofpdf.ImageOptions{ImageType: typeByMime(mimeType), ReadDpi: false, AllowNegativePosition: false}

	pdf.RegisterImageOptionsReader(name, opt, img)

	pdf.ImageOptions(name, x, y, w, h, false, opt, 0, "")

	return nil
}

func process(img io.Reader, orientation Orientation, page PageSize) (image.Image, io.Reader, bool, float64, float64, float64, float64) {
	i, img, err := decode(img)
	if err != nil {
		log.Fatalf("can't decode: %s", err)
	}

	img, err, rotated := rotate(i, orientation)
	if err != nil {
		log.Fatalf("can't rotate: %s", err)
	}

	i, img, err = decode(img)
	if err != nil {
		log.Fatalf("can't decode: %s", err)
	}

	var (
		x0, y0 float64
	)
	switch orientation {
	case Portrait:
		page, x0, y0 = pad(page, 20, 5, 5, 5)
	case Landscape:
		page, x0, y0 = pad(page, 5, 5, 5, 20)
	}

	x, y, w, h := place(i, page, orientation)

	return i, img, rotated, x0 + x, y0 + y, w, h
}

func rotate(img image.Image, orientation Orientation) (io.Reader, error, bool) {
	rotated := false

	g := gift.New()

	x := img.Bounds().Dx()
	y := img.Bounds().Dy()
	if (orientation == Landscape && y/x > 1) || (orientation == Portrait && y/x < 1) {
		g.Add(gift.Rotate270())
		rotated = true
	}

	dst := image.NewRGBA(g.Bounds(img.Bounds()))

	g.Draw(dst, img)

	b := &bytes.Buffer{}
	err := jpeg.Encode(b, dst, nil)
	if err != nil {
		return nil, err, false
	}

	return b, nil, rotated
}

func pad(p PageSize, t, r, b, l float64) (PageSize, float64, float64) {
	return PageSize{Width: p.Width - r - l, Height: p.Height - t - b, Units: p.Units}, l, t
}

func place(img image.Image, page PageSize, orientation Orientation) (float64, float64, float64, float64) {
	switch orientation {
	case Portrait:
		bw1, bw2, bw3, bw4 := placeByWidth(img, page)
		if bw1 < 0 || bw2 < 0 {
			return placeByHeight(img, page)
		}
		return bw1, bw2, bw3, bw4
	case Landscape:
		bw1, bw2, bw3, bw4 := placeByHeight(img, page)
		if bw1 < 0 || bw2 < 0 {
			return placeByWidth(img, page)
		}
		return bw1, bw2, bw3, bw4
	}

	return 0, 0, 0, 0
}

func placeByWidth(img image.Image, page PageSize) (float64, float64, float64, float64) {
	imgRat := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	nh := page.Width * imgRat
	return 0, (page.Height - nh) / 2, page.Width, nh
}

func placeByHeight(img image.Image, page PageSize) (float64, float64, float64, float64) {
	imgRat := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	nw := page.Height / imgRat
	return (page.Width - nw) / 2, 0, nw, page.Height
}

func decode(img io.Reader) (image.Image, io.Reader, error) {
	b, err := io.ReadAll(img)
	if err != nil {
		return nil, nil, err
	}

	i, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, nil, err
	}

	return i, bytes.NewReader(b), nil
}
