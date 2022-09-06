package rooms

import (
	"context"
	"time"
)

type Room struct {
	ID         string
	Name       string
	Square     *float64
	Level      *int
	CreatedAt  time.Time
	ModifiedAt time.Time
	HouseID    string
}

type Store interface {
	Save(context.Context, *Room) (*Room, error)
	Delete(context.Context, *Room) error
	List(ctx context.Context, spec Spec, limit, offset int) ([]*Room, error)
}

type Spec interface {
	Is(*Room) bool
}

// IDInSpec is specification that point Room has specified ID
type IDInSpec struct {
	IDs []string
}

func (s IDInSpec) Is(c *Room) bool {
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

// HouseIDInSpec is specification that point Room belongs specified House
type HouseIDInSpec struct {
	IDs []string
}

func (s HouseIDInSpec) Is(c *Room) bool {
	for _, id := range s.IDs {
		if c.HouseID == id {
			return true
		}
	}

	return false
}

func HouseIDIn(ids ...string) Spec {
	return HouseIDInSpec{IDs: ids}
}
