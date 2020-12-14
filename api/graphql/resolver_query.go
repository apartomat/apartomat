package graphql

import (
	"context"
	"crypto/md5"
	"fmt"
)

type queryResolver struct {
	*rootResolver
}

func (r *queryResolver) Profile(ctx context.Context) (UserProfileResult, error) {
	if userCtx := UserFromCtx(ctx); userCtx != nil {
		user, err := r.useCases.GetUserProfile.Do(ctx, userCtx.Email)
		if err != nil {
			return ServerError{Message: "can't get profile"}, nil
		}

		return UserProfile{
			Email: user.Email,
			Gravatar: &Gravatar{
				URL: fmt.Sprintf("https://www.gravatar.com/avatar/%x", md5.Sum([]byte(user.Email))),
			},
		}, nil
	}

	return Forbidden{}, nil
}

func (r *queryResolver) ShoppingList(ctx context.Context) (*ShoppingList, error) {
	return &ShoppingList{}, nil
}
