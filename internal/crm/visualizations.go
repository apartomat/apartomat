package crm

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/projects"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
	"path/filepath"
)

func (u *CRM) UploadVisualization(
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

	position, err := u.Visualizations.GetMaxSortingPosition(ctx, NotDeletedByProjectAndRoom(projectID, roomID))
	if err != nil {
		return nil, nil, fmt.Errorf("can't get max soring position: %w", err)
	}

	vis := NewVisualization(id, "", "", 0, position+1, projectID, file.ID, roomID)

	if err = u.Visualizations.Save(ctx, vis); err != nil {
		return nil, nil, err
	}

	return file, vis, nil
}

type VisualizationWithFile struct {
	Visualization *Visualization
	File          *files.File
}

func (u *CRM) UploadVisualizations(
	ctx context.Context,
	projectID string,
	uploads []Upload,
	roomID *string,
) ([]*VisualizationWithFile, error) {
	var (
		res = make([]*VisualizationWithFile, 0, len(uploads))
	)

	project, err := u.Projects.Get(ctx, projects.IDIn(projectID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUploadFile(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) files: %w", project.ID, ErrForbidden)
	}

	position, err := u.Visualizations.GetMaxSortingPosition(ctx, NotDeletedByProjectAndRoom(projectID, roomID))
	if err != nil {
		return nil, fmt.Errorf("can't get max soring position: %w", err)
	}

	for _, file := range uploads {
		position++

		fileID, err := GenerateNanoID()
		if err != nil {
			return nil, err
		}

		path := fmt.Sprintf("p/%s/%s%s", project.ID, fileID, filepath.Ext(file.Name))

		url, err := u.Uploader.Upload(ctx, file.Data, file.Size, path, file.MimeType)
		if err != nil {
			return nil, err
		}

		f := files.NewFile(fileID, file.Name, url, files.FileTypeVisualization, file.MimeType, file.Size, projectID)

		if err := u.Files.Save(ctx, f); err != nil {
			return nil, err
		}

		id, err := GenerateNanoID()
		if err != nil {
			return nil, err
		}

		vis := NewVisualization(id, "", "", 0, position, projectID, f.ID, roomID)

		if err := u.Visualizations.Save(ctx, vis); err != nil {
			return nil, err
		}

		res = append(res, &VisualizationWithFile{Visualization: vis, File: f})
	}

	return res, err
}

func (u *CRM) GetVisualizations(
	ctx context.Context,
	projectID string,
	spec Spec,
	limit, offset int,
) ([]*Visualization, error) {
	if ok, err := u.Acl.CanGetVisualizationsOfProjectID(ctx, auth.UserFromCtx(ctx), projectID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) visualizations: %w", projectID, ErrForbidden)
	}

	return u.Visualizations.List(ctx, And(spec, ProjectIDIn(projectID)), SortRoomAscPositionAsc, limit, offset)
}

func (u *CRM) CountVisualizations(
	ctx context.Context,
	projectID string,
) (int, error) {
	if ok, err := u.Acl.CanCountVisualizationsOfProjectID(ctx, auth.UserFromCtx(ctx), projectID); err != nil {
		return 0, err
	} else if !ok {
		return 0, fmt.Errorf("can't count project (id=%s) visualizations: %w", projectID, ErrForbidden)
	}

	var (
		spec = ProjectIDIn(projectID)
	)

	return u.Visualizations.Count(ctx, spec)
}

func (u *CRM) DeleteVisualizations(
	ctx context.Context,
	id []string,
) ([]*Visualization, error) {
	vis, err := u.Visualizations.List(ctx, IDIn(id...), SortDefault, len(id), 0)
	if err != nil {
		return nil, err
	}

	for _, v := range vis {
		if ok, err := u.Acl.CanDeleteVisualization(ctx, auth.UserFromCtx(ctx), v); err != nil {
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

func (u *CRM) GetVisualization(
	ctx context.Context,
	id string,
) (*Visualization, error) {
	visualization, err := u.Visualizations.Get(ctx, IDIn(id))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanGetVisualization(ctx, auth.UserFromCtx(ctx), visualization); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) visualizations: %w", visualization.ProjectID, ErrForbidden)
	}

	return visualization, nil
}
