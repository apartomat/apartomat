package postgres

import (
	"context"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/files"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/uptrace/bun"
	"time"
)

const (
	filesTableName = `apartomat.files`
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

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*File, error) {
	sql, args, err := selectBySpec(filesTableName, spec, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		recs = make([]record, 0)
	)

	if err := s.db.NewRaw(sql, args...).
		Scan(bunhook.WithQueryContext(ctx, "Files.List"), &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*File, error) {
	res, err := s.List(ctx, spec, SortDefault, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrFileNotFound
	}

	return res[0], nil
}

func (s *store) Count(ctx context.Context, spec Spec) (int, error) {
	sql, args, err := countBySpec(filesTableName, spec)
	if err != nil {
		return 0, err
	}

	var (
		c int
	)

	if err = s.db.NewRaw(sql, args).Scan(bunhook.WithQueryContext(ctx, "Files.Count"), &c); err != nil {
		return 0, err
	}

	return c, nil
}

func (s *store) Save(ctx context.Context, files ...*File) error {
	recs := toRecords(files)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(bunhook.WithQueryContext(ctx, "Files.Save"))

	return err
}

func (s *store) Delete(ctx context.Context, files ...*File) error {
	var (
		ids = make([]string, len(files))
	)

	for i, f := range files {
		ids[i] = f.ID
	}

	_, err := s.db.NewDelete().
		Model((*record)(nil)).
		Where(`id IN (?)`, bun.In(ids)).
		Exec(bunhook.WithQueryContext(ctx, "Files.Delete"))

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.albums,alias:a"`

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

func fromRecords(records []record) []*File {
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

func selectBySpec(tableName string, spec Spec, sort Sort, limit, offset int) (string, []interface{}, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return "", nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return "", nil, err
	}

	var (
		order = make([]exp.OrderedExpression, 0)
	)

	switch sort {
	case SortDefault:
		//
	}

	var (
		q = goqu.From(tableName).Where(expr).Limit(uint(limit)).Offset(uint(offset))
	)

	if len(order) > 0 {
		q = q.Order(order...)
	}

	return q.ToSQL()
}

func countBySpec(tableName string, spec Spec) (string, []interface{}, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return "", nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return "", nil, err
	}

	return goqu.Select(goqu.COUNT(goqu.Star())).From(tableName).Where(expr).ToSQL()
}
