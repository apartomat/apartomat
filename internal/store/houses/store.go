package houses

import "context"

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*House, error)
	Save(context.Context, ...*House) error
}
