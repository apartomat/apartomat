package postgres

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	roomsTableName = `apartomat.rooms`
)

type store struct {
	db *pg.DB
}

func NewStore(db *pg.DB) *store {
	return &store{db}
}

var (
	_ Store = (*store)(nil)
)

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]*Room, error) {
	qs, err := toRoomSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	orderExpr := goqu.I("created_at").Asc()

	sql, args, err := goqu.From(roomsTableName).
		Where(expr).
		Order(orderExpr).
		Limit(uint(limit)).
		Offset(uint(offset)).
		ToSQL()
	if err != nil {
		return nil, err
	}

	recs := make([]*record, 0)

	_, err = s.db.QueryContext(ctx, &recs, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRecord(recs), nil
}

func (s *store) Save(ctx context.Context, rooms ...*Room) error {
	recs := toRecords(rooms)

	_, err := s.db.ModelContext(ctx, &recs).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()

	return err
}

func (s *store) Delete(ctx context.Context, rooms ...*Room) error {
	var (
		ids = make([]string, len(rooms))
	)

	for i, r := range rooms {
		ids[i] = r.ID
	}

	_, err := s.db.ModelContext(ctx, (*record)(nil)).Where(`id IN (?)`, pg.In(ids)).Delete()

	return err
}

type record struct {
	tableName  struct{}  `pg:"apartomat.rooms"`
	ID         string    `pg:"id,pk"`
	Name       string    `pg:"name"`
	Square     *float64  `pg:"square"`
	Level      *int      `pg:"level"`
	CreatedAt  time.Time `pg:"created_at"`
	ModifiedAt time.Time `pg:"modified_at"`
	HouseID    string    `pg:"house_id"`
}

func toRecord(room *Room) *record {
	return &record{
		ID:         room.ID,
		Name:       room.Name,
		Square:     room.Square,
		Level:      room.Level,
		CreatedAt:  room.CreatedAt,
		ModifiedAt: room.ModifiedAt,
		HouseID:    room.HouseID,
	}
}

func toRecords(rooms []*Room) []*record {
	var (
		res = make([]*record, len(rooms))
	)

	for i, p := range rooms {
		res[i] = toRecord(p)
	}

	return res
}

func fromRecord(records []*record) []*Room {
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
