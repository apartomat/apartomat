package postgres

import (
	"errors"
	. "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type specQuery interface {
	Expression() (goqu.Expression, error)
}

func toSpecQuery(spec Spec) (specQuery, error) {
	if s, ok := spec.(specQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return idInSpecQuery{s}, nil
	case AlbumIDInSpec:
		return albumIdInSpecQuery{s}, nil
	case VersionInSpec:
		return versionInSpecQuery{s}, nil
	case VersionGteSpec:
		return versionGteSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown album files spec")
}

type andSpecQuery struct {
	spec AndSpec
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
	spec OrSpec
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

//

type idInSpecQuery struct {
	spec IDInSpec
}

func (s idInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.ID}, nil
}

//

type albumIdInSpecQuery struct {
	spec AlbumIDInSpec
}

func (s albumIdInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"album_id": s.spec.AlbumID}, nil
}

//

type versionInSpecQuery struct {
	spec VersionInSpec
}

func (s versionInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"version": s.spec.Version}, nil
}

//

type versionGteSpecQuery struct {
	spec VersionGteSpec
}

func (s versionGteSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"version": s.spec.Version}, nil
}

//

func selectBySpec(tableName string, spec Spec, sort Sort, limit, offset int) (string, []interface{}, error) {
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
		order = goqu.I("id").Asc()
	case SortVersionAsc:
		order = goqu.I("version").Asc()
	case SortVersionDesc:
		order = goqu.I("version").Desc()
	}

	return goqu.From(tableName).Where(expr).Limit(uint(limit)).Order(order).Offset(uint(offset)).ToSQL()
}

func countBySpec(tableName string, spec Spec) (string, []interface{}, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return "", nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return "", nil, err
	}

	return goqu.Select(goqu.COUNT(goqu.Star())).From(tableName).Where(expr).ToSQL()
}
