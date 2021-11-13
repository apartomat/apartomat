package store

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	workspaceTableName = `apartomat.workspaces`
)

type workspaceStore struct {
	pg *pgxpool.Pool
}

func NewWorkspaceStore(pg *pgxpool.Pool) *workspaceStore {
	return &workspaceStore{pg}
}

var (
	_ store.WorkspaceStore = (*workspaceStore)(nil)
)

func (s *workspaceStore) Save(ctx context.Context, workspace *store.Workspace) (*store.Workspace, error) {
	if workspace.CreatedAt.IsZero() {
		workspace.CreatedAt = time.Now()
	}

	if workspace.ModifiedAt.IsZero() {
		workspace.ModifiedAt = workspace.CreatedAt
	}

	q, args, err := InsertIntoWorkspaces().
		Columns("name", "is_active", "user_id", "created_at", "modified_at").
		Values(workspace.Name, workspace.IsActive, workspace.UserID, workspace.CreatedAt, workspace.ModifiedAt).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = s.pg.QueryRow(ctx, q, args...).Scan(&workspace.ID)
	if err != nil {
		return nil, err
	}

	return workspace, err
}

func (s *workspaceStore) List(ctx context.Context, q store.WorkspaceStoreQuery) ([]*store.Workspace, error) {
	sql, args, err := SelectFromWorkspaces(
		"id",
		"name",
		"is_active",
		"user_id",
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

	workspaces := make([]*store.Workspace, 0)

	for rows.Next() {
		workspace := new(store.Workspace)

		err := rows.Scan(
			&workspace.ID,
			&workspace.Name,
			&workspace.IsActive,
			&workspace.UserID,
			&workspace.CreatedAt,
			&workspace.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}

		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}

// workspaceInsertBuilder is insert builder shortcut for workspaces table
type workspaceInsertBuilder struct {
	sq.InsertBuilder
}

func InsertIntoWorkspaces() *workspaceInsertBuilder {
	return &workspaceInsertBuilder{
		sq.Insert(workspaceTableName).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id"),
	}
}

// workspaceSelectBuilder is insert builder shortcut for workspaces table
type workspaceSelectBuilder struct {
	sq.SelectBuilder
}

func SelectFromWorkspaces(columns ...string) *workspaceSelectBuilder {
	return &workspaceSelectBuilder{
		sq.Select(columns...).
			From(workspaceTableName).
			PlaceholderFormat(sq.Dollar),
	}
}

func (builder *workspaceSelectBuilder) Where(q store.WorkspaceStoreQuery) *workspaceSelectBuilder {
	if len(q.ID.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{`"id"`: q.ID.Eq})
	}

	if len(q.UserID.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{"user_id": q.UserID.Eq})
	}

	return builder
}

func (builder *workspaceSelectBuilder) Limit(n int) *workspaceSelectBuilder {
	if n != 0 {
		builder.SelectBuilder = builder.SelectBuilder.Limit(uint64(n))
	}

	return builder
}
