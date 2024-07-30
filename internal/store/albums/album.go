package albums

import (
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"time"
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

type AlbumPageCover struct {
	CoverID string
	FileID  string
	Rotate  float64
}

func (AlbumPageCover) IsAlbumPage() bool {
	return true
}

type AlbumPageCoverUploaded struct {
	FileID string
	Rotate float64
}

func (AlbumPageCoverUploaded) IsAlbumPage() bool {
	return true
}

type AlbumPageVisualization struct {
	VisualizationID string
	FileID          string
	Rotate          float64
}

func (AlbumPageVisualization) IsAlbumPage() bool {
	return true
}

func (album *Album) AddPageWithVisualization(vis *visualizations.Visualization) (AlbumPageVisualization, int) {
	var (
		page = AlbumPageVisualization{
			VisualizationID: vis.ID,
			FileID:          vis.FileID,
		}
	)

	album.Pages = append(album.Pages, page)

	return page, len(album.Pages) - 1
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
