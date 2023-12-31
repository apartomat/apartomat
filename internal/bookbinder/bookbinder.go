package bookbinder

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/disintegration/gift"
	"github.com/jung-kurt/gofpdf"
	"image"
	"image/jpeg"
	"io"
	"log"
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

func Bind(orientation Orientation, format Format, pages []string, images map[string]io.Reader) (io.Reader, error) {
	var (
		page = defaultSizes[format][orientation]

		pdf = gofpdf.New(orientation.String(), page.Units.String(), format.String(), "")

		res = &bytes.Buffer{}
	)

	for _, name := range pages {
		if _, ok := images[name]; !ok {
			continue
		}

		var (
			opt = gofpdf.ImageOptions{
				ImageType:             "jpg", // @todo
				ReadDpi:               false,
				AllowNegativePosition: false,
			}
		)

		pdf.AddPage()

		_, img, err := decode(images[name])
		if err != nil {
			log.Fatalf("can't decode file: %s", err)
		}

		_, img, _, x, y, w, h := process(img, orientation, page)

		name := fmt.Sprintf("%x", md5.Sum([]byte(name)))

		pdf.RegisterImageOptionsReader(name, opt, img)

		pdf.ImageOptions(name, x, y, w, h, false, opt, 0, "")
	}

	return res, pdf.Output(res)
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
