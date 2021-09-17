package graphql

import (
	"context"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
	"log"
	"time"
)

func (r *rootResolver) WorkspaceProjects() WorkspaceProjectsResolver {
	return &workspaceProjectsResolver{r}
}

type workspaceProjectsResolver struct {
	*rootResolver
}

func (r *workspaceProjectsResolver) List(ctx context.Context, obj *WorkspaceProjects) (WorkspaceProjectsListResult, error) {
	projects, err := r.useCases.GetWorkspaceProjects.Do(ctx, obj.Workspace.ID, 10, 0)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't resolve workspace (id=%d) projects: %s", obj.Workspace.ID, err)

		return ServerError{Message: "internal server error"}, nil
	}

	return WorkspaceProjectsList{Items: projectsToGraphQLWorkspaceProjects(projects)}, nil
}

func (r *workspaceProjectsResolver) Total(ctx context.Context, obj *WorkspaceProjects) (WorkspaceProjectsTotalResult, error) {
	return notImplementedYetError() //todo
}

func projectToGraphQLWorkspaceProject(project *store.Project) *WorkspaceProject {
	wp := &WorkspaceProject{
		ID:   project.ID,
		Name: project.Name,
	}

	if project.StartAt != nil {
		wp.Period = period(project.StartAt, project.EndAt)
	}

	return wp
}

func projectsToGraphQLWorkspaceProjects(projects []*store.Project) []*WorkspaceProject {
	result := make([]*WorkspaceProject, 0, len(projects))

	for _, u := range projects {
		result = append(result, projectToGraphQLWorkspaceProject(u))
	}

	return result
}

func period(start, end *time.Time) *string {
	if start == nil {
		return pstring("")
	}

	if end == nil {
		return pstring(fmt.Sprintf("%d", start.Year()))
	}

	var (
		per string

		mmap = map[time.Month]string{
			time.January:   "янв",
			time.February:  "фев",
			time.March:     "мар",
			time.April:     "апр",
			time.May:       "май",
			time.June:      "июн",
			time.July:      "июл",
			time.August:    "авг",
			time.September: "сен",
			time.October:   "окт",
			time.November:  "ноя",
			time.December:  "дек",
		}
	)

	switch {
	case start.Year() == end.Year() && start.Month() == end.Month():
		per = fmt.Sprintf("%s, %d", mmap[start.Month()], start.Year())
	case end.Year() > start.Year():
		per = fmt.Sprintf("%s-%s, %d", mmap[start.Month()], mmap[end.Month()], end.Year())
	default:
		per = fmt.Sprintf("%s-%s, %d", mmap[start.Month()], mmap[end.Month()], start.Year())
	}

	return pstring(per)
}
