package workspace_users

import (
	"context"
)

type Store interface {
	List(ctx context.Context, spec Spec, limit, offset int) ([]*WorkspaceUser, error)
	Save(context.Context, ...*WorkspaceUser) error
}
