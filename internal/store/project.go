package store

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"time"
)

type Project struct {
	ID          string
	Name        string
	Status      ProjectStatus
	StartAt     *time.Time
	EndAt       *time.Time
	CreatedAt   time.Time
	ModifiedAt  time.Time
	WorkspaceID string
}

type ProjectStatus string

const (
	ProjectStatusNew        ProjectStatus = "NEW"
	ProjectStatusInProgress ProjectStatus = "IN_PROGRESS"
	ProjectStatusDone       ProjectStatus = "DONE"
	ProjectStatusCanceled   ProjectStatus = "CANCELED"
)

type ProjectStore interface {
	Save(context.Context, *Project) (*Project, error)
	List(context.Context, ProjectStoreQuery) ([]*Project, error)
}

type ProjectStoreQuery struct {
	ID          expr.Str
	WorkspaceID expr.Str
	Status      ProjectStatusExpr
	Limit       int
	Offset      int
}

type ProjectStatusExpr struct {
	Eq []ProjectStatus
}
