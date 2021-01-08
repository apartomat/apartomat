package store

import (
	"context"
	"github.com/ztsu/apartomat/internal/pkg/expr"
	"time"
)

type WorkspaceUserRole string

const (
	WorkspaceUserRoleAdmin = "admin"
	WorkspaceUserRoleUser  = "user"
)

type WorkspaceUser struct {
	ID          int
	WorkspaceID int
	UserID      int
	Role        WorkspaceUserRole
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

type WorkspaceUserStore interface {
	Save(context.Context, *WorkspaceUser) (*WorkspaceUser, error)
	List(context.Context, WorkspaceUserStoreQuery) ([]*WorkspaceUser, error)
}

type WorkspaceUserStoreQuery struct {
	WorkspaceID expr.Int
	UserID      expr.Int
	Limit       int
	Offset      int
}
