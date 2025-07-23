package projectpages

import (
	"errors"
	"time"
)

var (
	ErrProjectPageIsPublic    = errors.New("project page is already public")
	ErrProjectPageIsNotPublic = errors.New("project page is already not public")
)

type ProjectPage struct {
	ID          string
	Status      Status
	URL         string
	Settings    Settings
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

type Settings struct {
	AllowVisualizations bool
	AllowAlbums         bool
}

func NewProjectPage(id, title, description, url string, status Status, settings Settings, ProjectID string) ProjectPage {
	now := time.Now()

	return ProjectPage{
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

func (p *ProjectPage) ToPublic() error {
	if Public().Is(p) {
		return ErrProjectPageIsPublic
	}

	p.Status = StatusPublic

	p.ModifiedAt = time.Now()

	return nil
}

func (p *ProjectPage) ToNotPublic() error {
	if NotPublic().Is(p) {
		return ErrProjectPageIsNotPublic
	}

	p.Status = StatusNotPublic

	p.ModifiedAt = time.Now()

	return nil
}

func (p *ProjectPage) Is(spec Spec) bool {
	return spec.Is(p)
}
