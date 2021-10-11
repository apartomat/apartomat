package store

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"time"
)

type Project struct {
	ID          int
	Name        string
	Status      ProjectStatus
	WorkspaceID int
	StartAt     *time.Time
	EndAt       *time.Time
	CreatedAt   time.Time
	ModifiedAt  time.Time
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
	ID          expr.Int
	WorkspaceID expr.Int
	Limit       int
	Offset      int
}
