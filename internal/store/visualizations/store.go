package visualizations

import "context"

type Store interface {
	Save(context.Context, ...*Visualization) error
	Delete(context.Context, *Visualization) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Visualization, error)
}
