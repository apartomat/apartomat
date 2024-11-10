package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) UploadFile(
	ctx context.Context,
	input UploadFileInput,
) (UploadFileResult, error) {
	pf, err := r.useCases.UploadFile(
		ctx,
		input.ProjectID,
		crm.Upload{
			Name:     input.Data.Filename,
			MimeType: input.Data.ContentType,
			Data:     input.Data.File,
			Size:     input.Data.Size,
		},
		toFileType(input.Type),
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(ctx, "can't upload file to project", slog.String("project", input.ProjectID), slog.Any("err", err))

		return serverError()
	}

	return FileUploaded{File: fileToGraphQL(pf)}, nil
}
