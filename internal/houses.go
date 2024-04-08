package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
	. "github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projects"
)

func (u *Apartomat) GetHouses(ctx context.Context, projectID string, limit, offset int) ([]*House, error) {
	if ok, err := u.Acl.CanGetHousesOfProjectID(ctx, auth.UserFromCtx(ctx), projectID); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) houses: %w", projectID, ErrForbidden)
	}

	return u.Houses.List(ctx, ProjectIDIn(projectID), SortDefault, limit, offset)
}

func (u *Apartomat) AddHouse(
	ctx context.Context,
	projectID string,
	city, address, housingComplex string,
) (*House, error) {
	project, err := u.Projects.Get(ctx, projects.IDIn(projectID))
	if ok, err := u.Acl.CanAddHouse(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add house to project (id=%s): %w", project.ID, ErrForbidden)
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, err
	}

	house := NewHouse(id, city, address, housingComplex, project.ID)

	if err := u.Houses.Save(ctx, house); err != nil {
		return nil, err
	}

	return house, nil
}

func (u *Apartomat) UpdateHouse(
	ctx context.Context,
	houseID string,
	city, address, housingComplex string,
) (*House, error) {
	house, err := u.Houses.Get(ctx, IDIn(houseID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateHouse(ctx, auth.UserFromCtx(ctx), house); err != nil {
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
