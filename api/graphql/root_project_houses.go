package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/houses"
	"log"
)

func (r *rootResolver) ProjectHouses() ProjectHousesResolver {
	return &projectHousesResolver{r}
}

type projectHousesResolver struct {
	*rootResolver
}

func (r *projectHousesResolver) List(
	ctx context.Context,
	obj *ProjectHouses,
	filter ProjectHousesFilter,
	limit int,
	offset int,
) (ProjectHousesListResult, error) {
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		log.Printf("can't resolve project houses: %s", errors.New("unknown project"))

		return serverError()
	} else {
		houses, err := r.useCases.GetHouses(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve project (id=%d) houses: %s", project.ID, err)

			return ServerError{Message: "internal server error"}, nil
		}

		return ProjectHousesList{Items: projectHousesToGraphQL(houses)}, nil
	}
}

func (r *projectHousesResolver) Total(
	ctx context.Context,
	obj *ProjectHouses,
	filter ProjectHousesFilter,
) (ProjectHousesTotalResult, error) {
	return notImplementedYetError()
}

func projectHousesToGraphQL(houses []*houses.House) []*House {
	result := make([]*House, 0, len(houses))

	for _, item := range houses {
		result = append(result, projectHouseToGraphQL(item))
	}

	return result
}

func projectHouseToGraphQL(house *houses.House) *House {
	return &House{
		ID:             house.ID,
		City:           house.City,
		Address:        house.Address,
		HousingComplex: house.HousingComplex,
		CreatedAt:      house.CreatedAt,
		ModifiedAt:     house.ModifiedAt,
	}
}
