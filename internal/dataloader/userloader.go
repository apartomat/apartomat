package dataloader

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

func NewUserLoaderConfig(ctx context.Context, userStore store.UserStore) UserLoaderConfig {
	return UserLoaderConfig{
		Fetch: func(keys []string) ([]*store.User, []error) {
			result := make([]*store.User, len(keys))
			errors := make([]error, len(keys))

			users, err := userStore.List(ctx, store.UserStoreQuery{ID: expr.Str{Eq: keys}})
			if err != nil {
				return nil, []error{err}
			}

			usersByID := make(map[string]*store.User)
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
