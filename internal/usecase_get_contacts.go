package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	. "github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/pkg/errors"
)

type GetContacts struct {
	projects store.ProjectStore
	contacts Store
	acl      *Acl
}

func NewGetContacts(
	projects store.ProjectStore,
	contacts Store,
	acl *Acl,
) *GetContacts {
	return &GetContacts{projects, contacts, acl}
}

func (u *GetContacts) Do(ctx context.Context, projectID int, limit, offset int) ([]*Contact, error) {
	projects, err := u.projects.List(ctx, store.ProjectStoreQuery{ID: expr.IntEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	project := projects[0]

	if ok, err := u.acl.CanGetProjectContacts(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't get project (id=%d) contacts", project.ID)
	}

	return u.contacts.List(ctx, ProjectIDIn(project.ID), limit, offset)
}
