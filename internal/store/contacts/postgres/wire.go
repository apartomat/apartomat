package postgres

import (
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStore,
	wire.Bind(new(contacts.Store), new(*store)),
)
