package postgres

import (
	"context"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
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

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]*Visualization, error) {
	sql, args, err := selectBySpec(`apartomat.visualizations`, spec, limit, offset)
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

	ID          string     `pg:"id,pk"`
	Name        string     `pg:"name"`
	Description string     `pg:"description"`
	Version     int        `pg:"version"`
	Status      string     `pg:"status"`
	CreatedAt   time.Time  `pg:"created_at"`
	ModifiedAt  time.Time  `pg:"modified_at"`
	DeletedAt   *time.Time `pg:"deleted_at"`
	ProjectID   string     `pg:"project_id"`
	FileID      string     `pg:"file_id"`
	RoomID      *string    `pg:"room_id"`
}

func toRecord(val *Visualization) record {
	return record{
		ID:          val.ID,
		Name:        val.Name,
		Description: val.Description,
		Version:     val.Version,
		Status:      string(val.Status),
		CreatedAt:   val.CreatedAt,
		ModifiedAt:  val.ModifiedAt,
		DeletedAt:   val.DeletedAt,
		ProjectID:   val.ProjectID,
		FileID:      val.FileID,
		RoomID:      val.RoomID,
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
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Version:     r.Version,
			Status:      VisualizationStatus(r.Status),
			CreatedAt:   r.CreatedAt,
			ModifiedAt:  r.ModifiedAt,
			DeletedAt:   r.DeletedAt,
			ProjectID:   r.ProjectID,
			FileID:      r.FileID,
			RoomID:      r.RoomID,
		}
	}

	return res
}
