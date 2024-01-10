package postgres

import (
	"errors"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
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

func selectBySpec(tablename string, spec Spec, limit, offset int) (string, []interface{}, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return "", nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return "", nil, err
	}

	return goqu.From(tablename).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
}
