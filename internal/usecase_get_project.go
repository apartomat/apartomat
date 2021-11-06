package apartomat

import (
	"context"
	. "github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type GetProject struct {
	projects store.ProjectStore
	acl      *Acl
}

func NewGetProject(
	projects store.ProjectStore,
	acl *Acl,
) *GetProject {
	return &GetProject{projects, acl}
}

func (u *GetProject) Do(ctx context.Context, id int) (*store.Project, error) {
	projects, err := u.projects.List(ctx, store.ProjectStoreQuery{ID: IntEq(id)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", id)
	}

	project := projects[0]

	if !u.acl.CanGetProject(ctx, UserFromCtx(ctx), project) {
		return nil, errors.Wrapf(ErrForbidden, "can't get project %d", id)
	}

	return project, nil
}
