package rooms

import "context"

type Store interface {
	Save(context.Context, ...*Room) error
	Delete(context.Context, ...*Room) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Room, error)
}
