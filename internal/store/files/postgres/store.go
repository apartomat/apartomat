package postgres

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/files"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	filesTableName = `apartomat.files`
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

func (s *store) Save(ctx context.Context, files ...*File) error {
	recs := toRecords(files)

	_, err := s.db.ModelContext(ctx, &recs).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()

	return err
}

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]*File, error) {
	qs, err := toQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(filesTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
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

func (s *store) Count(ctx context.Context, spec Spec) (int, error) {
	qs, err := toQuery(spec)
	if err != nil {
		return 0, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return 0, err
	}

	sql, args, err := goqu.Select(goqu.COUNT(goqu.Star())).From(filesTableName).Where(expr).ToSQL()
	if err != nil {
		return 0, err
	}

	var (
		c int
	)

	_, err = s.db.QueryOneContext(ctx, pg.Scan(&c), sql, args...)

	return c, err
}

type record struct {
	tableName  struct{}  `pg:"apartomat.files"`
	ID         string    `pg:"id,pk"`
	Name       string    `pg:"name"`
	URL        string    `pg:"url"`
	Type       string    `pg:"type"`
	MimeType   string    `pg:"mime_type"`
	CreatedAt  time.Time `pg:"created_at"`
	ModifiedAt time.Time `pg:"modified_at"`
	ProjectID  string    `pg:"project_id"`
}

func toRecord(file *File) *record {
	return &record{
		ID:         file.ID,
		Name:       file.Name,
		URL:        file.URL,
		Type:       string(file.Type),
		MimeType:   file.MimeType,
		CreatedAt:  file.CreatedAt,
		ModifiedAt: file.ModifiedAt,
		ProjectID:  file.ProjectID,
	}
}

func toRecords(projects []*File) []*record {
	var (
		res = make([]*record, len(projects))
	)

	for i, p := range projects {
		res[i] = toRecord(p)
	}

	return res
}

func fromRecords(records []*record) []*File {
	files := make([]*File, len(records))

	for i, r := range records {
		files[i] = &File{
			ID:         r.ID,
			Name:       r.Name,
			URL:        r.URL,
			Type:       FileType(r.Type),
			MimeType:   r.MimeType,
			CreatedAt:  r.CreatedAt,
			ModifiedAt: r.ModifiedAt,
			ProjectID:  r.ProjectID,
		}
	}

	return files
}
