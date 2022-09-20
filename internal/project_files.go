package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type GetProjectFilesFilter struct {
	Type store.ProjectFileTypeExpr
}

func (u *Apartomat) GetProjectFiles(
	ctx context.Context,
	projectID string,
	filter GetProjectFilesFilter,
	limit, offset int,
) ([]*store.ProjectFile, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanGetProjectFiles(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't get project (id=%d) files", project.ID)
	}

	p, err := u.ProjectFiles.List(
		ctx,
		store.ProjectFileStoreQuery{
			ProjectID: expr.StrEq(projectID),
			Type:      filter.Type,
			Limit:     limit,
			Offset:    offset,
		},
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *Apartomat) CanGetProjectFiles(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
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

func (u *Apartomat) CountProjectFiles(
	ctx context.Context,
	projectID string,
	filter GetProjectFilesFilter,
) (int, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID)})
	if err != nil {
		return 0, err
	}

	if len(projects) == 0 {
		return 0, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanCountProjectFiles(ctx, UserFromCtx(ctx), project); err != nil {
		return 0, err
	} else if !ok {
		return 0, errors.Wrapf(ErrForbidden, "can't count project (id=%d) files", project.ID)
	}

	return u.ProjectFiles.Count(
		ctx,
		store.ProjectFileStoreQuery{
			ProjectID: expr.StrEq(projectID),
			Type:      filter.Type,
		},
	)
}

func (u *Apartomat) CanCountProjectFiles(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
	return u.CanGetProjectFiles(ctx, subj, obj)
}
