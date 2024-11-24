package project_page

import (
	"context"
	"errors"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/public_sites"
	"github.com/apartomat/apartomat/internal/store/visualizations"
)

var (
	ErrForbidden = errors.New("forbidden")
)

type Service struct {
	Files          files.Store
	PublicSites    public_sites.Store
	Visualizations visualizations.Store
	Albums         albums.Store
	AlbumFiles     albumFiles.Store
	Houses         houses.Store
}

func NewService(
	filesStore files.Store,
	publicSitesStore public_sites.Store,
	visualizationsStore visualizations.Store,
	albumsStore albums.Store,
	albumsFilesStore albumFiles.Store,
	housesStore houses.Store,
) *Service {
	return &Service{
		Files:          filesStore,
		PublicSites:    publicSitesStore,
		Visualizations: visualizationsStore,
		Albums:         albumsStore,
		AlbumFiles:     albumsFilesStore,
		Houses:         housesStore,
	}
}

func (u *Service) GetProjectPage(ctx context.Context, id string) (*public_sites.PublicSite, error) {
	s, err := u.PublicSites.Get(ctx, public_sites.IDIn(id))
	if err != nil {
		return nil, err
	}

	if !s.Is(public_sites.Public()) {
		return nil, ErrForbidden
	}

	return s, nil
}
