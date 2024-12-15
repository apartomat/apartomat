package project_page

import (
	"context"
	"errors"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projectpage"
	"github.com/apartomat/apartomat/internal/store/visualizations"
)

var (
	ErrForbidden = errors.New("forbidden")
)

type Service struct {
	Files          files.Store
	ProjectPages   projectpage.Store
	Visualizations visualizations.Store
	Albums         albums.Store
	AlbumFiles     albumFiles.Store
	Houses         houses.Store
}

func NewService(
	filesStore files.Store,
	projectPages projectpage.Store,
	visualizationsStore visualizations.Store,
	albumsStore albums.Store,
	albumsFilesStore albumFiles.Store,
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

func (u *Service) GetProjectPage(ctx context.Context, id string) (*projectpage.ProjectPage, error) {
	s, err := u.ProjectPages.Get(ctx, projectpage.IDIn(id))
	if err != nil {
		return nil, err
	}

	if !s.Is(projectpage.Public()) {
		return nil, ErrForbidden
	}

	return s, nil
}
