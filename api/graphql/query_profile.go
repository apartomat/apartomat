package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/pkg/gravatar"
	"log"
)

func (r *queryResolver) Profile(ctx context.Context) (UserProfileResult, error) {
	if userCtx := apartomat.UserFromCtx(ctx); userCtx != nil {
		user, err := r.useCases.GetUserProfile(ctx, userCtx.ID)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't get profile (id=%d): %s", userCtx.ID, err)

			return ServerError{}, nil
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
