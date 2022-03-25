package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

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

	project := projects[0]

	if !u.CanGetProjectFiles(ctx, UserFromCtx(ctx), project) {
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

type GetProjectFilesFilter struct {
	Type store.ProjectFileTypeExpr
}

func (u *Apartomat) CanGetProjectFiles(ctx context.Context, subj *UserCtx, obj *store.Project) bool {
	return true
}
