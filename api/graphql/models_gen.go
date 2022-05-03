// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type AddContactResult interface {
	IsAddContactResult()
}

type ChangeProjectDatesResult interface {
	IsChangeProjectDatesResult()
}

type ChangeProjectStatusResult interface {
	IsChangeProjectStatusResult()
}

type ConfirmLoginLinkResult interface {
	IsConfirmLoginLinkResult()
}

type ConfirmLoginPinResult interface {
	IsConfirmLoginPinResult()
}

type CreateProjectResult interface {
	IsCreateProjectResult()
}

type DeleteContactResult interface {
	IsDeleteContactResult()
}

type Error interface {
	IsError()
}

type HouseRoomsListResult interface {
	IsHouseRoomsListResult()
}

type LoginByEmailResult interface {
	IsLoginByEmailResult()
}

type MenuResult interface {
	IsMenuResult()
}

type ProjectContactsListResult interface {
	IsProjectContactsListResult()
}

type ProjectContactsTotalResult interface {
	IsProjectContactsTotalResult()
}

type ProjectFilesListResult interface {
	IsProjectFilesListResult()
}

type ProjectFilesResult interface {
	IsProjectFilesResult()
}

type ProjectFilesTotalResult interface {
	IsProjectFilesTotalResult()
}

type ProjectHousesListResult interface {
	IsProjectHousesListResult()
}

type ProjectHousesTotalResult interface {
	IsProjectHousesTotalResult()
}

type ProjectResult interface {
	IsProjectResult()
}

type UpdateContactResult interface {
	IsUpdateContactResult()
}

