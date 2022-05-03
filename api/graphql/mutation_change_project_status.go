package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
	"time"
)

func (r *mutationResolver) ChangeProjectStatus(
	ctx context.Context,
	projectID string,
	status ProjectStatus,
) (ChangeProjectStatusResult, error) {
	time.Sleep(1 * time.Second)
	project, err := r.useCases.ChangeProjectStatus(
		ctx,
		projectID,
		toProjectStatus(status),
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't change project status: %s", err)

		return ServerError{Message: "can't change project status"}, nil
	}

	return ProjectStatusChanged{Project: projectToGraphQL(project)}, nil
}
