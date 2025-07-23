package project_page

import (
	"context"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/projectpages"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
)

func (u *Service) GetVisualizations(ctx context.Context, projectPageID string, limit, offset int) ([]*Visualization, error) {
	page, err := u.ProjectPages.Get(ctx, projectpages.IDIn(projectPageID))
	if err != nil {
		return nil, err
	}

	res, err := u.Visualizations.List(ctx, ProjectIDIn(page.ProjectID), SortRoomAscPositionAsc, limit, offset)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Service) GetVisualizationFile(ctx context.Context, fileID string) (*files.File, error) {
	return u.Files.Get(ctx, files.And(files.IDIn(fileID), files.FileTypeIn(files.FileTypeVisualization)))
}
