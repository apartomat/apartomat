package postgres

import (
	"context"
	"github.com/apartomat/apartomat/internal/postgres"
	. "github.com/apartomat/apartomat/internal/store/users"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	usersTableName = `apartomat.users`
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

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]*User, error) {
	qs, err := toQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(usersTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	rows := make([]*record, 0)

	_, err = s.db.QueryContext(postgres.WithQueryContext(ctx, "users.List"), &rows, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRecords(rows), nil
}

func (s *store) Save(ctx context.Context, users ...*User) error {
	recs := toRecords(users)

	_, err := s.db.ModelContext(postgres.WithQueryContext(ctx, "users.Save"), &recs).
		Returning("NULL").
		OnConflict("(id) DO UPDATE").
		Insert()

	return err
}

type record struct {
	tableName     struct{}  `pg:"apartomat.users"`
	ID            string    `pg:"id,pk"`
	Email         string    `pg:"email"`
	FullName      string    `pg:"full_name"`
	IsActive      bool      `pg:"is_active"`
	UserGravatart bool      `pg:"use_gravatar"`
	CreatedAt     time.Time `pg:"created_at"`
	ModifiedAt    time.Time `pg:"modified_at"`
}

func toRecord(user *User) *record {
	return &record{
		ID:            user.ID,
		Email:         user.Email,
		FullName:      user.FullName,
		IsActive:      user.IsActive,
		UserGravatart: user.UseGravatar,
		CreatedAt:     user.CreatedAt,
		ModifiedAt:    user.ModifiedAt,
	}
}

func toRecords(users []*User) []*record {
	var (
		res = make([]*record, len(users))
	)

	for i, u := range users {
		res[i] = toRecord(u)
	}

	return res
}

func fromRecords(records []*record) []*User {
	files := make([]*User, len(records))

	for i, r := range records {
		files[i] = &User{
			ID:          r.ID,
			Email:       r.Email,
			FullName:    r.FullName,
			IsActive:    r.IsActive,
			UseGravatar: r.UserGravatart,
			CreatedAt:   r.CreatedAt,
			ModifiedAt:  r.ModifiedAt,
		}
	}

	return files
}
