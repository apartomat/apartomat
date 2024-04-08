package dataloaders

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/users"
)

func fetchUsers(store Store) func(ctx context.Context, ids []string) ([]*User, []error) {
	return func(ctx context.Context, ids []string) ([]*User, []error) {
		res, err := store.List(ctx, IDIn(ids...), SortDefault, len(ids), 0)
		if err != nil {
			return nil, []error{err}
		}

		var (
			users  = make([]*User, len(ids))
			errors = make([]error, len(ids))
		)

	keys:
		for i, id := range ids {
			for _, u := range res {
				if u.ID == id {
					users[i] = u
					continue keys
				}
			}

			errors[i] = ErrUserNotFound
		}

		return users, errors
	}
}
