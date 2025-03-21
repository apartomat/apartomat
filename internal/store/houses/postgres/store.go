package postgres

import (
	"context"
	"time"

	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/houses"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/uptrace/bun"
)

const (
	housesTableName = `apartomat.houses`
)

type store struct {
	db *bun.DB
}

func NewStore(db *bun.DB) *store {
	return &store{db}
}

var (
	_ Store = (*store)(nil)
)

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*House, error) {
	sql, args, err := selectBySpec(housesTableName, spec, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		recs = make([]*record, 0)
	)

	if err := s.db.NewRaw(sql, args...).
		Scan(bunhook.WithQueryContext(ctx, "Houses.List"), &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*House, error) {
	res, err := s.List(ctx, spec, SortDefault, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrHouseNotFound
	}

	return res[0], nil
}

func (s *store) Save(ctx context.Context, houses ...*House) error {
	recs := toRecords(houses)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(bunhook.WithQueryContext(ctx, "Houses.Save"))

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.houses,alias:h"`

	ID             string    `bun:"id,pk"`
	City           string    `bun:"city"`
	Address        string    `bun:"address"`
	HousingComplex string    `bun:"housing_complex"`
	CreatedAt      time.Time `bun:"created_at"`
	ModifiedAt     time.Time `bun:"modified_at"`
	ProjectID      string    `bun:"project_id"`
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
		order = make([]exp.OrderedExpression, 0)
	)

	switch sort {
	case SortDefault:
		//
	}

	var (
		q = goqu.From(tableName).Where(expr).Limit(uint(limit)).Offset(uint(offset))
	)

	if len(order) > 0 {
		q = q.Order(order...)
	}

	return q.ToSQL()
}
