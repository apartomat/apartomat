package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ztsu/apartomat/internal/store"
	"time"
)

const (
	usersTableName = `apartomat.users`
)

type userStore struct {
	pg *pgxpool.Pool
}

func NewUserStore(pg *pgxpool.Pool) *userStore {
	return &userStore{pg}
}

var (
	_ store.UserStore = (*userStore)(nil)
)

func (s *userStore) Save(ctx context.Context, user *store.User) (*store.User, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	if user.ModifiedAt.IsZero() {
		user.ModifiedAt = user.CreatedAt
	}

	q, args, err := InsertIntoUsers().
		Columns("email", "full_name", "is_active", "created_at", "modified_at").
		Values(user.Email, user.FullName, user.IsActive, user.CreatedAt, user.ModifiedAt).
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

func (s *userStore) List(ctx context.Context, q store.UserStoreQuery) ([]*store.User, error) {
	sql, args, err := SelectFromUsers("id", "email", "full_name", "is_active", "created_at", "modified_at").
		Where(q).Limit(q.Limit).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.pg.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*store.User, 0)

	for rows.Next() {
		user := new(store.User)

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FullName,
			&user.IsActive,
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

// userInsertBuilder is insert builder shortcut for users table
type userInsertBuilder struct {
	sq.InsertBuilder
}

func InsertIntoUsers() *userInsertBuilder {
	return &userInsertBuilder{sq.Insert(usersTableName).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id")}
}

// userSelectBuilder is insert builder shortcut for users table
type userSelectBuilder struct {
	sq.SelectBuilder
}

func SelectFromUsers(columns ...string) *userSelectBuilder {
	return &userSelectBuilder{
		sq.Select(columns...).
			From(usersTableName).
			PlaceholderFormat(sq.Dollar),
	}
}

func (builder *userSelectBuilder) Where(q store.UserStoreQuery) *userSelectBuilder {
	if len(q.ID.Eq) > 0 {
		builder.SelectBuilder.Where(sq.Eq{"id": q.ID.Eq})
	}

	if len(q.Email.Eq) > 0 {
		builder.SelectBuilder.Where(sq.Eq{"email": q.Email.Eq})
	}

	return builder
}

func (builder *userSelectBuilder) Limit(n int) *userSelectBuilder {
	if n != 0 {
		builder.SelectBuilder = builder.SelectBuilder.Limit(uint64(n))
	}

	return builder
}
