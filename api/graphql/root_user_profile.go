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
	if obj.DefaultWorkspace == nil {
		return nil, nil
	}

	if obj.DefaultWorkspace.ID == "" {
		return obj.DefaultWorkspace, nil
	}

	w, err := r.useCases.GetWorkspace(ctx, obj.DefaultWorkspace.ID)
	if err != nil {
		log.Printf("can't get default workspace: %s", err)

		return nil, errors.New("internal server error")
	}

	return workspaceToGraphQL(w), nil
}
