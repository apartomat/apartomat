package postgres

import (
	"errors"
	"github.com/doug-martin/goqu/v9/exp"

	. "github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/doug-martin/goqu/v9"
)

type specQuery interface {
	Expression() (goqu.Expression, error)
}

func toSpecQuery(spec Spec) (specQuery, error) {
	if spec == nil {
		return nil, nil
	}

	if s, ok := spec.(specQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return idInSpecQuery{s}, nil
	case HouseIDInSpec:
		return houseIDInSpecQuery{s}, nil
	}

	return nil, errors.New("unknown spec")
}

type idInSpecQuery struct {
	spec IDInSpec
}

func (s idInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.IDs}, nil
}

type houseIDInSpecQuery struct {
	spec HouseIDInSpec
}

func (s houseIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"house_id": s.spec.IDs}, nil
}

func selectBySpec(tablename string, spec Spec, sort Sort, limit, offset int) (string, []interface{}, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return "", nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return "", nil, err
	}

	var (
		order exp.OrderedExpression
	)

	switch sort {
	case SortIDAsc:
		order = goqu.I("id").Asc()
	case SortIDDesc:
		order = goqu.I("id").Desc()
	case SortPositionAsc:
		order = goqu.I("sorting_position").Asc()
	case SortPositionDesc:
		order = goqu.I("sorting_position").Desc()
	}

	return goqu.From(tablename).Where(expr).Limit(uint(limit)).Offset(uint(offset)).Order(order).ToSQL()
}
