package contacts

import "context"

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Contact, error)
	Save(context.Context, ...*Contact) error
	Delete(context.Context, ...*Contact) error
}
