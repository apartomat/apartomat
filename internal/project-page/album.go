package project_page

import (
	"context"
	"errors"

	. "github.com/apartomat/apartomat/internal/store/albumfiles"
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/projectpages"
)

func (u *Service) GetAlbumAndFile(ctx context.Context, projectPageID string) (*albums.Album, *files.File, error) {
	page, err := u.ProjectPages.Get(ctx, projectpages.IDIn(projectPageID))
	if err != nil {
		return nil, nil, err
	}

	album, err := u.Albums.Get(ctx, albums.ProjectIDIn(page.ProjectID))
	if err != nil {
		return nil, nil, err
	}

	albumFile, err := u.AlbumFiles.GetLastVersion(
		ctx,
		And(
			AlbumIDIn(album.ID),
			StatusIn(StatusDone),
		),
	)
	if err != nil {
		return nil, nil, err
	}

	if albumFile.FileID == nil {
		return nil, nil, errors.New("albumFile has no FileID")
	}

	file, err := u.Files.Get(ctx, files.IDIn(*albumFile.FileID))
	if err != nil {
		return nil, nil, err
	}

	return album, file, nil
}
