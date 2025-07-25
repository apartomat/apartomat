package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/visualizations"
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
		slog.ErrorContext(ctx, "can't resolve project visualizations", slog.String("err", "unknown project"))

		return serverError()
	} else {
		res, err := r.crm.GetVisualizations(
			ctx,
			project.ID,
			filter.ToSpec(),
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(
				ctx,
				"can't resolve project visualizations",
				slog.String("project", project.ID),
				slog.Any("err", err),
			)

			return serverError()
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
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		slog.ErrorContext(
			ctx,
			"can't resolve project visualizations total",
			slog.Any("err", errors.New("unknown project")),
		)

		return serverError()
	} else {
		tot, err := r.crm.CountVisualizations(
			ctx,
			project.ID,
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(
				ctx,
				"can't resolve project visualizations total",
				slog.String("project", project.ID),
				slog.Any("err", err),
			)

			return serverError()
		}

		return ProjectVisualizationsTotal{Total: tot}, nil
	}
}
