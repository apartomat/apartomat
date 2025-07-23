package postgres

import (
	"github.com/apartomat/apartomat/internal/store/workspaces"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(workspaces.Store), new(*store)),
)
