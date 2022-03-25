package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/houses"
	. "github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/pkg/errors"
)

func (u *Apartomat) GetRooms(ctx context.Context, houseID string, limit, offset int) ([]*Room, error) {
	hh, err := u.Houses.List(ctx, houses.IDIn(houseID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(hh) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "house %s", houseID)
	}

	house := hh[0]

	projects, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(house.ProjectID)})
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, errors.Wrapf(ErrNotFound, "project %d", house.ProjectID)
	}

	project := projects[0]

	if ok, err := u.CanGetHouses(ctx, UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.Wrapf(ErrForbidden, "can't get project (id=%d) houses", project.ID)
	}

	return u.Rooms.List(ctx, HouseIDIn(houseID), limit, offset)
}
