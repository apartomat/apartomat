package graphql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) UploadVisualizations(
	ctx context.Context,
	projectID string,
	files []*graphql.Upload,
	roomID *string,
) (UploadVisualizationsResult, error) {
	var (
		uploads = make([]apartomat.Upload, len(files))
	)

	for i, f := range files {
		uploads[i] = apartomat.Upload{
			Name:     f.Filename,
			MimeType: f.ContentType,
			Data:     f.File,
			Size:     f.Size,
		}
	}

	res, err := r.useCases.UploadVisualizations(ctx, projectID, uploads, roomID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		r.logger.Error("can't upload file to project", zap.String("project", projectID), zap.Error(err))

		return nil, err
	}

	if len(res) != len(files) {
		return SomeVisualizationsUploaded{Visualizations: visualizationsWithFilesToGraphQL(res)}, nil
	}

	return VisualizationsUploaded{Visualizations: visualizationsWithFilesToGraphQL(res)}, nil

}

func visualizationsWithFilesToGraphQL(visualizations []*apartomat.VisualizationWithFile) []*Visualization {
	var (
		res = make([]*Visualization, len(visualizations))
	)

	for i, vis := range visualizations {
		res[i] = visualizationToGraphQL(vis.Visualization, vis.File)
	}

	return res
}
