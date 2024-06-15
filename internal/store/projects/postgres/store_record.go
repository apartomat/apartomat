package postgres

import (
	. "github.com/apartomat/apartomat/internal/store/projects"
	"github.com/uptrace/bun"
	"time"
)

type record struct {
	bun.BaseModel `bun:"table:apartomat.projects,alias:p"`

	ID          string     `pg:"id,pk"`
	Name        string     `pg:"name"`
	Status      string     `pg:"status"`
	StartAt     *time.Time `pg:"start_at"`
	EndAt       *time.Time `pg:"end_at"`
	CreatedAt   time.Time  `pg:"created_at"`
	ModifiedAt  time.Time  `pg:"modified_at"`
	WorkspaceID string     `pg:"workspace_id"`
}

func toRecord(project *Project) record {
	return record{
		ID:          project.ID,
		Name:        project.Name,
		Status:      string(project.Status),
		StartAt:     project.StartAt,
		EndAt:       project.EndAt,
		CreatedAt:   project.CreatedAt,
		ModifiedAt:  project.ModifiedAt,
		WorkspaceID: project.WorkspaceID,
	}
}

func toRecords(projects []*Project) []record {
	var (
		records = make([]record, len(projects))
	)

	for i, p := range projects {
		records[i] = toRecord(p)
	}

	return records
}

func fromRecords(records []record) []*Project {
	projects := make([]*Project, len(records))

	for i, r := range records {
		projects[i] = &Project{
			ID:          r.ID,
			Name:        r.Name,
			Status:      Status(r.Status),
			StartAt:     r.StartAt,
			EndAt:       r.EndAt,
			CreatedAt:   r.CreatedAt,
			ModifiedAt:  r.ModifiedAt,
			WorkspaceID: r.WorkspaceID,
		}
	}

	return projects
}
