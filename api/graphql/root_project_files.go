package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/files"
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
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		log.Printf("can't resolve project files: %s", errors.New("unknown project"))

		return serverError()
	} else {
		files, err := r.useCases.GetFiles(
			ctx,
			project.ID,
			toProjectFileTypes(filter.Type),
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve project (id=%s) files: %s", project.ID, err)

			return ServerError{Message: "internal server error"}, nil
		}

		return ProjectFilesList{Items: projectFilesToGraphQL(files)}, nil
	}
}

func projectFilesToGraphQL(projects []*files.File) []*File {
	result := make([]*File, 0, len(projects))

	for _, u := range projects {
		result = append(result, fileToGraphQL(u))
	}

	return result
}

func fileToGraphQL(file *files.File) *File {
	return &File{
		ID:       file.ID,
		Name:     file.Name,
		URL:      file.URL,
		Type:     projectFileTypeToGraphQL(file.Type),
		MimeType: file.MimeType,
	}
}

func (r *projectFilesResolver) Total(
	ctx context.Context,
	obj *ProjectFiles,
	filter ProjectFilesListFilter,
) (ProjectFilesTotalResult, error) {
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		log.Printf("can't resolve project files: %s", errors.New("unknown project"))

		return nil, errors.New("server error: can't resolver project files")
	} else {
		tot, err := r.useCases.CountFiles(
			ctx,
			project.ID,
			toProjectFileTypes(filter.Type),
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve project (id=%s) files: %s", project.ID, err)

			return nil, errors.New("server error: can't resolver project files")
		}

		return ProjectFilesTotal{Total: tot}, nil
	}
}

func projectFileTypeToGraphQL(t files.FileType) FileType {
	switch t {
	case files.FileTypeVisualization:
		return FileTypeVisualization
	case files.FileTypeNone:
		return FileTypeNone
	default:
		return FileTypeNone
	}
}

func toProjectFileType(t FileType) files.FileType {
	switch t {
	case FileTypeVisualization:
		return files.FileTypeVisualization
	case FileTypeNone:
		return files.FileTypeNone
	default:
		return files.FileTypeNone
	}
}

func toProjectFileTypes(l []FileType) []files.FileType {
	res := make([]files.FileType, len(l))

	for i, t := range l {
		res[i] = toProjectFileType(t)
	}

	return res
}
