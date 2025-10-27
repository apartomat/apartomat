package graphql

import (
	"context"
	"log/slog"

	"github.com/apartomat/apartomat/api/crm/graphql/dataloaders"
)

type splitCoverResolver struct {
	r *rootResolver
}

func (r *rootResolver) SplitCover() SplitCoverResolver {
	return &splitCoverResolver{r}
}

func (s *splitCoverResolver) Image(ctx context.Context, obj *SplitCover) (SplitCoverImageFileResult, error) {
	switch imageFile := obj.Image.(type) {
	case File:
		f, err := dataloaders.FromContext(ctx).Files.Load(ctx, imageFile.ID)
		if err != nil {
			slog.ErrorContext(ctx, "can't resolve split cover image file", slog.Any("err", err))
			return serverError()
		}

		return fileToGraphQL(f), nil
	case nil:
		return notFound()
	default:
		slog.ErrorContext(ctx, "unknown obj.Image type")

		return serverError()
	}
}
