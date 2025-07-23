package postgres

import (
	"github.com/apartomat/apartomat/internal/store/albumfiles"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(albumfiles.Store), new(*store)),
)
