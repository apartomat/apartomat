package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
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

		r.logger.Error("can't update house", zap.Error(err))

		return serverError()
	}

	return HouseUpdated{House: houseToGraphQL(house)}, nil
}
