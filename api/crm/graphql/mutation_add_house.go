package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/houses"
)

func (r *mutationResolver) AddHouse(
	ctx context.Context,
	projectID string,
	input AddHouseInput,
) (AddHouseResult, error) {
	contact, err := r.crm.AddHouse(
		ctx,
		projectID,
		input.City,
		input.Address,
		input.HousingComplex,
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(ctx, "can't add house", slog.Any("err", err))

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
