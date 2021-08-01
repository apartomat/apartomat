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

func (u *UploadProjectFile) Do(
	ctx context.Context,
	projectID int,
	name, contentType string,
	file io.Reader,
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

	path := fmt.Sprintf("p/%d/%s", project.ID, name)

	url, err := u.uploader.Upload(ctx, file, path, contentType)
	if err != nil {
		return nil, err
	}

	f := &store.ProjectFile{
		ProjectID: projectID,
		Name:      name,
		URL:       url,
		Type:      contentType,
	}

	f, err = u.files.Save(ctx, f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
