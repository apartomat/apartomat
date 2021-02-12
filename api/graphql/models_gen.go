// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
)

type ConfirmLoginResult interface {
	IsConfirmLoginResult()
}

//  Common types
type Error interface {
	IsError()
}

type LoginByEmailResult interface {
	IsLoginByEmailResult()
}

type UserProfileResult interface {
	IsUserProfileResult()
}

type WorkspaceProjectsListResult interface {
	IsWorkspaceProjectsListResult()
}

type WorkspaceProjectsTotalResult interface {
	IsWorkspaceProjectsTotalResult()
}

type WorkspaceResult interface {
	IsWorkspaceResult()
}

type WorkspaceUsersResult interface {
	IsWorkspaceUsersResult()
}

type CheckEmail struct {
	Email string `json:"email"`
}

func (CheckEmail) IsLoginByEmailResult() {}

type ExpiredToken struct {
	Message string `json:"message"`
}

func (ExpiredToken) IsConfirmLoginResult() {}
func (ExpiredToken) IsError()              {}

type Forbidden struct {
	Message string `json:"message"`
}

func (Forbidden) IsUserProfileResult()            {}
func (Forbidden) IsError()                        {}
func (Forbidden) IsWorkspaceResult()              {}
func (Forbidden) IsWorkspaceUsersResult()         {}
func (Forbidden) IsWorkspaceProjectsListResult()  {}
func (Forbidden) IsWorkspaceProjectsTotalResult() {}

type Gravatar struct {
	URL string `json:"url"`
}

type ID struct {
	ID int `json:"id"`
}

type InvalidEmail struct {
	Message string `json:"message"`
}

func (InvalidEmail) IsLoginByEmailResult() {}
func (InvalidEmail) IsError()              {}

type InvalidToken struct {
	Message string `json:"message"`
}

func (InvalidToken) IsConfirmLoginResult() {}
func (InvalidToken) IsError()              {}

type LoginConfirmed struct {
	Token string `json:"token"`
}

func (LoginConfirmed) IsConfirmLoginResult() {}

type NotFound struct {
	Message string `json:"message"`
}

func (NotFound) IsError()           {}
func (NotFound) IsWorkspaceResult() {}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ServerError struct {
	Message string `json:"message"`
}

func (ServerError) IsLoginByEmailResult()           {}
func (ServerError) IsConfirmLoginResult()           {}
func (ServerError) IsUserProfileResult()            {}
func (ServerError) IsError()                        {}
func (ServerError) IsWorkspaceResult()              {}
func (ServerError) IsWorkspaceUsersResult()         {}
func (ServerError) IsWorkspaceProjectsListResult()  {}
func (ServerError) IsWorkspaceProjectsTotalResult() {}

type ShoppinglistQuery struct {
	ProductOnPage *Product `json:"productOnPage"`
}

type UserProfile struct {
	ID               int        `json:"id"`
	Email            string     `json:"email"`
	Gravatar         *Gravatar  `json:"gravatar"`
	DefaultWorkspace *Workspace `json:"defaultWorkspace"`
}

func (UserProfile) IsUserProfileResult() {}

type Workspace struct {
	ID       int                  `json:"id"`
	Name     string               `json:"name"`
	Users    WorkspaceUsersResult `json:"users"`
	Projects *WorkspaceProjects   `json:"projects"`
}

func (Workspace) IsWorkspaceResult() {}

type WorkspaceProject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type WorkspaceProjects struct {
	Workspace *ID                          `json:"workspace"`
	List      WorkspaceProjectsListResult  `json:"list"`
	Total     WorkspaceProjectsTotalResult `json:"total"`
}

type WorkspaceProjectsList struct {
	Items []*WorkspaceProject `json:"items"`
}

func (WorkspaceProjectsList) IsWorkspaceProjectsListResult() {}

type WorkspaceProjectsTotal struct {
	Total int `json:"total"`
}

func (WorkspaceProjectsTotal) IsWorkspaceProjectsTotalResult() {}

type WorkspaceUser struct {
	ID      int                   `json:"id"`
	Role    WorkspaceUserRole     `json:"role"`
	Profile *WorkspaceUserProfile `json:"profile"`
}

type WorkspaceUserProfile struct {
	ID       int       `json:"id"`
	Email    string    `json:"email"`
	Gravatar *Gravatar `json:"gravatar"`
}

type WorkspaceUsers struct {
	Items []*WorkspaceUser `json:"items"`
}

func (WorkspaceUsers) IsWorkspaceUsersResult() {}

type WorkspaceUserRole string

const (
	WorkspaceUserRoleAdmin WorkspaceUserRole = "ADMIN"
	WorkspaceUserRoleUser  WorkspaceUserRole = "USER"
)

var AllWorkspaceUserRole = []WorkspaceUserRole{
	WorkspaceUserRoleAdmin,
	WorkspaceUserRoleUser,
}

func (e WorkspaceUserRole) IsValid() bool {
	switch e {
	case WorkspaceUserRoleAdmin, WorkspaceUserRoleUser:
		return true
	}
	return false
}

func (e WorkspaceUserRole) String() string {
	return string(e)
}

func (e *WorkspaceUserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkspaceUserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkspaceUserRole", str)
	}
	return nil
}

func (e WorkspaceUserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
