package workspaces

import (
	"context"
)

type Store interface {
	Save(context.Context, ...*Workspace) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Workspace, error)
}
