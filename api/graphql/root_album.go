package graphql

import (
	"context"
	"errors"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/apartomat/apartomat/internal/store/albums"
	"go.uber.org/zap"
)

func (r *rootResolver) Album() AlbumResolver { return &albumResolver{r} }

type albumResolver struct {
	*rootResolver
}

func (r *albumResolver) Pages(ctx context.Context, obj *Album) (AlbumPagesResult, error) {
	return obj.Pages, nil
}

func (r *albumResolver) Project(ctx context.Context, obj *Album) (AlbumProjectResult, error) {
	var (
		gp *Project
	)

	if pr, ok := obj.Project.(Project); ok {
		gp = &pr
	}

	if gp == nil {
		return serverError()
	}

	project, err := r.useCases.GetProject(ctx, gp.ID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		r.logger.Error("can't resolve project", zap.String("project", gp.ID), zap.Error(err))

		return nil, fmt.Errorf("can't resolve project (id=%s): %w", gp.ID, err)
	}

	return projectToGraphQL(project), nil
}

func (r *albumResolver) File(ctx context.Context, obj *Album) (AlbumRecentFileResult, error) {
	albumFile, file, err := r.useCases.GetAlbumRecentFile(ctx, obj.ID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, albums.ErrAlbumNotFound) {
			return notFound()
		}

		if errors.Is(err, albumFiles.ErrAlbumFileNotFound) {
			return notFound()
		}

		r.logger.Error("can't resolve recent album file", zap.String("project", obj.ID), zap.Error(err))

		return nil, fmt.Errorf("can't resolve album (id=%s) recent file: %w", obj.ID, err)
	}

	var (
		res = albumFileToGraphQL(albumFile)
	)

	if file != nil {
		res.File = fileToGraphQL(file)
	}

	return res, nil
}

func albumFileToGraphQL(file *albumFiles.AlbumFile) *AlbumFile {
	return &AlbumFile{
		ID:                  file.ID,
		Status:              albumFileStatusToGraphQL(file.Status),
		Version:             file.Version,
		GeneratingStartedAt: file.GeneratingStartedAt,
		GeneratingDoneAt:    file.GeneratingDoneAt,
	}
}

func albumFileStatusToGraphQL(status albumFiles.Status) AlbumFileStatus {
	switch status {
	case albumFiles.StatusNew:
		return AlbumFileStatusNew
	case albumFiles.StatusInProgress:
		return AlbumFileStatusGeneratingInProgress
	case albumFiles.StatusDone:
		return AlbumFileStatusGeneratingDone
	}

	return AlbumFileStatusNew
}
