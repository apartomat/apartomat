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

type AddHouseResult interface {
	IsAddHouseResult()
}

type AddRoomResult interface {
	IsAddRoomResult()
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

type DeleteRoomResult interface {
	IsDeleteRoomResult()
}

type Error interface {
	IsError()
	GetMessage() string
}

type HouseRoomsListResult interface {
	IsHouseRoomsListResult()
}

type LoginByEmailResult interface {
	IsLoginByEmailResult()
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

type ProjectVisualizationsListResult interface {
	IsProjectVisualizationsListResult()
}

type ProjectVisualizationsTotalResult interface {
	IsProjectVisualizationsTotalResult()
}

type UpdateContactResult interface {
	IsUpdateContactResult()
}

type UpdateHouseResult interface {
	IsUpdateHouseResult()
}

type UpdateRoomResult interface {
	IsUpdateRoomResult()
}

type UploadProjectFileResult interface {
	IsUploadProjectFileResult()
}

type UploadVisualizationResult interface {
	IsUploadVisualizationResult()
}

type UploadVisualizationsResult interface {
	IsUploadVisualizationsResult()
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

type WorkspaceUsersListResult interface {
	IsWorkspaceUsersListResult()
}

type WorkspaceUsersTotalResult interface {
	IsWorkspaceUsersTotalResult()
}

type AddContactInput struct {
	FullName string                 `json:"fullName"`
	Details  []*ContactDetailsInput `json:"details"`
}

type AddHouseInput struct {
	City           string `json:"city"`
	Address        string `json:"address"`
	HousingComplex string `json:"housingComplex"`
}

type AddRoomInput struct {
	Name   string   `json:"name"`
	Square *float64 `json:"square"`
	Level  *int     `json:"level"`
}

type AlreadyExists struct {
	Message string `json:"message"`
}

func (AlreadyExists) IsUploadProjectFileResult() {}

func (AlreadyExists) IsError()                {}
func (this AlreadyExists) GetMessage() string { return this.Message }

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
	Name        string     `json:"name"`
	StartAt     *time.Time `json:"startAt"`
	EndAt       *time.Time `json:"endAt"`
}

type Enums struct {
	Project *ProjectEnums `json:"project"`
}

type ExpiredToken struct {
	Message string `json:"message"`
}

func (ExpiredToken) IsConfirmLoginLinkResult() {}

func (ExpiredToken) IsError()                {}
func (this ExpiredToken) GetMessage() string { return this.Message }

func (ExpiredToken) IsConfirmLoginPinResult() {}

type Forbidden struct {
	Message string `json:"message"`
}

func (Forbidden) IsAddContactResult() {}

func (Forbidden) IsAddHouseResult() {}

func (Forbidden) IsAddRoomResult() {}

func (Forbidden) IsChangeProjectDatesResult() {}

func (Forbidden) IsChangeProjectStatusResult() {}

func (Forbidden) IsCreateProjectResult() {}

func (Forbidden) IsDeleteContactResult() {}

func (Forbidden) IsDeleteRoomResult() {}

func (Forbidden) IsUpdateContactResult() {}

func (Forbidden) IsUpdateHouseResult() {}

func (Forbidden) IsUpdateRoomResult() {}

func (Forbidden) IsUploadProjectFileResult() {}

func (Forbidden) IsUploadVisualizationResult() {}

func (Forbidden) IsUploadVisualizationsResult() {}

func (Forbidden) IsUserProfileResult() {}

func (Forbidden) IsProjectResult() {}

func (Forbidden) IsProjectContactsListResult() {}

func (Forbidden) IsProjectContactsTotalResult() {}

func (Forbidden) IsProjectHousesListResult() {}

func (Forbidden) IsProjectHousesTotalResult() {}

func (Forbidden) IsHouseRoomsListResult() {}

func (Forbidden) IsProjectVisualizationsListResult() {}

func (Forbidden) IsProjectVisualizationsTotalResult() {}

func (Forbidden) IsProjectFilesListResult() {}

func (Forbidden) IsProjectFilesTotalResult() {}

func (Forbidden) IsWorkspaceResult() {}

func (Forbidden) IsWorkspaceProjectsListResult() {}

func (Forbidden) IsWorkspaceProjectsTotalResult() {}

func (Forbidden) IsWorkspaceUsersListResult() {}

func (Forbidden) IsWorkspaceUsersTotalResult() {}

func (Forbidden) IsError()                {}
func (this Forbidden) GetMessage() string { return this.Message }

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

type HouseAdded struct {
	House *House `json:"house"`
}

func (HouseAdded) IsAddHouseResult() {}

type HouseRooms struct {
	List HouseRoomsListResult `json:"list"`
}

type HouseRoomsList struct {
	Items []*Room `json:"items"`
}

func (HouseRoomsList) IsHouseRoomsListResult() {}

type HouseUpdated struct {
	House *House `json:"house"`
}

func (HouseUpdated) IsUpdateHouseResult() {}

type ID struct {
	ID string `json:"id"`
}

type InvalidEmail struct {
	Message string `json:"message"`
}

func (InvalidEmail) IsLoginByEmailResult() {}

func (InvalidEmail) IsError()                {}
func (this InvalidEmail) GetMessage() string { return this.Message }

type InvalidPin struct {
	Message string `json:"message"`
}

func (InvalidPin) IsConfirmLoginPinResult() {}

func (InvalidPin) IsError()                {}
func (this InvalidPin) GetMessage() string { return this.Message }

type InvalidToken struct {
	Message string `json:"message"`
}

func (InvalidToken) IsConfirmLoginLinkResult() {}

func (InvalidToken) IsError()                {}
func (this InvalidToken) GetMessage() string { return this.Message }

func (InvalidToken) IsConfirmLoginPinResult() {}

type LinkSentByEmail struct {
	Email string `json:"email"`
}

func (LinkSentByEmail) IsLoginByEmailResult() {}

type LoginConfirmed struct {
	Token string `json:"token"`
}

func (LoginConfirmed) IsConfirmLoginLinkResult() {}

func (LoginConfirmed) IsConfirmLoginPinResult() {}

type NotFound struct {
	Message string `json:"message"`
}

func (NotFound) IsAddHouseResult() {}

func (NotFound) IsAddRoomResult() {}

func (NotFound) IsChangeProjectDatesResult() {}

func (NotFound) IsChangeProjectStatusResult() {}

func (NotFound) IsDeleteContactResult() {}

func (NotFound) IsDeleteRoomResult() {}

func (NotFound) IsUpdateContactResult() {}

func (NotFound) IsUpdateHouseResult() {}

func (NotFound) IsUpdateRoomResult() {}

func (NotFound) IsProjectResult() {}

func (NotFound) IsWorkspaceResult() {}

func (NotFound) IsError()                {}
func (this NotFound) GetMessage() string { return this.Message }

type PinSentByEmail struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (PinSentByEmail) IsLoginByEmailResult() {}

type Project struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Status         ProjectStatus          `json:"status"`
	StartAt        *time.Time             `json:"startAt"`
	EndAt          *time.Time             `json:"endAt"`
	Period         *string                `json:"period"`
	Contacts       *ProjectContacts       `json:"contacts"`
	Houses         *ProjectHouses         `json:"houses"`
	Visualizations *ProjectVisualizations `json:"visualizations"`
	Files          *ProjectFiles          `json:"files"`
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

type ProjectVisualizations struct {
	List  ProjectVisualizationsListResult  `json:"list"`
	Total ProjectVisualizationsTotalResult `json:"total"`
}

type ProjectVisualizationsList struct {
	Items []*Visualization `json:"items"`
}

func (ProjectVisualizationsList) IsProjectVisualizationsListResult() {}

type ProjectVisualizationsListFilter struct {
	RoomID []string `json:"roomID"`
}

type ProjectVisualizationsTotal struct {
	Total int `json:"total"`
}

func (ProjectVisualizationsTotal) IsProjectVisualizationsTotalResult() {}

type Room struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Square     *float64  `json:"square"`
	Level      *int      `json:"level"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type RoomAdded struct {
	Room *Room `json:"room"`
}

func (RoomAdded) IsAddRoomResult() {}

type RoomDeleted struct {
	Room *Room `json:"room"`
}

func (RoomDeleted) IsDeleteRoomResult() {}

type RoomUpdated struct {
	Room *Room `json:"room"`
}

func (RoomUpdated) IsUpdateRoomResult() {}

type ServerError struct {
	Message string `json:"message"`
}

func (ServerError) IsAddContactResult() {}

func (ServerError) IsAddHouseResult() {}

func (ServerError) IsChangeProjectDatesResult() {}

func (ServerError) IsChangeProjectStatusResult() {}

func (ServerError) IsConfirmLoginLinkResult() {}

func (ServerError) IsConfirmLoginPinResult() {}

func (ServerError) IsCreateProjectResult() {}

func (ServerError) IsDeleteContactResult() {}

func (ServerError) IsLoginByEmailResult() {}

func (ServerError) IsUpdateContactResult() {}

func (ServerError) IsUpdateHouseResult() {}

func (ServerError) IsUploadProjectFileResult() {}

func (ServerError) IsUploadVisualizationResult() {}

func (ServerError) IsUploadVisualizationsResult() {}

func (ServerError) IsUserProfileResult() {}

func (ServerError) IsProjectResult() {}

func (ServerError) IsProjectContactsListResult() {}

func (ServerError) IsProjectContactsTotalResult() {}

func (ServerError) IsProjectHousesListResult() {}

func (ServerError) IsProjectHousesTotalResult() {}

func (ServerError) IsHouseRoomsListResult() {}

func (ServerError) IsProjectVisualizationsListResult() {}

func (ServerError) IsProjectVisualizationsTotalResult() {}

func (ServerError) IsProjectFilesListResult() {}

func (ServerError) IsProjectFilesTotalResult() {}

func (ServerError) IsWorkspaceResult() {}

func (ServerError) IsWorkspaceProjectsListResult() {}

func (ServerError) IsWorkspaceProjectsTotalResult() {}

func (ServerError) IsWorkspaceUsersListResult() {}

func (ServerError) IsWorkspaceUsersTotalResult() {}

func (ServerError) IsError()                {}
func (this ServerError) GetMessage() string { return this.Message }

type SomeVisualizationsUploaded struct {
	Visualizations []*Visualization `json:"visualizations"`
}

func (SomeVisualizationsUploaded) IsUploadVisualizationsResult() {}

type UpdateContactInput struct {
	FullName string                 `json:"fullName"`
	Details  []*ContactDetailsInput `json:"details"`
}

type UpdateHouseInput struct {
	City           string `json:"city"`
	Address        string `json:"address"`
	HousingComplex string `json:"housingComplex"`
}

type UpdateRoomInput struct {
	Name   string   `json:"name"`
	Square *float64 `json:"square"`
	Level  *int     `json:"level"`
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

type Visualization struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Version     int          `json:"version"`
	CreatedAt   time.Time    `json:"createdAt"`
	ModifiedAt  time.Time    `json:"modifiedAt"`
	File        *ProjectFile `json:"file"`
	Room        *Room        `json:"room"`
}

type VisualizationUploaded struct {
	Visualization *Visualization `json:"visualization"`
}

func (VisualizationUploaded) IsUploadVisualizationResult() {}

type VisualizationsUploaded struct {
	Visualizations []*Visualization `json:"visualizations"`
}

func (VisualizationsUploaded) IsUploadVisualizationsResult() {}

type Workspace struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Projects *WorkspaceProjects `json:"projects"`
	Users    *WorkspaceUsers    `json:"users"`
}

func (Workspace) IsWorkspaceResult() {}

type WorkspaceProjects struct {
	List  WorkspaceProjectsListResult  `json:"list"`
	Total WorkspaceProjectsTotalResult `json:"total"`
}

type WorkspaceProjectsFilter struct {
	Status []ProjectStatus `json:"status"`
}

type WorkspaceProjectsList struct {
	Items []*Project `json:"items"`
}

func (WorkspaceProjectsList) IsWorkspaceProjectsListResult() {}

type WorkspaceProjectsTotal struct {
	Total int `json:"total"`
}

func (WorkspaceProjectsTotal) IsWorkspaceProjectsTotalResult() {}

type WorkspaceUser struct {
	ID        string            `json:"id"`
	Workspace *ID               `json:"workspace"`
	Role      WorkspaceUserRole `json:"role"`
	Profile   *UserProfile      `json:"profile"`
}

type WorkspaceUsers struct {
	List  WorkspaceUsersListResult  `json:"list"`
	Total WorkspaceUsersTotalResult `json:"total"`
}

type WorkspaceUsersFilter struct {
	Role []WorkspaceUserRole `json:"role"`
}

type WorkspaceUsersList struct {
	Items []*WorkspaceUser `json:"items"`
}

func (WorkspaceUsersList) IsWorkspaceUsersListResult() {}

type WorkspaceUsersTotal struct {
	Total int `json:"total"`
}

func (WorkspaceUsersTotal) IsWorkspaceUsersTotalResult() {}

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
