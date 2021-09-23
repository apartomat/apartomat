package store

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"time"
)

type ProjectFile struct {
	ID         int
	ProjectID  int
	Name       string
	URL        string
	Type       ProjectFileType
	MimeType   string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type ProjectFileType string

const (
	ProjectFileTypeNone  ProjectFileType = "NONE"
	ProjectFileTypeImage ProjectFileType = "IMAGE"
)

type ProjectFileStore interface {
	Save(context.Context, *ProjectFile) (*ProjectFile, error)
	List(context.Context, ProjectFileStoreQuery) ([]*ProjectFile, error)
}

type ProjectFileStoreQuery struct {
	ID        expr.Int
	ProjectID expr.Int
	Type      ProjectFileTypeExpr
	Limit     int
	Offset    int
}

type ProjectFileTypeExpr struct {
	Eq []ProjectFileType
}