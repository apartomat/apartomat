package workspace_users

import (
	"context"
)

type Store interface {
	Save(context.Context, ...*WorkspaceUser) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*WorkspaceUser, error)
}
