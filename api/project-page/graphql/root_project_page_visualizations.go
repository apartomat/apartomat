package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/visualizations"
)

func (r *rootResolver) ProjectPageVisualizations() ProjectPageVisualizationsResolver {
	return &projectPageVisualizationsResolver{r}
}

type projectPageVisualizationsResolver struct {
	*rootResolver
}

func (r *projectPageVisualizationsResolver) List(
	ctx context.Context,
	obj *ProjectPageVisualizations,
	filter ProjectPageVisualizationsFilter,
	limit int,
	offset int,
) (ProjectPageVisualizationsListResult, error) {
	publicSite, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(ProjectPage)
	if !ok {
		return ServerError{Message: "unknown project page"}, nil
	}

	res, err := r.projectPage.GetVisualizations(ctx, publicSite.ID, limit, offset)
	if err != nil {
		return nil, err
	}

	return VisualizationsList{Items: visualizationsToGraphQL(res)}, nil
}

func visualizationsToGraphQL(visualizations []*visualizations.Visualization) []*Visualization {
	var (
		res = make([]*Visualization, len(visualizations))
	)

	for i, vis := range visualizations {
		res[i] = visualizationToGraphQL(vis, nil)
	}

	return res
}

func visualizationToGraphQL(vis *visualizations.Visualization, file *files.File) *Visualization {
	if vis == nil {
		return nil
	}

	res := &Visualization{
		ID:          vis.ID,
		Name:        vis.Name,
		Description: vis.Description,
		File: &VisualizationFile{
			ID: vis.FileID,
		},
	}

	if vis.RoomID != nil {
		res.Room = &Room{
			ID: *vis.RoomID,
		}
	}

	return res
}

func (r *projectPageVisualizationsResolver) Total(
	ctx context.Context,
	obj *ProjectPageVisualizations,
	filter ProjectPageVisualizationsFilter,
) (ProjectPageVisualizationsTotalResult, error) {
	return notImplementedYetError()
}
