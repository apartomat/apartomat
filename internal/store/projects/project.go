package projects

import "time"

type Project struct {
	ID          string
	Name        string
	Status      Status
	StartAt     *time.Time
	EndAt       *time.Time
	CreatedAt   time.Time
	ModifiedAt  time.Time
	WorkspaceID string
}

type Status string

const (
	StatusNew        Status = "NEW"
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
	StatusCanceled   Status = "CANCELED"
)

func NewProject(id, name string, startAt, endAt *time.Time, workspaceID string) *Project {
	now := time.Now()

	return &Project{
		ID:          id,
		Name:        name,
		Status:      StatusNew,
		StartAt:     startAt,
		EndAt:       endAt,
		CreatedAt:   now,
		ModifiedAt:  now,
		WorkspaceID: workspaceID,
	}
}

func (p *Project) ChangeDates(startAt *time.Time, endAt *time.Time) {
	p.StartAt = startAt
	p.EndAt = endAt
	p.ModifiedAt = time.Now()
}

func (p *Project) ChangeStatus(status Status) {
	p.Status = status
	p.ModifiedAt = time.Now()
}
