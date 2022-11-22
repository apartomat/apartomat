package houses

import "context"

type Store interface {
	Save(context.Context, ...*House) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*House, error)
}
