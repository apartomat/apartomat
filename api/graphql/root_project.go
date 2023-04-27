package graphql

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store/projects"
	"time"
)

func (r *rootResolver) Project() ProjectResolver { return &projectResolver{r} }

type projectResolver struct {
	*rootResolver
}

func (r *projectResolver) Period(ctx context.Context, obj *Project, timezone *string) (*string, error) {
	return period(obj.StartAt, obj.EndAt, timezone)
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

func (r *projectResolver) Albums(ctx context.Context, obj *Project) (*ProjectAlbums, error) {
	return &ProjectAlbums{}, nil
}

func (r *projectResolver) Statuses(ctx context.Context, obj *Project) (*ProjectStatusDictionary, error) {
	return &ProjectStatusDictionary{
		Items: []*ProjectStatusDictionaryItem{
			{ProjectStatusNew, "Новый"},
			{ProjectStatusInProgress, "В работе"},
			{ProjectStatusDone, "Завершен"},
			{ProjectStatusCanceled, "Отменен"},
		},
	}, nil
}

func projectToGraphQL(p *projects.Project) *Project {
	if p == nil {
		return nil
	}

	return &Project{
		ID:      p.ID,
		Name:    p.Name,
		Status:  projectStatusToGraphQL(p.Status),
		StartAt: p.StartAt,
		EndAt:   p.EndAt,
	}
}

func projectStatusToGraphQL(s projects.Status) ProjectStatus {
	switch s {
	case projects.StatusNew:
		return ProjectStatusNew
	case projects.StatusInProgress:
		return ProjectStatusInProgress
	case projects.StatusDone:
		return ProjectStatusDone
	case projects.StatusCanceled:
		return ProjectStatusCanceled
	default:
		return ""
	}
}

func toProjectStatus(status ProjectStatus) projects.Status {
	switch status {
	case ProjectStatusNew:
		return projects.StatusNew
	case ProjectStatusInProgress:
		return projects.StatusInProgress
	case ProjectStatusDone:
		return projects.StatusDone
	case ProjectStatusCanceled:
		return projects.StatusCanceled
	default:
		return ""
	}
}

func toProjectStatuses(l []ProjectStatus) []projects.Status {
	res := make([]projects.Status, len(l))

	for i, status := range l {
		res[i] = toProjectStatus(status)
	}

	return res
}

func period(startAt, endAt *time.Time, timezone *string) (*string, error) {
	var (
		loc = time.UTC
	)

	if timezone != nil {
		l, err := time.LoadLocation(*timezone)
		if err != nil {
			return nil, err
		}

		loc = l
	}

	if startAt == nil {
		return newof(""), nil
	}

	startAtLoc := startAt.In(loc)

	if endAt == nil {
		return newof(fmt.Sprintf("%d", startAtLoc.Year())), nil
	}

	entAtLoc := endAt.In(loc)

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
	case startAtLoc.Year() == entAtLoc.Year() && startAtLoc.Month() == entAtLoc.Month():
		per = fmt.Sprintf("%s, %d", mmap[startAtLoc.Month()], startAtLoc.Year())
	case entAtLoc.Year() > startAtLoc.Year():
		per = fmt.Sprintf("%s-%s, %d", mmap[startAtLoc.Month()], mmap[entAtLoc.Month()], entAtLoc.Year())
	default:
		per = fmt.Sprintf("%s-%s, %d", mmap[startAtLoc.Month()], mmap[entAtLoc.Month()], startAtLoc.Year())
	}

	return newof(per), nil
}

func newof[T any](val T) *T {
	return &val
}
