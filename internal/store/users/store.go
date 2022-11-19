package users

import (
	"context"
)

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*User, error)
	Save(context.Context, ...*User) error
}
