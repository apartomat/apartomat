package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/projectpage"
)

func (r *mutationResolver) MakeProjectNotPublic(
	ctx context.Context,
	projectID string,
) (MakeProjectNotPublicResult, error) {
	ps, err := r.useCases.MakeProjectNotPublic(ctx, projectID)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, projectpage.ErrProjectPageNotFound) {
			return notFound()
		}

		if errors.Is(err, projectpage.ErrProjectPageIsNotPublic) {
			return ProjectIsAlreadyNotPublic{}, nil
		}

		slog.ErrorContext(ctx, "can't make project not public", slog.Any("err", err))

		return serverError()
	}

	return ProjectMadeNotPublic{ProjectPage: projectPageToGraphQL(ps)}, nil
}
