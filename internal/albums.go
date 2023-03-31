package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
	. "github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
)

func (u *Apartomat) CreateAlbum(
	ctx context.Context,
	projectID string,
	name string,
) (*Album, error) {
	prjs, err := u.Projects.List(ctx, projects.IDIn(projectID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanCreateAlbum(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't create album in project (id=%s): %w", project.ID, ErrForbidden)
	}

	id, err := NewNanoID()
	if err != nil {
		return nil, err
	}

	album := NewAlbum(id, name, project.ID)

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return album, nil
}

func (u *Apartomat) CanCreateAlbum(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(obj.WorkspaceID),
			workspace_users.UserIDIn(subj.ID),
		),
		1,
		0,
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}

func (u *Apartomat) GetAlbums(
	ctx context.Context,
	projectID string,
	limit, offset int,
) ([]*Album, error) {
	prjs, err := u.Projects.List(ctx, projects.IDIn(projectID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanGetAlbums(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) albums: %w", project.ID, ErrForbidden)
	}

	var (
		spec = ProjectIDIn(projectID)
	)

	p, err := u.Albums.List(ctx, spec, limit, offset)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *Apartomat) CanGetAlbums(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}

func (u *Apartomat) GetAlbum(
	ctx context.Context,
	id string,
) (*Album, error) {

	albums, err := u.Albums.List(ctx, IDIn(id), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(albums) == 0 {
		return nil, fmt.Errorf("album (id=%s): %w", id, ErrNotFound)
	}

	var (
		album = albums[0]
	)

	prjs, err := u.Projects.List(ctx, projects.IDIn(album.ProjectID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", album.ProjectID, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanGetAlbums(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) album (id=%s): %w", project.ID, album.ID, ErrForbidden)
	}

	return album, nil
}

func (u *Apartomat) CountAlbums(
	ctx context.Context,
	projectID string,
) (int, error) {
	prjs, err := u.Projects.List(ctx, projects.IDIn(projectID), 1, 0)
	if err != nil {
		return 0, err
	}

	if len(prjs) == 0 {
		return 0, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanCountAlbums(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return 0, err
	} else if !ok {
		return 0, fmt.Errorf("can't get project (id=%s) albums: %w", project.ID, ErrForbidden)
	}

	var (
		spec = ProjectIDIn(projectID)
	)

	return u.Albums.Count(ctx, spec)
}

func (u *Apartomat) CanCountAlbums(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}

func (u *Apartomat) DeleteAlbum(ctx context.Context, id string) (*Album, error) {
	albums, err := u.Albums.List(ctx, IDIn(id), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(albums) == 0 {
		return nil, fmt.Errorf("album (id=%s): %w", id, ErrNotFound)
	}

	var (
		album = albums[0]
	)

	if ok, err := u.CanDeleteAlbum(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't delete album (id=%s): %w", album.ID, ErrForbidden)
	}

	err = u.Albums.Delete(ctx, album)

	return album, err
}

func (u *Apartomat) CanDeleteAlbum(ctx context.Context, subj *auth.UserCtx, obj *Album) (bool, error) {
	prjs, err := u.Projects.List(ctx, projects.IDIn(obj.ProjectID), 1, 0)
	if err != nil {
		return false, err
	}

	if len(prjs) == 0 {
		return false, nil
	}

	var (
		project = prjs[0]
	)

	return u.isProjectUser(ctx, subj, project)
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
	albums, err := u.Albums.List(ctx, IDIn(albumID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(albums) == 0 {
		return nil, fmt.Errorf("album (id=%s): %w", albumID, ErrNotFound)
	}

	var (
		album = albums[0]
	)

	if ok, err := u.CanAddPageToAlbum(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add visualization to album (id=%s): %w", album.ID, ErrForbidden)
	}

	list, err := u.Visualizations.List(ctx, visualizations.IDIn(visualizationID...), len(visualizationID), 0)
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

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Apartomat) CanAddPageToAlbum(ctx context.Context, subj *auth.UserCtx, obj *Album) (bool, error) {
	prjs, err := u.Projects.List(ctx, projects.IDIn(obj.ProjectID), 1, 0)
	if err != nil {
		return false, err
	}

	if len(prjs) == 0 {
		return false, nil
	}

	var (
		project = prjs[0]
	)

	return u.isProjectUser(ctx, subj, project)
}

func (u *Apartomat) ChangeAlbumPageSize(
	ctx context.Context,
	albumID string,
	size PageSize,
) (*Album, error) {
	list, err := u.Albums.List(ctx, IDIn(albumID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("album (id=%s): %w", albumID, ErrNotFound)
	}

	var (
		album = list[0]
	)

	if ok, err := u.CanChangeAlbumSettings(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't change album (id=%s) settings: %w", album.ID, ErrForbidden)
	}

	album.ChangePageSize(size)

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
	list, err := u.Albums.List(ctx, IDIn(albumID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("album (id=%s): %w", albumID, ErrNotFound)
	}

	var (
		album = list[0]
	)

	if ok, err := u.CanChangeAlbumSettings(ctx, auth.UserFromCtx(ctx), album); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't change album (id=%s) settings: %w", album.ID, ErrForbidden)
	}

	album.ChangePageOrientation(orientation)

	if err := u.Albums.Save(ctx, album); err != nil {
		return nil, err
	}

	return album, nil
}

func (u *Apartomat) CanChangeAlbumSettings(ctx context.Context, subj *auth.UserCtx, obj *Album) (bool, error) {
	prjs, err := u.Projects.List(ctx, projects.IDIn(obj.ProjectID), 1, 0)
	if err != nil {
		return false, err
	}

	if len(prjs) == 0 {
		return false, nil
	}

	var (
		project = prjs[0]
	)

	return u.isProjectUser(ctx, subj, project)
}
