package store

import (
	"context"
	"github.com/ztsu/apartomat/internal/pkg/expr"
	"time"
)

type Workspace struct {
	ID         int
	Name       string
	IsActive   bool
	UserID     int
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type WorkspaceStore interface {
	Save(context.Context, *Workspace) (*Workspace, error)
	List(context.Context, WorkspaceStoreQuery) ([]*Workspace, error)
}

type WorkspaceStoreQuery struct {
	ID     expr.Int
	UserID expr.Int
	Limit  int
	Offset int
}
