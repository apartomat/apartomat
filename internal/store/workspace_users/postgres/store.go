package postgres

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/workspace_users"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	workspaceUsersTableName = `apartomat.workspace_users`
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

func (s *store) Save(ctx context.Context, users ...*WorkspaceUser) error {
	recs := toRecords(users)

	_, err := s.db.ModelContext(ctx, &recs).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()

	return err
}

func (s *store) List(
	ctx context.Context,
	spec Spec,
	limit,
	offset int,
) ([]*WorkspaceUser, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(workspaceUsersTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	rows := make([]*record, 0)

	_, err = s.db.QueryContext(ctx, &rows, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRecords(rows), nil
}

type record struct {
	tableName   struct{}  `pg:"apartomat.workspace_users"`
	ID          string    `pg:"id,pk"`
	WorkspaceID string    `pg:"workspace_id"`
	UserID      string    `pg:"user_id"`
	Role        string    `pg:"role"`
	CreatedAt   time.Time `pg:"created_at"`
	ModifiedAt  time.Time `pg:"modified_at"`
}

func toRecord(user *WorkspaceUser) *record {
	return &record{
		ID:          user.ID,
		WorkspaceID: user.WorkspaceID,
		UserID:      user.WorkspaceID,
		Role:        string(user.Role),
		CreatedAt:   user.CreatedAt,
		ModifiedAt:  user.ModifiedAt,
	}
}

func toRecords(users []*WorkspaceUser) []*record {
	var (
		res = make([]*record, len(users))
	)

	for i, p := range users {
		res[i] = toRecord(p)
	}

	return res
}

func fromRecords(records []*record) []*WorkspaceUser {
	users := make([]*WorkspaceUser, len(records))

	for i, r := range records {
		users[i] = &WorkspaceUser{
			ID:          r.ID,
			WorkspaceID: r.WorkspaceID,
			UserID:      r.UserID,
			Role:        WorkspaceUserRole(r.Role),
			CreatedAt:   r.CreatedAt,
			ModifiedAt:  r.ModifiedAt,
		}
	}

	return users
}
