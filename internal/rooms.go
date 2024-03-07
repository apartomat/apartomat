package apartomat

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/store/houses"
	projects "github.com/apartomat/apartomat/internal/store/projects"
	. "github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
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

	if ok, err := u.CanGetRooms(ctx, auth.UserFromCtx(ctx), house); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get rooms of house (id=%s): %w", houseID, ErrForbidden)
	}

	return u.Rooms.List(ctx, HouseIDIn(houseID), SortPositionAsc, limit, offset)
}

func (u *Apartomat) CanGetRooms(ctx context.Context, subj *auth.UserCtx, obj *houses.House) (bool, error) {
	if subj == nil {
		return false, nil
	}

	var (
		project *projects.Project
	)

	if p, err := u.Projects.List(ctx, projects.IDIn(obj.ProjectID), 1, 0); err != nil {
		return false, err
	} else if len(p) == 0 {
		return false, fmt.Errorf("project (id=%s): %w", obj.ProjectID, ErrNotFound)
	} else {
		project = p[0]
	}

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

	if ok, err := u.CanAddRoom(ctx, auth.UserFromCtx(ctx), house); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't add room to house (id=%s): %w", house.ID, ErrForbidden)
	}

	id, err := GenerateNanoID()
	if err != nil {
		return nil, err
	}

	var (
		room = NewRoom(id, name, square, level, houseID)
	)

	if err := u.Rooms.Save(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

func (u *Apartomat) CanAddRoom(ctx context.Context, subj *auth.UserCtx, obj *houses.House) (bool, error) {
	return u.CanGetRooms(ctx, subj, obj)
}

func (u *Apartomat) UpdateRoom(
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

	if ok, err := u.CanUpdateRoom(ctx, auth.UserFromCtx(ctx), room); err != nil {
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

func (u *Apartomat) CanUpdateRoom(ctx context.Context, subj *auth.UserCtx, obj *Room) (bool, error) {
	if subj == nil || obj == nil {
		return false, nil
	}

	var (
		house   *houses.House
		project *projects.Project
	)

	if h, err := u.Houses.List(ctx, houses.IDIn(obj.HouseID), 1, 0); err != nil {
		return false, err
	} else {
		if len(h) == 0 {
			return false, fmt.Errorf("house (id=%s): %w", obj.HouseID, ErrNotFound)
		}

		house = h[0]
	}

	if p, err := u.Projects.List(ctx, projects.IDIn(house.ProjectID), 1, 0); err != nil {
		return false, err
	} else if len(p) == 0 {
		return false, fmt.Errorf("project (id=%s): %w", house.ProjectID, ErrNotFound)
	} else {
		project = p[0]
	}

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

func (u *Apartomat) DeleteRoom(ctx context.Context, roomID string) (*Room, error) {
	room, err := u.Rooms.Get(ctx, IDIn(roomID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.CanDeleteRoom(ctx, auth.UserFromCtx(ctx), room); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't delete room (id=%s): %w", room.ID, ErrForbidden)
	}

	err = u.Rooms.Delete(ctx, room)

	return room, err
}

func (u *Apartomat) CanDeleteRoom(ctx context.Context, subj *auth.UserCtx, obj *Room) (bool, error) {
	return u.CanUpdateRoom(ctx, subj, obj)
}

func (u *Apartomat) MoveRoomToPosition(ctx context.Context, roomID string, position int) (*Room, error) {
	room, err := u.Rooms.Get(ctx, IDIn(roomID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.CanUpdateRoom(ctx, auth.UserFromCtx(ctx), room); err != nil {
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
