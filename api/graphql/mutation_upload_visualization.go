package graphql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"go.uber.org/zap"
)

func (r *mutationResolver) UploadVisualization(
	ctx context.Context,
	projectID string,
	file graphql.Upload,
	roomID *string,
) (UploadVisualizationResult, error) {
	uploaded, vis, err := r.useCases.UploadVisualization(
		ctx,
		projectID,
		apartomat.Upload{
			Name:     file.Filename,
			MimeType: file.ContentType,
			Data:     file.File,
			Size:     file.Size,
		},
		roomID,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		r.logger.Error("can't upload file to project", zap.String("project", projectID), zap.Error(err))

		return serverError()
	}

	return VisualizationUploaded{Visualization: visualizationToGraphQL(vis, uploaded)}, nil
}

func visualizationToGraphQL(vis *visualizations.Visualization, file *files.File) *Visualization {
	if vis == nil {
		return nil
	}

	res := &Visualization{
		ID:          vis.ID,
		Name:        vis.Name,
		Description: vis.Description,
		Version:     vis.Version,
		Status:      visualizationStatusToGraphQL(vis.Status),
		CreatedAt:   vis.CreatedAt,
		ModifiedAt:  vis.ModifiedAt,
		File: &File{
			ID: vis.FileID,
		},
	}

	if file != nil {
		res.File = fileToGraphQL(file)
	}

	if vis.RoomID != nil {
		res.Room = &Room{
			ID: *vis.RoomID,
		}
	}

	return res
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

func visualizationStatusToGraphQL(status visualizations.VisualizationStatus) VisualizationStatus {
	switch status {
	case visualizations.VisualizationStatusApproved:
		return VisualizationStatusApproved
	case visualizations.VisualizationStatusDeleted:
		return VisualizationStatusDeleted
	default:
		return VisualizationStatusUnknown
	}
}
