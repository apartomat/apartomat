package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/projects"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
	"path/filepath"
)

func (u *Apartomat) UploadVisualization(
	ctx context.Context,
	projectID string,
	upload Upload,
	roomID *string,
) (*files.File, *Visualization, error) {
	file, err := u.UploadFile(ctx, projectID, upload, files.FileTypeVisualization)
	if err != nil {
		return nil, nil, err
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, nil, err
	}

	vis := NewVisualization(id, projectID, file.ID, roomID)

	if err = u.Visualizations.Save(ctx, vis); err != nil {
		return nil, nil, err
	}

	return file, vis, nil
}

type VisualizationWithFile struct {
	Visualization *Visualization
	File          *files.File
}

func (u *Apartomat) UploadVisualizations(
	ctx context.Context,
	projectID string,
	uploads []Upload,
	roomID *string,
) ([]*VisualizationWithFile, error) {
	var (
		res = make([]*VisualizationWithFile, 0, len(uploads))
	)

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

	if ok, err := u.CanUploadFile(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) files: %w", project.ID, ErrForbidden)
	}

	for _, file := range uploads {
		fileID, err := GenerateNanoID()
		if err != nil {
			return nil, err
		}

		path := fmt.Sprintf("p/%s/%s%s", project.ID, fileID, filepath.Ext(file.Name))

		url, err := u.Uploader.Upload(ctx, file.Data, file.Size, path, file.MimeType)
		if err != nil {
			return nil, err
		}

		f := files.NewFile(fileID, file.Name, url, files.FileTypeVisualization, file.MimeType, projectID)

		if err := u.Files.Save(ctx, f); err != nil {
			return nil, err
		}

		id, err := GenerateNanoID()
		if err != nil {
			return nil, err
		}

		vis := NewVisualization(id, projectID, f.ID, roomID)

		if err := u.Visualizations.Save(ctx, vis); err != nil {
			return nil, err
		}

		res = append(res, &VisualizationWithFile{Visualization: vis, File: f})
	}

	return res, err
}

func (u *Apartomat) GetVisualizations(
	ctx context.Context,
	projectID string,
	spec Spec,
	limit, offset int,
) ([]*Visualization, error) {
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

	if ok, err := u.CanGetVisualizations(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) visualizations: %w", project.ID, ErrForbidden)
	}

	return u.Visualizations.List(ctx, And(spec, ProjectIDIn(project.ID)), limit, offset)
}

func (u *Apartomat) CanGetVisualizations(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}

func (u *Apartomat) DeleteVisualizations(
	ctx context.Context,
	id []string,
) ([]*Visualization, error) {
	vis, err := u.Visualizations.List(ctx, IDIn(id...), len(id), 0)
	if err != nil {
		return nil, err
	}

	for _, v := range vis {
		if ok, err := u.CanDeleteVisualization(ctx, auth.UserFromCtx(ctx), v); err != nil {
			return nil, err
		} else if !ok {
			return nil, fmt.Errorf("can't delete visualization (id=%s): %w", v.ID, ErrForbidden)
		}

		v.Delete()
	}

	if err := u.Visualizations.Save(ctx, vis...); err != nil {
		return nil, err
	}

	return vis, err
}

func (u *Apartomat) CanDeleteVisualization(ctx context.Context, subj *auth.UserCtx, obj *Visualization) (bool, error) {
	var (
		project *projects.Project
	)

	prjs, err := u.Projects.List(ctx, projects.IDIn(obj.ProjectID), 1, 0)
	if err != nil {
		return false, err
	}

	if len(prjs) > 0 {
		project = prjs[0]
	}

	return u.isProjectUser(ctx, subj, project)
}

func (u *Apartomat) GetVisualization(
	ctx context.Context,
	id string,
) (*Visualization, error) {
	res, err := u.Visualizations.List(ctx, IDIn(id), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("visualization (id=%s): %w", id, ErrNotFound)
	}

	var (
		visualization = res[0]
	)

	prjs, err := u.Projects.List(ctx, projects.IDIn(visualization.ProjectID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", visualization.ProjectID, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanGetVisualizations(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) visualizations: %w", project.ID, ErrForbidden)
	}

	return visualization, nil
}
