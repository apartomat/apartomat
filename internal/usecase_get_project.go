package apartomat

import (
	"context"
	. "github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

func (u *Apartomat) GetProject(ctx context.Context, id int) (*store.Project, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: IntEq(id)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", id)
	}

	project := projects[0]

	if !u.CanGetProject(ctx, UserFromCtx(ctx), project) {
		return nil, errors.Wrapf(ErrForbidden, "can't get project %d", id)
	}

	return project, nil
}

func (u *Apartomat) CanGetProject(ctx context.Context, subj *UserCtx, obj *store.Project) bool {
	// todo check subj is workspace owner or admin
	return true
}
