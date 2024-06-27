package visualizations

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrVisualizationNotFound = fmt.Errorf("visualization: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
	SortIDAsc
	SortIDDesc
	SortPositionAsc
	SortPositionDesc
	SortRoomAscPositionAsc
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Visualization, error)
	Get(ctx context.Context, spec Spec) (*Visualization, error)
	GetMaxSortingPosition(ctx context.Context, spec Spec) (int, error)
	Save(context.Context, ...*Visualization) error
	Delete(context.Context, ...*Visualization) error
}
