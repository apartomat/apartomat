package postgres

import (
	"errors"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/doug-martin/goqu/v9"
)

type specQuery interface {
	Expression() (goqu.Expression, error)
}

func toSpecQuery(spec contacts.Spec) (specQuery, error) {
	if s, ok := spec.(specQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case contacts.IDInSpec:
		return contactIDInSpecQuery{s}, nil
	case contacts.ProjectIDInSpec:
		return contactProjectIDInSpecQuery{s}, nil
	case contacts.AndSpec:
		return andSpecQuery{spec: s}, nil
	case contacts.OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown contacts spec")
}

type andSpecQuery struct {
	spec contacts.AndSpec
}

func (s andSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toSpecQuery(spec); err != nil {
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
	spec contacts.OrSpec
}

func (s orSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toSpecQuery(spec); err != nil {
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

type contactIDInSpecQuery struct {
	spec contacts.IDInSpec
}

func (s contactIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.IDs}, nil
}

type contactProjectIDInSpecQuery struct {
	spec contacts.ProjectIDInSpec
}

func (s contactProjectIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"project_id": s.spec.IDs}, nil
}
