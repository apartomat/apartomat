package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/houses"
	"log"
)

func (r *mutationResolver) AddHouse(
	ctx context.Context,
	projectID string,
	input AddHouseInput,
) (AddHouseResult, error) {
	contact, err := r.useCases.AddHouse(
		ctx,
		projectID,
		input.City,
		input.Address,
		input.HousingComplex,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't add house: %s", err)

		return ServerError{Message: "can't add house"}, nil
	}

	return HouseAdded{House: houseToGraphQL(contact)}, nil
}

func houseToGraphQL(house *houses.House) *House {
	return &House{
		ID:             house.ID,
		City:           house.City,
		Address:        house.Address,
		HousingComplex: house.HousingComplex,
		CreatedAt:      house.CreatedAt,
		ModifiedAt:     house.ModifiedAt,
	}
}
