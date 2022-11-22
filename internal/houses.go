package apartomat

import (
	"context"
	"fmt"
	. "github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
)

func (u *Apartomat) GetHouses(ctx context.Context, projectID string, limit, offset int) ([]*House, error) {
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

	if ok, err := u.CanGetHouses(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) houses: %w", project.ID, ErrForbidden)
	}

	return u.Houses.List(ctx, ProjectIDIn(project.ID), limit, offset)
}

func (u *Apartomat) CanGetHouses(ctx context.Context, subj *UserCtx, obj *projects.Project) (bool, error) {
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

func (u *Apartomat) AddHouse(
	ctx context.Context,
	projectID string,
	city, address, housingComplex string,
) (*House, error) {
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

	if ok, err := u.CanAddHouse(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add house to project (id=%s): %w", project.ID, ErrForbidden)
	}

	id, err := NewNanoID()
	if err != nil {
		return nil, err
	}

	house := New(id, city, address, housingComplex, project.ID)

	if err := u.Houses.Save(ctx, house); err != nil {
		return nil, err
	}

	return house, nil
}

func (u *Apartomat) CanAddHouse(ctx context.Context, subj *UserCtx, obj *projects.Project) (bool, error) {
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

func (u *Apartomat) UpdateHouse(
	ctx context.Context,
	houseID string,
	city, address, housingComplex string,
) (*House, error) {
	houses, err := u.Houses.List(ctx, IDIn(houseID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(houses) == 0 {
		return nil, fmt.Errorf("house (id=%s): %w", houseID, ErrNotFound)
	}

	var (
		house = houses[0]
	)

	if ok, err := u.CanUpdateHouse(ctx, UserFromCtx(ctx), house); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update house (id=%s): %w", houseID, ErrForbidden)
	}

	house.Change(city, address, housingComplex)

	if err := u.Houses.Save(ctx, house); err != nil {
		return nil, err
	}

	return house, nil
}

func (u *Apartomat) CanUpdateHouse(ctx context.Context, subj *UserCtx, obj *House) (bool, error) {
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
