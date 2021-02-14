package main

import (
	"github.com/docopt/docopt-go"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"path/filepath"
)

const usage = `binder

Book binder.

Usage:
	album <format> [<orientation>] [--dir=DIR]

Options:
	--dir=DIR  Directory name with images [default: ./]
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

var (
	defaultSizes map[Format]map[Orientation]struct {
		left, top     float64
		width, height float64
	}
)

func init() {
	defaultSizes = map[Format]map[Orientation]struct {
		left, top     float64
		width, height float64
	}{
		A3: {
			// 297x420
			Portrait: {
				left:   20,
				top:    5,
				width:  272,
				height: 410,
			},
			// 420x297
			Landscape: {
				left:   20,
				top:    5,
				width:  395,
				height: 287,
			},
		},
		A4: {
			// 210x297
			Portrait: {
				left:   20,
				top:    5,
				width:  185,
				height: 287,
			},
			// 297x210
			Landscape: {
				left:   20,
				top:    5,
				width:  272,
				height: 200,
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
	dirname, _ := opts.String("--dir")

	var (
		format Format      = A4
		orient Orientation = Landscape
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
		orient = Landscape
	case "portrait":
		orient = Portrait
	default:
		if len(orientationStr) > 0 {
			log.Fatal("not valid orientation (should be landscape or portrait)")
		}
	}

	path, err := filepath.Abs(dirname)
	if err != nil {
		log.Fatalf("can't find out directory: %s", err)
	}

	println(path)

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

	pdf := gofpdf.New(orient.String(), "mm", format.String(), "")

	println(orient.String(), format.String())

	x := defaultSizes[format][orient].left
	y := defaultSizes[format][orient].top
	w := defaultSizes[format][orient].width
	h := defaultSizes[format][orient].height

	println(x,y,w,h)

	for _, path := range images {
		var (
			opt gofpdf.ImageOptions
		)

		opt.ImageType = "jpg"

		pdf.AddPage()

		pdf.ImageOptions(path, x, y, w, h, false, opt, 0, "")
	}

	if err := pdf.OutputFileAndClose("book.pdf"); err != nil {
		log.Fatal("can't write pdf file")
	}

	log.Print("Done")
}