type UploadProjectFileResult interface {
	IsUploadProjectFileResult()
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

type AddContactInput struct {
	FullName string                 `json:"fullName"`
	Details  []*ContactDetailsInput `json:"details"`
}

type AlreadyExists struct {
	Message string `json:"message"`
}

func (AlreadyExists) IsUploadProjectFileResult() {}
func (AlreadyExists) IsError()                   {}

type ChangeProjectDatesInput struct {
	StartAt *time.Time `json:"startAt"`
	EndAt   *time.Time `json:"endAt"`
}

type Contact struct {
	ID         string            `json:"id"`
	FullName   string            `json:"fullName"`
	Photo      string            `json:"photo"`
	Details    []*ContactDetails `json:"details"`
	CreatedAt  time.Time         `json:"createdAt"`
	ModifiedAt time.Time         `json:"modifiedAt"`
}

type ContactAdded struct {
	Contact *Contact `json:"contact"`
}

func (ContactAdded) IsAddContactResult() {}

type ContactDeleted struct {
	Contact *Contact `json:"contact"`
}

func (ContactDeleted) IsDeleteContactResult() {}

type ContactDetails struct {
	Type  ContactType `json:"type"`
	Value string      `json:"value"`
}

type ContactDetailsInput struct {
	Type  ContactType `json:"type"`
	Value string      `json:"value"`
}

type ContactUpdated struct {
	Contact *Contact `json:"contact"`
}

func (ContactUpdated) IsUpdateContactResult() {}

type CreateProjectInput struct {
	WorkspaceID string     `json:"workspaceId"`
	Title       string     `json:"title"`
	StartAt     *time.Time `json:"startAt"`
	EndAt       *time.Time `json:"endAt"`
}

type ExpiredToken struct {
	Message string `json:"message"`
}

func (ExpiredToken) IsConfirmLoginLinkResult() {}
func (ExpiredToken) IsError()                  {}
func (ExpiredToken) IsConfirmLoginPinResult()  {}

type FilesScreen struct {
	Project ProjectResult `json:"project"`
	Menu    MenuResult    `json:"menu"`
}

type Forbidden struct {
	Message string `json:"message"`
}

func (Forbidden) IsChangeProjectStatusResult()    {}
func (Forbidden) IsAddContactResult()             {}
func (Forbidden) IsChangeProjectDatesResult()     {}
func (Forbidden) IsCreateProjectResult()          {}
func (Forbidden) IsDeleteContactResult()          {}
func (Forbidden) IsUpdateContactResult()          {}
func (Forbidden) IsUploadProjectFileResult()      {}
func (Forbidden) IsUserProfileResult()            {}
func (Forbidden) IsProjectResult()                {}
func (Forbidden) IsProjectFilesListResult()       {}
func (Forbidden) IsProjectFilesTotalResult()      {}
func (Forbidden) IsProjectFilesResult()           {}
func (Forbidden) IsProjectContactsListResult()    {}
func (Forbidden) IsProjectContactsTotalResult()   {}
func (Forbidden) IsProjectHousesListResult()      {}
func (Forbidden) IsProjectHousesTotalResult()     {}
func (Forbidden) IsHouseRoomsListResult()         {}
func (Forbidden) IsWorkspaceResult()              {}
func (Forbidden) IsWorkspaceUsersResult()         {}
func (Forbidden) IsWorkspaceProjectsListResult()  {}
func (Forbidden) IsWorkspaceProjectsTotalResult() {}
func (Forbidden) IsError()                        {}

type Gravatar struct {
	URL string `json:"url"`
}

type House struct {
	ID             string      `json:"id"`
	City           string      `json:"city"`
	Address        string      `json:"address"`
	HousingComplex string      `json:"housingComplex"`
	CreatedAt      time.Time   `json:"createdAt"`
	ModifiedAt     time.Time   `json:"modifiedAt"`
	Rooms          *HouseRooms `json:"rooms"`
}

type HouseRooms struct {
	List HouseRoomsListResult `json:"list"`
}

type HouseRoomsList struct {
	Items []*Room `json:"items"`
}

func (HouseRoomsList) IsHouseRoomsListResult() {}

type ID struct {
	ID string `json:"id"`
}

type InvalidEmail struct {
	Message string `json:"message"`
}

func (InvalidEmail) IsLoginByEmailResult() {}
func (InvalidEmail) IsError()              {}

type InvalidPin struct {
	Message string `json:"message"`
}

func (InvalidPin) IsConfirmLoginPinResult() {}
func (InvalidPin) IsError()                 {}

type InvalidToken struct {
	Message string `json:"message"`
}

func (InvalidToken) IsConfirmLoginLinkResult() {}
func (InvalidToken) IsError()                  {}
func (InvalidToken) IsConfirmLoginPinResult()  {}

type LinkSentByEmail struct {
	Email string `json:"email"`
}

func (LinkSentByEmail) IsLoginByEmailResult() {}

type LoginConfirmed struct {
	Token string `json:"token"`
}

func (LoginConfirmed) IsConfirmLoginLinkResult() {}
func (LoginConfirmed) IsConfirmLoginPinResult()  {}

type MenuItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type MenuItems struct {
	Items []*MenuItem `json:"items"`
}

func (MenuItems) IsMenuResult() {}

type NotFound struct {
	Message string `json:"message"`
}

func (NotFound) IsChangeProjectStatusResult() {}
func (NotFound) IsChangeProjectDatesResult()  {}
func (NotFound) IsDeleteContactResult()       {}
func (NotFound) IsUpdateContactResult()       {}
func (NotFound) IsProjectResult()             {}
func (NotFound) IsWorkspaceResult()           {}
func (NotFound) IsError()                     {}

type PinSentByEmail struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (PinSentByEmail) IsLoginByEmailResult() {}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Project struct {
	ID       string           `json:"id"`
	Title    string           `json:"title"`
	Status   ProjectStatus    `json:"status"`
	StartAt  *time.Time       `json:"startAt"`
	EndAt    *time.Time       `json:"endAt"`
	Files    *ProjectFiles    `json:"files"`
	Contacts *ProjectContacts `json:"contacts"`
	Houses   *ProjectHouses   `json:"houses"`
}

func (Project) IsProjectResult() {}

type ProjectContacts struct {
	List  ProjectContactsListResult  `json:"list"`
	Total ProjectContactsTotalResult `json:"total"`
}

type ProjectContactsFilter struct {
	Type []ContactType `json:"type"`
}

type ProjectContactsList struct {
	Items []*Contact `json:"items"`
}

func (ProjectContactsList) IsProjectContactsListResult() {}

type ProjectContactsTotal struct {
	Total int `json:"total"`
}

func (ProjectContactsTotal) IsProjectContactsTotalResult() {}

type ProjectCreated struct {
	Project *Project `json:"project"`
}

func (ProjectCreated) IsCreateProjectResult() {}

type ProjectDatesChanged struct {
	Project *Project `json:"project"`
}

func (ProjectDatesChanged) IsChangeProjectDatesResult() {}

type ProjectEnums struct {
	Status *ProjectStatusEnum `json:"status"`
}

type ProjectFile struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	URL      string          `json:"url"`
	Type     ProjectFileType `json:"type"`
	MimeType string          `json:"mimeType"`
}

