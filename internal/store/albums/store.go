package albums

import "context"

type Store interface {
	Save(context.Context, ...*Album) error
	Delete(context.Context, ...*Album) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Album, error)
	Count(context.Context, Spec) (int, error)
}
