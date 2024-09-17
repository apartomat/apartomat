package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/houses"
	"go.uber.org/zap"
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
			return forbidden()
		}

		r.logger.Error("can't add house", zap.Error(err))

		return serverError()
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
