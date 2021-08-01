package apartomat

import (
	"context"
	"github.com/apartomat/apartomat/internal/store"
)

type Acl struct {
}

func NewAcl() *Acl {
	return &Acl{}
}

func (acl *Acl) CanConfirmLogin(ctx context.Context, subj *UserCtx, obj string) bool {
	return true
}

func (acl *Acl) CanGetWorkspaceUsers(ctx context.Context, subj *UserCtx, obj *store.Workspace) bool {
	// todo check subj is workspace owner or admin
	return true
}

func (acl *Acl) CanGetWorkspaceUserProfile(ctx context.Context, subj *UserCtx, obj struct{ WorkspaceID, UserID int }) bool {
	// todo check subj has access to workspace
	return true
}

func (acl *Acl) CanGetWorkspaceProjects(ctx context.Context, subj *UserCtx, obj *store.Workspace) bool {
	// todo check subj is workspace owner or admin
	return true
}

func (acl *Acl) CanGetProject(ctx context.Context, subj *UserCtx, obj *store.Project) bool {
	// todo check subj is workspace owner or admin
	return true
}

func (acl *Acl) CanGetProjectFiles(ctx context.Context, subj *UserCtx, obj *store.Project) bool {
	return true
}

func (acl *Acl) CanUploadProjectFile(ctx context.Context, subj *UserCtx, obj *store.Project) bool {
	return true
}
