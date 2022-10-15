package postgres

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	visualizationsTableName = `apartomat.visualizations`
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

func (s *store) Save(ctx context.Context, visualization *Visualization) (*Visualization, error) {
	if visualization.CreatedAt.IsZero() {
		visualization.CreatedAt = time.Now()
	}

	if visualization.ModifiedAt.IsZero() {
		visualization.ModifiedAt = visualization.CreatedAt
	}

	rec := toVisualizationsRecord(visualization)

	_, err := s.db.ModelContext(ctx, rec).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()
	if err != nil {
		return nil, err
	}

	return visualization, nil
}

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]*Visualization, error) {
	qs, err := toVisualizationSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(visualizationsTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	rows := make([]*visualizationsRecord, 0)

	_, err = s.db.QueryContext(ctx, &rows, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromHousesRecords(rows), nil
}

func (s *store) Delete(ctx context.Context, visualization *Visualization) error {
	_, err := s.db.ModelContext(ctx, (*visualizationsRecord)(nil)).Where(`id = ?`, visualization.ID).Delete()
	if err != nil {
		return err
	}

	return nil
}

type visualizationsRecord struct {
	tableName     struct{}  `pg:"apartomat.visualizations,alias:visualizations"`
	ID            string    `pg:"id,pk"`
	Name          string    `pg:"name,use_zero"`
	Description   string    `pg:"description,use_zero"`
	Version       int       `pg:"version,use_zero"`
	CreatedAt     time.Time `pg:"created_at"`
	ModifiedAt    time.Time `pg:"modified_at"`
	ProjectID     string    `pg:"project_id"`
	ProjectFileID string    `pg:"project_file_id"`
	RoomID        *string   `pg:"room_id"`
}

func toVisualizationsRecord(visualization *Visualization) *visualizationsRecord {
	return &visualizationsRecord{
		ID:            visualization.ID,
		Name:          visualization.Name,
		Description:   visualization.Description,
		Version:       visualization.Version,
		CreatedAt:     visualization.CreatedAt,
		ModifiedAt:    visualization.ModifiedAt,
		ProjectID:     visualization.ProjectID,
		ProjectFileID: visualization.ProjectFileID,
		RoomID:        visualization.RoomID,
	}
}

func fromHousesRecords(records []*visualizationsRecord) []*Visualization {
	visualizations := make([]*Visualization, len(records))

	for i, r := range records {
		visualizations[i] = &Visualization{
			ID:            r.ID,
			Name:          r.Name,
			Description:   r.Description,
			Version:       r.Version,
			CreatedAt:     r.CreatedAt,
			ModifiedAt:    r.ModifiedAt,
			ProjectID:     r.ProjectID,
			ProjectFileID: r.ProjectFileID,
			RoomID:        r.RoomID,
		}
	}

	return visualizations
}
