package postgres

import (
	"context"
	"github.com/apartomat/apartomat/internal/postgres"
	. "github.com/apartomat/apartomat/internal/store/workspaces"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	workspacesTableName = `apartomat.workspaces`
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

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Workspace, error) {
	qs, err := toQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(workspacesTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	rows := make([]*record, 0)

	_, err = s.db.QueryContext(postgres.WithQueryContext(ctx, "workspaces.List"), &rows, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRecords(rows), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*Workspace, error) {
	res, err := s.List(ctx, spec, SortDefault, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrWorkspaceNotFound
	}

	return res[0], nil
}

func (s *store) Save(ctx context.Context, workspaces ...*Workspace) error {
	recs := toRecords(workspaces)

	_, err := s.db.ModelContext(postgres.WithQueryContext(ctx, "workspaces.Save"), &recs).
		Returning("NULL").
		OnConflict("(id) DO UPDATE").
		Insert()

	return err
}

type record struct {
	tableName  struct{}  `pg:"apartomat.workspaces"`
	ID         string    `pg:"id,pk"`
	Name       string    `pg:"name"`
	IsActive   bool      `pg:"is_active"`
	CreatedAt  time.Time `pg:"created_at"`
	ModifiedAt time.Time `pg:"modified_at"`
	UserID     string    `pg:"user_id"`
}

func toRecord(workspace *Workspace) *record {
	return &record{
		ID:         workspace.ID,
		Name:       workspace.Name,
		IsActive:   workspace.IsActive,
		CreatedAt:  workspace.CreatedAt,
		ModifiedAt: workspace.ModifiedAt,
		UserID:     workspace.UserID,
	}
}

func toRecords(workspaces []*Workspace) []*record {
	var (
		res = make([]*record, len(workspaces))
	)

	for i, u := range workspaces {
		res[i] = toRecord(u)
	}

	return res
}

func fromRecords(records []*record) []*Workspace {
	files := make([]*Workspace, len(records))

	for i, r := range records {
		files[i] = &Workspace{
			ID:         r.ID,
			Name:       r.Name,
			IsActive:   r.IsActive,
			CreatedAt:  r.CreatedAt,
			ModifiedAt: r.ModifiedAt,
			UserID:     r.UserID,
		}
	}

	return files
}
