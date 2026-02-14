package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *queryResolver) SplitCoverFormDefaults(ctx context.Context, albumID string) (SplitCoverFormDefaultsResult, error) {
	defaults, err := r.crm.GetSplitCoverFormDefaults(ctx, albumID)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(
			ctx,
			"can't get split cover form defaults",
			slog.String("album", albumID),
			slog.Any("err", err),
		)

		return serverError()
	}

	return &SplitCoverFormDefaults{
		City:   defaults.City,
		Year:   defaults.Year,
		WithQR: defaults.WithQr,
	}, nil
}
