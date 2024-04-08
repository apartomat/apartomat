package dataloaders

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/workspaces"
)

func fetchWorkspaces(store Store) func(ctx context.Context, ids []string) ([]*Workspace, []error) {
	return func(ctx context.Context, ids []string) ([]*Workspace, []error) {
		res, err := store.List(ctx, IDIn(ids...), SortDefault, len(ids), 0)
		if err != nil {
			return nil, []error{err}
		}

		var (
			workspaces = make([]*Workspace, len(ids))
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

			errors[i] = ErrWorkspaceNotFound
		}

		return workspaces, errors
	}
}
