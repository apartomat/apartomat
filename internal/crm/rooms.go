package crm

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/apartomat/apartomat/internal/store/houses"
	. "github.com/apartomat/apartomat/internal/store/rooms"
)

func (u *CRM) GetRooms(ctx context.Context, houseID string, limit, offset int) ([]*Room, error) {
	house, err := u.Houses.Get(ctx, houses.IDIn(houseID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanGetRooms(ctx, auth.UserFromCtx(ctx), house); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get rooms of house (id=%s): %w", houseID, ErrForbidden)
	}

	return u.Rooms.List(ctx, HouseIDIn(houseID), SortPositionAsc, limit, offset)
}

func (u *CRM) AddRoom(
	ctx context.Context,
	houseID string,
	name string,
	square *float64,
	level *int,
) (*Room, error) {
	house, err := u.Houses.Get(ctx, houses.IDIn(houseID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanAddRoom(ctx, auth.UserFromCtx(ctx), house); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add room to house (id=%s): %w", house.ID, ErrForbidden)
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, err
	}

	var (
		room = NewRoom(id, name, square, level, 0, houseID)
	)

	if err := u.Rooms.Save(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

func (u *CRM) UpdateRoom(
	ctx context.Context,
	roomID string,
	name string,
	square *float64,
	level *int,
) (*Room, error) {
	room, err := u.Rooms.Get(ctx, IDIn(roomID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateRoom(ctx, auth.UserFromCtx(ctx), room); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update room (id=%s): %w", room.ID, ErrForbidden)
	}

	room.Name = name
	room.Square = square
	room.Level = level

	if err := u.Rooms.Save(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

func (u *CRM) DeleteRoom(ctx context.Context, roomID string) (*Room, error) {
	room, err := u.Rooms.Get(ctx, IDIn(roomID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanDeleteRoom(ctx, auth.UserFromCtx(ctx), room); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't delete room (id=%s): %w", room.ID, ErrForbidden)
	}

	err = u.Rooms.Delete(ctx, room)

	return room, err
}

func (u *CRM) MoveRoomToPosition(ctx context.Context, roomID string, position int) (*Room, error) {
	room, err := u.Rooms.Get(ctx, IDIn(roomID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateRoom(ctx, auth.UserFromCtx(ctx), room); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update room (id=%s): %w", room.ID, ErrForbidden)
	}

	var (
		asc = room.SortingPosition < position
	)

	room.MoveToPosition(position)

	if err := u.Rooms.Save(ctx, room); err != nil {
		return nil, err
	}

	if err := u.Rooms.Reorder(ctx, room.HouseID, asc); err != nil {
		return nil, err
	}

	return room, nil
}
