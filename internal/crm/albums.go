package crm

import (
	"context"
	"errors"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files"
	. "github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/visualizations"
)

func (u *CRM) CreateAlbum(
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

func (u *CRM) GetAlbums(
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

func (u *CRM) GetAlbum(
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

func (u *CRM) CountAlbums(
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

func (u *CRM) DeleteAlbum(ctx context.Context, id string) (*Album, error) {
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

func (u *CRM) AddVisualizationsToAlbum(
	ctx context.Context,
	albumID string,
	visualizationID []string,
) ([]AlbumPageVisualization, int, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, 0, err
	}

	if ok, err := u.Acl.CanAddPageToAlbum(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, 0, err
	} else if !ok {
		return nil, 0, fmt.Errorf("can't add visualization to album (id=%s): %w", album.ID, ErrForbidden)
	}

	list, err := u.Visualizations.List(
		ctx,
		visualizations.IDIn(visualizationID...),
		visualizations.SortDefault,
		len(visualizationID),
		0,
	)
	if err != nil {
		return nil, 0, err
	}

	if len(list) != len(visualizationID) {
		return nil, 0, fmt.Errorf("visualization (id=%s): %w", visualizationID, ErrNotFound)
	}

	var (
		res = make([]AlbumPageVisualization, len(visualizationID))
		num *int
	)

visLoop:
	for i, id := range visualizationID {
		for _, vis := range list {
			if vis.ID == id {
				page, n := album.AddPageWithVisualization(vis)
				if num == nil {
					num = &n
				}

				res[i] = page

				continue visLoop
			}
		}
	}

	album.UpVersion()

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, 0, err
	}

	if num == nil {
		return res, 0, nil
	}

	return res, *num, nil
}

func (u *CRM) ChangeAlbumPageSize(
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

func (u *CRM) ChangeAlbumPageOrientation(
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

func (u *CRM) GetAlbumRecentFile(ctx context.Context, albumID string) (*albumFiles.AlbumFile, *files.File, error) {
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

func (u *CRM) StartGenerateAlbumFile(ctx context.Context, albumID string) (*albumFiles.AlbumFile, *files.File, error) {
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

func (u *CRM) UploadAlbumCover(
	ctx context.Context,
	albumID string,
	upload Upload,
) (*files.File, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, err
	}

	file, err := u.UploadFile(ctx, album.ProjectID, upload, files.FileTypeAlbumCover)
	if err != nil {
		return nil, err
	}

	album.Pages = append(
		[]AlbumPage{AlbumPageCoverUploaded{FileID: file.ID}},
		album.Pages...,
	)

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return file, nil
}
