package store

import (
	"context"
	"github.com/ztsu/apartomat/internal/pkg/expr"
	"time"
)

type Project struct {
	ID          int
	Name        string
	IsActive    bool
	WorkspaceID int
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

type ProjectStore interface {
	Save(context.Context, *Project) (*Project, error)
	List(context.Context, ProjectStoreQuery) ([]*Project, error)
}

type ProjectStoreQuery struct {
	ID          expr.Int
	WorkspaceID expr.Int
	Limit       int
	Offset      int
}
