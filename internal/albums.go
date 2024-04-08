package apartomat

import (
	"context"
	"errors"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files"
	. "github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/visualizations"
)

func (u *Apartomat) CreateAlbum(
	ctx context.Context,
	projectID string,
	name string,
) (*Album, error) {
	project, err := u.Projects.Get(ctx, projects.IDIn(projectID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanCreateAlbum(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't create album in project (id=%s): %w", project.ID, ErrForbidden)
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, err
	}

	album := NewAlbum(id, name, Settings{PageOrientation: Portrait, PageSize: A4}, project.ID)

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return album, nil
}

func (u *Apartomat) GetAlbums(
	ctx context.Context,
	projectID string,
	limit, offset int,
) ([]*Album, error) {
	if ok, err := u.Acl.CanGetAlbumsOfProjectID(ctx, auth.UserFromCtx(ctx), projectID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) albums: %w", projectID, ErrForbidden)
	}

	var (
		spec = ProjectIDIn(projectID)
	)

	res, err := u.Albums.List(ctx, spec, SortDefault, limit, offset)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Apartomat) GetAlbum(
	ctx context.Context,
	id string,
) (*Album, error) {
	album, err := u.Albums.Get(ctx, IDIn(id))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanGetAlbum(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) album (id=%s): %w", album.ProjectID, album.ID, ErrForbidden)
	}

	return album, nil
}

func (u *Apartomat) CountAlbums(
	ctx context.Context,
	projectID string,
) (int, error) {
	if ok, err := u.Acl.CanCountAlbumsOfProjectID(ctx, auth.UserFromCtx(ctx), projectID); err != nil {
		return 0, err
	} else if !ok {
		return 0, fmt.Errorf("can't get project (id=%s) albums: %w", projectID, ErrForbidden)
	}

	var (
		spec = ProjectIDIn(projectID)
	)

	return u.Albums.Count(ctx, spec)
}

func (u *Apartomat) DeleteAlbum(ctx context.Context, id string) (*Album, error) {
	album, err := u.Albums.Get(ctx, IDIn(id))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanDeleteAlbum(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't delete album (id=%s): %w", album.ID, ErrForbidden)
	}

	err = u.Albums.Delete(ctx, album)

	return album, err
}

type VisualizationWithPosition struct {
	Position      int
	Visualization *visualizations.Visualization
}

func (u *Apartomat) AddVisualizationsToAlbum(
	ctx context.Context,
	albumID string,
	visualizationID []string,
) ([]VisualizationWithPosition, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanAddPageToAlbum(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add visualization to album (id=%s): %w", album.ID, ErrForbidden)
	}

	list, err := u.Visualizations.List(
		ctx,
		visualizations.IDIn(visualizationID...),
		visualizations.SortDefault,
		len(visualizationID),
		0,
	)
	if err != nil {
		return nil, err
	}

	if len(list) != len(visualizationID) {
		return nil, fmt.Errorf("visualization (id=%s): %w", visualizationID, ErrNotFound)
	}

	var (
		res = make([]VisualizationWithPosition, len(visualizationID))
	)

visLoop:
	for i, id := range visualizationID {
		for _, vis := range list {
			if vis.ID == id {
				pos, err := album.AddPageWithVisualization(vis)
				if err != nil {
					return nil, err
				}

				res[i] = VisualizationWithPosition{
					Position:      pos,
					Visualization: vis,
				}

				continue visLoop
			}
		}
	}

	album.UpVersion()

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Apartomat) ChangeAlbumPageSize(
	ctx context.Context,
	albumID string,
	size PageSize,
) (*Album, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanChangeAlbumSettings(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't change album (id=%s) settings: %w", album.ID, ErrForbidden)
	}

	album.ChangePageSize(size)

	album.UpVersion()

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return album, nil
}

func (u *Apartomat) ChangeAlbumPageOrientation(
	ctx context.Context,
	albumID string,
	orientation PageOrientation,
) (*Album, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanChangeAlbumSettings(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't change album (id=%s) settings: %w", album.ID, ErrForbidden)
	}

	album.ChangePageOrientation(orientation)

	album.UpVersion()

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return album, nil
}

func (u *Apartomat) GetAlbumRecentFile(ctx context.Context, albumID string) (*albumFiles.AlbumFile, *files.File, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, nil, err
	}

	if ok, err := u.Acl.CanGetAlbumFile(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, nil, err
	} else if !ok {
		return nil, nil, fmt.Errorf("can't get album (id=%s) recent file: %w", album.ID, ErrForbidden)
	}

	albumFile, err := u.AlbumFiles.GetLastVersion(ctx, albumFiles.And(albumFiles.AlbumIDIn(albumID)))
	if err != nil {
		return nil, nil, err
	}

	if albumFile != nil && albumFile.FileID != nil {
		file, err := u.Files.Get(ctx, files.IDIn(*albumFile.FileID))
		if err != nil {
			return nil, nil, err
		}

		return albumFile, file, err
	}

	return albumFile, nil, err
}

var (
	ErrAlbumFileVersionExisted = errors.New("album file version existed")
)

func (u *Apartomat) StartGenerateAlbumFile(ctx context.Context, albumID string) (*albumFiles.AlbumFile, *files.File, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, nil, err
	}

	if ok, err := u.Acl.CanGetAlbumFile(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, nil, err
	} else if !ok {
		return nil, nil, fmt.Errorf("can't generate album (id=%s) file: %w", album.ID, ErrForbidden)
	}

	existedFile, err := u.AlbumFiles.GetLastVersion(ctx, albumFiles.And(albumFiles.AlbumIDIn(albumID)))
	if err != nil && !errors.Is(err, albumFiles.ErrAlbumFileNotFound) {
		return nil, nil, err
	}

	if existedFile != nil && existedFile.Version >= album.Version {
		return nil, nil, ErrAlbumFileVersionExisted
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, nil, err
	}

	af := albumFiles.NewAlbumFile(id, albumFiles.StatusNew, album.ID, album.Version)

	if err := u.AlbumFiles.Save(ctx, af); err != nil {
		return nil, nil, err
	}

	return af, nil, err
}
