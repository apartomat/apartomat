package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"go.uber.org/zap"
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
			return forbidden()
		}

		r.logger.Error("can't add contact", zap.Error(err))

		return serverError()
	}

	return ContactAdded{Contact: contactToGraphQL(contact)}, nil
}

func contactsDetailsFromGraphQL(input []*ContactDetailsInput) []contacts.Details {
	details := make([]contacts.Details, len(input))

	for i, d := range input {
		details[i] = contacts.Details{
			Type:  contactTypeFromGraphQL(d.Type),
			Value: d.Value,
		}
	}

	return details
}
