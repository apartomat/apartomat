package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) ChangeProjectDates(
	ctx context.Context,
	projectID string,
	input ChangeProjectDatesInput,
) (ChangeProjectDatesResult, error) {
	project, err := r.useCases.ChangeProjectDates(
		ctx,
		projectID,
		input.StartAt,
		input.EndAt,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't change project dates: %s", err)

		return ServerError{Message: "can't change project dates"}, nil
	}

	return ProjectDatesChanged{Project: projectToGraphQL(project)}, nil
}
