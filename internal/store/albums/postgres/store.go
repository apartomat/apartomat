package postgres

import (
	"context"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/albums"
	"github.com/uptrace/bun"
	"time"
)

const (
	albumsTableName = `apartomat.albums`
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

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Album, error) {
	sql, args, err := selectBySpec(albumsTableName, spec, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		recs = make([]record, 0)
	)

	if err := s.db.NewRaw(sql, args...).
		Scan(bunhook.WithQueryContext(ctx, "Albums.List"), &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*Album, error) {
	res, err := s.List(ctx, spec, SortDefault, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrAlbumNotFound
	}

	return res[0], nil
}

func (s *store) Count(ctx context.Context, spec Spec) (int, error) {
	sql, args, err := countBySpec(albumsTableName, spec)
	if err != nil {
		return 0, err
	}

	var (
		c int
	)

	if err = s.db.NewRaw(sql, args).Scan(bunhook.WithQueryContext(ctx, "Albums.Count"), &c); err != nil {
		return 0, err
	}

	return c, nil
}

func (s *store) Save(ctx context.Context, albums ...*Album) error {
	recs := toRecords(albums)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(bunhook.WithQueryContext(ctx, "Albums.Save"))

	return err
}

func (s *store) Delete(ctx context.Context, albums ...*Album) error {
	var (
		ids = make([]string, len(albums))
	)

	for i, f := range albums {
		ids[i] = f.ID
	}

	_, err := s.db.NewDelete().
		Model((*record)(nil)).
		Where(`id IN (?)`, bun.In(ids)).
		Exec(bunhook.WithQueryContext(ctx, "Albums.Delete"))

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.albums,alias:a"`

	ID         string         `bun:"id,pk"`
	Name       string         `bun:"name"`
	Version    int            `bun:"version"`
	Settings   settingsRecord `bun:"settings"`
	Pages      []pageRecord   `bun:"pages"`
	CreatedAt  time.Time      `bun:"created_at"`
	ModifiedAt time.Time      `bun:"modified_at"`
	ProjectID  string         `bun:"project_id"`
}

type settingsRecord struct {
	PageOrientation string `json:"pageOrientation"`
	PageSize        string `json:"pageSize"`
}

type pageRecord struct {
	VisualizationID string `json:"visualization_id"`
	FileID          string `json:"file_id"`
}

func toRecord(album *Album) record {
	return record{
		ID:         album.ID,
		Name:       album.Name,
		Version:    album.Version,
		Settings:   toSettingsRecord(album.Settings),
		Pages:      toPageRecords(album.Pages),
		CreatedAt:  album.CreatedAt,
		ModifiedAt: album.ModifiedAt,
		ProjectID:  album.ProjectID,
	}
}

func toRecords(albums []*Album) []record {
	var (
		res = make([]record, len(albums))
	)

	for i, c := range albums {
		res[i] = toRecord(c)
	}

	return res
}

func fromRecords(recs []record) []*Album {
	var (
		res = make([]*Album, len(recs))
	)

	for i, rec := range recs {
		res[i] = &Album{
			ID:         rec.ID,
			Name:       rec.Name,
			Version:    rec.Version,
			Settings:   fromSettingsRecord(rec.Settings),
			Pages:      fromPageRecords(rec.Pages),
			CreatedAt:  rec.CreatedAt,
			ModifiedAt: rec.ModifiedAt,
			ProjectID:  rec.ProjectID,
		}
	}

	return res
}

func toSettingsRecord(settings Settings) settingsRecord {
	return settingsRecord{
		PageOrientation: string(settings.PageOrientation),
		PageSize:        string(settings.PageSize),
	}
}

func toPageRecords(pages []AlbumPageVisualization) []pageRecord {
	var (
		res = make([]pageRecord, len(pages))
	)

	for i, p := range pages {
		res[i] = pageRecord{
			VisualizationID: p.VisualizationID,
			FileID:          p.FileID,
		}
	}

	return res
}

func fromSettingsRecord(rec settingsRecord) Settings {
	return Settings{
		PageOrientation: PageOrientation(rec.PageOrientation),
		PageSize:        PageSize(rec.PageSize),
	}
}

func fromPageRecords(recs []pageRecord) []AlbumPageVisualization {
	var (
		res = make([]AlbumPageVisualization, len(recs))
	)

	for i, rec := range recs {
		res[i] = AlbumPageVisualization{
			VisualizationID: rec.VisualizationID,
			FileID:          rec.FileID,
		}
	}

	return res
}
