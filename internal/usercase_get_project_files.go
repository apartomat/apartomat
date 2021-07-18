package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
)

type GetProjectFiles struct {
	projects     store.ProjectStore
	projectFiles store.ProjectFileStore
	acl          *Acl
}

func NewGetProjectFiles(
	projects store.ProjectStore,
	projectFiles store.ProjectFileStore,
	acl *Acl,
) *GetProjectFiles {
	return &GetProjectFiles{projects, projectFiles, acl}
}

func (u *GetProjectFiles) Do(ctx context.Context, projectID, limit, offset int) ([]*store.ProjectFile, error) {
	projects, err := u.projects.List(ctx, store.ProjectStoreQuery{ID: expr.IntEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	project := projects[0]

	if !u.acl.CanGetProjectFiles(ctx, UserFromCtx(ctx), project) {
		return nil, errors.Wrapf(ErrForbidden, "can't get project (id=%d) files", project.ID)
	}

	p, err := u.projectFiles.List(ctx, store.ProjectFileStoreQuery{ProjectID: expr.IntEq(projectID)})
	if err != nil {
		return nil, err
	}

	return p, nil
}
