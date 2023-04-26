package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
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

		r.logger.Error("can't update contact", zap.Error(err))

		return serverError()
	}

	return ContactUpdated{Contact: contactToGraphQL(contact)}, nil
}
