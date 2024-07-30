package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
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

func (u *Apartomat) GetFiles(
	ctx context.Context,
	projectID string,
	fileType []FileType,
	limit, offset int,
) ([]*File, error) {
	if ok, err := u.Acl.CanGetFilesOfProjectID(ctx, auth.UserFromCtx(ctx), projectID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) files: %w", projectID, ErrForbidden)
	}

	p, err := u.Files.List(
		ctx,
		And(ProjectIDIn(projectID), FileTypeIn(fileType...)),
		SortDefault,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *Apartomat) CountFiles(
	ctx context.Context,
	projectID string,
	fileType []FileType,
) (int, error) {
	if ok, err := u.Acl.CanCountFilesOfProjectID(ctx, auth.UserFromCtx(ctx), projectID); err != nil {
		return 0, err
	} else if !ok {
		return 0, fmt.Errorf("can't get project (id=%s) files: %w", projectID, ErrForbidden)
	}

	return u.Files.Count(
		ctx,
		And(ProjectIDIn(projectID), FileTypeIn(fileType...)),
	)
}

func (u *Apartomat) UploadFile(
	ctx context.Context,
	projectID string,
	upload Upload,
	fileType FileType,
) (*File, error) {
	project, err := u.Projects.Get(ctx, projects.IDIn(projectID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUploadFile(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) files: %w", project.ID, ErrForbidden)
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("p/%s/%s%s", project.ID, id, filepath.Ext(upload.Name))

	url, err := u.Uploader.Upload(ctx, upload.Data, upload.Size, path, upload.MimeType)
	if err != nil {
		return nil, err
	}

	f := NewFile(id, upload.Name, url, fileType, upload.MimeType, projectID)

	if err := u.Files.Save(ctx, f); err != nil {
		return nil, err
	}

	return f, nil
}

func (u *Apartomat) GetFile(
	ctx context.Context,
	id string,
) (*File, error) {
	f, err := u.Files.Get(ctx, IDIn(id))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanGetFilesOfProjectID(ctx, auth.UserFromCtx(ctx), f.ProjectID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get file (id=%s): %w", id, ErrForbidden)
	}

	return f, nil
}
