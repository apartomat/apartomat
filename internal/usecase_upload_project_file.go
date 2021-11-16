package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
	"io"
)

type Upload struct {
	Name     string
	Type     store.ProjectFileType
	MimeType string
	Data     io.Reader
}

func (u *Apartomat) UploadFile(
	ctx context.Context,
	projectID int,
	upload Upload,
) (*store.ProjectFile, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.IntEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	project := projects[0]

	if !u.CanUploadProjectFile(ctx, UserFromCtx(ctx), project) {
		return nil, errors.Wrapf(ErrForbidden, "can't upload file to project (id=%d)", project.ID)
	}

	path := fmt.Sprintf("p/%d/%s", project.ID, upload.Name)

	url, err := u.Uploader.Upload(ctx, upload.Data, path, upload.MimeType)
	if err != nil {
		return nil, err
	}

	f := &store.ProjectFile{
		ProjectID: projectID,
		Name:      upload.Name,
		URL:       url,
		Type:      upload.Type,
		MimeType:  upload.MimeType,
	}

	f, err = u.ProjectFiles.Save(ctx, f)
	if err != nil {
		if errors.Is(err, store.ErrAlreadyExists) {
			return nil, errors.Wrapf(ErrAlreadyExists, "%s", upload.Name)
		}

		return nil, err
	}

	return f, nil
}

func (u *Apartomat) CanUploadProjectFile(ctx context.Context, subj *UserCtx, obj *store.Project) bool {
	return true
}
