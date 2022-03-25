package store

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/expr"
	"time"
)

type User struct {
	ID          string
	Email       string
	FullName    string
	IsActive    bool
	UseGravatar bool
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

type UserStore interface {
	Save(context.Context, *User) (*User, error)
	List(context.Context, UserStoreQuery) ([]*User, error)
}

type UserStoreQuery struct {
	ID          expr.Str
	Email       expr.Str
	WorkspaceID expr.Str
	Sort        []UserStoreQuerySort
	Limit       int
	Offset      int
}

type UserStoreQuerySort int

const (
	UserStoreQuerySortID UserStoreQuerySort = iota
	UserStoreQuerySortIDDesc
)
