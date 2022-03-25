package store

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"time"
)

type Workspace struct {
	ID         string
	Name       string
	IsActive   bool
	CreatedAt  time.Time
	ModifiedAt time.Time
	UserID     string
}

type WorkspaceStore interface {
	Save(context.Context, *Workspace) (*Workspace, error)
	List(context.Context, WorkspaceStoreQuery) ([]*Workspace, error)
}

type WorkspaceStoreQuery struct {
	ID     expr.Str
	UserID expr.Str
	Limit  int
	Offset int
}
