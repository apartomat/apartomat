package projects

import (
	"context"
)

type Store interface {
	Save(context.Context, ...*Project) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Project, error)
}
