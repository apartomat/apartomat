package graphql

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/store/files"
)

func (r *rootResolver) Visualization() VisualizationResolver {
	return &visualizationResolver{r}
}

type visualizationResolver struct {
	*rootResolver
}

func (r *visualizationResolver) File(ctx context.Context, obj *Visualization) (VisualizationFileResult, error) {
	if f, ok := obj.File.(*VisualizationFile); ok {
		f, err := r.projectPage.GetVisualizationFile(ctx, f.ID)
		if err != nil {
			return ServerError{"failed resolve visualization file"}, nil
		}

		return fileToGraphQL(f), nil
	}

	return ServerError{"file is not VisualizationFile"}, nil
}

func fileToGraphQL(file *files.File) *VisualizationFile {
	return &VisualizationFile{
		ID:       file.ID,
		URL:      file.URL,
		MimeType: file.MimeType,
	}
}

func (r *visualizationResolver) Room(ctx context.Context, obj *Visualization) (VisualizationRoomResult, error) {
	return nil, errors.New("not implemented yet")
}
