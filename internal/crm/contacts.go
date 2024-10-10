package crm

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	. "github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/apartomat/apartomat/internal/store/projects"
)

type AddContactParams struct {
	FullName string
	Photo    string
	Details  []Details
}

func (u *CRM) GetContacts(ctx context.Context, projectID string, limit, offset int) ([]*Contact, error) {
	if ok, err := u.Acl.CanGetContactsOfProjectID(ctx, auth.UserFromCtx(ctx), projectID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) contacts: %w", projectID, ErrForbidden)
	}

	return u.Contacts.List(ctx, ProjectIDIn(projectID), SortDefault, limit, offset)
}

func (u *CRM) AddContact(ctx context.Context, projectID string, params AddContactParams) (*Contact, error) {
	project, err := u.Projects.Get(ctx, projects.IDIn(projectID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanAddContact(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add contact to project (id=%s): %w", project.ID, ErrForbidden)
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, err
	}

	contact := NewContact(id, params.FullName, params.Photo, params.Details, projectID)

	if err := u.Contacts.Save(ctx, contact); err != nil {
		return nil, err
	}

	return contact, nil
}

type UpdateContactParams struct {
	FullName string
	Photo    string
	Details  []Details
}

func (u *CRM) UpdateContact(ctx context.Context, contactID string, params UpdateContactParams) (*Contact, error) {
	contact, err := u.Contacts.Get(ctx, IDIn(contactID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateContact(ctx, auth.UserFromCtx(ctx), contact); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update contact (id=%s): %w", contact.ID, ErrForbidden)
	}

	contact.Change(params.FullName, params.Photo, params.Details)

	if err := u.Contacts.Save(ctx, contact); err != nil {
		return nil, err
	}

	return contact, nil
}

func (u *CRM) DeleteContact(ctx context.Context, contactID string) (*Contact, error) {
	contact, err := u.Contacts.Get(ctx, IDIn(contactID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanDeleteContact(ctx, auth.UserFromCtx(ctx), contact); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't delete contact (id=%s): %w", contact.ID, ErrForbidden)
	}

	err = u.Contacts.Delete(ctx, contact)

	return contact, err
}
