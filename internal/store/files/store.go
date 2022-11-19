package files

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
)

type Store interface {
	Save(context.Context, ...*File) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*File, error)
	Count(context.Context, Spec) (int, error)
}

type ProjectFileStoreQuery struct {
	ID        expr.Str
	ProjectID expr.Str
	Type      ProjectFileTypeExpr
	Limit     int
	Offset    int
}

type ProjectFileTypeExpr struct {
	Eq []FileType
}