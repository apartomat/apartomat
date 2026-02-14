package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/api/crm/graphql/dataloaders"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/albumfiles"
	"github.com/apartomat/apartomat/internal/store/albums"
)

func (r *rootResolver) Album() AlbumResolver { return &albumResolver{r} }

type albumResolver struct {
	*rootResolver
}

func (r *albumResolver) Pages(ctx context.Context, obj *Album) (AlbumPagesResult, error) {
	return obj.Pages, nil
}

func (r *albumResolver) Project(ctx context.Context, obj *Album) (AlbumProjectResult, error) {
	if p, ok := obj.Project.(Project); ok {
		project, err := r.crm.GetProject(ctx, p.ID)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			if errors.Is(err, crm.ErrNotFound) {
				return notFound()
			}

			slog.ErrorContext(ctx, "can't resolve project", slog.String("project", p.ID), slog.Any("err", err))

			return serverError()
		}

		return projectToGraphQL(project), nil
	}

	slog.ErrorContext(ctx, "can't resolve album project: obj.Project is not a Project")

	return serverError()
}

func (r *albumResolver) File(ctx context.Context, obj *Album) (AlbumRecentFileResult, error) {
	albumFile, file, err := r.crm.GetAlbumRecentFile(ctx, obj.ID)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, albums.ErrAlbumNotFound) {
			return notFound()
		}

		if errors.Is(err, albumfiles.ErrAlbumFileNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't resolve recent album file", slog.String("project", obj.ID), slog.Any("err", err))

		return serverError()
	}

	var (
		res = albumFileToGraphQL(albumFile)
	)

	if file != nil {
		res.File = fileToGraphQL(file)
	}

	return res, nil
}

func (r *albumResolver) Cover(ctx context.Context, obj *Album) (AlbumCoverResult, error) {
	if pages, ok := obj.Pages.(*AlbumPages); ok {
		if len(pages.Items) == 0 {
			return notFound()
		}

		if p, ok := pages.Items[0].(*AlbumPageCover); ok {
			switch c := p.Cover.(type) {
			case *CoverUploaded:
				if f, ok := c.File.(File); ok {
					file, err := dataloaders.FromContext(ctx).Files.Load(ctx, f.ID)
					if err != nil {
						slog.ErrorContext(ctx, "can't load cover uploaded file", slog.String("file", f.ID), slog.Any("err", err))
						return serverError()
					}

					return fileToGraphQL(file), nil
				}
			case *SplitCover:
				if f, ok := c.Image.(File); ok {
					file, err := dataloaders.FromContext(ctx).Files.Load(ctx, f.ID)
					if err != nil {
						slog.ErrorContext(ctx, "can't load split cover image file", slog.String("file", f.ID), slog.Any("err", err))
						return serverError()
					}

					return fileToGraphQL(file), nil
				}
			}
		}

		return notFound()
	}

	slog.ErrorContext(ctx, "obj.Pages is not a *AlbumPages")

	return serverError()
}

func albumFileToGraphQL(file *albumfiles.AlbumFile) *AlbumFile {
	return &AlbumFile{
		ID:                  file.ID,
		Status:              albumFileStatusToGraphQL(file.Status),
		Version:             file.Version,
		GeneratingStartedAt: file.GeneratingStartedAt,
		GeneratingDoneAt:    file.GeneratingDoneAt,
	}
}

func albumFileStatusToGraphQL(status albumfiles.Status) AlbumFileStatus {
	switch status {
	case albumfiles.StatusNew:
		return AlbumFileStatusNew
	case albumfiles.StatusInProgress:
		return AlbumFileStatusGeneratingInProgress
	case albumfiles.StatusDone:
		return AlbumFileStatusGeneratingDone
	}

	return AlbumFileStatusNew
}
