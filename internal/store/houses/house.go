package houses

import (
	"time"
)

type House struct {
	ID             string
	City           string
	Address        string
	HousingComplex string
	CreatedAt      time.Time
	ModifiedAt     time.Time
	ProjectID      string
}

func New(id, city, address, housingComplex, projectID string) *House {
	now := time.Now()

	return &House{
		ID:             id,
		City:           city,
		Address:        address,
		HousingComplex: housingComplex,
		CreatedAt:      now,
		ModifiedAt:     now,
		ProjectID:      projectID,
	}
}

func (h *House) Change(city, address, housingComplex string) {
	h.City = city
	h.Address = address
	h.HousingComplex = housingComplex
	h.ModifiedAt = time.Now()
}
