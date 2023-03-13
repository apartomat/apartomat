package albums

import "context"

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Album, error)
	Save(context.Context, ...*Album) error
	Delete(context.Context, ...*Album) error
	Count(context.Context, Spec) (int, error)
}
