package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) MakeProjectNotPublic(
	ctx context.Context,
	projectID string,
) (MakeProjectNotPublicResult, error) {
	ps, err := r.useCases.MakeProjectNotPublic(ctx, projectID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		r.logger.Error("can't make project not public", zap.Error(err))

		return serverError()
	}

	return ProjectMadeNotPublic{PublicSite: publicSiteToGraphQL(ps)}, nil
}
