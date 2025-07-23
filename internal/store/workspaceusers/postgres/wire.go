package postgres

import (
	"github.com/apartomat/apartomat/internal/store/workspaceusers"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(workspaceusers.Store), new(*store)),
)
