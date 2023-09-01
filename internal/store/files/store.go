package files

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrFileNotFound = fmt.Errorf("album file: %w", store.ErrNotFound)
)

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*File, error)
	Get(ctx context.Context, spec Spec) (*File, error)
	Save(context.Context, ...*File) error
	Count(context.Context, Spec) (int, error)
}
