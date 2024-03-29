package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/houses"
	"go.uber.org/zap"
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
		r.logger.Error("can't resolve project houses", zap.Error(errors.New("unknown project")))

		return serverError()
	} else {
		items, err := r.useCases.GetHouses(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return forbidden()
			}

			r.logger.Error(
				"can't resolve project houses",
				zap.String("project", project.ID),
				zap.Error(err),
			)

			return serverError()
		}

		return ProjectHousesList{Items: projectHousesToGraphQL(items)}, nil
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
