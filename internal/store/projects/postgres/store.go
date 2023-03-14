package postgres

import (
	"context"
	"github.com/apartomat/apartomat/internal/postgres"
	. "github.com/apartomat/apartomat/internal/store/projects"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	projectsTableName = `apartomat.projects`
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

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]*Project, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(projectsTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	rows := make([]*record, 0)

	_, err = s.db.QueryContext(postgres.WithQueryContext(ctx, "projects.List"), &rows, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRecords(rows), nil
}

func (s *store) Save(ctx context.Context, projects ...*Project) error {
	recs := toRecords(projects)

	_, err := s.db.ModelContext(postgres.WithQueryContext(ctx, "projects.Save"), &recs).
		Returning("NULL").
		OnConflict("(id) DO UPDATE").
		Insert()

	return err
}

type record struct {
	tableName   struct{}   `pg:"apartomat.projects"`
	ID          string     `pg:"id,pk"`
	Name        string     `pg:"name"`
	Status      string     `pg:"status"`
	StartAt     *time.Time `pg:"start_at"`
	EndAt       *time.Time `pg:"end_at"`
	CreatedAt   time.Time  `pg:"created_at"`
	ModifiedAt  time.Time  `pg:"modified_at"`
	WorkspaceID string     `pg:"workspace_id"`
}

func toRecord(project *Project) *record {
	return &record{
		ID:          project.ID,
		Name:        project.Name,
		Status:      string(project.Status),
		StartAt:     project.StartAt,
		EndAt:       project.EndAt,
		CreatedAt:   project.CreatedAt,
		ModifiedAt:  project.ModifiedAt,
		WorkspaceID: project.WorkspaceID,
	}
}

func toRecords(projects []*Project) []*record {
	var (
		res = make([]*record, len(projects))
	)

	for i, p := range projects {
		res[i] = toRecord(p)
	}

	return res
}

func fromRecords(records []*record) []*Project {
	visualizations := make([]*Project, len(records))

	for i, r := range records {
		visualizations[i] = &Project{
			ID:          r.ID,
			Name:        r.Name,
			Status:      Status(r.Status),
			StartAt:     r.StartAt,
			EndAt:       r.EndAt,
			CreatedAt:   r.CreatedAt,
			ModifiedAt:  r.ModifiedAt,
			WorkspaceID: r.WorkspaceID,
		}
	}

	return visualizations
}
