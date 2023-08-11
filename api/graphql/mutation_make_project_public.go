package graphql

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/store/public_sites"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) MakeProjectPublic(
	ctx context.Context,
	projectID string,
) (MakeProjectPublicResult, error) {
	ps, err := r.useCases.MakeProjectPublic(ctx, projectID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, public_sites.ErrPublicSiteNotFound) {
			return notFound()
		}

		if errors.Is(err, public_sites.ErrPublicSiteIsPublic) {
			return ProjectIsAlreadyPublic{}, nil
		}

		r.logger.Error("can't make project public", zap.Error(err))

		return serverError()
	}

	return ProjectMadePublic{PublicSite: publicSiteToGraphQL(ps)}, nil
}
