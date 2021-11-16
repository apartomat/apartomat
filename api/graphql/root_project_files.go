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

func (r *projectFilesResolver) List(
	ctx context.Context,
	obj *ProjectFiles,
	filter ProjectFilesListFilter,
	limit int,
	offset int,
) (ProjectFilesListResult, error) {
	files, err := r.useCases.GetProjectFiles(
		ctx,
		obj.Project.ID,
		apartomat.GetProjectFilesFilter{
			Type: store.ProjectFileTypeExpr{Eq: toProjectFileTypes(filter.Type)},
		},
		limit,
		offset,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't resolve project (id=%d) files: %s", obj.Project.ID, err)

		return ServerError{Message: "internal server error"}, nil
	}

	return ProjectFilesList{Items: projectFilesToGraphQL(files)}, nil
}

func projectFilesToGraphQL(projects []*store.ProjectFile) []*ProjectFile {
	result := make([]*ProjectFile, 0, len(projects))

	for _, u := range projects {
		result = append(result, projectFileToGraphQL(u))
	}

	return result
}

func projectFileToGraphQL(file *store.ProjectFile) *ProjectFile {
	return &ProjectFile{
		ID:       file.ID,
		Name:     file.Name,
		URL:      file.URL,
		Type:     projectFileTypeToGraphQL(file.Type),
		MimeType: file.MimeType,
	}
}

func (r *projectFilesResolver) Total(ctx context.Context, obj *ProjectFiles) (ProjectFilesTotalResult, error) {
	return ProjectFilesTotal{Total: 99}, nil // todo
}

func projectFileTypeToGraphQL(t store.ProjectFileType) ProjectFileType {
	switch t {
	case store.ProjectFileTypeVisualization:
		return ProjectFileTypeVisualization
	case store.ProjectFileTypeNone:
		return ProjectFileTypeNone
	default:
		return ProjectFileTypeNone
	}
}

func toProjectFileType(t ProjectFileType) store.ProjectFileType {
	switch t {
	case ProjectFileTypeVisualization:
		return store.ProjectFileTypeVisualization
	case ProjectFileTypeNone:
		return store.ProjectFileTypeNone
	default:
		return store.ProjectFileTypeNone
	}
}

func toProjectFileTypes(l []ProjectFileType) []store.ProjectFileType {
	res := make([]store.ProjectFileType, len(l))

	for i, t := range l {
		res[i] = toProjectFileType(t)
	}

	return res
}