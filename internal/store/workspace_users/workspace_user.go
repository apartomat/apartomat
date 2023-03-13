package workspace_users

import "time"

const (
	WorkspaceUserRoleAdmin = "admin"
	WorkspaceUserRoleUser  = "user"
)

type WorkspaceUserRole string

type WorkspaceUser struct {
	ID          string
	WorkspaceID string
	UserID      string
	Role        WorkspaceUserRole
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

func NewWorkspaceUser(id string, role WorkspaceUserRole, workspaceID, userID string) *WorkspaceUser {
	now := time.Now()

	return &WorkspaceUser{
		ID:          id,
		Role:        role,
		CreatedAt:   now,
		ModifiedAt:  now,
		WorkspaceID: workspaceID,
		UserID:      userID,
	}
}
