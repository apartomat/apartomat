package dataloaders

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/rooms"
)

func fetchRooms(store Store) func(ctx context.Context, ids []string) ([]*Room, []error) {
	return func(ctx context.Context, ids []string) ([]*Room, []error) {
		res, err := store.List(ctx, IDIn(ids...), SortDefault, len(ids), 0)
		if err != nil {
			return nil, []error{err}
		}

		var (
			workspaces = make([]*Room, len(ids))
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

			errors[i] = ErrRoomNotFound
		}

		return workspaces, errors
	}
}
