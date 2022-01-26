package graphql

import (
	"context"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/pkg/gravatar"
	"log"
)

func (r *queryResolver) Profile(ctx context.Context) (UserProfileResult, error) {
	if userCtx := apartomat.UserFromCtx(ctx); userCtx != nil {
		user, err := r.useCases.GetUserProfile(ctx, userCtx.Email)
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

	return forbidden()
}
