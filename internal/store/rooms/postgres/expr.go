package postgres

import (
	"errors"

	. "github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/doug-martin/goqu/v9"
)

type specQuery interface {
	Expression() (goqu.Expression, error)
}

func toRoomSpecQuery(spec Spec) (specQuery, error) {
	if s, ok := spec.(specQuery); ok {
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
