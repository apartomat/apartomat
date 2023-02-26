package store

import (
	"context"
	"errors"
	. "github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	roomsTableName = `apartomat.rooms`
)

type roomsStore struct {
	db *pg.DB
}

func NewRoomsStore(db *pg.DB) *roomsStore {
	return &roomsStore{db}
}

var (
	_ Store = (*roomsStore)(nil)
)

func (s *roomsStore) Save(ctx context.Context, room *Room) (*Room, error) {
	if room.CreatedAt.IsZero() {
		room.CreatedAt = time.Now()
	}

	if room.ModifiedAt.IsZero() {
		room.ModifiedAt = room.CreatedAt
	}

	rec := toRoomsRecord(room)

	_, err := s.db.ModelContext(ctx, rec).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *roomsStore) Delete(ctx context.Context, contact *Room) error {
	_, err := s.db.ModelContext(ctx, (*roomsRecord)(nil)).Where(`id = ?`, contact.ID).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (s *roomsStore) List(ctx context.Context, spec Spec, limit, offset int) ([]*Room, error) {
	qs, err := toRoomSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	orderExpr := goqu.I("created_at").Asc()

	sql, args, err := goqu.From(roomsTableName).
		Where(expr).
		Order(orderExpr).
		Limit(uint(limit)).
		Offset(uint(offset)).
		ToSQL()
	if err != nil {
		return nil, err
	}

	recs := make([]*roomsRecord, 0)

	_, err = s.db.QueryContext(ctx, &recs, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRoomsRecords(recs), nil
}

type roomSpecQuery interface {
	Expression() (goqu.Expression, error)
}

func toRoomSpecQuery(spec Spec) (roomSpecQuery, error) {
	if s, ok := spec.(roomSpecQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return roomIDInSpecQuery{s}, nil
	case HouseIDInSpec:
		return roomHouseIDInSpecQuery{s}, nil
	}

	return nil, errors.New("unknown spec")
}

type roomIDInSpecQuery struct {
	spec IDInSpec
}

func (s roomIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.IDs}, nil
}

type roomHouseIDInSpecQuery struct {
	spec HouseIDInSpec
}

func (s roomHouseIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"house_id": s.spec.IDs}, nil
}

type roomsRecord struct {
	tableName  struct{}  `pg:"apartomat.rooms,alias:rooms"`
	ID         string    `pg:"id,pk"`
	Name       string    `pg:"name"`
	Square     *float64  `pg:"square"`
	Level      *int      `pg:"level"`
	CreatedAt  time.Time `pg:"created_at"`
	ModifiedAt time.Time `pg:"modified_at"`
	HouseID    string    `pg:"house_id"`
}

func toRoomsRecord(room *Room) *roomsRecord {
	return &roomsRecord{
		ID:         room.ID,
		Name:       room.Name,
		Square:     room.Square,
		Level:      room.Level,
		CreatedAt:  room.CreatedAt,
		ModifiedAt: room.ModifiedAt,
		HouseID:    room.HouseID,
	}
}

func fromRoomsRecords(records []*roomsRecord) []*Room {
	rooms := make([]*Room, len(records))

	for i, r := range records {
		rooms[i] = &Room{
			ID:         r.ID,
			Name:       r.Name,
			Square:     r.Square,
			Level:      r.Level,
			CreatedAt:  r.CreatedAt,
			ModifiedAt: r.ModifiedAt,
			HouseID:    r.HouseID,
		}
	}

	return rooms
}
