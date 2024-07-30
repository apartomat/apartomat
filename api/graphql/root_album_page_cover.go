package graphql

import (
	"context"
)

func (r *rootResolver) AlbumPageCover() AlbumPageCoverResolver {
	return &albumPageCoverResolver{r}
}

type albumPageCoverResolver struct {
	*rootResolver
}

func (r *albumPageCoverResolver) Cover(
	ctx context.Context,
	obj *AlbumPageCover,
) (AlbumPageCoverResult, error) {
	var (
		cov interface{} = obj.Cover
	)

	switch c := cov.(type) {
	case *CoverUploaded:
		return c, nil
	case *Cover:
		return notImplementedYetError()
	default:
		r.logger.Error("unknown obj type")
		return serverError()
	}
}
