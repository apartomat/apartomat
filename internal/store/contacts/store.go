package contacts

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrContactNotFound = fmt.Errorf("contact: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Contact, error)
	Get(ctx context.Context, spec Spec) (*Contact, error)
	Save(context.Context, ...*Contact) error
	Delete(context.Context, ...*Contact) error
}
