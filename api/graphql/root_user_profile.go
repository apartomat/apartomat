package graphql

import (
	"context"
	"errors"
	"go.uber.org/zap"
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
		r.logger.Error("can't get default workspace", zap.Error(err))

		return nil, errors.New("internal server error")
	}

	return workspaceToGraphQL(w), nil
}
