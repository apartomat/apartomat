package rooms

import "context"

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Room, error)
	Save(context.Context, ...*Room) error
	Delete(context.Context, ...*Room) error
}
