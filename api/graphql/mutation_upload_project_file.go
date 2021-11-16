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
			Type:     toProjectFileType(input.Type),
			MimeType: input.File.ContentType,
			Data:     input.File.File,
		},
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrAlreadyExists) {
			return AlreadyExists{}, nil
		}

		log.Printf("can't upload file to project (id=%d): %s", input.ProjectID, err)

		return serverError()
	}

	return ProjectFileUploaded{File: projectFileToGraphQL(pf)}, nil
}
