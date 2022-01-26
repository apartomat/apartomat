package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	. "github.com/apartomat/apartomat/internal/store/houses"
	"github.com/pkg/errors"
)

func (u *Apartomat) GetHouses(ctx context.Context, projectID int, limit, offset int) ([]*House, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.IntEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	project := projects[0]

	if ok, err := u.CanGetHouses(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't get project (id=%d) houses", project.ID)
	}

	return u.Houses.List(ctx, ProjectIDIn(project.ID), limit, offset)
}

func (u *Apartomat) CanGetHouses(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.IntEq(obj.WorkspaceID), UserID: expr.IntEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}
