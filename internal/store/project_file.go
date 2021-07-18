package store

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"time"
)

type ProjectFile struct {
	ID         int
	Name       string
	URL        string
	Type       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type ProjectFileStore interface {
	Save(context.Context, *ProjectFile) (*ProjectFile, error)
	List(context.Context, ProjectFileStoreQuery) ([]*ProjectFile, error)
}

type ProjectFileStoreQuery struct {
	ID        expr.Int
	ProjectID expr.Int
	Limit     int
	Offset    int
}
