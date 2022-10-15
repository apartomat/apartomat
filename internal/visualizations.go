package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/pkg/errors"
	"path/filepath"
)

func (u *Apartomat) UploadVisualization(
	ctx context.Context,
	projectID string,
	upload Upload,
	roomID *string,
) (*store.ProjectFile, *Visualization, error) {
	file, err := u.UploadFile(ctx, projectID, upload, store.ProjectFileTypeVisualization)
	if err != nil {
		return nil, nil, err
	}

	id, err := NewNanoID()
	if err != nil {
		return nil, nil, err
	}

	vis := &Visualization{
		ID:            id,
		ProjectID:     projectID,
		ProjectFileID: file.ID,
		RoomID:        roomID,
	}

	vis, err = u.Visualizations.Save(ctx, vis)
	if err != nil {
		return nil, nil, err
	}

	return file, vis, err
}

type VisualizationWithFile struct {
	Visualization *Visualization
	File          *store.ProjectFile
}

func (u *Apartomat) UploadVisualizations(
	ctx context.Context,
	projectID string,
	files []Upload,
	roomID *string,
) ([]*VisualizationWithFile, error) {
	var (
		res = make([]*VisualizationWithFile, 0, len(files))
	)

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

	for _, file := range files {
		fileID, err := NewNanoID()
		if err != nil {
			return nil, err
		}

		path := fmt.Sprintf("p/%s/%s%s", project.ID, fileID, filepath.Ext(file.Name))

		url, err := u.Uploader.Upload(ctx, file.Data, file.Size, path, file.MimeType)
		if err != nil {
			return nil, err
		}

		f := &store.ProjectFile{
			ID:        fileID,
			ProjectID: projectID,
			Name:      file.Name,
			URL:       url,
			Type:      store.ProjectFileTypeVisualization,
			MimeType:  file.MimeType,
		}

		f, err = u.ProjectFiles.Save(ctx, f)
		if err != nil {
			return nil, err
		}

		visID, err := NewNanoID()
		if err != nil {
			return nil, err
		}

		vis := &Visualization{
			ID:            visID,
			ProjectID:     projectID,
			ProjectFileID: f.ID,
			RoomID:        roomID,
		}

		vis, err = u.Visualizations.Save(ctx, vis)
		if err != nil {
			return nil, err
		}

		res = append(res, &VisualizationWithFile{Visualization: vis, File: f})
	}

	return res, err
}

func (u *Apartomat) GetVisualizations(
	ctx context.Context,
	projectID string,
	limit, offset int,
) ([]*Visualization, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanGetVisualizations(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't get project (id=%s) visualizations", project.ID)
	}

	return u.Visualizations.List(ctx, ProjectIDIn(project.ID), limit, offset)
}

func (u *Apartomat) CanGetVisualizations(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}
