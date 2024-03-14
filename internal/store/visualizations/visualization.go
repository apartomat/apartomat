package visualizations

import (
	"time"
)

type Visualization struct {
	ID              string
	Name            string
	Description     string
	Version         int
	Status          VisualizationStatus
	SortingPosition int
	CreatedAt       time.Time
	ModifiedAt      time.Time
	DeletedAt       *time.Time
	ProjectID       string
	FileID          string
	RoomID          *string
}

type VisualizationStatus string

const (
	VisualizationStatusUnknown  VisualizationStatus = "UNKNOWN"
	VisualizationStatusApproved VisualizationStatus = "APPROVED"
	VisualizationStatusDeleted  VisualizationStatus = "DELETED"
)

func NewVisualization(id, projectID, projectFileID string, roomID *string) *Visualization {
	var (
		now = time.Now()
	)

	return &Visualization{
		ID:          id,
		Name:        "",
		Description: "",
		Version:     0,
		Status:      VisualizationStatusUnknown,
		CreatedAt:   now,
		ModifiedAt:  now,
		DeletedAt:   nil,
		ProjectID:   projectID,
		FileID:      projectFileID,
		RoomID:      roomID,
	}
}

func (v *Visualization) Delete() {
	var (
		now = time.Now()
	)

	v.DeletedAt = &now
	v.Status = VisualizationStatusDeleted
}