type ProjectFileUploaded struct {
	File *ProjectFile `json:"file"`
}

func (ProjectFileUploaded) IsUploadProjectFileResult() {}

type ProjectFiles struct {
	List  ProjectFilesListResult  `json:"list"`
	Total ProjectFilesTotalResult `json:"total"`
}

func (ProjectFiles) IsProjectFilesResult() {}

type ProjectFilesList struct {
	Items []*ProjectFile `json:"items"`
}

func (ProjectFilesList) IsProjectFilesListResult() {}

type ProjectFilesListFilter struct {
	Type []ProjectFileType `json:"type"`
}

type ProjectFilesTotal struct {
	Total int `json:"total"`
}

func (ProjectFilesTotal) IsProjectFilesTotalResult() {}

type ProjectHouses struct {
	List  ProjectHousesListResult  `json:"list"`
	Total ProjectHousesTotalResult `json:"total"`
}

type ProjectHousesFilter struct {
	ID []string `json:"ID"`
}

type ProjectHousesList struct {
	Items []*House `json:"items"`
}

func (ProjectHousesList) IsProjectHousesListResult() {}

type ProjectHousesTotal struct {
	Total int `json:"total"`
}

func (ProjectHousesTotal) IsProjectHousesTotalResult() {}

type ProjectScreen struct {
	Project ProjectResult `json:"project"`
	Menu    MenuResult    `json:"menu"`
	Enums   *ProjectEnums `json:"enums"`
}

type ProjectStatusChanged struct {
	Project *Project `json:"project"`
}

func (ProjectStatusChanged) IsChangeProjectStatusResult() {}

type ProjectStatusEnum struct {
	Items []*ProjectStatusEnumItem `json:"items"`
}

type ProjectStatusEnumItem struct {
	Key   ProjectStatus `json:"key"`
	Value string        `json:"value"`
}

type Room struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Square     *float64  `json:"square"`
	Design     bool      `json:"design"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type ScreenQuery struct {
	Version string         `json:"version"`
	Files   *FilesScreen   `json:"files"`
	Project *ProjectScreen `json:"project"`
	Spec    *SpecScreen    `json:"spec"`
}

type ServerError struct {
	Message string `json:"message"`
}

func (ServerError) IsChangeProjectStatusResult()    {}
func (ServerError) IsAddContactResult()             {}
func (ServerError) IsChangeProjectDatesResult()     {}
func (ServerError) IsConfirmLoginLinkResult()       {}
func (ServerError) IsConfirmLoginPinResult()        {}
func (ServerError) IsCreateProjectResult()          {}
func (ServerError) IsDeleteContactResult()          {}
func (ServerError) IsLoginByEmailResult()           {}
func (ServerError) IsUpdateContactResult()          {}
func (ServerError) IsUploadProjectFileResult()      {}
func (ServerError) IsUserProfileResult()            {}
func (ServerError) IsProjectResult()                {}
func (ServerError) IsProjectFilesListResult()       {}
func (ServerError) IsProjectFilesTotalResult()      {}
func (ServerError) IsProjectFilesResult()           {}
func (ServerError) IsProjectContactsListResult()    {}
func (ServerError) IsProjectContactsTotalResult()   {}
func (ServerError) IsProjectHousesListResult()      {}
func (ServerError) IsProjectHousesTotalResult()     {}
func (ServerError) IsHouseRoomsListResult()         {}
func (ServerError) IsMenuResult()                   {}
func (ServerError) IsWorkspaceResult()              {}
func (ServerError) IsWorkspaceUsersResult()         {}
func (ServerError) IsWorkspaceProjectsListResult()  {}
func (ServerError) IsWorkspaceProjectsTotalResult() {}
func (ServerError) IsError()                        {}

type ShoppinglistQuery struct {
	ProductOnPage *Product `json:"productOnPage"`
}

type SpecScreen struct {
	Project ProjectResult `json:"project"`
	Menu    MenuResult    `json:"menu"`
}

type UpdateContactInput struct {
	FullName string                 `json:"fullName"`
	Details  []*ContactDetailsInput `json:"details"`
}

type UploadProjectFileInput struct {
	ProjectID string          `json:"projectId"`
	Type      ProjectFileType `json:"type"`
	File      graphql.Upload  `json:"file"`
}

type UserProfile struct {
	ID               string     `json:"id"`
	Email            string     `json:"email"`
	FullName         string     `json:"fullName"`
	Abbr             string     `json:"abbr"`
	Gravatar         *Gravatar  `json:"gravatar"`
	DefaultWorkspace *Workspace `json:"defaultWorkspace"`
}

func (UserProfile) IsUserProfileResult() {}

type Workspace struct {
	ID       string               `json:"id"`
	Name     string               `json:"name"`
	Users    WorkspaceUsersResult `json:"users"`
	Projects *WorkspaceProjects   `json:"projects"`
}

func (Workspace) IsWorkspaceResult() {}

type WorkspaceProject struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Status  ProjectStatus `json:"status"`
	StartAt *time.Time    `json:"startAt"`
	EndAt   *time.Time    `json:"endAt"`
	Period  *string       `json:"period"`
}

type WorkspaceProjects struct {
	Workspace *ID                          `json:"workspace"`
	List      WorkspaceProjectsListResult  `json:"list"`
	Total     WorkspaceProjectsTotalResult `json:"total"`
}

type WorkspaceProjectsFilter struct {
	Status []ProjectStatus `json:"status"`
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
	ID        string                `json:"id"`
	Workspace *ID                   `json:"workspace"`
	Role      WorkspaceUserRole     `json:"role"`
	Profile   *WorkspaceUserProfile `json:"profile"`
}

type WorkspaceUserProfile struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"fullName"`
	Abbr     string    `json:"abbr"`
	Gravatar *Gravatar `json:"gravatar"`
}

