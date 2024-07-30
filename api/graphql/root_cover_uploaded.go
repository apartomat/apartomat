package graphql

import (
	"context"
	"github.com/apartomat/apartomat/internal/dataloaders"
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

	r.logger.Error("can't resolver uploaded cover file: obj.File is not a *File")

	return serverError()
}
