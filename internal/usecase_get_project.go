package apartomat

import (
	"context"
	"fmt"
	. "github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
)

func (u *Apartomat) GetProject(ctx context.Context, id string) (*store.Project, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: StrEq(id)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", id, ErrNotFound)
	}

	project := projects[0]

	if !u.CanGetProject(ctx, UserFromCtx(ctx), project) {
		return nil, fmt.Errorf("can't get project (id=%s): %w", id, ErrForbidden)
	}

	return project, nil
}

func (u *Apartomat) CanGetProject(ctx context.Context, subj *UserCtx, obj *store.Project) bool {
	// todo check subj is workspace owner or admin
	return true
}
