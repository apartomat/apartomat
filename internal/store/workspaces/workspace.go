package workspaces

import "time"

type Workspace struct {
	ID         string
	Name       string
	IsActive   bool
	CreatedAt  time.Time
	ModifiedAt time.Time
	UserID     string
}

func New(id, name string, isActive bool, userID string) *Workspace {
	now := time.Now()

	return &Workspace{
		ID:         id,
		Name:       name,
		IsActive:   isActive,
		CreatedAt:  now,
		ModifiedAt: now,
		UserID:     userID,
	}
}
