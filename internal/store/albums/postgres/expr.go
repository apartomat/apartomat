package postgres

import (
	"errors"
	. "github.com/apartomat/apartomat/internal/store/albums"
	"github.com/doug-martin/goqu/v9"
)

type specQuery interface {
	Expression() (goqu.Expression, error)
}

func toQuery(spec Spec) (specQuery, error) {
	if s, ok := spec.(specQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return idInSpecQuery{s}, nil
	case ProjectIDInSpec:
		return projectIDInSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown albums spec")
}

type andSpecQuery struct {
	spec AndSpec
}

func (s andSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toQuery(spec); err != nil {
			return nil, err
		} else {
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
		if ps, err := toQuery(spec); err != nil {
			return nil, err
		} else {
			expr, err := ps.Expression()
			if err != nil {
				return nil, err
			}

			exs = append(exs, expr)
		}
	}

	return goqu.Or(exs...), nil
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
