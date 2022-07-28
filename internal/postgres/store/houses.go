package store

import (
	"context"
	"errors"
	. "github.com/apartomat/apartomat/internal/store/houses"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
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

func (s *housesStore) Save(ctx context.Context, house *House) (*House, error) {
	if house.CreatedAt.IsZero() {
		house.CreatedAt = time.Now()
	}

	if house.ModifiedAt.IsZero() {
		house.ModifiedAt = house.CreatedAt
	}

	rec := toHousesRecord(house)

	_, err := s.db.ModelContext(ctx, rec).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		return nil, err
	}

	return house, nil
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

	houses := make([]*housesRecord, 0)

	_, err = s.db.QueryContext(ctx, &houses, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromHousesRecords(houses), nil
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

type housesRecord struct {
	tableName      struct{}  `pg:"apartomat.houses,alias:houses"`
	ID             string    `pg:"id,pk"`
	City           string    `pg:"city,use_zero"`
	Address        string    `pg:"address,use_zero"`
	HousingComplex string    `pg:"housing_complex,use_zero"`
	CreatedAt      time.Time `pg:"created_at"`
	ModifiedAt     time.Time `pg:"modified_at"`
	ProjectID      string    `pg:"project_id"`
}

func toHousesRecord(house *House) *housesRecord {
	return &housesRecord{
		ID:             house.ID,
		City:           house.City,
		Address:        house.Address,
		HousingComplex: house.HousingComplex,
		CreatedAt:      house.CreatedAt,
		ModifiedAt:     house.ModifiedAt,
		ProjectID:      house.ProjectID,
	}
}

func fromHousesRecords(records []*housesRecord) []*House {
	houses := make([]*House, len(records))

	for i, r := range records {
		houses[i] = &House{
			ID:             r.ID,
			City:           r.City,
			Address:        r.Address,
			HousingComplex: r.HousingComplex,
			CreatedAt:      r.CreatedAt,
			ModifiedAt:     r.ModifiedAt,
			ProjectID:      r.ProjectID,
		}
	}

	return houses
}
