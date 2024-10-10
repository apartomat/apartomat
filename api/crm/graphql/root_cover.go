package graphql

import (
	"context"
	"log/slog"

	"github.com/apartomat/apartomat/api/crm/graphql/dataloaders"
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

	slog.ErrorContext(ctx, "can't resolver cover file: obj.File is not a *File")

	return serverError()
}
