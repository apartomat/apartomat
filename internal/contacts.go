package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	. "github.com/apartomat/apartomat/internal/store/contacts"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pkg/errors"
	"strconv"
)

type AddContactParams struct {
	FullName string
	Photo    string
	Details  []Details
}

func (u *Apartomat) AddContact(ctx context.Context, projectID string, params AddContactParams) (*Contact, error) {
	prid, err := strconv.Atoi(projectID)
	if err != nil {
		return nil, err
	}

	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.IntEq(prid)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %s", projectID)
	}

	project := projects[0]

	if ok, err := u.CanAddContact(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't add contact to project (id=%d)", project.ID)
	}

	id, err := gonanoid.New(21)
	if err != nil {
		return nil, err
	}

	contact := &Contact{
		ID:        id,
		FullName:  params.FullName,
		Photo:     params.Photo,
		Details:   params.Details,
		ProjectID: prid,
	}

	return u.Contacts.Save(ctx, contact)
}

func (u *Apartomat) CanAddContact(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
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
