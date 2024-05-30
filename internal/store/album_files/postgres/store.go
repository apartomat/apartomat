package postgres

import (
	"context"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
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

	if err := s.db.NewRaw(sql, args...).Scan(bunhook.WithQueryContext(ctx, "AlbumFiles.List"), &recs); err != nil {
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

	if err = s.db.NewRaw(sql, args).Scan(bunhook.WithQueryContext(ctx, "AlbumFiles.Count"), &c); err != nil {
		return 0, err
	}

	return c, nil
}

func (s *store) Save(ctx context.Context, files ...*AlbumFile) error {
	recs := toRecords(files)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(bunhook.WithQueryContext(ctx, "AlbumFiles.Save"))

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
		Exec(bunhook.WithQueryContext(ctx, "AlbumFiles.Delete"))

	return err
}
