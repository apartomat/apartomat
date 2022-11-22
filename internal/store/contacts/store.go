package contacts

import "context"

type Store interface {
	Save(context.Context, ...*Contact) error
	Delete(context.Context, ...*Contact) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Contact, error)
}
