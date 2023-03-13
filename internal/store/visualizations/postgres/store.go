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

func (s *store) Save(ctx context.Context, visualizations ...*Visualization) error {
	recs := toVisualizationsRecords(visualizations)

	_, err := s.db.ModelContext(ctx, &recs).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()

	return err
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

	return fromVisualizationsRecords(rows), nil
}

func (s *store) Delete(ctx context.Context, visualization *Visualization) error {
	_, err := s.db.ModelContext(ctx, (*visualizationsRecord)(nil)).Where(`id = ?`, visualization.ID).Delete()
	if err != nil {
		return err
	}

	return nil
}

type visualizationsRecord struct {
	tableName   struct{}   `pg:"apartomat.visualizations"`
	ID          string     `pg:"id,pk"`
	Name        string     `pg:"name,use_zero"`
	Description string     `pg:"description,use_zero"`
	Version     int        `pg:"version,use_zero"`
	Status      string     `pg:"status"`
	CreatedAt   time.Time  `pg:"created_at"`
	ModifiedAt  time.Time  `pg:"modified_at"`
	DeletedAt   *time.Time `pg:"deleted_at"`
	ProjectID   string     `pg:"project_id"`
	FileID      string     `pg:"file_id"`
	RoomID      *string    `pg:"room_id"`
}

func toVisualizationsRecord(visualization *Visualization) *visualizationsRecord {
	return &visualizationsRecord{
		ID:          visualization.ID,
		Name:        visualization.Name,
		Description: visualization.Description,
		Version:     visualization.Version,
		Status:      string(visualization.Status),
		CreatedAt:   visualization.CreatedAt,
		ModifiedAt:  visualization.ModifiedAt,
		DeletedAt:   visualization.DeletedAt,
		ProjectID:   visualization.ProjectID,
		FileID:      visualization.FileID,
		RoomID:      visualization.RoomID,
	}
}

func toVisualizationsRecords(visualizations []*Visualization) []*visualizationsRecord {
	var (
		res = make([]*visualizationsRecord, len(visualizations))
	)

	for i, v := range visualizations {
		res[i] = toVisualizationsRecord(v)
	}

	return res
}

func fromVisualizationsRecords(records []*visualizationsRecord) []*Visualization {
	visualizations := make([]*Visualization, len(records))

	for i, r := range records {
		visualizations[i] = &Visualization{
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

	return visualizations
}
