package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) UploadProjectFile(
	ctx context.Context,
	input UploadProjectFileInput,
) (UploadProjectFileResult, error) {
	pf, err := r.useCases.UploadFile(
		ctx,
		input.ProjectID,
		apartomat.Upload{
			Name:     input.File.Filename,
			MimeType: input.File.ContentType,
			Data:     input.File.File,
			Size:     input.File.Size,
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

	return ProjectFileUploaded{File: projectFileToGraphQL(pf)}, nil
}
