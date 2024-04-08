package postgres

import (
	"context"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/workspace_users"
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

func (s *store) List(
	ctx context.Context,
	spec Spec,
	sort Sort,
	limit,
	offset int,
) ([]*WorkspaceUser, error) {
	sql, args, err := selectBySpec(spec, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		recs = make([]record, 0)
	)

	if err := s.db.NewRaw(sql, args...).Scan(bunhook.WithQueryContext(ctx, "WorkspaceUsers.List"), &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*WorkspaceUser, error) {
	res, err := s.List(ctx, spec, SortDefault, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrWorkspaceUserNotFound
	}

	return res[0], nil
}

func (s *store) Save(ctx context.Context, users ...*WorkspaceUser) error {
	var (
		recs = toRecords(users)
	)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(bunhook.WithQueryContext(ctx, "WorkspaceUsers.Save"))

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.workspace_users,alias:wu"`

	ID          string    `bun:"id,pk"`
	WorkspaceID string    `bun:"workspace_id"`
	UserID      string    `bun:"user_id"`
	Role        string    `bun:"role"`
	CreatedAt   time.Time `bun:"created_at"`
	ModifiedAt  time.Time `bun:"modified_at"`
}

func toRecord(user *WorkspaceUser) record {
	return record{
		ID:          user.ID,
		WorkspaceID: user.WorkspaceID,
		UserID:      user.UserID,
		Role:        string(user.Role),
		CreatedAt:   user.CreatedAt,
		ModifiedAt:  user.ModifiedAt,
	}
}

func toRecords(users []*WorkspaceUser) []record {
	var (
		res = make([]record, len(users))
	)

	for i, p := range users {
		res[i] = toRecord(p)
	}

	return res
}

func fromRecords(records []record) []*WorkspaceUser {
	users := make([]*WorkspaceUser, len(records))

	for i, r := range records {
		users[i] = &WorkspaceUser{
			ID:          r.ID,
			Role:        WorkspaceUserRole(r.Role),
			CreatedAt:   r.CreatedAt,
			ModifiedAt:  r.ModifiedAt,
			WorkspaceID: r.WorkspaceID,
			UserID:      r.UserID,
		}
	}

	return users
}
