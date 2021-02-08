package graphql

import (
	"context"
	"github.com/ztsu/apartomat/internal/pkg/gravatar"
)

func (r *rootResolver) WorkspaceUser() WorkspaceUserResolver {
	return &workspaceUserResolver{r}
}

type workspaceUserResolver struct {
	*rootResolver
}

func (r *workspaceUserResolver) Profile(ctx context.Context, obj *WorkspaceUser) (*WorkspaceUserProfile, error) {
	user, err := r.useCases.GetWorkspaceUserProfile.Do(ctx, 0, obj.ID)
	if err != nil {
		return nil, err
	}

	return &WorkspaceUserProfile{
		ID:       obj.ID,
		Email:    user.Email,
		Gravatar: &Gravatar{URL: gravatar.Url(user.Email)},
	}, nil
}
