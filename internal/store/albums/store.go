package albums

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrAlbumNotFound = fmt.Errorf("album: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Album, error)
	Get(ctx context.Context, spec Spec) (*Album, error)
	Count(context.Context, Spec) (int, error)
	Save(context.Context, ...*Album) error
	Delete(context.Context, ...*Album) error
}
