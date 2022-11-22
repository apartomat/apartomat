package apartomat

import (
	"context"
	"fmt"
	. "github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
)

type AddContactParams struct {
	FullName string
	Photo    string
	Details  []Details
}

func (u *Apartomat) AddContact(ctx context.Context, projectID string, params AddContactParams) (*Contact, error) {
	prjs, err := u.Projects.List(ctx, projects.IDIn(projectID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanAddContact(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add contact to project (id=%s): %w", project.ID, ErrForbidden)
	}

	id, err := NewNanoID()
	if err != nil {
		return nil, err
	}

	contact := &Contact{
		ID:        id,
		FullName:  params.FullName,
		Photo:     params.Photo,
		Details:   params.Details,
		ProjectID: projectID,
	}

	return u.Contacts.Save(ctx, contact)
}

func (u *Apartomat) CanAddContact(ctx context.Context, subj *UserCtx, obj *projects.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(obj.WorkspaceID),
			workspace_users.UserIDIn(subj.ID),
		),
		1,
		0,
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}

func (u *Apartomat) GetContacts(ctx context.Context, projectID string, limit, offset int) ([]*Contact, error) {
	prjs, err := u.Projects.List(ctx, projects.IDIn(projectID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanGetContacts(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) contacts: %w", project.ID, ErrForbidden)
	}

	return u.Contacts.List(ctx, ProjectIDIn(project.ID), limit, offset)
}

func (u *Apartomat) CanGetContacts(ctx context.Context, subj *UserCtx, obj *projects.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(obj.WorkspaceID),
			workspace_users.UserIDIn(subj.ID),
		),
		1,
		0,
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
		return nil, fmt.Errorf("contact (id=%s): %w", contactID, ErrNotFound)
	}

	var (
		contact = contacts[0]
	)

	if ok, err := u.CanUpdateContact(ctx, UserFromCtx(ctx), contact); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update contact (id=%s): %w", contact.ID, ErrForbidden)
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

	prjs, err := u.Projects.List(ctx, projects.IDIn(obj.ProjectID), 1, 0)
	if err != nil {
		return false, err
	}

	if len(prjs) == 0 {
		return false, fmt.Errorf("project (id=%s): %w", obj.ProjectID, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(project.WorkspaceID),
			workspace_users.UserIDIn(subj.ID),
		),
		1,
		0,
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
		return nil, fmt.Errorf("contact (id=%s): %w", contactID, ErrNotFound)
	}

	var (
		contact = contacts[0]
	)

	if ok, err := u.CanDeleteContact(ctx, UserFromCtx(ctx), contact); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't delete contact (id=%s): %w", contact.ID, ErrForbidden)
	}

	err = u.Contacts.Delete(ctx, contact)

	return contact, err
}

func (u *Apartomat) CanDeleteContact(ctx context.Context, subj *UserCtx, obj *Contact) (bool, error) {
	if subj == nil {
		return false, nil
	}

	prjs, err := u.Projects.List(ctx, projects.IDIn(obj.ProjectID), 1, 0)
	if err != nil {
		return false, err
	}

	if len(prjs) == 0 {
		return false, nil
	}

	var (
		project = prjs[0]
	)

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(project.WorkspaceID),
			workspace_users.UserIDIn(subj.ID),
		),
		1,
		0,
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
