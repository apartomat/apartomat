package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/files"
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
		slog.ErrorContext(ctx, "can't resolve project files", slog.String("err", "unknown project"))

		return serverError()
	} else {
		items, err := r.crm.GetFiles(
			ctx,
			project.ID,
			toFileTypes(filter.Type),
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(ctx, "can't resolve project files", slog.String("project", project.ID), slog.Any("err", err))

			return serverError()
		}

		return ProjectFilesList{Items: projectFilesToGraphQL(items)}, nil
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
		Type:     fileTypeToGraphQL(file.Type),
		MimeType: file.MimeType,
		Size:     int(file.Size),
	}
}

func (r *projectFilesResolver) Total(
	ctx context.Context,
	obj *ProjectFiles,
	filter ProjectFilesListFilter,
) (ProjectFilesTotalResult, error) {
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		slog.ErrorContext(ctx, "can't resolve project files", slog.String("err", "unknown project"))

		return serverError()
	} else {
		tot, err := r.crm.CountFiles(
			ctx,
			project.ID,
			toFileTypes(filter.Type),
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(ctx, "can't resolve project files", slog.String("project", project.ID), slog.Any("err", err))

			return serverError()
		}

		return ProjectFilesTotal{Total: tot}, nil
	}
}

func fileTypeToGraphQL(t files.FileType) FileType {
	switch t {
	case files.FileTypeNone:
		return FileTypeNone
	case files.FileTypeVisualization:
		return FileTypeVisualization
	case files.FileTypeAlbum:
		return FileTypeAlbum
	default:
		return FileTypeNone
	}
}

func toFileType(t FileType) files.FileType {
	switch t {
	case FileTypeVisualization:
		return files.FileTypeVisualization
	case FileTypeNone:
		return files.FileTypeNone
	default:
		return files.FileTypeNone
	}
}

func toFileTypes(l []FileType) []files.FileType {
	res := make([]files.FileType, len(l))

	for i, t := range l {
		res[i] = toFileType(t)
	}

	return res
}
