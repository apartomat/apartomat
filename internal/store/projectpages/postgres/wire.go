package postgres

import (
	"github.com/apartomat/apartomat/internal/store/projectpages"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(projectpages.Store), new(*store)),
)
