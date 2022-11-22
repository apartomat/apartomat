package files

import (
	"context"
)

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*File, error)
	Count(context.Context, Spec) (int, error)
	Save(context.Context, ...*File) error
}
