package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
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
			return forbidden()
		}

		slog.ErrorContext(ctx, "can't update contact", slog.Any("err", err))

		return serverError()
	}

	return ContactUpdated{Contact: contactToGraphQL(contact)}, nil
}
