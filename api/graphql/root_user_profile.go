package graphql

import (
	"context"
	"errors"
	"log"
)

type userProfileResolver struct {
	*rootResolver
}

func (r *rootResolver) UserProfile() UserProfileResolver { return &userProfileResolver{r} }

func (r *userProfileResolver) DefaultWorkspace(ctx context.Context, obj *UserProfile) (*Workspace, error) {
	w, err := r.useCases.GetDefaultWorkspace(ctx, obj.ID)
	if err != nil {
		log.Printf("can't get default workspace: %s", err)
		return nil, errors.New("internal server error")
	}

	return workspaceToGraphQL(w), nil
}
