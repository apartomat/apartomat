package store

import (
	"context"
	"errors"
	. "github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
)

const (
	contactsTableName = `apartomat.contacts`
)

type contactsStore struct {
	db *pg.DB
}

func NewContactsStore(db *pg.DB) *contactsStore {
	return &contactsStore{db}
}

var (
	_ Store = (*contactsStore)(nil)
)

func (s *contactsStore) Save(context.Context, *Contact) (*Contact, error) {
	return nil, errors.New("ContactsStore.Save not implemented yet")

}

func (s *contactsStore) List(ctx context.Context, spec Spec, limit, offset int) ([]*Contact, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(contactsTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	contacts := make([]*Contact, 0)

	_, err = s.db.QueryContext(ctx, &contacts, sql, args...)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

//

type specQuery interface {
	Expression() (goqu.Expression, error)
}

func toSpecQuery(spec Spec) (specQuery, error) {
	if s, ok := spec.(specQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return contactIDInSpecQuery{s}, nil
	case ProjectIDInSpec:
		return contactProjectIDInSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown spec")
}

//
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

//
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
type contactIDInSpecQuery struct {
	spec IDInSpec
}

func (s contactIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.IDs}, nil
}

//
type contactProjectIDInSpecQuery struct {
	spec ProjectIDInSpec
}

func (s contactProjectIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"project_id": s.spec.IDs}, nil
}
