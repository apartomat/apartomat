package project_page

import (
	"context"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/public_sites"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
)

func (u *Service) GetVisualizations(ctx context.Context, publicSiteID string, limit, offset int) ([]*Visualization, error) {
	site, err := u.PublicSites.Get(ctx, public_sites.IDIn(publicSiteID))
	if err != nil {
		return nil, err
	}

	res, err := u.Visualizations.List(ctx, ProjectIDIn(site.ProjectID), SortRoomAscPositionAsc, limit, offset)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Service) GetVisualizationFile(ctx context.Context, fileID string) (*files.File, error) {
	return u.Files.Get(ctx, files.And(files.IDIn(fileID), files.FileTypeIn(files.FileTypeVisualization)))
}
