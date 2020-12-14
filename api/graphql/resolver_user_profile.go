package graphql

import (
	"context"
	"github.com/ztsu/apartomat/internal/store"
)

type userProfileResolver struct {
	*rootResolver
}

func (r *userProfileResolver) DefaultWorkspace(ctx context.Context, obj *UserProfile) (*Workspace, error) {
	w, err := r.useCases.GetDefaultWorkspace.Do(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return workspaceToGraphQL(w), nil
}

func workspaceToGraphQL(w *store.Workspace) *Workspace {
	return &Workspace{
		ID:   w.ID,
		Name: w.Name,
	}
}
