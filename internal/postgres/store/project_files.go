package store

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	projectFilesTableName = `apartomat.project_files`
)

type projectFileStore struct {
	pg *pgxpool.Pool
}

func NewProjectFileStore(pg *pgxpool.Pool) *projectFileStore {
	return &projectFileStore{pg}
}

var (
	_ store.ProjectFileStore = (*projectFileStore)(nil)
)

func (s *projectFileStore) Save(ctx context.Context, file *store.ProjectFile) (*store.ProjectFile, error) {
	if file.CreatedAt.IsZero() {
		file.CreatedAt = time.Now()
	}

	if file.ModifiedAt.IsZero() {
		file.ModifiedAt = file.CreatedAt
	}

	q, args, err := InsertIntoProjectFiles().
		Columns("id", "project_id", "name", "url", "type", "mime_type", "created_at", "modified_at").
		Values(file.ID, file.ProjectID, file.Name, file.URL, file.Type, file.MimeType, file.CreatedAt, file.ModifiedAt).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = s.pg.QueryRow(ctx, q, args...).Scan(&file.ID)
	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok && pgerr.ConstraintName == "project_files_ukey" {
			return nil, store.ErrAlreadyExists
		}

		return nil, err
	}

	return file, err
}

func (s *projectFileStore) List(ctx context.Context, q store.ProjectFileStoreQuery) ([]*store.ProjectFile, error) {
	sql, args, err := SelectFromProjectFiles(
		"id",
		"project_id",
		"name",
		"url",
		"type",
		"mime_type",
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

	files := make([]*store.ProjectFile, 0)

	for rows.Next() {
		file := new(store.ProjectFile)

		err := rows.Scan(
			&file.ID,
			&file.ProjectID,
			&file.Name,
			&file.URL,
			&file.Type,
			&file.MimeType,
			&file.CreatedAt,
			&file.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func (s *projectFileStore) Count(ctx context.Context, q store.ProjectFileStoreQuery) (int, error) {
	sql, args, err := SelectFromProjectFiles("count(id)").Where(q).Limit(1).ToSql()
	if err != nil {
		return 0, err
	}

	row := s.pg.QueryRow(ctx, sql, args...)

	var (
		c = 0
	)

	if err := row.Scan(&c); err != nil {
		return 0, err
	}

	return c, nil
}

//
type projectFilesInsertBuilder struct {
	sq.InsertBuilder
}

func InsertIntoProjectFiles() *projectFilesInsertBuilder {
	return &projectFilesInsertBuilder{
		sq.Insert(projectFilesTableName).PlaceholderFormat(sq.Dollar).Suffix("RETURNING id"),
	}
}

//
type projectFilesSelectBuilder struct {
	sq.SelectBuilder
}

func SelectFromProjectFiles(columns ...string) *projectFilesSelectBuilder {
	return &projectFilesSelectBuilder{
		sq.Select(columns...).
			From(projectFilesTableName).
			PlaceholderFormat(sq.Dollar),
	}
}

func (builder *projectFilesSelectBuilder) Where(q store.ProjectFileStoreQuery) *projectFilesSelectBuilder {
	if len(q.ID.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{`"id"`: q.ID.Eq})
	}

	if len(q.ProjectID.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{"project_id": q.ProjectID.Eq})
	}

	if len(q.Type.Eq) > 0 {
		builder.SelectBuilder = builder.SelectBuilder.Where(sq.Eq{"type": q.Type.Eq})
	}

	return builder
}

func (builder *projectFilesSelectBuilder) Limit(n int) *projectFilesSelectBuilder {
	if n != 0 {
		builder.SelectBuilder = builder.SelectBuilder.Limit(uint64(n))
	}

	return builder
}
