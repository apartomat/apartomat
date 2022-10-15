package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/pkg/errors"
	"log"
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
			return Forbidden{}, nil
		}

		log.Printf("can't upload file to project (id=%d): %s", projectID, err)

		return nil, err
	}

	return VisualizationUploaded{Visualization: visualizationToGraphQL(vis, uploaded)}, nil
}

func visualizationToGraphQL(vis *visualizations.Visualization, file *store.ProjectFile) *Visualization {
	if vis == nil {
		return nil
	}

	res := &Visualization{
		ID:          vis.ID,
		Name:        vis.Name,
		Description: vis.Description,
		Version:     vis.Version,
		CreatedAt:   vis.CreatedAt,
		ModifiedAt:  vis.ModifiedAt,
		File: &ProjectFile{
			ID: vis.ProjectFileID,
		},
	}

	if file != nil {
		res.File = projectFileToGraphQL(file)
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
