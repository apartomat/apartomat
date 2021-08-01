package apartomat

import "github.com/pkg/errors"

var (
	ErrForbidden     = errors.New("forbidden")
	ErrNotFound      = errors.New("not found")
)
