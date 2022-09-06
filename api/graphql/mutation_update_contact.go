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
	input UpdateContactInput,
) (UpdateContactResult, error) {
	contact, err := r.useCases.UpdateContact(
		ctx,
		contactID,
		apartomat.UpdateContactParams{
			FullName: input.FullName,
			Details:  contactsDetailsFromGraphQL(input.Details),
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
