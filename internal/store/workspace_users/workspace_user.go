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
