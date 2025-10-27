package postgres

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/albums"
	"github.com/uptrace/bun"
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

	return fromRecords(recs)
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
	Pages      pageRecords    `bun:"pages"`
	CreatedAt  time.Time      `bun:"created_at"`
	ModifiedAt time.Time      `bun:"modified_at"`
	ProjectID  string         `bun:"project_id"`
}

type settingsRecord struct {
	PageOrientation string `json:"pageOrientation"`
	PageSize        string `json:"pageSize"`
}

type pageRecordType string

const (
	pageRecordTypeSplitCover    = "SPLIT_COVER"
	pageRecordTypeCoverUploaded = "COVER_UPLOADED"
	pageRecordTypeVisualization = "VISUALIZATION"
)

type pageRecord struct {
	ID     string         `json:"id,omitempty"`
	Type   pageRecordType `json:"type"`
	Rotate float64        `json:"rotate,omitempty"`
}

type splitCoverPageRecord struct {
	pageRecord
	Title     string `json:"title,omitempty"`
	Subtitle  string `json:"sub_title,omitempty"`
	ImgFileID string `json:"img_file_id,omitempty"`
	WithQR    bool   `json:"qr,omitempty"`
	City      string `json:"city,omitempty"`
	Year      int    `json:"year,omitempty"`
}

type coverUploadedPageRecord struct {
	pageRecord
	FileID string `json:"file_id,omitempty"`
}

type visualizationPageRecord struct {
	pageRecord
	VisualizationID string `json:"visualization_id,omitempty"`
	FileID          string `json:"file_id,omitempty"`
}

type pageRecordInterface interface {
	GetID() string
	GetType() pageRecordType
	GetRotate() float64
}

type pageRecords []pageRecordInterface

func (prs pageRecords) Value() (driver.Value, error) {
	if len(prs) == 0 {
		return "[]", nil
	}

	bytes, err := json.Marshal(prs)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

func (prs *pageRecords) Scan(value interface{}) error {
	if value == nil {
		*prs = nil
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return nil
	}

	var rawRecords []json.RawMessage
	if err := json.Unmarshal(bytes, &rawRecords); err != nil {
		return err
	}

	*prs = make([]pageRecordInterface, len(rawRecords))

	for i, rawRecord := range rawRecords {
		var baseRecord struct {
			Type pageRecordType `json:"type"`
		}
		if err := json.Unmarshal(rawRecord, &baseRecord); err != nil {
			return err
		}

		switch baseRecord.Type {
		case pageRecordTypeSplitCover:
			var rec splitCoverPageRecord
			if err := json.Unmarshal(rawRecord, &rec); err != nil {
				return err
			}
			(*prs)[i] = rec
		case pageRecordTypeCoverUploaded:
			var rec coverUploadedPageRecord
			if err := json.Unmarshal(rawRecord, &rec); err != nil {
				return err
			}
			(*prs)[i] = rec
		case pageRecordTypeVisualization:
			var rec visualizationPageRecord
			if err := json.Unmarshal(rawRecord, &rec); err != nil {
				return err
			}
			(*prs)[i] = rec
		default:
			// For backward compatibility, treat unknown types as visualization
			var rec visualizationPageRecord
			if err := json.Unmarshal(rawRecord, &rec); err != nil {
				return err
			}
			(*prs)[i] = rec
		}
	}

	return nil
}

func (p pageRecord) GetID() string {
	return p.ID
}

func (p pageRecord) GetType() pageRecordType {
	return p.Type
}

func (p pageRecord) GetRotate() float64 {
	return p.Rotate
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

func fromRecords(recs []record) ([]*Album, error) {
	var (
		res = make([]*Album, len(recs))
	)

	for i, rec := range recs {
		pages, err := fromPageRecords(rec.Pages)
		if err != nil {
			return nil, err
		}

		res[i] = &Album{
			ID:         rec.ID,
			Name:       rec.Name,
			Version:    rec.Version,
			Settings:   fromSettingsRecord(rec.Settings),
			Pages:      pages,
			CreatedAt:  rec.CreatedAt,
			ModifiedAt: rec.ModifiedAt,
			ProjectID:  rec.ProjectID,
		}
	}

	return res, nil
}

func toSettingsRecord(settings Settings) settingsRecord {
	return settingsRecord{
		PageOrientation: string(settings.PageOrientation),
		PageSize:        string(settings.PageSize),
	}
}

func toPageRecords(pages []AlbumPage) pageRecords {
	var (
		res = make([]pageRecordInterface, len(pages))
	)

	for i, p := range pages {
		switch p.(type) {
		case AlbumPageSplitCover:
			var (
				page = (p).(AlbumPageSplitCover)
			)
			res[i] = splitCoverPageRecord{
				pageRecord: pageRecord{
					ID:     page.ID,
					Type:   pageRecordTypeSplitCover,
					Rotate: page.Rotate,
				},
				Title:     page.Title,
				Subtitle:  *page.Subtitle,
				ImgFileID: page.ImgFileID,
				WithQR:    page.WithQR,
				City:      *page.City,
				Year:      *page.Year,
			}
		case AlbumPageCoverUploaded:
			var (
				page = (p).(AlbumPageCoverUploaded)
			)
			res[i] = coverUploadedPageRecord{
				pageRecord: pageRecord{
					ID:     page.ID,
					Type:   pageRecordTypeCoverUploaded,
					Rotate: page.Rotate,
				},
				FileID: page.FileID,
			}
		case AlbumPageVisualization:
			var (
				page = (p).(AlbumPageVisualization)
			)
			res[i] = visualizationPageRecord{
				pageRecord: pageRecord{
					ID:     page.ID,
					Type:   pageRecordTypeVisualization,
					Rotate: page.Rotate,
				},
				VisualizationID: page.VisualizationID,
				FileID:          page.FileID,
			}
		}
	}

	return pageRecords(res)
}

func fromSettingsRecord(rec settingsRecord) Settings {
	return Settings{
		PageOrientation: PageOrientation(rec.PageOrientation),
		PageSize:        PageSize(rec.PageSize),
	}
}

func fromPageRecords(recs pageRecords) ([]AlbumPage, error) {
	var (
		res = make([]AlbumPage, len(recs))
	)

	for i, rec := range recs {
		switch rec.GetType() {
		case pageRecordTypeSplitCover:
			if splitCoverRec, ok := rec.(splitCoverPageRecord); ok {
				res[i] = AlbumPageSplitCover{
					ID:        splitCoverRec.ID,
					Rotate:    splitCoverRec.Rotate,
					Title:     splitCoverRec.Title,
					Subtitle:  &splitCoverRec.Subtitle,
					ImgFileID: splitCoverRec.ImgFileID,
					WithQR:    splitCoverRec.WithQR,
					City:      &splitCoverRec.City,
					Year:      &splitCoverRec.Year,
				}
			}
		case pageRecordTypeCoverUploaded:
			if coverUploadedRec, ok := rec.(coverUploadedPageRecord); ok {
				res[i] = AlbumPageCoverUploaded{
					ID:     coverUploadedRec.ID,
					FileID: coverUploadedRec.FileID,
					Rotate: coverUploadedRec.Rotate,
				}
			}
		case pageRecordTypeVisualization:
			if visualizationRec, ok := rec.(visualizationPageRecord); ok {
				res[i] = AlbumPageVisualization{
					ID:              visualizationRec.ID,
					VisualizationID: visualizationRec.VisualizationID,
					FileID:          visualizationRec.FileID,
					Rotate:          visualizationRec.Rotate,
				}
			}
		default:
			return nil, errors.New("unknown pageRecordType")
		}
	}

	return res, nil
}
