package albums

import (
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"time"
)

type Album struct {
	ID         string
	Name       string
	Version    int
	CreatedAt  time.Time
	ModifiedAt time.Time
	ProjectID  string
	Pages      []AlbumPageVisualization
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

type AlbumPageVisualization struct {
	VisualizationID string
	FileID          string
}
