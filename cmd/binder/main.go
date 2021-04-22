package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/disintegration/gift"
	"github.com/docopt/docopt-go"
	"github.com/jung-kurt/gofpdf"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const usage = `binder

Book binder.

Usage:
	binder <format> <orientation> <path>
`

type Orientation string

const (
	Landscape Orientation = "landscape"
	Portrait  Orientation = "portrait"
)

func (o Orientation) String() string {
	switch o {
	case Landscape:
		return "L"
	case Portrait:
		return "P"
	default:
		return "P"
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

func sz(s PageSize) (left float64, top float64, width float64, height float64, units PageSizeUnit) {
	return 0, 0, s.Width, s.Height, s.Units
}

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

func main() {
	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		log.Fatalf("can't parse arguments: %s", err)
	}

	formatStr, _ := opts.String("<format>")
	orientationStr, _ := opts.String("<orientation>")
	dirname, _ := opts.String("<path>")

	var (
		format      = A4
		orientation = Landscape
	)

	switch formatStr {
	case "a3":
	case "A3":
		format = A3
	case "a4":
	case "A4":
		format = A4
	default:
		if len(formatStr) > 0 {
			log.Fatal("not valid format (should be A3 or A4)")
		}
	}

	switch orientationStr {
	case "landscape":
		orientation = Landscape
	case "portrait":
		orientation = Portrait
	default:
		if len(orientationStr) > 0 {
			log.Fatal("not valid orientation (should be landscape or portrait)")
		}
	}

	path, err := filepath.Abs(dirname)
	if err != nil {
		log.Fatalf("can't find out directory: %s", err)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("can't read dir %s: %s", path, err)
	}

	images := make([]string, 0)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		switch filepath.Ext(f.Name()) {
		case ".jpeg", ".jpg", ".png":
			images = append(images, filepath.Join(path, f.Name()))
		default:
			continue
		}
	}

	if len(images) == 0 {
		log.Fatal("no files")
	}

	page := defaultSizes[format][orientation]

	pdf := gofpdf.New(orientation.String(), page.Units.String(), format.String(), "")

	for _, path := range images {
		var (
			opt gofpdf.ImageOptions
		)

		pdf.AddPage()

		opt.ImageType = "jpg" // @todo

		f, err := os.Open(path)
		if err != nil {
			log.Fatalf("can't read file: %s", err)
		}

		orig, img, err := decode(f)
		if err != nil {
			log.Fatalf("can't decode file: %s", err)
		}

		outp, img, rt, x, y, w, h := process(img, orientation, page)

		err = f.Close()
		if err != nil {
			log.Printf("can't close file: %s\n", err)
		}

		name := fmt.Sprintf("%x", md5.Sum([]byte(path)))

		pdf.RegisterImageOptionsReader(name, opt, img)
		pdf.ImageOptions(name, x, y, w, h, false, opt, 0, "")

		//

		msg := fmt.Sprintf(
			"%s: orig: %dx%d, rotate: %+v, outp: %dx%d, place: %f,%f; %f,%f",
			path,
			orig.Bounds().Dx(),
			orig.Bounds().Dy(),
			rt,
			outp.Bounds().Dx(),
			outp.Bounds().Dy(),
			x, y,
			w, h,
		)

		pdf.SetFont("Arial", "", 12)
		pdf.Cell(100, 10, msg)

		log.Print(msg)
	}

	if err := pdf.OutputFileAndClose("book.pdf"); err != nil {
		log.Fatal("can't write pdf file; %s", err)
	}

	log.Print("Done")
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

func place(img image.Image, page PageSize, orientation Orientation) (float64, float64, float64, float64) {
	//imgRat := imgRat(img)
	//pageRat := pageRat(page)

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

//func placeByHeight(img image.Image, page PageSize) (float64, float64, float64, float64) {
	//imgRat := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	//nw := page.Height * imgRat
	//return (page.Width - nw) / 2, 0, nw, page.Height
//}

func placeByHeight(img image.Image, page PageSize) (float64, float64, float64, float64) {
	imgRat := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	nw := page.Height / imgRat
	return (page.Width - nw) / 2, 0, nw, page.Height
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

	x, y, w, h := place(i, page, orientation)

	return i, img, rotated, x, y, w, h
}

func decode(img io.Reader) (image.Image, io.Reader, error) {
	b, err := ioutil.ReadAll(img)
	if err != nil {
		return nil, nil, err
	}

	i, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, nil, err
	}

	return i, bytes.NewReader(b), nil
}

func imgRat(img image.Image) float64 {
	return float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
}

func pageRat(page PageSize) float64 {
	return float64(page.Height) / float64(page.Width)
}
