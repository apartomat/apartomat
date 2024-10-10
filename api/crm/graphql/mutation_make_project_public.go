package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/public_sites"
)

func (r *mutationResolver) MakeProjectPublic(
	ctx context.Context,
	projectID string,
) (MakeProjectPublicResult, error) {
	ps, err := r.useCases.MakeProjectPublic(ctx, projectID)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, public_sites.ErrPublicSiteNotFound) {
			return notFound()
		}

		if errors.Is(err, public_sites.ErrPublicSiteIsPublic) {
			return ProjectIsAlreadyPublic{}, nil
		}

		slog.ErrorContext(ctx, "can't make project public", slog.Any("err", err))

		return serverError()
	}

	return ProjectMadePublic{PublicSite: publicSiteToGraphQL(ps)}, nil
}
