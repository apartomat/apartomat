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
	ID     expr.Int
	Email  expr.Str
	Sort   []UserStoreQuerySort
	Limit  *uint64
	Offset *uint64
}

type UserStoreQuerySort int

const (
	UserStoreQuerySortID UserStoreQuerySort = iota
	UserStoreQuerySortIDDesc
)
