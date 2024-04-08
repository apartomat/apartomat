package houses

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrHouseNotFound = fmt.Errorf("house: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*House, error)
	Get(ctx context.Context, spec Spec) (*House, error)
	Save(context.Context, ...*House) error
}
