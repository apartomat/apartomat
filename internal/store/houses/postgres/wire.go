package postgres

import (
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(houses.Store), new(*store)),
)
