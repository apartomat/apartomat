package public_sites

import (
	"errors"
	"time"
)

var (
	ErrPublicSiteIsPublic    = errors.New("project is already public")
	ErrPublicSiteIsNotPublic = errors.New("project is already not public")
)

type PublicSite struct {
	ID          string
	Status      Status
	URL         string
	Settings    PublicSiteSettings
	Title       string
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
	ProjectID   string
}

type Status string

const (
	StatusPublic    Status = "PUBLIC"
	StatusNotPublic Status = "NOT_PUBLIC"
)

type PublicSiteSettings struct {
	AllowVisualizations bool
	AllowAlbums         bool
}

func NewPublicSite(id, title, description, url string, status Status, settings PublicSiteSettings, ProjectID string) PublicSite {
	now := time.Now()

	return PublicSite{
		ID:          id,
		Status:      status,
		URL:         url,
		Settings:    settings,
		Title:       title,
		Description: description,
		CreatedAt:   now,
		ModifiedAt:  now,
		ProjectID:   ProjectID,
	}
}

func (ps *PublicSite) ToPublic() error {
	if Public().Is(ps) {
		return ErrPublicSiteIsPublic
	}

	ps.Status = StatusPublic

	return nil
}

func (ps *PublicSite) ToNotPublic() error {
	if NotPublic().Is(ps) {
		return ErrPublicSiteIsNotPublic
	}

	ps.Status = StatusNotPublic

	return nil
}

func (ps *PublicSite) Is(spec Spec) bool {
	return spec.Is(ps)
}
