package visualizations

import (
	"time"
)

type Visualization struct {
	ID            string
	Name          string
	Description   string
	Version       int
	CreatedAt     time.Time
	ModifiedAt    time.Time
	ProjectID     string
	ProjectFileID string
	RoomID        *string
}
