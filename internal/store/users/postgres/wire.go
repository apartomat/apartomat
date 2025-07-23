package postgres

import (
	"github.com/apartomat/apartomat/internal/store/users"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(users.Store), new(*store)),
)
