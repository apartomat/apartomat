package store

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"time"
)

type ProjectFile struct {
	ID         string
	Name       string
	URL        string
	Type       ProjectFileType
	MimeType   string
	CreatedAt  time.Time
	ModifiedAt time.Time
	ProjectID  string
}

type ProjectFileType string

const (
	ProjectFileTypeNone          ProjectFileType = "NONE"
	ProjectFileTypeVisualization ProjectFileType = "VISUALIZATION"
)

type ProjectFileStore interface {
	Save(context.Context, *ProjectFile) (*ProjectFile, error)
	List(context.Context, ProjectFileStoreQuery) ([]*ProjectFile, error)
	Count(context.Context, ProjectFileStoreQuery) (int, error)
}

type ProjectFileStoreQuery struct {
	ID        expr.Str
	ProjectID expr.Str
	Type      ProjectFileTypeExpr
	Limit     int
	Offset    int
}

type ProjectFileTypeExpr struct {
	Eq []ProjectFileType
}
