package postgres

import (
	"errors"
	. "github.com/apartomat/apartomat/internal/store/workspace_users"
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
	case UserIDInSpec:
		return userIDInSpecQuery{s}, nil
	case WorkspaceIDInSpec:
		return workspaceIDInSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown workspace users spec")
}

type idInSpecQuery struct {
	spec IDInSpec
}

func (s idInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.ID}, nil
}

type workspaceIDInSpecQuery struct {
	spec WorkspaceIDInSpec
}

func (s workspaceIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"workspace_id": s.spec.WorkspaceID}, nil
}

type userIDInSpecQuery struct {
	spec UserIDInSpec
}

func (s userIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"user_id": s.spec.UserID}, nil
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
		order = append(order, goqu.I("wu.id").Asc())
	}

	var (
		q = goqu.From(goqu.T("workspace_users").Schema("apartomat").As("wu")).Select("wu.*")
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
