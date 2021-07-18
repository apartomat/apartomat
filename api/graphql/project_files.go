package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"log"
)

func (r *rootResolver) ProjectFiles() ProjectFilesResolver {
	return &projectFilesResolver{r}
}

type projectFilesResolver struct {
	*rootResolver
}

func (r *projectFilesResolver) List(ctx context.Context, obj *ProjectFiles) (ProjectFilesListResult, error) {
	files, err := r.useCases.GetProjectFiles.Do(ctx, obj.Project.ID, 10, 0)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't resolve project (id=%d) files: %s", obj.Project.ID, err)

		return ServerError{Message: "internal server error"}, nil
	}

	return ProjectFilesList{Items: filesToGraphQLProjectFiles(files)}, nil
}

func filesToGraphQLProjectFiles(projects []*store.ProjectFile) []*ProjectFile {
	result := make([]*ProjectFile, 0, len(projects))

	for _, u := range projects {
		result = append(result, fileToGraphQLProjectFile(u))
	}

	return result
}

func fileToGraphQLProjectFile(file *store.ProjectFile) *ProjectFile {
	return &ProjectFile{
		ID:   file.ID,
		Name: file.Name,
		URL:  file.URL,
		Type: file.Type,
	}
}

func (r *projectFilesResolver) Total(ctx context.Context, obj *ProjectFiles) (ProjectFilesTotalResult, error) {
	return notImplementedYetError() // todo
}
