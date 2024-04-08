package users

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrUserNotFound = fmt.Errorf("user: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*User, error)
	Get(ctx context.Context, spec Spec) (*User, error)
	Save(context.Context, ...*User) error
}
