package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) UpdateHouse(
	ctx context.Context,
	houseID string,
	input UpdateHouseInput,
) (UpdateHouseResult, error) {
	contact, err := r.useCases.UpdateHouse(
		ctx,
		houseID,
		input.City,
		input.Address,
		input.HousingComplex,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't update house: %s", err)

		return ServerError{Message: "can't update house"}, nil
	}

	return HouseUpdated{House: houseToGraphQL(contact)}, nil
}
