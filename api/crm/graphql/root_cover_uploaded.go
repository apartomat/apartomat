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
			slog.ErrorContext(ctx, "can't resolve uploaded cover file", slog.String("file", f.ID), slog.Any("err", err))
			return serverError()
		}

		return fileToGraphQL(file), nil
	}

	slog.ErrorContext(ctx, "can't resolve uploaded cover file: obj.File is not a File")

	return serverError()
}
