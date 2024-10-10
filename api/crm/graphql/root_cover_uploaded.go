package graphql

import (
	"context"
	"log/slog"

	"github.com/apartomat/apartomat/api/crm/graphql/dataloaders"
)

func (r *rootResolver) CoverUploaded() CoverUploadedResolver {
	return &coverUploadedResolver{r}
}

type coverUploadedResolver struct {
	*rootResolver
}

func (r *coverUploadedResolver) File(ctx context.Context, obj *CoverUploaded) (CoverFileResult, error) {
	if obj.File == nil {
		return notFound()
	}

	if f, ok := obj.File.(File); ok {
		file, err := dataloaders.FromContext(ctx).Files.Load(ctx, f.ID)
		if err != nil {
			return nil, err
		}

		return fileToGraphQL(file), nil
	}

	slog.ErrorContext(ctx, "can't resolver uploaded cover file: obj.File is not a *File")

	return serverError()
}
