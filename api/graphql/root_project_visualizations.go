package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"log"
)

func (r *rootResolver) ProjectVisualizations() ProjectVisualizationsResolver {
	return &projectVisualizationsResolver{r}
}

type projectVisualizationsResolver struct {
	*rootResolver
}

func (r *projectVisualizationsResolver) List(
	ctx context.Context,
	obj *ProjectVisualizations,
	filter ProjectVisualizationsListFilter,
	limit int,
	offset int,
) (ProjectVisualizationsListResult, error) {
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		log.Printf("can't resolve project visualizations: %s", errors.New("unknown project"))

		return serverError()
	} else {
		res, err := r.useCases.GetVisualizations(
			ctx,
			project.ID,
			filter.ToSpec(),
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve project (id=%s) visualizations: %s", project.ID, err)

			return ServerError{Message: "internal server error"}, nil
		}

		return ProjectVisualizationsList{Items: visualizationsToGraphQL(res)}, nil
	}
}

func (f ProjectVisualizationsListFilter) ToSpec() visualizations.Spec {
	var (
		s visualizations.Spec
	)

	if f.RoomID != nil && len(f.RoomID.Eq) > 0 {
		s = visualizations.And(s, visualizations.RoomIDIn(f.RoomID.Eq...))
	}

	if f.Status != nil && len(f.Status.Eq) > 0 {
		s = visualizations.And(s, visualizations.StatusIn(mapf(f.Status.Eq, visualizationStatusFromGraphQL)...))
	}

	return s
}

func visualizationStatusFromGraphQL(status VisualizationStatus) visualizations.VisualizationStatus {
	switch status {
	case VisualizationStatusApproved:
		return visualizations.VisualizationStatusApproved
	case VisualizationStatusDeleted:
		return visualizations.VisualizationStatusDeleted
	default:
		return visualizations.VisualizationStatusUnknown
	}
}

func mapf[T any, V any](vals []T, mapfn func(T) V) []V {
	var (
		res = make([]V, len(vals))
	)

	for i, val := range vals {
		res[i] = mapfn(val)
	}

	return res
}

func (r *projectVisualizationsResolver) Total(
	ctx context.Context,
	obj *ProjectVisualizations,
	filter ProjectVisualizationsListFilter,
) (ProjectVisualizationsTotalResult, error) {
	return notImplementedYetError()
}
