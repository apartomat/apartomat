package store

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	workspaceUserTableName = `apartomat.workspace_users`
)

type workspaceUserStore struct {
	pg *pgxpool.Pool
}

func NewWorkspaceUserStore(pg *pgxpool.Pool) *workspaceUserStore {
	return &workspaceUserStore{pg}
}

var (
	_ store.WorkspaceUserStore = (*workspaceUserStore)(nil)
)

func (s *workspaceUserStore) Save(ctx context.Context, user *store.WorkspaceUser) (*store.WorkspaceUser, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	if user.ModifiedAt.IsZero() {
		user.ModifiedAt = user.CreatedAt
	}

	q, args, err := InsertIntoWorkspaceUsers().
		Columns("id", "workspace_id", "user_id", "role", "created_at", "modified_at").
		Values(user.ID, user.WorkspaceID, user.UserID, user.Role, user.CreatedAt, user.ModifiedAt).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = s.pg.QueryRow(ctx, q, args...).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *workspaceUserStore) List(ctx context.Context, q store.WorkspaceUserStoreQuery) ([]*store.WorkspaceUser, error) {
	sql, args, err := SelectFromWorkspaceUsers(
		"id",
		"workspace_id",
		"user_id",
		"role",
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

	users := make([]*store.WorkspaceUser, 0)

	for rows.Next() {
		user := new(store.WorkspaceUser)

		err := rows.Scan(
			&user.ID,
			&user.WorkspaceID,
			&user.UserID,
			&user.Role,
			&user.CreatedAt,
			&user.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// workspaceInsertBuilder is insert builder shortcut for "workspace_users" table
type workspaceUserInsertBuilder struct {
	sq.InsertBuilder
}

func InsertIntoWorkspaceUsers() *workspaceUserInsertBuilder {
	return &workspaceUserInsertBuilder{
		sq.Insert(workspaceUserTableName).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id"),
	}
}

// workspaceUserSelectBuilder is insert builder shortcut for workspaces table
type workspaceUserSelectBuilder struct {
	sq.SelectBuilder
}

func SelectFromWorkspaceUsers(columns ...string) *workspaceUserSelectBuilder {
	return &workspaceUserSelectBuilder{
		sq.Select(columns...).
			From(workspaceUserTableName).
			PlaceholderFormat(sq.Dollar),
	}
}

func (builder *workspaceUserSelectBuilder) Where(q store.WorkspaceUserStoreQuery) *workspaceUserSelectBuilder {
	if len(q.WorkspaceID.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{`"workspace_id"`: q.WorkspaceID.Eq})
	}

	if len(q.UserID.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{"user_id": q.UserID.Eq})
	}

	return builder
}

func (builder *workspaceUserSelectBuilder) Limit(n int) *workspaceUserSelectBuilder {
	if n != 0 {
		builder.SelectBuilder = builder.SelectBuilder.Limit(uint64(n))
	}

	return builder
}
