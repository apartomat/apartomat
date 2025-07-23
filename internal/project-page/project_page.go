package project_page

import (
	"context"
	"errors"

	"github.com/apartomat/apartomat/internal/store/albumfiles"
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projectpages"
	"github.com/apartomat/apartomat/internal/store/visualizations"
)

var (
	ErrForbidden = errors.New("forbidden")
)

type Service struct {
	Files          files.Store
	ProjectPages   projectpages.Store
	Visualizations visualizations.Store
	Albums         albums.Store
	AlbumFiles     albumfiles.Store
	Houses         houses.Store
}

func NewService(
	filesStore files.Store,
	projectPages projectpages.Store,
	visualizationsStore visualizations.Store,
	albumsStore albums.Store,
	albumsFilesStore albumfiles.Store,
	housesStore houses.Store,
) *Service {
	return &Service{
		Files:          filesStore,
		ProjectPages:   projectPages,
		Visualizations: visualizationsStore,
		Albums:         albumsStore,
		AlbumFiles:     albumsFilesStore,
		Houses:         housesStore,
	}
}

func (u *Service) GetProjectPage(ctx context.Context, id string) (*projectpages.ProjectPage, error) {
	s, err := u.ProjectPages.Get(ctx, projectpages.IDIn(id))
	if err != nil {
		return nil, err
	}

	if !s.Is(projectpages.Public()) {
		return nil, ErrForbidden
	}

	return s, nil
}
