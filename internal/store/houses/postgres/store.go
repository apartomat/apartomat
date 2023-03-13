package postgres

import (
	"context"
	"time"

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

func NewStore(db *pg.DB) *housesStore {
	return &housesStore{db}
}

var (
	_ Store = (*housesStore)(nil)
)

func (s *housesStore) Save(ctx context.Context, houses ...*House) error {
	recs := toRecords(houses)

	_, err := s.db.ModelContext(ctx, &recs).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()

	return err
}

func (s *housesStore) List(ctx context.Context, spec Spec, limit, offset int) ([]*House, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	orderExpr := goqu.I("created_at").Asc()

	sql, args, err := goqu.From(housesTableName).
		Where(expr).
		Order(orderExpr).
		Limit(uint(limit)).
		Offset(uint(offset)).
		ToSQL()
	if err != nil {
		return nil, err
	}

	houses := make([]*record, 0)

	_, err = s.db.QueryContext(ctx, &houses, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRecords(houses), nil
}

type record struct {
	tableName      struct{}  `pg:"apartomat.houses"`
	ID             string    `pg:"id,pk"`
	City           string    `pg:"city,use_zero"`
	Address        string    `pg:"address,use_zero"`
	HousingComplex string    `pg:"housing_complex,use_zero"`
	CreatedAt      time.Time `pg:"created_at"`
	ModifiedAt     time.Time `pg:"modified_at"`
	ProjectID      string    `pg:"project_id"`
}

func toRecord(house *House) *record {
	return &record{
		ID:             house.ID,
		City:           house.City,
		Address:        house.Address,
		HousingComplex: house.HousingComplex,
		CreatedAt:      house.CreatedAt,
		ModifiedAt:     house.ModifiedAt,
		ProjectID:      house.ProjectID,
	}
}

func toRecords(houses []*House) []*record {
	var (
		res = make([]*record, len(houses))
	)

	for i, p := range houses {
		res[i] = toRecord(p)
	}

	return res
}

func fromRecords(records []*record) []*House {
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
