package store

import (
	"context"
	"errors"
	. "github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
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

func (s *roomsStore) Save(context.Context, *Room) (*Room, error) {
	return nil, errors.New("RoomsStore.Save not implemented yet")
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

	sql, args, err := goqu.From(roomsTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	contacts := make([]*Room, 0)

	_, err = s.db.QueryContext(ctx, &contacts, sql, args...)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

type roomSpecQuery interface {
	Expression() (goqu.Expression, error)
}

func toRoomSpecQuery(spec Spec) (roomSpecQuery, error) {
	if s, ok := spec.(roomSpecQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case HouseIDInSpec:
		return roomHouseIDInSpecQuery{s}, nil
	}

	return nil, errors.New("unknown spec")
}

//
type roomHouseIDInSpecQuery struct {
	spec HouseIDInSpec
}

func (s roomHouseIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"house_id": s.spec.IDs}, nil
}
