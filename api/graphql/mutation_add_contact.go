package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"log"
)

func (r *mutationResolver) AddContact(
	ctx context.Context,
	projectID string,
	input AddContactInput,
) (AddContactResult, error) {
	contact, err := r.useCases.AddContact(
		ctx,
		projectID,
		apartomat.AddContactParams{
			FullName: input.FullName,
			Details:  contactsDetailsFromGraphQL(input.Details),
		},
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't add contact: %s", err)

		return ServerError{Message: "can't add contact"}, nil
	}

	return ContactAdded{Contact: projectContactToGraphQL(contact)}, nil
}

func contactsDetailsFromGraphQL(input []*AddContactDetailsInput) []contacts.Details {
	details := make([]contacts.Details, len(input))

	for i, d := range input {
		details[i] = contacts.Details{
			Type:  projectContactTypeFromGraphQL(d.Type),
			Value: d.Value,
		}
	}

	return details
}
