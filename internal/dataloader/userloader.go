package dataloader

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/users"
)

func NewUserLoaderConfig(ctx context.Context, userStore Store) UserLoaderConfig {
	return UserLoaderConfig{
		Fetch: func(keys []string) ([]*User, []error) {
			result := make([]*User, len(keys))
			errors := make([]error, len(keys))

			users, err := userStore.List(ctx, IDIn(keys...), len(keys), 0)
			if err != nil {
				return nil, []error{err}
			}

			usersByID := make(map[string]*User)
			for _, u := range users {
				usersByID[u.ID] = u
			}

			for i, id := range keys {
				result[i] = usersByID[id]
				errors[i] = err
			}

			return result, errors
		},
	}

}
