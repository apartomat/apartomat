package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/contacts"
)

func (r *mutationResolver) AddContact(
	ctx context.Context,
	projectID string,
	input AddContactInput,
) (AddContactResult, error) {
	contact, err := r.crm.AddContact(
		ctx,
		projectID,
		crm.AddContactParams{
			FullName: input.FullName,
			Details:  contactsDetailsFromGraphQL(input.Details),
		},
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(ctx, "can't add contact", slog.Any("err", err))

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
