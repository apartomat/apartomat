package postgres

import (
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(albums.Store), new(*store)),
)
