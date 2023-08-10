package postgres

import (
	"context"
	"time"

	. "github.com/apartomat/apartomat/internal/store/public_sites"
	"github.com/uptrace/bun"
)

const (
	publicSitesTableName = `apartomat.public_sites`
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

func (s *store) List(ctx context.Context, spec Spec, limit, offset int) ([]PublicSite, error) {
	sql, args, err := selectBySpec(publicSitesTableName, spec, limit, offset)
	if err != nil {
		return nil, err
	}

	recs := make([]record, 0)

	if err := s.db.NewRaw(sql, args...).Scan(ctx, &recs); err != nil {
		return nil, err
	}

	return fromRecords(recs), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*PublicSite, error) {
	res, err := s.List(ctx, spec, 1, 0)
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
		Exec(ctx)

	return err
}

type record struct {
	bun.BaseModel `bun:"table:apartomat.public_sites,alias:ps"`

	ID         string         `bun:"id,pk"`
	Status     string         `bun:"status"`
	URL        string         `bun:"url"`
	Settings   settingsRecord `bun:"settings,type:jsonb"`
	CreatedAt  time.Time      `bun:"created_at"`
	ModifiedAt time.Time      `bun:"modified_at"`
	ProjectID  string         `bun:"project_id"`
}

type settingsRecord struct {
	AllowVisualizations bool `json:"allowVisualizations"`
	AllowAlbums         bool `json:"allowAlbums"`
}

func toRecord(site PublicSite) record {
	return record{
		ID:         site.ID,
		Status:     string(site.Status),
		URL:        site.URL,
		Settings:   toSettingsRecord(site.Settings),
		CreatedAt:  site.CreatedAt,
		ModifiedAt: site.ModifiedAt,
		ProjectID:  site.ProjectID,
	}
}

func toRecords(sites []PublicSite) []record {
	var (
		res = make([]record, len(sites))
	)

	for i, p := range sites {
		res[i] = toRecord(p)
	}

	return res
}

func toSettingsRecord(settings PublicSiteSettings) settingsRecord {
	return settingsRecord{
		AllowVisualizations: settings.AllowVisualizations,
		AllowAlbums:         settings.AllowAlbums,
	}
}

func fromRecords(records []record) []PublicSite {
	sites := make([]PublicSite, len(records))

	for i, rec := range records {
		sites[i] = PublicSite{
			ID:         rec.ID,
			Status:     Status(rec.Status),
			URL:        rec.URL,
			Settings:   fromSettingsRecord(rec.Settings),
			CreatedAt:  time.Time{},
			ModifiedAt: time.Time{},
			ProjectID:  rec.ProjectID,
		}
	}

	return sites
}

func fromSettingsRecord(rec settingsRecord) PublicSiteSettings {
	return PublicSiteSettings{
		AllowVisualizations: rec.AllowVisualizations,
		AllowAlbums:         rec.AllowAlbums,
	}
}
