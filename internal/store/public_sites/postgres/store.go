package postgres

import (
	"context"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	"time"

	. "github.com/apartomat/apartomat/internal/store/public_sites"
	"github.com/uptrace/bun"
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

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]PublicSite, error) {
	sql, args, err := selectBySpec(`apartomat.public_sites`, spec, limit, offset)
	if err != nil {
		return nil, err
	}

	recs := make([]record, 0)

	if err := s.db.NewRaw(sql, args...).Scan(bunhook.WithQueryContext(ctx, "PublicSites.List"), &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*PublicSite, error) {
	res, err := s.List(ctx, spec, SortDefault, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrPublicSiteNotFound
	}

	return &res[0], nil
}

func (s *store) Save(ctx context.Context, sites ...PublicSite) error {
	recs := toRecords(sites)

	_, err := s.db.NewInsert().Model(&recs).
		Returning("NULL").
		On("CONFLICT (id) DO UPDATE").
		Exec(bunhook.WithQueryContext(ctx, "PublicSites.Save"))

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.public_sites,alias:ps"`

	ID          string         `bun:"id,pk"`
	Status      string         `bun:"status"`
	URL         string         `bun:"url"`
	Settings    settingsRecord `bun:"settings,type:jsonb"`
	Title       string         `bun:"title"`
	Description string         `bun:"description"`
	CreatedAt   time.Time      `bun:"created_at"`
	ModifiedAt  time.Time      `bun:"modified_at"`
	ProjectID   string         `bun:"project_id"`
}

type settingsRecord struct {
	AllowVisualizations bool `json:"allowVisualizations"`
	AllowAlbums         bool `json:"allowAlbums"`
}

func toRecord(val PublicSite) record {
	return record{
		ID:          val.ID,
		Status:      string(val.Status),
		URL:         val.URL,
		Settings:    toSettingsRecord(val.Settings),
		Title:       val.Title,
		Description: val.Description,
		CreatedAt:   val.CreatedAt,
		ModifiedAt:  val.ModifiedAt,
		ProjectID:   val.ProjectID,
	}
}

func toRecords(vals []PublicSite) []record {
	var (
		res = make([]record, len(vals))
	)

	for i, p := range vals {
		res[i] = toRecord(p)
	}

	return res
}

func toSettingsRecord(val PublicSiteSettings) settingsRecord {
	return settingsRecord{
		AllowVisualizations: val.AllowVisualizations,
		AllowAlbums:         val.AllowAlbums,
	}
}

func fromRecords(records []record) []PublicSite {
	var (
		res = make([]PublicSite, len(records))
	)

	for i, rec := range records {
		res[i] = PublicSite{
			ID:          rec.ID,
			Status:      Status(rec.Status),
			URL:         rec.URL,
			Settings:    fromSettingsRecord(rec.Settings),
			Title:       rec.Title,
			Description: rec.Description,
			CreatedAt:   time.Time{},
			ModifiedAt:  time.Time{},
			ProjectID:   rec.ProjectID,
		}
	}

	return res
}

func fromSettingsRecord(rec settingsRecord) PublicSiteSettings {
	return PublicSiteSettings{
		AllowVisualizations: rec.AllowVisualizations,
		AllowAlbums:         rec.AllowAlbums,
	}
}
