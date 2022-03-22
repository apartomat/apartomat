package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) UpdateContact(
	ctx context.Context,
	contactID string,
	data UpdateContactInput,
) (UpdateContactResult, error) {
	contact, err := r.useCases.UpdateContact(
		ctx,
		contactID,
		apartomat.UpdateContactParams{
			FullName: data.FullName,
			Details:  contactsDetailsFromGraphQL(data.Details),
		},
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't update contact: %s", err)

		return ServerError{Message: "can't update contact"}, nil
	}

	return ContactUpdated{Contact: contactToGraphQL(contact)}, nil
}
