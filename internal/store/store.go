package store

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
)

type Store struct {
	WorkspaceUsers WorkspaceUserStore
}
