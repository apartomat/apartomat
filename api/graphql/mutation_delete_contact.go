package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
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

		log.Printf("can't delete contact: %s", err)

		return ServerError{Message: "can't delete contact"}, nil
	}

	return ContactDeleted{Contact: contactToGraphQL(contact)}, nil
}
