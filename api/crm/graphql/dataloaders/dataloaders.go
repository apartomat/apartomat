package dataloaders

import (
	"github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/apartomat/apartomat/internal/store/users"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	"github.com/vikstrous/dataloadgen"
	"time"
)

const (
	defaultDataLoaderWait = 16 * time.Millisecond
)

type DataLoaders struct {
	Files      *dataloadgen.Loader[string, *files.File]
	Rooms      *dataloadgen.Loader[string, *rooms.Room]
	Users      *dataloadgen.Loader[string, *users.User]
	Workspaces *dataloadgen.Loader[string, *workspaces.Workspace]
}

func NewDataLoaders(
	filesStore files.Store,
	roomsStore rooms.Store,
	usersStore users.Store,
	workspacesStore workspaces.Store,
) *DataLoaders {
	return &DataLoaders{
		Files:      dataloadgen.NewLoader(fetchFiles(filesStore), dataloadgen.WithWait(defaultDataLoaderWait)),
		Rooms:      dataloadgen.NewLoader(fetchRooms(roomsStore), dataloadgen.WithWait(defaultDataLoaderWait)),
		Users:      dataloadgen.NewLoader(fetchUsers(usersStore), dataloadgen.WithWait(defaultDataLoaderWait)),
		Workspaces: dataloadgen.NewLoader(fetchWorkspaces(workspacesStore), dataloadgen.WithWait(defaultDataLoaderWait)),
	}
}
