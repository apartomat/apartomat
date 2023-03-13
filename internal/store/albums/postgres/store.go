package postgres

import (
	"context"
	. "github.com/apartomat/apartomat/internal/store/albums"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	albumsTableName = `apartomat.albums`
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

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]*Album, error) {
	qs, err := toQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	orderExpr := goqu.I("created_at").Asc()

	q := goqu.From(albumsTableName).Where(expr).Order(orderExpr)

	sql, args, err := q.Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	records := make([]*record, 0)

	_, err = s.db.QueryContext(ctx, &records, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRecords(records), nil
}

func (s *store) Save(ctx context.Context, albums ...*Album) error {
	recs := toRecords(albums)

	_, err := s.db.ModelContext(ctx, &recs).Returning("NULL").OnConflict("(id) DO UPDATE").Insert()

	return err
}

func (s *store) Delete(ctx context.Context, albums ...*Album) error {
	var (
		ids = make([]string, len(albums))
	)

	for i, c := range albums {
		ids[i] = c.ID
	}

	_, err := s.db.ModelContext(ctx, (*record)(nil)).Where(`id IN (?)`, pg.In(ids)).Delete()

	return err
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

	sql, args, err := goqu.Select(goqu.COUNT(goqu.Star())).From(albumsTableName).Where(expr).ToSQL()
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
	tableName  struct{}       `pg:"apartomat.albums"`
	ID         string         `pg:"id,pk"`
	Name       string         `pg:"name"`
	Version    int            `pg:"version"`
	Settings   settingsRecord `pg:"settings"`
	Pages      []pageRecord   `pg:"pages"`
	CreatedAt  time.Time      `pg:"created_at"`
	ModifiedAt time.Time      `pg:"modified_at"`
	ProjectID  string         `pg:"project_id"`
}

type settingsRecord struct {
	Orientation string `pg:"orientation"`
	Format      string `pg:"format"`
}

type pageRecord struct {
	VisualizationID string `pg:"visualization_id"`
	FileID          string `pg:"file_id"`
}

func toRecord(album *Album) *record {
	return &record{
		ID:         album.ID,
		Name:       album.Name,
		Version:    album.Version,
		Pages:      toPageRecords(album.Pages),
		CreatedAt:  album.CreatedAt,
		ModifiedAt: album.ModifiedAt,
		ProjectID:  album.ProjectID,
	}
}

func toRecords(albums []*Album) []*record {
	var (
		res = make([]*record, len(albums))
	)

	for i, c := range albums {
		res[i] = toRecord(c)
	}

	return res
}

func fromRecords(records []*record) []*Album {
	albums := make([]*Album, len(records))

	for i, r := range records {
		albums[i] = &Album{
			ID:         r.ID,
			Name:       r.Name,
			Pages:      fromPageRecords(r.Pages),
			CreatedAt:  r.CreatedAt,
			ModifiedAt: r.ModifiedAt,
			ProjectID:  r.ProjectID,
		}
	}

	return albums
}

func toPageRecords(pages []AlbumPageVisualization) []pageRecord {
	var (
		records = make([]pageRecord, len(pages))
	)

	for i, p := range pages {
		records[i] = pageRecord{
			VisualizationID: p.VisualizationID,
			FileID:          p.FileID,
		}
	}

	return records
}

func fromPageRecords(records []pageRecord) []AlbumPageVisualization {
	var (
		pages = make([]AlbumPageVisualization, len(records))
	)

	for i, r := range records {
		pages[i] = AlbumPageVisualization{
			VisualizationID: r.VisualizationID,
			FileID:          r.FileID,
		}
	}

	return pages
}
