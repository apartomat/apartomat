package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
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

		r.logger.Error("can't delete contact", zap.Error(err))

		return serverError()
	}

	return ContactDeleted{Contact: contactToGraphQL(contact)}, nil
}
