package apartomat

import "context"

func (acl *Acl) CanGetWorkspaceUserProfile(ctx context.Context, subj *UserCtx, obj struct{ WorkspaceID, UserID int }) bool {
	// todo check subj has access to workspace
	return true
}
