package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (DeleteContactResult, error) {
	contact, err := r.crm.DeleteContact(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't delete contact", slog.Any("err", err))

		return serverError()
	}

	return ContactDeleted{Contact: contactToGraphQL(contact)}, nil
}
