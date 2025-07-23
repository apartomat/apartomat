package workspaceusers

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrWorkspaceUserNotFound = fmt.Errorf("workspace user: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*WorkspaceUser, error)
	Get(ctx context.Context, spec Spec) (*WorkspaceUser, error)
	Save(context.Context, ...*WorkspaceUser) error
}
