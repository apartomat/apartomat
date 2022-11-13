package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/houses"
	. "github.com/apartomat/apartomat/internal/store/rooms"
)

func (u *Apartomat) GetRooms(ctx context.Context, houseID string, limit, offset int) ([]*Room, error) {
	var (
		house *houses.House
	)

	if h, err := u.Houses.List(ctx, houses.IDIn(houseID), 1, 0); err != nil {
		return nil, err
	} else {
		if len(h) == 0 {
			return nil, fmt.Errorf("house (id=%s): %w", houseID, ErrNotFound)
		}

		house = h[0]
	}

	if ok, err := u.CanGetRooms(ctx, UserFromCtx(ctx), house); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get rooms of house (id=%s): %w", houseID, ErrForbidden)
	}

	return u.Rooms.List(ctx, HouseIDIn(houseID), limit, offset)
}

func (u *Apartomat) CanGetRooms(ctx context.Context, subj *UserCtx, obj *houses.House) (bool, error) {
	if subj == nil {
		return false, nil
	}

	var (
		project *store.Project
	)

	if p, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(obj.ProjectID)}); err != nil {
		return false, err
	} else if len(p) == 0 {
		return false, fmt.Errorf("project (id=%s): %w", obj.ProjectID, ErrNotFound)
	} else {
		project = p[0]
	}

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

func (u *Apartomat) AddRoom(
	ctx context.Context,
	houseID string,
	name string,
	square *float64,
	level *int,
) (*Room, error) {
	var (
		house *houses.House
	)

	if h, err := u.Houses.List(ctx, houses.IDIn(houseID), 1, 0); err != nil {
		return nil, err
	} else {
		if len(h) == 0 {
			return nil, fmt.Errorf("house (id=%s): %w", houseID, ErrNotFound)
		}

		house = h[0]
	}

	if ok, err := u.CanAddRoom(ctx, UserFromCtx(ctx), house); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add room to house (id=%s): %w", house.ID, ErrForbidden)
	}

	id, err := NewNanoID()
	if err != nil {
		return nil, err
	}

	room := &Room{
		ID:      id,
		Name:    name,
		Square:  square,
		Level:   level,
		HouseID: houseID,
	}

	return u.Rooms.Save(ctx, room)
}

func (u *Apartomat) CanAddRoom(ctx context.Context, subj *UserCtx, obj *houses.House) (bool, error) {
	return u.CanGetRooms(ctx, subj, obj)
}

func (u *Apartomat) UpdateRoom(
	ctx context.Context,
	roomID string,
	name string,
	square *float64,
	level *int,
) (*Room, error) {
	var (
		room *Room
	)

	if r, err := u.Rooms.List(ctx, IDIn(roomID), 1, 0); err != nil {
		return nil, err
	} else {
		if len(r) == 0 {
			return nil, fmt.Errorf("room (id=%s): %w", roomID, ErrNotFound)
		}

		room = r[0]
	}

	if ok, err := u.CanUpdateRoom(ctx, UserFromCtx(ctx), room); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update room (id=%s): %w", room.ID, ErrForbidden)
	}

	room = &Room{
		ID:      room.ID,
		Name:    name,
		Square:  square,
		Level:   level,
		HouseID: room.HouseID,
	}

	return u.Rooms.Save(ctx, room)
}

func (u *Apartomat) CanUpdateRoom(ctx context.Context, subj *UserCtx, obj *Room) (bool, error) {
	if subj == nil || obj == nil {
		return false, nil
	}

	var (
		house   *houses.House
		project *store.Project
	)

	if h, err := u.Houses.List(ctx, houses.IDIn(obj.HouseID), 1, 0); err != nil {
		return false, err
	} else {
		if len(h) == 0 {
			return false, fmt.Errorf("house (id=%s): %w", obj.HouseID, ErrNotFound)
		}

		house = h[0]
	}

	if p, err := u.Projects.List(ctx, store.ProjectStoreQuery{ID: expr.StrEq(house.ProjectID)}); err != nil {
		return false, err
	} else if len(p) == 0 {
		return false, fmt.Errorf("project (id=%s): %w", house.ProjectID, ErrNotFound)
	} else {
		project = p[0]
	}

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

func (u *Apartomat) DeleteRoom(ctx context.Context, roomID string) (*Room, error) {
	rooms, err := u.Rooms.List(ctx, IDIn(roomID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(rooms) == 0 {
		return nil, fmt.Errorf("room (id=%s): %w", roomID, ErrNotFound)
	}

	var (
		room = rooms[0]
	)

	if ok, err := u.CanDeleteRoom(ctx, UserFromCtx(ctx), room); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't delete room (id=%s): %w", room.ID, ErrForbidden)
	}

	err = u.Rooms.Delete(ctx, room)

	return room, err
}

func (u *Apartomat) CanDeleteRoom(ctx context.Context, subj *UserCtx, obj *Room) (bool, error) {
	return u.CanUpdateRoom(ctx, subj, obj)
}
