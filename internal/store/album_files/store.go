package album_files

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrAlbumFileNotFound = fmt.Errorf("album file: %w", store.ErrNotFound)
)

type Sort int

const (
	SortIDAsc Sort = iota
	SortIDDesc
	SortVersionAsc
	SortVersionDesc
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*AlbumFile, error)
	Get(ctx context.Context, spec Spec) (*AlbumFile, error)
	GetMaxVersion(ctx context.Context, spec Spec) (*AlbumFile, error)
	Count(context.Context, Spec) (int, error)
	Save(context.Context, ...*AlbumFile) error
	Delete(context.Context, ...*AlbumFile) error
}
