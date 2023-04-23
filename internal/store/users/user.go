package users

import "time"

type User struct {
	ID                 string
	Email              string
	FullName           string
	IsActive           bool
	UseGravatar        bool
	DefaultWorkspaceID *string
	CreatedAt          time.Time
	ModifiedAt         time.Time
}

func NewUser(id, email, fullName string, isActive, useGravatar bool) *User {
	var (
		now = time.Now()
	)

	return &User{
		ID:          id,
		Email:       email,
		FullName:    fullName,
		IsActive:    isActive,
		UseGravatar: useGravatar,
		CreatedAt:   now,
		ModifiedAt:  now,
	}
}
