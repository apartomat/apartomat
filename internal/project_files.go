package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"path/filepath"
)

type GetProjectFilesFilter struct {
	Type store.ProjectFileTypeExpr
}

func (u *Apartomat) GetProjectFiles(
	ctx context.Context,
	projectID string,
	filter GetProjectFilesFilter,
	limit, offset int,
) ([]*store.ProjectFile, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanGetProjectFiles(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) files: %w", project.ID, ErrForbidden)
	}

	p, err := u.ProjectFiles.List(
		ctx,
		store.ProjectFileStoreQuery{
			ProjectID: expr.StrEq(projectID),
			Type:      filter.Type,
			Limit:     limit,
			Offset:    offset,
		},
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *Apartomat) CanGetProjectFiles(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}

func (u *Apartomat) CountProjectFiles(
	ctx context.Context,
	projectID string,
	filter GetProjectFilesFilter,
) (int, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID)})
	if err != nil {
		return 0, err
	}

	if len(projects) == 0 {
		return 0, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanCountProjectFiles(ctx, UserFromCtx(ctx), project); err != nil {
		return 0, err
	} else if !ok {
		return 0, fmt.Errorf("can't get project (id=%s) files: %w", project.ID, ErrForbidden)
	}

	return u.ProjectFiles.Count(
		ctx,
		store.ProjectFileStoreQuery{
			ProjectID: expr.StrEq(projectID),
			Type:      filter.Type,
		},
	)
}

func (u *Apartomat) CanCountProjectFiles(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}

func (u *Apartomat) UploadFile(
	ctx context.Context,
	projectID string,
	upload Upload,
	fileType store.ProjectFileType,
) (*store.ProjectFile, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	project := projects[0]

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

	f := &store.ProjectFile{
		ID:        id,
		ProjectID: projectID,
		Name:      upload.Name,
		URL:       url,
		Type:      fileType,
		MimeType:  upload.MimeType,
	}

	f, err = u.ProjectFiles.Save(ctx, f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (u *Apartomat) CanUploadProjectFile(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}
