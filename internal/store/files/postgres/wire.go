package postgres

import (
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(files.Store), new(*store)),
)
