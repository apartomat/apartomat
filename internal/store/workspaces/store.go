package workspaces

import (
	"context"
)

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Workspace, error)
	Save(context.Context, ...*Workspace) error
}
