package albums

import (
	"time"

	"github.com/apartomat/apartomat/internal/store/visualizations"
)

type Album struct {
	ID         string
	Name       string
	Version    int
	Settings   Settings
	Pages      []AlbumPage
	CreatedAt  time.Time
	ModifiedAt time.Time
	ProjectID  string
}

func NewAlbum(id, name string, settings Settings, projectID string) *Album {
	now := time.Now()

	return &Album{
		ID:         id,
		Name:       name,
		Version:    0,
		Settings:   settings,
		CreatedAt:  now,
		ModifiedAt: now,
		ProjectID:  projectID,
	}
}

type Settings struct {
	PageOrientation PageOrientation
	PageSize        PageSize
}

type PageOrientation string

const (
	Landscape PageOrientation = "LANDSCAPE"
	Portrait  PageOrientation = "PORTRAIT"
)

type PageSize string

const (
	A3 PageSize = "A3"
	A4 PageSize = "A4"
)

type AlbumPage interface {
	IsAlbumPage() bool
}

type AlbumPageSplitCover struct {
	ID        string
	Rotate    float64
	Title     string
	Subtitle  *string
	ImgFileID string
	WithQR    bool
	City      *string
	Year      *int
}

func (AlbumPageSplitCover) IsAlbumPage() bool {
	return true
}

type AlbumPageCoverUploaded struct {
	ID     string
	Rotate float64
	FileID string
}

func (AlbumPageCoverUploaded) IsAlbumPage() bool {
	return true
}

type AlbumPageVisualization struct {
	ID              string
	Rotate          float64
	VisualizationID string
	FileID          string
}

func (AlbumPageVisualization) IsAlbumPage() bool {
	return true
}

func (album *Album) AddVisualizationPageWithID(
	vis *visualizations.Visualization,
	pageID string,
) (AlbumPageVisualization, int) {
	var (
		page = AlbumPageVisualization{
			ID:              pageID,
			VisualizationID: vis.ID,
			FileID:          vis.FileID,
		}
	)

	album.Pages = append(album.Pages, page)

	return page, len(album.Pages) - 1
}

func (album *Album) AddUploadedCoverPageWithID(
	fileID string,
	pageID string,
) (AlbumPageCoverUploaded, int) {
	var (
		page = AlbumPageCoverUploaded{ID: pageID, FileID: fileID}
	)

	album.Pages = append(
		[]AlbumPage{page},
		album.Pages...,
	)

	return page, 0
}

func (album *Album) AddSplitCoverPageWithID(
	pageID string,
	title string,
	subtitle *string,
	imgFileID string,
	withQR bool,
	city *string,
	year *int,
) (AlbumPageSplitCover, int) {
	var (
		page = AlbumPageSplitCover{
			ID:        pageID,
			Title:     title,
			Subtitle:  subtitle,
			ImgFileID: imgFileID,
			WithQR:    withQR,
			City:      city,
			Year:      year,
		}
	)

	album.Pages = append(
		[]AlbumPage{page},
		album.Pages...,
	)

	return page, 0
}

func (album *Album) ChangePageSize(size PageSize) {
	album.Settings.PageSize = size
}

func (album *Album) ChangePageOrientation(orientation PageOrientation) {
	album.Settings.PageOrientation = orientation
}

func (album *Album) UpVersion() {
	album.Version++
}
