package postgres

import (
	"github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(rooms.Store), new(*store)),
)
