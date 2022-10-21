package graphql

import (
	"context"
	"fmt"
	"github.com/apartomat/apartomat/internal/store"
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

func projectToGraphQL(p *store.Project) *Project {
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
