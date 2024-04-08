package public_sites

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrPublicSiteNotFound = fmt.Errorf("public site: %w", store.ErrNotFound)
)

type Sort int

const (
	SortDefault Sort = iota
)

type Store interface {
	List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]PublicSite, error)
	Get(ctx context.Context, spec Spec) (*PublicSite, error)
	Save(context.Context, ...PublicSite) error
}
