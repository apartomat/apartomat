package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/public_sites"
	"log/slog"
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

		if errors.Is(err, public_sites.ErrPublicSiteNotFound) {
			return notFound()
		}

		if errors.Is(err, public_sites.ErrPublicSiteIsNotPublic) {
			return ProjectIsAlreadyNotPublic{}, nil
		}

		slog.ErrorContext(ctx, "can't make project not public", slog.Any("err", err))

		return serverError()
	}

	return ProjectMadeNotPublic{PublicSite: publicSiteToGraphQL(ps)}, nil
}
