package visualizations

import "context"

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Visualization, error)
	Save(context.Context, ...*Visualization) error
	Delete(context.Context, ...*Visualization) error
}
