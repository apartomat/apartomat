package postgres

import (
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(projects.Store), new(*store)),
)
