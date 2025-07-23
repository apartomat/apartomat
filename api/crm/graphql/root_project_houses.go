package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/houses"
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
		slog.ErrorContext(ctx, "can't resolve project houses", slog.String("err", "unknown project"))

		return serverError()
	} else {
		items, err := r.crm.GetHouses(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(
				ctx,
				"can't resolve project houses",
				slog.String("project", project.ID),
				slog.Any("err", err),
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
