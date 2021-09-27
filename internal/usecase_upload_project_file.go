package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
	"io"
)

type UploadProjectFile struct {
	projects store.ProjectStore
	files    store.ProjectFileStore
	acl      *Acl
	uploader ImageUploader
}

func NewUploadProjectFile(
	projects store.ProjectStore,
	files store.ProjectFileStore,
	acl *Acl,
	uploader ImageUploader,
) *UploadProjectFile {
	return &UploadProjectFile{projects: projects, files: files, acl: acl, uploader: uploader}
}

type Upload struct {
	Name     string
	MimeType string
	Data     io.Reader
}

func (u *UploadProjectFile) Do(
	ctx context.Context,
	projectID int,
	upload Upload,
) (*store.ProjectFile, error) {
	projects, err := u.projects.List(ctx, store.ProjectStoreQuery{ID: expr.IntEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	project := projects[0]

	if !u.acl.CanUploadProjectFile(ctx, UserFromCtx(ctx), project) {
		return nil, errors.Wrapf(ErrForbidden, "can't upload file to project (id=%d)", project.ID)
	}

	path := fmt.Sprintf("p/%d/%s", project.ID, upload.Name)

	url, err := u.uploader.Upload(ctx, upload.Data, path, upload.MimeType)
	if err != nil {
		return nil, err
	}

	f := &store.ProjectFile{
		ProjectID: projectID,
		Name:      upload.Name,
		URL:       url,
		Type:      "NONE",
		MimeType:  upload.MimeType,
	}

	f, err = u.files.Save(ctx, f)
	if err != nil {
		if errors.Is(err, store.ErrAlreadyExists) {
			return nil, errors.Wrapf(ErrAlreadyExists, "%s", upload.Name)
		}

		return nil, err
	}

	return f, nil
}
