package dataloaders

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/files"
)

func fetchFiles(store Store) func(ctx context.Context, ids []string) ([]*File, []error) {
	return func(ctx context.Context, ids []string) ([]*File, []error) {
		res, err := store.List(ctx, IDIn(ids...), SortDefault, len(ids), 0)
		if err != nil {
			return nil, []error{err}
		}

		var (
			workspaces = make([]*File, len(ids))
			errors     = make([]error, len(ids))
		)

	keys:
		for i, id := range ids {
			for _, w := range res {
				if w.ID == id {
					workspaces[i] = w
					continue keys
				}
			}

			errors[i] = ErrFileNotFound
		}

		return workspaces, errors
	}
}
