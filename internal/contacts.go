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

type UpdateContactParams struct {
	FullName string
	Photo    string
	Details  []Details
}

func (u *Apartomat) UpdateContact(ctx context.Context, contactID string, params UpdateContactParams) (*Contact, error) {
	contacts, err := u.Contacts.List(ctx, IDIn(contactID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(contacts) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "contact %s", contactID)
	}

	var (
		contact = contacts[0]
	)

	if ok, err := u.CanUpdateContact(ctx, UserFromCtx(ctx), contact); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't update contact (id=%s)", contact.ID)
	}

	contact = &Contact{
		ID:        contact.ID,
		FullName:  params.FullName,
		Photo:     params.Photo,
		Details:   params.Details,
		ProjectID: contact.ProjectID,
	}

	return u.Contacts.Save(ctx, contact)
}

func (u *Apartomat) CanUpdateContact(ctx context.Context, subj *UserCtx, obj *Contact) (bool, error) {
	if subj == nil {
		return false, nil
	}

	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.IntEq(obj.ProjectID)})
	if err != nil {
		return false, err
	}

	if len(projects) == 0 {
		return false, errors.Wrapf(ErrNotFound, "project %s", obj.ProjectID)
	}

	var (
		project = projects[0]
	)

	wu, err := u.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.IntEq(project.WorkspaceID), UserID: expr.IntEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}

func (u *Apartomat) DeleteContact(ctx context.Context, contactID string) (*Contact, error) {
	contacts, err := u.Contacts.List(ctx, IDIn(contactID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(contacts) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "contact %s", contactID)
	}

	var (
		contact = contacts[0]
	)

	if ok, err := u.CanDeleteContact(ctx, UserFromCtx(ctx), contact); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't delete contact (id=%s)", contact.ID)
	}

	err = u.Contacts.Delete(ctx, contact)

	return contact, err
}

func (u *Apartomat) CanDeleteContact(ctx context.Context, subj *UserCtx, obj *Contact) (bool, error) {
	if subj == nil {
		return false, nil
	}

	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.IntEq(obj.ProjectID), Limit: 1, Offset: 0})

	if len(projects) == 0 {
		return false, nil
	}

	var (
		project = projects[0]
	)

	wu, err := u.WorkspaceUsers.List(
		ctx,
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.IntEq(project.WorkspaceID), UserID: expr.IntEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	var (
		workspace = wu[0]
	)

	return workspace.UserID == subj.ID, nil
}
