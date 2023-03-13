package projects

import (
	"context"
)

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Project, error)
	Save(context.Context, ...*Project) error
}
