package graphql

import (
	"context"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/pkg/gravatar"
	"github.com/apartomat/apartomat/internal/store"
	"log"
)

func (r *rootResolver) UserProfile() UserProfileResolver { return &userProfileResolver{r} }

type userProfileResolver struct {
	*rootResolver
}

func (r *userProfileResolver) DefaultWorkspace(ctx context.Context, obj *UserProfile) (*Workspace, error) {
	w, err := r.useCases.GetDefaultWorkspace.Do(ctx, obj.ID)
	if err != nil {
		log.Printf("can't get default workspace: %s", err)
		return nil, nil
	}

	return workspaceToGraphQL(w), nil
}

func workspaceToGraphQL(w *store.Workspace) *Workspace {
	return &Workspace{
		ID:   w.ID,
		Name: w.Name,
	}
}

//

func (r *queryResolver) Profile(ctx context.Context) (UserProfileResult, error) {
	if userCtx := apartomat.UserFromCtx(ctx); userCtx != nil {
		user, err := r.useCases.GetUserProfile.Do(ctx, userCtx.Email)
		if err != nil {
			log.Printf("can't get profile for email=%s: %s", userCtx.Email, err)
			return ServerError{Message: "can't get profile"}, nil
		}

		return UserProfile{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
			Abbr:     abbr(user.FullName),
			Gravatar: &Gravatar{
				URL: gravatar.Url(user.Email),
			},
		}, nil
	}

	return Forbidden{}, nil
}
