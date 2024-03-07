package rooms

import (
	"time"
)

type Room struct {
	ID              string
	Name            string
	Square          *float64
	Level           *int
	SortingPosition int
	CreatedAt       time.Time
	ModifiedAt      time.Time
	HouseID         string
}

func NewRoom(id, name string, square *float64, level *int, houseID string) *Room {
	var (
		now = time.Now()
	)

	return &Room{
		ID:         id,
		Name:       name,
		Square:     square,
		Level:      level,
		CreatedAt:  now,
		ModifiedAt: now,
		HouseID:    houseID,
	}
}

func (r *Room) MoveToPosition(position int) {
	r.SortingPosition = position
	r.ModifiedAt = time.Now()
}
