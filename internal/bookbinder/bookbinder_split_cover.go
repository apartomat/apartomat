package bookbinder

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"image"
	"strings"

	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/jung-kurt/gofpdf"
)

const (
	splitCoverTitleFontPx    = 48.0
	splitCoverSubtitleFontPx = 20.0
	splitCoverMetaFontPx     = 16.0

	pxPerInch = 96.0
	ptPerInch = 72.0
)

func (b *Binder) addSplitCoverPage(
	ctx context.Context,
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	orientation Orientation,
	pageSize PageSize,
) error {
	pdf.AddPage()

	switch orientation {
	case Landscape:
		return splitCoverLandscape(ctx, pdf, page, pageSize, b.downloadFile)
	case Portrait:
		return fmt.Errorf("split cover portrait not implemented yet")
	default:
		return fmt.Errorf("unsupported orientation %q", orientation)
	}
}

func pxToPt(px float64) float64 {
	return px * ptPerInch / pxPerInch
}

func splitCoverLandscape(
	ctx context.Context,
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	pageSize PageSize,
	download func(context.Context, string) ([]byte, string, error),
) error {
	imgBytes, mimeType, err := download(ctx, page.ImgFileID)
	if err != nil {
		return err
	}

	if err := splitCoverRight(pdf, page, pageSize, imgBytes, mimeType); err != nil {
		return err
	}

	splitCoverLeft(pdf, page, pageSize)

	return nil
}

func splitCoverRight(
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	pageSize PageSize,
	imgBytes []byte,
	mimeType string,
) error {
	opt := gofpdf.ImageOptions{ImageType: typeByMime(mimeType), ReadDpi: false, AllowNegativePosition: false}

	i, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return fmt.Errorf("can't decode split cover image: %w", err)
	}

	// Unique name from page data
	pageBytes, err := json.Marshal(page)
	if err != nil {
		return err
	}
	name := fmt.Sprintf("split-%x", md5.Sum(pageBytes))
	// Register using a fresh reader (Register may consume the reader)
	pdf.RegisterImageOptionsReader(name, opt, bytes.NewReader(imgBytes))

	// Target rect for image: entire right half without any inner margins
	targetW := pageSize.Width / 2.0
	targetH := pageSize.Height

	rightX := pageSize.Width / 2.0

	// Compute COVER fit (may crop). Scale so that image fully covers target
	srcW := float64(i.Bounds().Dx())
	srcH := float64(i.Bounds().Dy())
	ratioW := targetW / srcW
	ratioH := targetH / srcH
	scale := ratioW
	if ratioH > scale {
		scale = ratioH
	}

	drawW := srcW * scale
	drawH := srcH * scale

	// Center inside right half; allow negative offsets (image may bleed outside)
	x := rightX + (targetW-drawW)/2.0
	y := (pageSize.Height - drawH) / 2.0

	pdf.ImageOptions(name, x, y, drawW, drawH, false, opt, 0, "")

	// Ensure no bleed into the left half: paint left area white over any overflow
	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(0, 0, rightX, pageSize.Height, "F")

	return nil
}

func splitCoverLeft(
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	pageSize PageSize,
) {

	splitCoverLeftTitle(pdf, page, pageSize)
	splitCoverLeftSubtitle(pdf, page, pageSize)
	splitCoverLeftCity(pdf, page, pageSize)
	splitCoverLeftYear(pdf, page, pageSize)
}

func splitCoverLeftTitle(
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	pageSize PageSize,
) {
	fontSize := pxToPt(splitCoverTitleFontPx)
	pdf.SetFont("Arsenal", "", fontSize)

	width := pdf.GetStringWidth(page.Title)

	x := (pageSize.Width/2.0 - width) / 2.0
	y := pageSize.Height/2.0 - pageSize.Height*0.07

	pdf.Text(x, y, page.Title)
}

func splitCoverLeftSubtitle(
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	pageSize PageSize,
) {
	var (
		lineHeight = 1.3
		marginTop  = 4.0
	)

	if page.Subtitle == nil || *page.Subtitle == "" {
		return
	}

	fontSize := pxToPt(splitCoverSubtitleFontPx)
	pdf.SetFont("Arsenal", "", fontSize)

	_, fontUnitSize := pdf.GetFontSize()

	startY := pageSize.Height/2.0 - pageSize.Height*0.07 + fontUnitSize + marginTop

	lines := strings.Split(*page.Subtitle, "\n")

	for i, subStr := range lines {
		width := pdf.GetStringWidth(subStr)
		x := (pageSize.Width/2.0 - width) / 2.0
		y := startY + float64(i)*fontUnitSize*lineHeight
		pdf.Text(x, y, subStr)
	}
}

func splitCoverLeftCity(
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	pageSize PageSize,
) {
	if page.City == nil || *page.City == "" {
		return
	}

	fontSize := pxToPt(splitCoverMetaFontPx)
	width := pageSize.Width / 2.0
	y := pageSize.Height - pxToPt(16.0)

	pdf.SetFont("Arsenal", "", fontSize)
	cityStr := *page.City
	cityWidth := pdf.GetStringWidth(cityStr)
	cityX := (width - cityWidth) / 2.0
	pdf.Text(cityX, y, cityStr)
}

func splitCoverLeftYear(
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	pageSize PageSize,
) {
	if page.Year == nil {
		return
	}

	fontSize := pxToPt(splitCoverMetaFontPx)
	width := pageSize.Width / 2.0
	y := pageSize.Height - pxToPt(8)

	pdf.SetFont("Arsenal", "", fontSize)
	yearStr := fmt.Sprintf("•%d•", *page.Year)
	yearWidth := pdf.GetStringWidth(yearStr)
	yearX := (width - yearWidth) / 2.0
	pdf.Text(yearX, y, yearStr)
}
