package postgres

import (
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(visualizations.Store), new(*store)),
)
