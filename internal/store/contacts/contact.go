package contacts

import (
	"time"
)

type Contact struct {
	ID         string
	FullName   string
	Photo      string
	Details    []Details
	CreatedAt  time.Time
	ModifiedAt time.Time
	ProjectID  string
}

func NewContact(id, fullName, photo string, details []Details, projectID string) *Contact {
	now := time.Now()

	return &Contact{
		ID:         id,
		FullName:   fullName,
		Photo:      photo,
		Details:    details,
		CreatedAt:  now,
		ModifiedAt: now,
		ProjectID:  projectID,
	}
}

func (contact *Contact) Change(fullName, photo string, details []Details) {
	contact.FullName = fullName
	contact.Photo = photo
	contact.Details = details
}

type Details struct {
	Type  Type
	Value string
}

type Type string

const (
	TypeInstagram Type = "INSTAGRAM"
	TypePhone     Type = "PHONE"
	TypeEmail     Type = "EMAIL"
	TypeWhatsApp  Type = "WHATSAPP"
	TypeTelegram  Type = "TELEGRAM"
	TypeUnknown   Type = "UNKNOWN"
)
