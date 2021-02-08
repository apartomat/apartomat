package store

import (
	"context"
	"github.com/ztsu/apartomat/internal/pkg/expr"
	"time"
)

type User struct {
	ID         int
	Email      string
	FullName   string
	IsActive   bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type UserStore interface {
	Save(context.Context, *User) (*User, error)
	List(context.Context, UserStoreQuery) ([]*User, error)
}

type UserStoreQuery struct {
	ID          expr.Int
	Email       expr.Str
	WorkspaceID expr.Int
	Sort        []UserStoreQuerySort
	Limit       int
	Offset      int
}

type UserStoreQuerySort int

const (
	UserStoreQuerySortID UserStoreQuerySort = iota
	UserStoreQuerySortIDDesc
)
