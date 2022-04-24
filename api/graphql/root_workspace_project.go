package graphql

import (
	"context"
	"fmt"
	"time"
)

type workspaceProjectResolver struct {
	*rootResolver
}

func (r *rootResolver) WorkspaceProject() WorkspaceProjectResolver {
	return &workspaceProjectResolver{r}
}

func (r *workspaceProjectResolver) Period(ctx context.Context, obj *WorkspaceProject, timezone *string) (*string, error) {
	return period(obj.StartAt, obj.EndAt, timezone)
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
		return pstring(""), nil
	}

	startAtLoc := startAt.In(loc)

	if endAt == nil {
		return pstring(fmt.Sprintf("%d", startAtLoc.Year())), nil
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

	return pstring(per), nil
}
