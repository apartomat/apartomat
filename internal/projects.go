package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
	"time"
)

func (u *Apartomat) ChangeProjectStatus(ctx context.Context, projectID string, status store.ProjectStatus) (*store.Project, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID), Limit: 1, Offset: 0})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %s", projectID)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanUpdateProject(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't update project (id=%s)", project.ID)
	}

	if project.Status != status {
		project.Status = status
		return u.Projects.Save(ctx, project)
	}

	return project, nil
}

func (u *Apartomat) ChangeProjectDates(ctx context.Context, projectID string, startAt, endAt *time.Time) (*store.Project, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID), Limit: 1, Offset: 0})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %s", projectID)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanUpdateProject(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't update project (id=%s)", project.ID)
	}

	project.StartAt = startAt
	project.EndAt = endAt

	return u.Projects.Save(ctx, project)
}

func (u *Apartomat) CanUpdateProject(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.StrEq(obj.WorkspaceID), UserID: expr.StrEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}
