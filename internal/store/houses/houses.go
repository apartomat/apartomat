package houses

import (
	"context"
	"time"
)

type House struct {
	ID             string
	City           string
	Address        string
	HousingComplex string
	CreatedAt      time.Time
	ModifiedAt     time.Time
	ProjectID      int
}

type Store interface {
	Save(context.Context, *House) (*House, error)
	List(ctx context.Context, spec Spec, limit, offset int) ([]*House, error)
}

type Spec interface {
	Is(*House) bool
}

// IDInSpec is specification that point House has specified ID
type IDInSpec struct {
	IDs []string
}

func (s IDInSpec) Is(c *House) bool {
	for _, id := range s.IDs {
		if c.ID == id {
			return true
		}
	}

	return false
}

func IDIn(ids ...string) Spec {
	return IDInSpec{IDs: ids}
}

// ProjectIDInSpec is specification that point House belongs specified Project
type ProjectIDInSpec struct {
	IDs []int
}

func (s ProjectIDInSpec) Is(c *House) bool {
	for _, id := range s.IDs {
		if c.ProjectID == id {
			return true
		}
	}

	return false
}

func ProjectIDIn(ids ...int) Spec {
	return ProjectIDInSpec{IDs: ids}
}
