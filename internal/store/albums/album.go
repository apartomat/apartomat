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
	Pages      []AlbumPageVisualization
	CreatedAt  time.Time
	ModifiedAt time.Time
	ProjectID  string
}

func NewAlbum(id, name, projectID string) *Album {
	now := time.Now()

	return &Album{
		ID:         id,
		Name:       name,
		Version:    0,
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

type AlbumPageVisualization struct {
	VisualizationID string
	FileID          string
}

func (album *Album) AddPageWithVisualization(vis *visualizations.Visualization) (int, error) {
	album.Pages = append(
		album.Pages,
		AlbumPageVisualization{
			VisualizationID: vis.ID,
			FileID:          vis.FileID,
		},
	)

	return len(album.Pages) - 1, nil
}

func (album *Album) ChangePageSize(size PageSize) {
	album.Settings.PageSize = size
}

func (album *Album) ChangePageOrientation(orientation PageOrientation) {
	album.Settings.PageOrientation = orientation
}
