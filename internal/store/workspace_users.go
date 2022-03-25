package store

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"time"
)

type WorkspaceUserRole string

const (
	WorkspaceUserRoleAdmin = "admin"
	WorkspaceUserRoleUser  = "user"
)

type WorkspaceUser struct {
	ID          string
	WorkspaceID string
	UserID      string
	Role        WorkspaceUserRole
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

type WorkspaceUserStore interface {
	Save(context.Context, *WorkspaceUser) (*WorkspaceUser, error)
	List(context.Context, WorkspaceUserStoreQuery) ([]*WorkspaceUser, error)
}

type WorkspaceUserStoreQuery struct {
	WorkspaceID expr.Str
	UserID      expr.Str
	Limit       int
	Offset      int
}
