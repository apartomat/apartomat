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
	if spec == nil {
		return nil, nil
	}

	if s, ok := spec.(specQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return houseIDInSpecQuery{s}, nil
	case ProjectIDInSpec:
		return houseProjectIDInSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown spec")
}

type houseIDInSpecQuery struct {
	spec IDInSpec
}

func (s houseIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.IDs}, nil
}

type houseProjectIDInSpecQuery struct {
	spec ProjectIDInSpec
}

func (s houseProjectIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"project_id": s.spec.IDs}, nil
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
