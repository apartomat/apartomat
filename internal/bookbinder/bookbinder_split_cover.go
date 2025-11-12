package bookbinder

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/jung-kurt/gofpdf"
	"image"
)

func (b *Binder) addSplitCoverPage(
	ctx context.Context,
	pdf *gofpdf.Fpdf,
	page albums.AlbumPageSplitCover,
	orientation Orientation,
	pageSize PageSize,
) error {
	pdf.AddPage()

	// Layout constants (mm) tuned to match reference spacing
	leftMargin := 22.0
	rightMargin := 12.0
	topMargin := 30.0
	bottomMargin := 26.0
	gutter := 0.0 // no space between halves

	totalW := pageSize.Width
	totalH := pageSize.Height

	// Compute halves
	halfW := (totalW - leftMargin - rightMargin - gutter) / 2.0
	leftX := leftMargin
	rightX := leftMargin + halfW + gutter

	imgBytes, mimeType, err := b.downloadFile(ctx, page.ImgFileID)

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
	targetW := halfW
	targetH := totalH

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
	y := (targetH - drawH) / 2.0

	pdf.ImageOptions(name, x, y, drawW, drawH, false, opt, 0, "")

	// Ensure no bleed into the left half: paint left area white over any overflow
	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(0, 0, rightX, totalH, "F")

	// 2) Left half: title, subtitle (centered horizontally AND vertically as a block),
	// bottom city and year (two lines, centered horizontally)
	// Fonts
	titleFontSize := 28.0
	subtitleFontSize := 18.0
	metaFontSize := 12.0

	// Text area for left half
	leftTextX := leftX
	leftTextW := halfW
	topTextY := topMargin
	bottomTextY := totalH - bottomMargin

	// Compute vertical centering for title/subtitle block
	blockGap := subtitleFontSize * 0.8
	hasSubtitle := page.Subtitle != nil && *page.Subtitle != ""
	blockHeight := titleFontSize
	if hasSubtitle {
		blockHeight += blockGap + subtitleFontSize
	}
	leftRectH := totalH - topMargin - bottomMargin
	blockStartY := topTextY + (leftRectH-blockHeight)/2.0 - 5.0 // lift a bit above exact center

	// Title (centered horizontally)
	pdf.SetFont("Arsenal", "", titleFontSize)
	titleStr := page.Title
	titleWidth := pdf.GetStringWidth(titleStr)
	titleX := leftTextX + (leftTextW-titleWidth)/2.0
	titleY := blockStartY + titleFontSize
	pdf.Text(titleX, titleY, titleStr)

	// Subtitle (if any), centered under title
	if hasSubtitle {
		pdf.SetFont("Arsenal", "", subtitleFontSize)
		subStr := *page.Subtitle
		subWidth := pdf.GetStringWidth(subStr)
		subX := leftTextX + (leftTextW-subWidth)/2.0
		subY := titleY + blockGap
		pdf.Text(subX, subY, subStr)
	}

	// Bottom meta: city and year, two lines at the very bottom of left half
	pdf.SetFont("Arsenal", "", metaFontSize)
	// Line height
	lh := metaFontSize * 1.4
	metaY2 := bottomTextY
	metaY1 := metaY2 - lh

	// City on the first line (if present), centered horizontally
	if page.City != nil && *page.City != "" {
		cityStr := *page.City
		cityWidth := pdf.GetStringWidth(cityStr)
		cityX := leftTextX + (leftTextW-cityWidth)/2.0
		pdf.Text(cityX, metaY1, cityStr)
	}

	// Year on the second line (if present), centered horizontally
	if page.Year != nil {
		yearStr := fmt.Sprintf("%d", *page.Year)
		yearWidth := pdf.GetStringWidth(yearStr)
		yearX := leftTextX + (leftTextW-yearWidth)/2.0
		pdf.Text(yearX, metaY2, yearStr)
	}

	return nil
}
