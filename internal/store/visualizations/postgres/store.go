package postgres

import (
	"context"
	"time"

	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/uptrace/bun"
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

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Visualization, error) {
	sql, args, err := selectBySpec(spec, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		recs = make([]record, 0)
	)

	if err := s.db.NewRaw(sql, args...).Scan(bunhook.WithQueryContext(ctx, "Visualizations.List"), &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*Visualization, error) {
	res, err := s.List(ctx, spec, SortIDAsc, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrVisualizationNotFound
	}

	return res[0], nil
}

func (s *store) GetMaxSortingPosition(ctx context.Context, spec Spec) (int, error) {
	var (
		num int
	)

	sql, args, err := selectMaxSoringPosition(spec)
	if err != nil {
		return 0, err
	}

	if err := s.db.NewRaw(sql, args...).
		Scan(bunhook.WithQueryContext(ctx, "Visualizations.GetMaxSortingPosition"), &num); err != nil {
		return 0, err
	}

	return num, err
}

func (s *store) Count(ctx context.Context, spec Spec) (int, error) {
	sql, args, err := countBySpec(spec)
	if err != nil {
		return 0, err
	}

	var (
		c int
	)

	if err = s.db.NewRaw(sql, args).Scan(bunhook.WithQueryContext(ctx, "Visualizations.Count"), &c); err != nil {
		return 0, err
	}

	return c, nil
}

func (s *store) Save(ctx context.Context, visualizations ...*Visualization) error {
	var (
		recs = toRecords(visualizations)
	)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(bunhook.WithQueryContext(ctx, "Visualizations.Save"))

	return err
}

func (s *store) Delete(ctx context.Context, visualizations ...*Visualization) error {
	var (
		recs = toRecords(visualizations)
	)

	_, err := s.db.NewDelete().Model(&recs).WherePK().Exec(bunhook.WithQueryContext(ctx, "Visualizations.Delete"))

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.visualizations,alias:v"`

	ID              string     `bun:"id,pk"`
	Name            string     `bun:"name"`
	Description     string     `bun:"description"`
	Version         int        `bun:"version"`
	Status          string     `bun:"status"`
	SortingPosition int        `bun:"sorting_position"`
	CreatedAt       time.Time  `bun:"created_at"`
	ModifiedAt      time.Time  `bun:"modified_at"`
	DeletedAt       *time.Time `bun:"deleted_at"`
	ProjectID       string     `bun:"project_id"`
	FileID          string     `bun:"file_id"`
	RoomID          *string    `bun:"room_id"`
}

func toRecord(val *Visualization) record {
	return record{
		ID:              val.ID,
		Name:            val.Name,
		Description:     val.Description,
		Version:         val.Version,
		Status:          string(val.Status),
		SortingPosition: val.SortingPosition,
		CreatedAt:       val.CreatedAt,
		ModifiedAt:      val.ModifiedAt,
		DeletedAt:       val.DeletedAt,
		ProjectID:       val.ProjectID,
		FileID:          val.FileID,
		RoomID:          val.RoomID,
	}
}

func toRecords(vals []*Visualization) []record {
	var (
		res = make([]record, len(vals))
	)

	for i, v := range vals {
		res[i] = toRecord(v)
	}

	return res
}

func fromRecords(records []record) []*Visualization {
	var (
		res = make([]*Visualization, len(records))
	)

	for i, r := range records {
		res[i] = &Visualization{
			ID:              r.ID,
			Name:            r.Name,
			Description:     r.Description,
			Version:         r.Version,
			Status:          VisualizationStatus(r.Status),
			SortingPosition: r.SortingPosition,
			CreatedAt:       r.CreatedAt,
			ModifiedAt:      r.ModifiedAt,
			DeletedAt:       r.DeletedAt,
			ProjectID:       r.ProjectID,
			FileID:          r.FileID,
			RoomID:          r.RoomID,
		}
	}

	return res
}
