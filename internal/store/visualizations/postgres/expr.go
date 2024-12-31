package postgres

import (
	"errors"

	. "github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var (
	tableVisualizations = goqu.T("visualizations").Schema("apartomat").As("v")
	tableRooms          = goqu.T("rooms").Schema("apartomat").As("r")
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
	case ProjectIDInSpec:
		return projectIDInSpecQuery{s}, nil
	case RoomIDInSpec:
		return roomIDInSpecQuery{s}, nil
	case StatusInSpec:
		return statusInSpecQuery{s}, nil
	case StatusNotInSpec:
		return statusNotInSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown spec")
}

type idInSpecQuery struct {
	spec IDInSpec
}

func (s idInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.ID}, nil
}

type projectIDInSpecQuery struct {
	spec ProjectIDInSpec
}

func (s projectIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"project_id": s.spec.ProjectID}, nil
}

type roomIDInSpecQuery struct {
	spec RoomIDInSpec
}

func (s roomIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"room_id": s.spec.RoomID}, nil
}

type statusInSpecQuery struct {
	spec StatusInSpec
}

func (s statusInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"status": s.spec.Status}, nil
}

type statusNotInSpecQuery struct {
	spec StatusNotInSpec
}

func (s statusNotInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"status": goqu.Op{"notIn": s.spec.Status}}, nil
}

type andSpecQuery struct {
	spec AndSpec
}

func (s andSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toSpecQuery(spec); err != nil {
			return nil, err
		} else if ps != nil {
			expr, err := ps.Expression()
			if err != nil {
				return nil, err
			}

			exs = append(exs, expr)
		}
	}

	return goqu.And(exs...), nil
}

type orSpecQuery struct {
	spec OrSpec
}

func (s orSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toSpecQuery(spec); err != nil {
			return nil, err
		} else if ps != nil {
			expr, err := ps.Expression()
			if err != nil {
				return nil, err
			}

			exs = append(exs, expr)
		}
	}

	return goqu.Or(exs...), nil
}

func specToExpr(spec Spec) (goqu.Expression, error) {
	sq, err := toSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := sq.Expression()
	if err != nil {
		return nil, err
	}

	return expr, nil
}

func selectBySpec(spec Spec, sort Sort, limit, offset int) (string, []interface{}, error) {
	expr, err := specToExpr(spec)
	if err != nil {
		return "", nil, err
	}

	type LeftJoin struct {
		Table goqu.Expression
		Cond  exp.JoinCondition
	}

	var (
		order = make([]exp.OrderedExpression, 0)

		join interface{}
	)

	switch sort {
	case SortDefault:
	case SortIDAsc:
		order = append(order, goqu.I("v.id").Asc())
	case SortIDDesc:
		order = append(order, goqu.I("v.id").Desc())
	case SortPositionAsc:
		order = append(order, goqu.I("v.sorting_position").Asc())
	case SortPositionDesc:
		order = append(order, goqu.I("v.sorting_position").Desc())
	case SortRoomAscPositionAsc:
		join = &LeftJoin{
			Table: tableRooms,
			Cond:  goqu.On(goqu.Ex{"r.id": goqu.I("v.room_id")}),
		}

		order = append(order, goqu.I("r.sorting_position").Asc(), goqu.I("v.sorting_position").Asc())
	}

	var (
		q = goqu.From(tableVisualizations).Select("v.*")
	)

	if join != nil {
		if j, ok := join.(*LeftJoin); ok && j != nil {
			q = q.LeftJoin(j.Table, j.Cond)
		}
	}

	q = q.Where(expr).Limit(uint(limit)).Offset(uint(offset))

	if len(order) > 0 {
		q = q.Order(order...)
	}

	return q.ToSQL()
}

func selectMaxSoringPosition(spec Spec) (string, []interface{}, error) {
	expr, err := specToExpr(spec)
	if err != nil {
		return "", nil, err
	}

	var (
		q = goqu.From(tableVisualizations).Select(goqu.MAX("sorting_position"))
	)

	return q.Where(expr).ToSQL()
}

func countBySpec(spec Spec) (string, []interface{}, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return "", nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return "", nil, err
	}

	return goqu.Select(goqu.COUNT(goqu.Star())).From(tableVisualizations).Where(expr).ToSQL()
}
