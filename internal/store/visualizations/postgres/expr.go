package postgres

import (
	"errors"

	. "github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
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

func selectBySpec(spec Spec, sort Sort, limit, offset int) (string, []interface{}, error) {
	sq, err := toSpecQuery(spec)
	if err != nil {
		return "", nil, err
	}

	expr, err := sq.Expression()
	if err != nil {
		return "", nil, err
	}

	type Join struct {
		Table goqu.Expression
		Cond  exp.JoinCondition
	}

	var (
		order = make([]exp.OrderedExpression, 0)

		join *Join
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
		join = &Join{
			Table: goqu.T("rooms").Schema("apartomat").As("r"),
			Cond:  goqu.On(goqu.Ex{"r.id": goqu.I("v.room_id")}),
		}

		order = append(order, goqu.I("r.sorting_position").Asc(), goqu.I("v.sorting_position").Asc())
	}

	var (
		q = goqu.From(goqu.T("visualizations").Schema("apartomat").As("v")).Select("v.*")
	)

	if join != nil {
		q = q.Join(join.Table, join.Cond)
	}

	q = q.Where(expr).Limit(uint(limit)).Offset(uint(offset))

	if len(order) > 0 {
		q = q.Order(order...)
	}

	return q.ToSQL()
}
