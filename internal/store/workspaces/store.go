package workspaces

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrWorkspaceNotFound = fmt.Errorf("workspace: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Workspace, error)
	Get(ctx context.Context, spec Spec) (*Workspace, error)
	Save(context.Context, ...*Workspace) error
}
