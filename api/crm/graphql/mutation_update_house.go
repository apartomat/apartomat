package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
)

func (r *mutationResolver) UpdateHouse(
	ctx context.Context,
	houseID string,
	input UpdateHouseInput,
) (UpdateHouseResult, error) {
	house, err := r.useCases.UpdateHouse(
		ctx,
		houseID,
		input.City,
		input.Address,
		input.HousingComplex,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't update house", slog.Any("err", err))

		return serverError()
	}

	return HouseUpdated{House: houseToGraphQL(house)}, nil
}
