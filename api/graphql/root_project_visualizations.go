package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
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
		visualizations, err := r.useCases.GetVisualizations(
			ctx,
			project.ID,
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

		return ProjectVisualizationsList{Items: visualizationsToGraphQL(visualizations)}, nil
	}
}

func (r *projectVisualizationsResolver) Total(
	ctx context.Context,
	obj *ProjectVisualizations,
	filter ProjectVisualizationsListFilter,
) (ProjectVisualizationsTotalResult, error) {
	return notImplementedYetError()
}
