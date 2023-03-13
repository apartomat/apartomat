package visualizations

import (
	"time"
)

type Visualization struct {
	ID          string
	Name        string
	Description string
	Version     int
	Status      VisualizationStatus
	CreatedAt   time.Time
	ModifiedAt  time.Time
	DeletedAt   *time.Time
	ProjectID   string
	FileID      string
	RoomID      *string
}

type VisualizationStatus string

const (
	VisualizationStatusUnknown  VisualizationStatus = "UNKNOWN"
	VisualizationStatusApproved VisualizationStatus = "APPROVED"
	VisualizationStatusDeleted  VisualizationStatus = "DELETED"
)

func NewVisualization(id, projectID, projectFileID string, roomID *string) *Visualization {
	now := time.Now()

	return &Visualization{
		ID:          id,
		Name:        "",
		Description: "",
		Version:     0,
		Status:      "",
		CreatedAt:   now,
		ModifiedAt:  now,
		DeletedAt:   nil,
		ProjectID:   projectID,
		FileID:      projectFileID,
		RoomID:      roomID,
	}
}

func (v *Visualization) Delete() {
	now := time.Now()

	v.DeletedAt = &now
	v.Status = VisualizationStatusDeleted
}
