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

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Room, error) {
	sql, args, err := selectBySpec(`apartomat.rooms`, spec, sort, limit, offset)
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

func (s *store) Get(ctx context.Context, spec Spec) (*Room, error) {
	res, err := s.List(ctx, spec, SortIDAsc, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrRoomNotFound
	}

	return res[0], nil
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

func (s *store) Reorder(ctx context.Context, houseID string, asc bool) error {
	var (
		sort = "DESC"
	)

	if asc {
		sort = "ASC"
	}

	_, err := s.db.NewRaw(`
UPDATE apartomat.rooms r1
	SET
	sorting_position = r2.pos,
	modified_at = now()
FROM (
	SELECT
		r2.*,
		row_number() OVER (ORDER BY sorting_position, modified_at ?) AS pos
	FROM apartomat.rooms r2 WHERE r2.house_id = ?
) r2 WHERE r1.id = r2.id`, bun.Safe(sort), houseID).Exec(bunhook.WithQueryContext(ctx, "Rooms.Reorder"))

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.rooms,alias:r"`

	ID              string    `bun:"id,pk"`
	Name            string    `bun:"name"`
	Square          *float64  `bun:"square"`
	Level           *int      `bun:"level"`
	SortingPosition int       `bun:"sorting_position"`
	CreatedAt       time.Time `bun:"created_at"`
	ModifiedAt      time.Time `bun:"modified_at"`
	HouseID         string    `bun:"house_id"`
}

func toRecord(val *Room) record {
	return record{
		ID:              val.ID,
		Name:            val.Name,
		Square:          val.Square,
		Level:           val.Level,
		SortingPosition: val.SortingPosition,
		CreatedAt:       val.CreatedAt,
		ModifiedAt:      val.ModifiedAt,
		HouseID:         val.HouseID,
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
			ID:              r.ID,
			Name:            r.Name,
			Square:          r.Square,
			Level:           r.Level,
			SortingPosition: r.SortingPosition,
			CreatedAt:       r.CreatedAt,
			ModifiedAt:      r.ModifiedAt,
			HouseID:         r.HouseID,
		}
	}

	return res
}
