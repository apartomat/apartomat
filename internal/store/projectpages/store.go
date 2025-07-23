package projectpages

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrProjectPageNotFound = fmt.Errorf("project page: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]ProjectPage, error)
	Get(ctx context.Context, spec Spec) (*ProjectPage, error)
	Save(context.Context, ...ProjectPage) error
}
