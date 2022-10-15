package graphql

import (
	"context"
	"github.com/apartomat/apartomat/internal/store"
)

func (r *rootResolver) Project() ProjectResolver { return &projectResolver{r} }

type projectResolver struct {
	*rootResolver
}

func (r *projectResolver) Contacts(ctx context.Context, obj *Project) (*ProjectContacts, error) {
	return &ProjectContacts{}, nil
}

func (r *projectResolver) Files(ctx context.Context, obj *Project) (*ProjectFiles, error) {
	return &ProjectFiles{}, nil
}

func (r *projectResolver) Houses(ctx context.Context, obj *Project) (*ProjectHouses, error) {
	return &ProjectHouses{}, nil
}

func (r *projectResolver) Visualizations(ctx context.Context, obj *Project) (*ProjectVisualizations, error) {
	return &ProjectVisualizations{}, nil
}

func projectToGraphQL(p *store.Project) *Project {
	if p == nil {
		return nil
	}

	return &Project{
		ID:      p.ID,
		Title:   p.Name,
		Status:  projectStatusToGraphQL(p.Status),
		StartAt: p.StartAt,
		EndAt:   p.EndAt,
	}
}

func projectStatusToGraphQL(s store.ProjectStatus) ProjectStatus {
	switch s {
	case store.ProjectStatusNew:
		return ProjectStatusNew
	case store.ProjectStatusInProgress:
		return ProjectStatusInProgress
	case store.ProjectStatusDone:
		return ProjectStatusDone
	case store.ProjectStatusCanceled:
		return ProjectStatusCanceled
	default:
		return ""
	}
}

func toProjectStatus(status ProjectStatus) store.ProjectStatus {
	switch status {
	case ProjectStatusNew:
		return store.ProjectStatusNew
	case ProjectStatusInProgress:
		return store.ProjectStatusInProgress
	case ProjectStatusDone:
		return store.ProjectStatusDone
	case ProjectStatusCanceled:
		return store.ProjectStatusCanceled
	default:
		return ""
	}
}

func toProjectStatuses(l []ProjectStatus) []store.ProjectStatus {
	res := make([]store.ProjectStatus, len(l))

	for i, status := range l {
		res[i] = toProjectStatus(status)
	}

	return res
}
