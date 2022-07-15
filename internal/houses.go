package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	. "github.com/apartomat/apartomat/internal/store/houses"
	"github.com/pkg/errors"
	"time"
)

func (u *Apartomat) GetHouses(ctx context.Context, projectID string, limit, offset int) ([]*House, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	project := projects[0]

	if ok, err := u.CanGetHouses(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't get project (id=%s) houses", project.ID)
	}

	return u.Houses.List(ctx, ProjectIDIn(project.ID), limit, offset)
}

func (u *Apartomat) CanGetHouses(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
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

func (u *Apartomat) AddHouse(
	ctx context.Context,
	projectID string,
	city, address, housingComplex string,
) (*House, error) {
	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(projectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", projectID)
	}

	project := projects[0]

	if ok, err := u.CanAddHouse(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't add house to project (id=%s)", project.ID)
	}

	id, err := NewNanoID()
	if err != nil {
		return nil, err
	}

	house := &House{
		ID:             id,
		City:           city,
		Address:        address,
		HousingComplex: housingComplex,
		ProjectID:      project.ID,
	}

	return u.Houses.Save(ctx, house)
}

func (u *Apartomat) CanAddHouse(ctx context.Context, subj *UserCtx, obj *store.Project) (bool, error) {
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
		return nil, errors.Wrapf(ErrNotFound, "house (id=%s)", houseID)
	}

	var (
		house = houses[0]
	)

	if ok, err := u.CanUpdateHouse(ctx, UserFromCtx(ctx), house); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't update house (id=%s)", houseID)
	}

	house = &House{
		ID:             house.ID,
		City:           city,
		Address:        address,
		HousingComplex: housingComplex,
		ProjectID:      house.ProjectID,
		CreatedAt:      house.CreatedAt,
		ModifiedAt:     time.Now(),
	}

	return u.Houses.Save(ctx, house)
}

func (u *Apartomat) CanUpdateHouse(ctx context.Context, subj *UserCtx, obj *House) (bool, error) {
	if subj == nil {
		return false, nil
	}

	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(obj.ProjectID)})
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
		store.WorkspaceUserStoreQuery{WorkspaceID: expr.StrEq(project.WorkspaceID), UserID: expr.StrEq(subj.ID)},
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}
