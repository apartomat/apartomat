package rooms

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrRoomNotFound = fmt.Errorf("room: %w", store.ErrNotFound)
)

type Sort int

const (
	SortIDAsc Sort = iota
	SortIDDesc
	SortPositionAsc
	SortPositionDesc
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Room, error)
	Get(ctx context.Context, spec Spec) (*Room, error)
	Save(context.Context, ...*Room) error
	Delete(context.Context, ...*Room) error
	Reorder(ctx context.Context, houseID string, asc bool) error
}
