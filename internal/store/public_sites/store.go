package public_sites

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
)

var (
	ErrPublicSiteNotFound = fmt.Errorf("public site: %w", store.ErrNotFound)
)

type Store interface {
	Get(ctx context.Context, spec Spec) (*PublicSite, error)
	List(ctx context.Context, spec Spec, limit, offset int) ([]PublicSite, error)
	Save(context.Context, ...PublicSite) error
}
