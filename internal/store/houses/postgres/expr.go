package postgres

import (
	"errors"
	. "github.com/apartomat/apartomat/internal/store/houses"
	"github.com/doug-martin/goqu/v9"
)

type specQuery interface {
	Expression() (goqu.Expression, error)
}

func toSpecQuery(spec Spec) (specQuery, error) {
	if s, ok := spec.(specQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return houseIDInSpecQuery{s}, nil
	case ProjectIDInSpec:
		return houseProjectIDInSpecQuery{s}, nil
	}

	return nil, errors.New("unknown spec")
}

type houseIDInSpecQuery struct {
	spec IDInSpec
}

func (s houseIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.IDs}, nil
}

//
type houseProjectIDInSpecQuery struct {
	spec ProjectIDInSpec
}

func (s houseProjectIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"project_id": s.spec.IDs}, nil
}
