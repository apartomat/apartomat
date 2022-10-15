package postgres

import (
	. "github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
)

type visualizationSpecQuery interface {
	Expression() (goqu.Expression, error)
}

func toVisualizationSpecQuery(spec Spec) (visualizationSpecQuery, error) {
	if s, ok := spec.(visualizationSpecQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return visualizationIDInSpecQuery{s}, nil
	case ProjectIDInSpec:
		return visualizationProjectIDInSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown spec")
}

type visualizationIDInSpecQuery struct {
	spec IDInSpec
}

func (s visualizationIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.ID}, nil
}

type visualizationProjectIDInSpecQuery struct {
	spec ProjectIDInSpec
}

func (s visualizationProjectIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"project_id": s.spec.ProjectID}, nil
}

//
type andSpecQuery struct {
	spec AndSpec
}

func (s andSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toVisualizationSpecQuery(spec); err != nil {
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

//
type orSpecQuery struct {
	spec OrSpec
}

func (s orSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toVisualizationSpecQuery(spec); err != nil {
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
