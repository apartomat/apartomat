package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
)

func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (DeleteContactResult, error) {
	contact, err := r.useCases.DeleteContact(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't delete contact", slog.Any("err", err))

		return serverError()
	}

	return ContactDeleted{Contact: contactToGraphQL(contact)}, nil
}
