package projects

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrProjectNotFound = fmt.Errorf("project: %w", store.ErrNotFound)
)

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Project, error)
	Get(ctx context.Context, spec Spec) (*Project, error)
	Save(context.Context, ...*Project) error
}
