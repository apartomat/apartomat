package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	projectTableName = `apartomat.projects`
)

type projectStore struct {
	pg *pgxpool.Pool
}

func NewProjectStore(pg *pgxpool.Pool) *projectStore {
	return &projectStore{pg}
}

var (
	_ store.ProjectStore = (*projectStore)(nil)
)

func (s *projectStore) Save(ctx context.Context, project *store.Project) (*store.Project, error) {
	if project.CreatedAt.IsZero() {
		project.CreatedAt = time.Now()
	}

	if project.ModifiedAt.IsZero() {
		project.ModifiedAt = project.CreatedAt
	}

	q, args, err := InsertIntoProjects().
		Columns(
			"name",
			"is_active",
			"workspace_id",
			"start_at",
			"end_at",
			"created_at",
			"modified_at",
		).
		Values(project.Name, project.IsActive, project.WorkspaceID, project.StartAt, project.EndAt, project.CreatedAt, project.ModifiedAt).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = s.pg.QueryRow(ctx, q, args...).Scan(&project.ID)
	if err != nil {
		return nil, err
	}

	return project, err
}

func (s *projectStore) List(ctx context.Context, q store.ProjectStoreQuery) ([]*store.Project, error) {
	sql, args, err := SelectFromProjects(
		"id",
		"name",
		"is_active",
		"workspace_id",
		"start_at",
		"end_at",
		"created_at",
		"modified_at",
	).Where(q).Limit(q.Limit).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.pg.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	projects := make([]*store.Project, 0)

	for rows.Next() {
		project := new(store.Project)

		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.IsActive,
			&project.WorkspaceID,
			&project.StartAt,
			&project.EndAt,
			&project.CreatedAt,
			&project.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

//
type projectInsertBuilder struct {
	sq.InsertBuilder
}

func InsertIntoProjects() *projectInsertBuilder {
	return &projectInsertBuilder{
		sq.Insert(projectTableName).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id"),
	}
}

//
type projectSelectBuilder struct {
	sq.SelectBuilder
}

func SelectFromProjects(columns ...string) *projectSelectBuilder {
	return &projectSelectBuilder{
		sq.Select(columns...).
			From(projectTableName).
			PlaceholderFormat(sq.Dollar),
	}
}

func (builder *projectSelectBuilder) Where(q store.ProjectStoreQuery) *projectSelectBuilder {
	if len(q.ID.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{`"id"`: q.ID.Eq})
	}

	if len(q.WorkspaceID.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{"workspace_id": q.WorkspaceID.Eq})
	}

	return builder
}

func (builder *projectSelectBuilder) Limit(n int) *projectSelectBuilder {
	if n != 0 {
		builder.SelectBuilder = builder.SelectBuilder.Limit(uint64(n))
	}

	return builder
}
