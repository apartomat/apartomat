package postgres

import (
	"errors"

	. "github.com/apartomat/apartomat/internal/store/users"
	"github.com/doug-martin/goqu/v9"
)

type query interface {
	Expression() (goqu.Expression, error)
}

func toQuery(spec Spec) (query, error) {
	if s, ok := spec.(query); ok {
		return s, nil
	}

	if spec == nil {
		return allSpecQuery{}, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return idInSpecQuery{s}, nil
	case EmailInSpec:
		return emailInSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown user spec")
}

type allSpecQuery struct{}

func (s allSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{}, nil
}

type idInSpecQuery struct {
	spec IDInSpec
}

func (s idInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.ID}, nil
}

type emailInSpecQuery struct {
	spec EmailInSpec
}

func (s emailInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"email": s.spec.Email}, nil
}

type andSpecQuery struct {
	spec AndSpec
}

func (s andSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toQuery(spec); err != nil {
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
		if ps, err := toQuery(spec); err != nil {
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
