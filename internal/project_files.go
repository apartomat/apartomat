package apartomat

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	. "github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/projects"
)

type Upload struct {
	Name     string
	MimeType string
	Data     io.Reader
	Size     int64
}

func (u *Apartomat) GetProjectFiles(
	ctx context.Context,
	projectID string,
	fileType []FileType,
	limit, offset int,
) ([]*File, error) {
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

	if ok, err := u.CanGetProjectFiles(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) files: %w", project.ID, ErrForbidden)
	}

	p, err := u.Files.List(
		ctx,
		And(ProjectIDIn(projectID), FileTypeIn(fileType...)),
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *Apartomat) CanGetProjectFiles(ctx context.Context, subj *UserCtx, obj *projects.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}

func (u *Apartomat) CountProjectFiles(
	ctx context.Context,
	projectID string,
	fileType []FileType,
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

	if ok, err := u.CanCountProjectFiles(ctx, UserFromCtx(ctx), project); err != nil {
		return 0, err
	} else if !ok {
		return 0, fmt.Errorf("can't get project (id=%s) files: %w", project.ID, ErrForbidden)
	}

	return u.Files.Count(
		ctx,
		And(ProjectIDIn(projectID), FileTypeIn(fileType...)),
	)
}

func (u *Apartomat) CanCountProjectFiles(ctx context.Context, subj *UserCtx, obj *projects.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}

func (u *Apartomat) UploadFile(
	ctx context.Context,
	projectID string,
	upload Upload,
	fileType FileType,
) (*File, error) {
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

	if ok, err := u.CanUploadProjectFile(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) files: %w", project.ID, ErrForbidden)
	}

	id, err := NewNanoID()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("p/%s/%s%s", project.ID, id, filepath.Ext(upload.Name))

	url, err := u.Uploader.Upload(ctx, upload.Data, upload.Size, path, upload.MimeType)
	if err != nil {
		return nil, err
	}

	f := New(id, upload.Name, url, fileType, upload.MimeType, projectID)

	if err := u.Files.Save(ctx, f); err != nil {
		return nil, err
	}

	return f, nil
}

func (u *Apartomat) CanUploadProjectFile(ctx context.Context, subj *UserCtx, obj *projects.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}
