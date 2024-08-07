package graphql

import (
	"context"
	"github.com/apartomat/apartomat/internal/dataloaders"
)

func (r *rootResolver) Cover() CoverResolver {
	return &coverResolver{r}
}

type coverResolver struct {
	*rootResolver
}

func (r *coverResolver) File(ctx context.Context, obj *Cover) (CoverFileResult, error) {
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

	r.logger.Error("can't resolver cover file: obj.File is not a *File")

	return serverError()
}
