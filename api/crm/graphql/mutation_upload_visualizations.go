package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) UploadVisualizations(
	ctx context.Context,
	projectID string,
	files []*graphql.Upload,
	roomID *string,
) (UploadVisualizationsResult, error) {
	var (
		uploads = make([]crm.Upload, len(files))
	)

	for i, f := range files {
		uploads[i] = crm.Upload{
			Name:     f.Filename,
			MimeType: f.ContentType,
			Data:     f.File,
			Size:     f.Size,
		}
	}

	res, err := r.crm.UploadVisualizations(ctx, projectID, uploads, roomID)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(ctx, "can't upload file to project", slog.String("project", projectID), slog.Any("err", err))

		return nil, err
	}

	if len(res) != len(files) {
		return SomeVisualizationsUploaded{Visualizations: visualizationsWithFilesToGraphQL(res)}, nil
	}

	return VisualizationsUploaded{Visualizations: visualizationsWithFilesToGraphQL(res)}, nil

}

func visualizationsWithFilesToGraphQL(visualizations []*crm.VisualizationWithFile) []*Visualization {
	var (
		res = make([]*Visualization, len(visualizations))
	)

	for i, vis := range visualizations {
		res[i] = visualizationToGraphQL(vis.Visualization, vis.File)
	}

	return res
}
