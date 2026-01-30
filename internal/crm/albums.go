package crm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/apartomat/apartomat/internal/store/albumfiles"
	. "github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/projectpages"
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
				page, n := album.AddVisualizationPageWithID(vis, MustGenerateNanoID())
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

	album.AddUploadedCoverPageWithID(file.ID, MustGenerateNanoID())

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return file, nil
}

type SplitCoverForAddToAlbum struct {
	Title     string
	Subtitle  *string
	ImgFileID string
	WithQR    bool
	City      *string
	Year      *int
}

type SplitCoverFormDefaults struct {
	City   *string
	Year   int
	WithQr bool
}

func (u *CRM) GetSplitCoverFormDefaults(ctx context.Context, albumID string) (*SplitCoverFormDefaults, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		if errors.Is(err, ErrAlbumNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	if ok, err := u.Acl.CanAddPageToAlbum(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get split cover form defaults for album (id=%s): %w", album.ID, ErrForbidden)
	}

	year := time.Now().Year()

	var city *string
	if houses, err := u.GetHouses(ctx, album.ProjectID, 1, 0); err == nil && len(houses) > 0 && houses[0].City != "" {
		city = &houses[0].City
	}

	withQr := false
	if page, err := u.GetProjectPage(ctx, album.ProjectID); err == nil {
		withQr = page.Is(projectpages.Public())
	}

	return &SplitCoverFormDefaults{
		City:   city,
		Year:   year,
		WithQr: withQr,
	}, nil
}

func (u *CRM) AddSplitCoverToAlbum(
	ctx context.Context,
	albumID string,
	cover SplitCoverForAddToAlbum,
) (*Album, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanAddPageToAlbum(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add split cover to album (id=%s): %w", album.ID, ErrForbidden)
	}

	file, err := u.Files.Get(ctx, files.And(files.IDIn(cover.ImgFileID), files.ProjectIDIn(album.ProjectID)))
	if err != nil {
		if errors.Is(err, files.ErrFileNotFound) {
			return nil, fmt.Errorf("image file (id=%s) doesn't belong to album project (id=%s): %w", cover.ImgFileID, album.ProjectID, ErrForbidden)

		}

		return nil, err
	}

	album.AddSplitCoverPageWithID(
		MustGenerateNanoID(),
		cover.Title,
		cover.Subtitle,
		file.ID,
		cover.WithQR,
		cover.City,
		cover.Year,
	)

	album.UpVersion()

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return album, nil
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

func (u *CRM) GetAlbumRecentFile(ctx context.Context, albumID string) (*albumfiles.AlbumFile, *files.File, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, nil, err
	}

	if ok, err := u.Acl.CanGetAlbumFile(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, nil, err
	} else if !ok {
		return nil, nil, fmt.Errorf("can't get album (id=%s) recent file: %w", album.ID, ErrForbidden)
	}

	albumFile, err := u.AlbumFiles.GetLastVersion(ctx, albumfiles.And(albumfiles.AlbumIDIn(albumID)))
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

func (u *CRM) StartGenerateAlbumFile(ctx context.Context, albumID string) (*albumfiles.AlbumFile, *files.File, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, nil, err
	}

	if ok, err := u.Acl.CanGetAlbumFile(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, nil, err
	} else if !ok {
		return nil, nil, fmt.Errorf("can't generate album (id=%s) file: %w", album.ID, ErrForbidden)
	}

	existedFile, err := u.AlbumFiles.GetLastVersion(ctx, albumfiles.And(albumfiles.AlbumIDIn(albumID)))
	if err != nil && !errors.Is(err, albumfiles.ErrAlbumFileNotFound) {
		return nil, nil, err
	}

	if existedFile != nil && existedFile.Version >= album.Version {
		return nil, nil, ErrAlbumFileVersionExisted
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, nil, err
	}

	af := albumfiles.NewAlbumFile(id, albumfiles.StatusNew, album.ID, album.Version)

	if err := u.AlbumFiles.Save(ctx, af); err != nil {
		return nil, nil, err
	}

	return af, nil, err
}

func (u *CRM) DeleteAlbumPage(
	ctx context.Context,
	albumID string,
	pageNumber int,
) (*AlbumPage, error) {
	album, err := u.Albums.Get(ctx, IDIn(albumID))
	if err != nil {
		return nil, err
	}

	if len(album.Pages) < pageNumber+1 {
		return nil, fmt.Errorf("there is no page (number=%d) in album (id=%s) : %w", pageNumber, album.ID, ErrForbidden)
	}

	var (
		deleted = album.Pages[pageNumber]
	)

	album.Pages = append(album.Pages[:pageNumber], album.Pages[pageNumber+1:]...)

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return &deleted, nil
}
