package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) UpdateContact(
	ctx context.Context,
	contactID string,
	input UpdateContactInput,
) (UpdateContactResult, error) {
	contact, err := r.crm.UpdateContact(
		ctx,
		contactID,
		crm.UpdateContactParams{
			FullName: input.FullName,
			Details:  contactsDetailsFromGraphQL(input.Details),
		},
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(ctx, "can't update contact", slog.Any("err", err))

		return serverError()
	}

	return ContactUpdated{Contact: contactToGraphQL(contact)}, nil
}
