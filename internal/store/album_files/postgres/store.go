package postgres

import (
	"context"
	"time"

	. "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/uptrace/bun"
)

const (
	albumFilesTableName = `apartomat.album_files`
)

type store struct {
	db *bun.DB
}

func NewStore(db *bun.DB) Store {
	return &store{db}
}

var (
	_ Store = (*store)(nil)
)

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*AlbumFile, error) {
	sql, args, err := selectBySpec(albumFilesTableName, spec, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		recs = make([]record, 0)
	)

	if err := s.db.NewRaw(sql, args...).Scan(ctx, &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*AlbumFile, error) {
	res, err := s.List(ctx, spec, SortIDAsc, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrAlbumFileNotFound
	}

	return res[0], nil
}

func (s *store) GetLastVersion(ctx context.Context, spec Spec) (*AlbumFile, error) {
	res, err := s.List(ctx, spec, SortVersionDesc, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrAlbumFileNotFound
	}

	return res[0], nil
}

func (s *store) Count(ctx context.Context, spec Spec) (int, error) {
	sql, args, err := countBySpec(albumFilesTableName, spec)
	if err != nil {
		return 0, err
	}

	var (
		c int
	)

	if err = s.db.NewRaw(sql, args).Scan(ctx, &c); err != nil {
		return 0, err
	}

	return c, nil
}

func (s *store) Save(ctx context.Context, files ...*AlbumFile) error {
	recs := toRecords(files)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx)

	return err
}

func (s *store) Delete(ctx context.Context, files ...*AlbumFile) error {
	var (
		ids = make([]string, len(files))
	)

	for i, f := range files {
		ids[i] = f.ID
	}

	_, err := s.db.NewDelete().
		Model((*record)(nil)).
		Where(`id IN (?)`, bun.In(ids)).
		Exec(ctx)

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.album_files,alias:af"`

	ID                  string     `bun:"id,pk"`
	Status              string     `bun:"status"`
	AlbumID             string     `bun:"album_id"`
	Version             int        `bun:"version"`
	FileID              *string    `bun:"file_id"`
	GeneratingStartedAt *time.Time `bun:"generating_started_at"`
	GeneratingDoneAt    *time.Time `bun:"generating_done_at"`
	CreatedAt           time.Time  `bun:"created_at"`
	ModifiedAt          time.Time  `bun:"modified_at"`
}

func toRecord(file *AlbumFile) record {
	return record{
		ID:                  file.ID,
		Status:              string(file.Status),
		AlbumID:             file.AlbumID,
		Version:             file.Version,
		FileID:              file.FileID,
		GeneratingStartedAt: file.GeneratingStartedAt,
		GeneratingDoneAt:    file.GeneratingDoneAt,
		CreatedAt:           file.CreatedAt,
		ModifiedAt:          file.ModifiedAt,
	}
}

func toRecords(files []*AlbumFile) []record {
	var (
		res = make([]record, len(files))
	)

	for i, p := range files {
		res[i] = toRecord(p)
	}

	return res
}

func fromRecords(records []record) []*AlbumFile {
	var (
		files = make([]*AlbumFile, len(records))
	)

	for i, rec := range records {
		files[i] = &AlbumFile{
			ID:                  rec.ID,
			Status:              Status(rec.Status),
			AlbumID:             rec.AlbumID,
			Version:             rec.Version,
			FileID:              rec.FileID,
			GeneratingStartedAt: rec.GeneratingStartedAt,
			GeneratingDoneAt:    rec.GeneratingDoneAt,
			CreatedAt:           rec.CreatedAt,
			ModifiedAt:          rec.ModifiedAt,
		}
	}

	return files
}
