package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) UploadFile(
	ctx context.Context,
	input UploadFileInput,
) (UploadFileResult, error) {
	pf, err := r.useCases.UploadFile(
		ctx,
		input.ProjectID,
		apartomat.Upload{
			Name:     input.Data.Filename,
			MimeType: input.Data.ContentType,
			Data:     input.Data.File,
			Size:     input.Data.Size,
		},
		toProjectFileType(input.Type),
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't upload file to project (id=%s): %s", input.ProjectID, err)

		return serverError()
	}

	return FileUploaded{File: fileToGraphQL(pf)}, nil
}
