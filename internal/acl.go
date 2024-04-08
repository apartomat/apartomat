package apartomat

import (
	"context"
	"fmt"
	"time"

	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/store/albums"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projects"
	"github.com/apartomat/apartomat/internal/store/rooms"
	"github.com/apartomat/apartomat/internal/store/users"
	"github.com/apartomat/apartomat/internal/store/visualizations"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	lru "github.com/hashicorp/golang-lru/v2/expirable"
)

type Acl struct {
	wuCache  *lru.LRU[string, *workspace_users.WorkspaceUser]
	prjCache *lru.LRU[string, *projects.Project]
	hCache   *lru.LRU[string, *houses.House]

	workspaceUsers workspace_users.Store
	projects       projects.Store
	houses         houses.Store
}

func NewAcl(
	workspaceUsersStore workspace_users.Store,
	projectsStore projects.Store,
	housesStore houses.Store,
) *Acl {
	return &Acl{
		wuCache:  lru.NewLRU[string, *workspace_users.WorkspaceUser](100, nil, time.Minute),
		prjCache: lru.NewLRU[string, *projects.Project](100, nil, time.Minute),

		workspaceUsers: workspaceUsersStore,
		projects:       projectsStore,
		houses:         housesStore,
	}
}

func (acl *Acl) CanGetAlbums(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetAlbumsOfProjectID(ctx context.Context, subj *auth.UserCtx, projectID string) (bool, error) {
	prj, err := acl.getProject(ctx, projectID)
	if err != nil {
		return false, err
	}

	return acl.CanGetAlbums(ctx, subj, prj)
}

func (acl *Acl) CanCountAlbums(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanCountAlbumsOfProjectID(ctx context.Context, subj *auth.UserCtx, projectID string) (bool, error) {
	prj, err := acl.getProject(ctx, projectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanGetAlbum(ctx context.Context, subj *auth.UserCtx, obj *albums.Album) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanCreateAlbum(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanDeleteAlbum(ctx context.Context, subj *auth.UserCtx, obj *albums.Album) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanAddPageToAlbum(ctx context.Context, subj *auth.UserCtx, obj *albums.Album) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanChangeAlbumSettings(ctx context.Context, subj *auth.UserCtx, obj *albums.Album) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanGetAlbumFile(ctx context.Context, subj *auth.UserCtx, obj *albums.Album) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanGetContacts(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetContactsOfProjectID(ctx context.Context, subj *auth.UserCtx, projectID string) (bool, error) {
	prj, err := acl.getProject(ctx, projectID)
	if err != nil {
		return false, err
	}

	return acl.CanGetContacts(ctx, subj, prj)
}

func (acl *Acl) CanAddContact(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanUpdateContact(ctx context.Context, subj *auth.UserCtx, obj *contacts.Contact) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanDeleteContact(ctx context.Context, subj *auth.UserCtx, obj *contacts.Contact) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanGetFiles(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetFilesOfProjectID(ctx context.Context, subj *auth.UserCtx, projectID string) (bool, error) {
	prj, err := acl.getProject(ctx, projectID)
	if err != nil {
		return false, err
	}

	return acl.CanGetFiles(ctx, subj, prj)
}

func (acl *Acl) CanCountFiles(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanCountFilesOfProjectID(ctx context.Context, subj *auth.UserCtx, projectID string) (bool, error) {
	prj, err := acl.getProject(ctx, projectID)
	if err != nil {
		return false, err
	}

	return acl.CanCountFiles(ctx, subj, prj)
}

func (acl *Acl) CanUploadFile(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetHouses(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetHousesOfProjectID(ctx context.Context, subj *auth.UserCtx, projectID string) (bool, error) {
	prj, err := acl.getProject(ctx, projectID)
	if err != nil {
		return false, err
	}

	return acl.CanGetHouses(ctx, subj, prj)
}

func (acl *Acl) CanAddHouse(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanUpdateHouse(ctx context.Context, subj *auth.UserCtx, obj *houses.House) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanCreateProject(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	return acl.isWorkspaceUser(ctx, subj, obj)
}

func (acl *Acl) CanUpdateProject(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetProject(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetPublicSite(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetPublicSiteOfProjectID(ctx context.Context, subj *auth.UserCtx, projectID string) (bool, error) {
	prj, err := acl.getProject(ctx, projectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanGetRooms(ctx context.Context, subj *auth.UserCtx, obj *houses.House) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanAddRoom(ctx context.Context, subj *auth.UserCtx, obj *houses.House) (bool, error) {
	return acl.CanGetRooms(ctx, subj, obj)
}

func (acl *Acl) CanUpdateRoom(ctx context.Context, subj *auth.UserCtx, obj *rooms.Room) (bool, error) {
	house, err := acl.getHouse(ctx, obj.HouseID)
	if err != nil {
		return false, err
	}

	prj, err := acl.getProject(ctx, house.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanDeleteRoom(ctx context.Context, subj *auth.UserCtx, obj *rooms.Room) (bool, error) {
	return acl.CanUpdateRoom(ctx, subj, obj)
}

func (acl *Acl) CanGetUserProfile(ctx context.Context, subj *auth.UserCtx, obj *users.User) (bool, error) {
	if subj == nil {
		return false, nil
	}

	return subj.ID == obj.ID, nil
}

func (acl *Acl) CanGetVisualizations(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	return acl.isProjectUser(ctx, subj, obj)
}

func (acl *Acl) CanGetVisualizationsOfProjectID(ctx context.Context, subj *auth.UserCtx, projectID string) (bool, error) {
	prj, err := acl.getProject(ctx, projectID)
	if err != nil {
		return false, err
	}

	return acl.CanGetVisualizations(ctx, subj, prj)
}

func (acl *Acl) CanGetVisualization(ctx context.Context, subj *auth.UserCtx, obj *visualizations.Visualization) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanDeleteVisualization(ctx context.Context, subj *auth.UserCtx, obj *visualizations.Visualization) (bool, error) {
	prj, err := acl.getProject(ctx, obj.ProjectID)
	if err != nil {
		return false, err
	}

	return acl.isProjectUser(ctx, subj, prj)
}

func (acl *Acl) CanGetWorkspace(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	return acl.isWorkspaceUser(ctx, subj, obj)
}

func (acl *Acl) CanGetWorkspaceUsersOfWorkspaceID(ctx context.Context, subj *auth.UserCtx, workspaceID string) (bool, error) {
	return acl.IsWorkspaceUserOfWorkspaceID(ctx, subj, workspaceID)
}

func (acl *Acl) CanGetWorkspaceProjectsOfWorkspaceID(ctx context.Context, subj *auth.UserCtx, workspaceID string) (bool, error) {
	return acl.IsWorkspaceUserOfWorkspaceID(ctx, subj, workspaceID)
}

func (acl *Acl) CanInviteUsersToWorkspace(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	return acl.isWorkspaceUserAndRoleIn(ctx, subj, obj, workspace_users.WorkspaceUserRoleAdmin)
}

func (acl *Acl) isWorkspaceUser(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := acl.getWorkspaceUser(ctx, obj.ID, subj.ID)
	if err != nil {
		return false, err
	}

	return wu.UserID == subj.ID, nil
}

func (acl *Acl) IsWorkspaceUserOfWorkspaceID(ctx context.Context, subj *auth.UserCtx, workspaceID string) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := acl.getWorkspaceUser(ctx, workspaceID, subj.ID)
	if err != nil {
		return false, err
	}

	return wu.UserID == subj.ID, nil
}

func (acl *Acl) isWorkspaceUserAndRoleIn(
	ctx context.Context,
	subj *auth.UserCtx,
	obj *workspaces.Workspace,
	roles ...workspace_users.WorkspaceUserRole,
) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := acl.getWorkspaceUser(ctx, obj.ID, subj.ID)
	if err != nil {
		return false, err
	}

	return workspace_users.And(
		workspace_users.UserIDIn(subj.ID),
		workspace_users.RoleIn(roles...),
	).Is(wu), nil
}

func (acl *Acl) isProjectUser(ctx context.Context, subj *auth.UserCtx, obj *projects.Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := acl.getWorkspaceUser(ctx, obj.WorkspaceID, subj.ID)
	if err != nil {
		return false, err
	}

	return wu.UserID == subj.ID, nil
}

func (acl *Acl) isProjectUserAndRoleIn(
	ctx context.Context,
	subj *auth.UserCtx,
	obj *projects.Project,
	roles ...workspace_users.WorkspaceUserRole,
) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := acl.getWorkspaceUser(ctx, obj.WorkspaceID, subj.ID)
	if err != nil {
		return false, err
	}

	return workspace_users.And(
		workspace_users.UserIDIn(subj.ID),
		workspace_users.RoleIn(roles...),
	).Is(wu), nil
}

func (acl *Acl) getHouse(ctx context.Context, houseID string) (*houses.House, error) {
	var (
		key = houseID
	)

	if h, ok := acl.hCache.Get(key); ok {
		return h, nil
	}

	house, err := acl.houses.Get(ctx, houses.IDIn(houseID))
	if err != nil {
		return nil, err
	}

	acl.hCache.Add(key, house)

	return house, nil
}

func (acl *Acl) getProject(ctx context.Context, projectID string) (*projects.Project, error) {
	var (
		key = projectID
	)

	if prj, ok := acl.prjCache.Get(key); ok {
		return prj, nil
	}

	prj, err := acl.projects.Get(ctx, projects.IDIn(projectID))
	if err != nil {
		return nil, err
	}

	acl.prjCache.Add(key, prj)

	return prj, nil
}

func (acl *Acl) getWorkspaceUser(ctx context.Context, workspaceID, userID string) (*workspace_users.WorkspaceUser, error) {
	var (
		key = fmt.Sprintf("%s_%s", workspaceID, userID)
	)

	if wu, ok := acl.wuCache.Get(key); ok {
		return wu, nil
	}

	wu, err := acl.workspaceUsers.Get(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(workspaceID),
			workspace_users.UserIDIn(userID),
		),
	)
	if err != nil {
		return nil, err
	}

	acl.wuCache.Add(key, wu)

	return wu, nil
}
