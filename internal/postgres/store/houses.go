package store

import (
	"context"
	"errors"
	. "github.com/apartomat/apartomat/internal/store/houses"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
)

const (
	housesTableName = `apartomat.houses`
)

type housesStore struct {
	db *pg.DB
}

func NewHousesStore(db *pg.DB) *housesStore {
	return &housesStore{db}
}

var (
	_ Store = (*housesStore)(nil)
)

func (s *housesStore) Save(context.Context, *House) (*House, error) {
	return nil, errors.New("HousesStore.Save not implemented yet")
}

func (s *housesStore) List(ctx context.Context, spec Spec, limit, offset int) ([]*House, error) {
	qs, err := toHouseSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(housesTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	contacts := make([]*House, 0)

	_, err = s.db.QueryContext(ctx, &contacts, sql, args...)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

type houseSpecQuery interface {
	Expression() (goqu.Expression, error)
}

func toHouseSpecQuery(spec Spec) (houseSpecQuery, error) {
	if s, ok := spec.(houseSpecQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return houseIDInSpecQuery{s}, nil
	case ProjectIDInSpec:
		return houseProjectIDInSpecQuery{s}, nil
	}

	return nil, errors.New("unknown spec")
}

type houseIDInSpecQuery struct {
	spec IDInSpec
}

func (s houseIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.IDs}, nil
}

//
type houseProjectIDInSpecQuery struct {
	spec ProjectIDInSpec
}

func (s houseProjectIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"project_id": s.spec.IDs}, nil
}