type WorkspaceUsers struct {
	Items []*WorkspaceUser `json:"items"`
}

func (WorkspaceUsers) IsWorkspaceUsersResult() {}

type ContactType string

const (
	ContactTypeInstagram ContactType = "INSTAGRAM"
	ContactTypePhone     ContactType = "PHONE"
	ContactTypeEmail     ContactType = "EMAIL"
	ContactTypeWhatsapp  ContactType = "WHATSAPP"
	ContactTypeTelegram  ContactType = "TELEGRAM"
	ContactTypeUnknown   ContactType = "UNKNOWN"
)

var AllContactType = []ContactType{
	ContactTypeInstagram,
	ContactTypePhone,
	ContactTypeEmail,
	ContactTypeWhatsapp,
	ContactTypeTelegram,
	ContactTypeUnknown,
}

func (e ContactType) IsValid() bool {
	switch e {
	case ContactTypeInstagram, ContactTypePhone, ContactTypeEmail, ContactTypeWhatsapp, ContactTypeTelegram, ContactTypeUnknown:
		return true
	}
	return false
}

func (e ContactType) String() string {
	return string(e)
}

func (e *ContactType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContactType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContactType", str)
	}
	return nil
}

func (e ContactType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProjectFileType string

const (
	ProjectFileTypeNone          ProjectFileType = "NONE"
	ProjectFileTypeVisualization ProjectFileType = "VISUALIZATION"
)

var AllProjectFileType = []ProjectFileType{
	ProjectFileTypeNone,
	ProjectFileTypeVisualization,
}

func (e ProjectFileType) IsValid() bool {
	switch e {
	case ProjectFileTypeNone, ProjectFileTypeVisualization:
		return true
	}
	return false
}

func (e ProjectFileType) String() string {
	return string(e)
}

func (e *ProjectFileType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectFileType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectFileType", str)
	}
	return nil
}

func (e ProjectFileType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProjectStatus string

const (
	ProjectStatusNew        ProjectStatus = "NEW"
	ProjectStatusInProgress ProjectStatus = "IN_PROGRESS"
	ProjectStatusDone       ProjectStatus = "DONE"
	ProjectStatusCanceled   ProjectStatus = "CANCELED"
)

var AllProjectStatus = []ProjectStatus{
	ProjectStatusNew,
	ProjectStatusInProgress,
	ProjectStatusDone,
	ProjectStatusCanceled,
}

func (e ProjectStatus) IsValid() bool {
	switch e {
	case ProjectStatusNew, ProjectStatusInProgress, ProjectStatusDone, ProjectStatusCanceled:
		return true
	}
	return false
}

func (e ProjectStatus) String() string {
	return string(e)
}

func (e *ProjectStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectStatus", str)
	}
	return nil
}

func (e ProjectStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

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
