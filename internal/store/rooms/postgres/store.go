package postgres

import (
	"context"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/uptrace/bun"
	"time"
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

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]*Room, error) {
	sql, args, err := selectBySpec(`apartomat.rooms`, spec, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		recs = make([]record, 0)
	)

	if err := s.db.NewRaw(sql, args...).Scan(bunhook.WithQueryContext(ctx, "Rooms.List"), &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Save(ctx context.Context, rooms ...*Room) error {
	var (
		recs = toRecords(rooms)
	)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(bunhook.WithQueryContext(ctx, "Rooms.Save"))

	return err
}

func (s *store) Delete(ctx context.Context, rooms ...*Room) error {
	var (
		recs = toRecords(rooms)
	)

	_, err := s.db.NewDelete().Model(&recs).WherePK().Exec(bunhook.WithQueryContext(ctx, "Rooms.Delete"))

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.rooms,alias:r"`

	ID         string    `pg:"id,pk"`
	Name       string    `pg:"name"`
	Square     *float64  `pg:"square"`
	Level      *int      `pg:"level"`
	CreatedAt  time.Time `pg:"created_at"`
	ModifiedAt time.Time `pg:"modified_at"`
	HouseID    string    `pg:"house_id"`
}

func toRecord(val *Room) record {
	return record{
		ID:         val.ID,
		Name:       val.Name,
		Square:     val.Square,
		Level:      val.Level,
		CreatedAt:  val.CreatedAt,
		ModifiedAt: val.ModifiedAt,
		HouseID:    val.HouseID,
	}
}

func toRecords(vals []*Room) []record {
	var (
		res = make([]record, len(vals))
	)

	for i, p := range vals {
		res[i] = toRecord(p)
	}

	return res
}

func fromRecords(records []record) []*Room {
	var (
		res = make([]*Room, len(records))
	)

	for i, r := range records {
		res[i] = &Room{
			ID:         r.ID,
			Name:       r.Name,
			Square:     r.Square,
			Level:      r.Level,
			CreatedAt:  r.CreatedAt,
			ModifiedAt: r.ModifiedAt,
			HouseID:    r.HouseID,
		}
	}

	return res
}