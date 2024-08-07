// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type AcceptInviteResult interface {
	IsAcceptInviteResult()
}

type AddContactResult interface {
	IsAddContactResult()
}

type AddHouseResult interface {
	IsAddHouseResult()
}

type AddRoomResult interface {
	IsAddRoomResult()
}

type AddVisualizationsToAlbumResult interface {
	IsAddVisualizationsToAlbumResult()
}

type AlbumCoverResult interface {
	IsAlbumCoverResult()
}

type AlbumFileGenerated interface {
	IsAlbumFileGenerated()
}

type AlbumPage interface {
	IsAlbumPage()
	GetNumber() int
	GetRotate() float64
}

type AlbumPageCoverResult interface {
	IsAlbumPageCoverResult()
}

type AlbumPageVisualizationResult interface {
	IsAlbumPageVisualizationResult()
}

type AlbumPagesResult interface {
	IsAlbumPagesResult()
}

type AlbumProjectResult interface {
	IsAlbumProjectResult()
}

type AlbumRecentFileResult interface {
	IsAlbumRecentFileResult()
}

type AlbumResult interface {
	IsAlbumResult()
}

type ChangeAlbumPageOrientationResult interface {
	IsChangeAlbumPageOrientationResult()
}

type ChangeAlbumPageSizeResult interface {
	IsChangeAlbumPageSizeResult()
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

type CoverFileResult interface {
	IsCoverFileResult()
}

type CreateAlbumResult interface {
	IsCreateAlbumResult()
}

type CreateProjectResult interface {
	IsCreateProjectResult()
}

type DeleteAlbumResult interface {
	IsDeleteAlbumResult()
}

type DeleteContactResult interface {
	IsDeleteContactResult()
}

type DeleteRoomResult interface {
	IsDeleteRoomResult()
}

type DeleteVisualizationsResult interface {
	IsDeleteVisualizationsResult()
}

type Error interface {
	IsError()
	GetMessage() string
}

type GenerateAlbumFileResult interface {
	IsGenerateAlbumFileResult()
}

type HouseRoomsListResult interface {
	IsHouseRoomsListResult()
}

type InviteUserToWorkspaceResult interface {
	IsInviteUserToWorkspaceResult()
}

type LoginByEmailResult interface {
	IsLoginByEmailResult()
}

type MakeProjectNotPublicResult interface {
	IsMakeProjectNotPublicResult()
}

type MakeProjectPublicResult interface {
	IsMakeProjectPublicResult()
}

type MoveRoomToPositionResult interface {
	IsMoveRoomToPositionResult()
}

type ProjectAlbumsListResult interface {
	IsProjectAlbumsListResult()
}

type ProjectAlbumsTotalResult interface {
	IsProjectAlbumsTotalResult()
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

type ProjectPublicSite interface {
	IsProjectPublicSite()
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

type UploadFileResult interface {
	IsUploadFileResult()
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
	Square *float64 `json:"square,omitempty"`
	Level  *int     `json:"level,omitempty"`
}

type Album struct {
	ID       string                `json:"id"`
	Name     string                `json:"name"`
	Version  int                   `json:"version"`
	Project  AlbumProjectResult    `json:"project"`
	Settings *AlbumSettings        `json:"settings"`
	Pages    AlbumPagesResult      `json:"pages"`
	File     AlbumRecentFileResult `json:"file,omitempty"`
	Cover    AlbumCoverResult      `json:"cover"`
}

func (Album) IsAlbumResult() {}

type AlbumCreated struct {
	Album *Album `json:"album"`
}

func (AlbumCreated) IsCreateAlbumResult() {}

type AlbumDeleted struct {
	Album *Album `json:"album"`
}

func (AlbumDeleted) IsDeleteAlbumResult() {}

type AlbumFile struct {
	ID                  string          `json:"id"`
	Status              AlbumFileStatus `json:"status"`
	Version             int             `json:"version"`
	File                *File           `json:"file,omitempty"`
	GeneratingStartedAt *time.Time      `json:"generatingStartedAt,omitempty"`
	GeneratingDoneAt    *time.Time      `json:"generatingDoneAt,omitempty"`
}

func (AlbumFile) IsAlbumRecentFileResult() {}

func (AlbumFile) IsAlbumFileGenerated() {}

type AlbumFileGenerationStarted struct {
	File *AlbumFile `json:"file"`
}

func (AlbumFileGenerationStarted) IsGenerateAlbumFileResult() {}

type AlbumPageCover struct {
	Number int                  `json:"number"`
	Rotate float64              `json:"rotate"`
	Cover  AlbumPageCoverResult `json:"cover"`
}

func (AlbumPageCover) IsAlbumPage()            {}
func (this AlbumPageCover) GetNumber() int     { return this.Number }
func (this AlbumPageCover) GetRotate() float64 { return this.Rotate }

type AlbumPageOrientationChanged struct {
	Album *Album `json:"album"`
}

func (AlbumPageOrientationChanged) IsChangeAlbumPageOrientationResult() {}

type AlbumPageSizeChanged struct {
	Album *Album `json:"album"`
}

func (AlbumPageSizeChanged) IsChangeAlbumPageSizeResult() {}

type AlbumPageVisualization struct {
	Number        int                          `json:"number"`
	Rotate        float64                      `json:"rotate"`
	Visualization AlbumPageVisualizationResult `json:"visualization"`
}

func (AlbumPageVisualization) IsAlbumPage()            {}
func (this AlbumPageVisualization) GetNumber() int     { return this.Number }
func (this AlbumPageVisualization) GetRotate() float64 { return this.Rotate }

type AlbumPages struct {
	Items []AlbumPage `json:"items"`
}

func (AlbumPages) IsAlbumPagesResult() {}

type AlbumSettings struct {
	PageSize        PageSize        `json:"pageSize"`
	PageOrientation PageOrientation `json:"pageOrientation"`
}

type AlreadyExists struct {
	Message string `json:"message"`
}

func (AlreadyExists) IsUploadFileResult() {}

func (AlreadyExists) IsError()                {}
func (this AlreadyExists) GetMessage() string { return this.Message }

type AlreadyInWorkspace struct {
	Message string `json:"message"`
}

func (AlreadyInWorkspace) IsAcceptInviteResult() {}

func (AlreadyInWorkspace) IsInviteUserToWorkspaceResult() {}

func (AlreadyInWorkspace) IsError()                {}
func (this AlreadyInWorkspace) GetMessage() string { return this.Message }

type ChangeProjectDatesInput struct {
	StartAt *time.Time `json:"startAt,omitempty"`
	EndAt   *time.Time `json:"endAt,omitempty"`
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

type Cover struct {
	ID   string          `json:"id"`
	File CoverFileResult `json:"file"`
}

func (Cover) IsAlbumPageCoverResult() {}

type CoverUploaded struct {
	File CoverFileResult `json:"file"`
}

func (CoverUploaded) IsAlbumPageCoverResult() {}

type CreateAlbumSettingsInput struct {
	PageSize    PageSize        `json:"pageSize"`
	Orientation PageOrientation `json:"orientation"`
}

type CreateProjectInput struct {
	WorkspaceID string     `json:"workspaceId"`
	Name        string     `json:"name"`
	StartAt     *time.Time `json:"startAt,omitempty"`
	EndAt       *time.Time `json:"endAt,omitempty"`
}

type ExpiredToken struct {
	Message string `json:"message"`
}

func (ExpiredToken) IsAcceptInviteResult() {}

func (ExpiredToken) IsConfirmLoginLinkResult() {}

func (ExpiredToken) IsError()                {}
func (this ExpiredToken) GetMessage() string { return this.Message }

func (ExpiredToken) IsConfirmLoginPinResult() {}

type File struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	URL      string   `json:"url"`
	Type     FileType `json:"type"`
	MimeType string   `json:"mimeType"`
}

func (File) IsCoverFileResult() {}

func (File) IsAlbumCoverResult() {}

type FileUploaded struct {
	File *File `json:"file"`
}

func (FileUploaded) IsUploadFileResult() {}

type Forbidden struct {
	Message string `json:"message"`
}

func (Forbidden) IsAddContactResult() {}

func (Forbidden) IsAddHouseResult() {}

func (Forbidden) IsAddRoomResult() {}

func (Forbidden) IsAddVisualizationsToAlbumResult() {}

func (Forbidden) IsChangeAlbumPageOrientationResult() {}

func (Forbidden) IsChangeAlbumPageSizeResult() {}

func (Forbidden) IsChangeProjectDatesResult() {}

func (Forbidden) IsChangeProjectStatusResult() {}

func (Forbidden) IsCreateAlbumResult() {}

func (Forbidden) IsCreateProjectResult() {}

func (Forbidden) IsDeleteAlbumResult() {}

func (Forbidden) IsDeleteContactResult() {}

func (Forbidden) IsDeleteRoomResult() {}

func (Forbidden) IsDeleteVisualizationsResult() {}

func (Forbidden) IsGenerateAlbumFileResult() {}

func (Forbidden) IsInviteUserToWorkspaceResult() {}

func (Forbidden) IsMakeProjectNotPublicResult() {}

func (Forbidden) IsMakeProjectPublicResult() {}

func (Forbidden) IsMoveRoomToPositionResult() {}

func (Forbidden) IsUpdateContactResult() {}

func (Forbidden) IsUpdateHouseResult() {}

func (Forbidden) IsUpdateRoomResult() {}

func (Forbidden) IsUploadFileResult() {}

func (Forbidden) IsUploadVisualizationResult() {}

func (Forbidden) IsUploadVisualizationsResult() {}

func (Forbidden) IsAlbumResult() {}

func (Forbidden) IsAlbumProjectResult() {}

func (Forbidden) IsAlbumPageCoverResult() {}

func (Forbidden) IsCoverFileResult() {}

func (Forbidden) IsAlbumRecentFileResult() {}

func (Forbidden) IsAlbumCoverResult() {}

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

func (Forbidden) IsProjectAlbumsListResult() {}

func (Forbidden) IsProjectAlbumsTotalResult() {}

func (Forbidden) IsWorkspaceResult() {}

func (Forbidden) IsWorkspaceProjectsListResult() {}

func (Forbidden) IsWorkspaceProjectsTotalResult() {}

func (Forbidden) IsWorkspaceUsersListResult() {}

func (Forbidden) IsWorkspaceUsersTotalResult() {}

func (Forbidden) IsError()                {}
func (this Forbidden) GetMessage() string { return this.Message }

func (Forbidden) IsAlbumFileGenerated() {}

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

func (InvalidToken) IsAcceptInviteResult() {}

func (InvalidToken) IsConfirmLoginLinkResult() {}

func (InvalidToken) IsError()                {}
func (this InvalidToken) GetMessage() string { return this.Message }

func (InvalidToken) IsConfirmLoginPinResult() {}

type InviteAccepted struct {
	Token string `json:"token"`
}

func (InviteAccepted) IsAcceptInviteResult() {}

type InviteSent struct {
	To string `json:"to"`
	//  token lifetime in seconds
	TokenExpiration int `json:"tokenExpiration"`
}

func (InviteSent) IsInviteUserToWorkspaceResult() {}

type LinkSentByEmail struct {
	Email string `json:"email"`
}

func (LinkSentByEmail) IsLoginByEmailResult() {}

type LoginConfirmed struct {
	Token string `json:"token"`
}

func (LoginConfirmed) IsConfirmLoginLinkResult() {}

func (LoginConfirmed) IsConfirmLoginPinResult() {}

type Mutation struct {
}

type NotFound struct {
	Message string `json:"message"`
}

func (NotFound) IsAddHouseResult() {}

func (NotFound) IsAddRoomResult() {}

func (NotFound) IsChangeAlbumPageOrientationResult() {}

func (NotFound) IsChangeAlbumPageSizeResult() {}

func (NotFound) IsChangeProjectDatesResult() {}

func (NotFound) IsChangeProjectStatusResult() {}

func (NotFound) IsDeleteAlbumResult() {}

func (NotFound) IsDeleteContactResult() {}

func (NotFound) IsDeleteRoomResult() {}

func (NotFound) IsDeleteVisualizationsResult() {}

func (NotFound) IsGenerateAlbumFileResult() {}

func (NotFound) IsInviteUserToWorkspaceResult() {}

func (NotFound) IsMakeProjectNotPublicResult() {}

func (NotFound) IsMakeProjectPublicResult() {}

func (NotFound) IsMoveRoomToPositionResult() {}

func (NotFound) IsUpdateContactResult() {}

func (NotFound) IsUpdateHouseResult() {}

func (NotFound) IsUpdateRoomResult() {}

func (NotFound) IsAlbumResult() {}

func (NotFound) IsAlbumProjectResult() {}

func (NotFound) IsAlbumPageCoverResult() {}

func (NotFound) IsCoverFileResult() {}

func (NotFound) IsAlbumPageVisualizationResult() {}

func (NotFound) IsAlbumRecentFileResult() {}

func (NotFound) IsAlbumCoverResult() {}

func (NotFound) IsProjectResult() {}

func (NotFound) IsProjectPublicSite() {}

func (NotFound) IsWorkspaceResult() {}

func (NotFound) IsError()                {}
func (this NotFound) GetMessage() string { return this.Message }

func (NotFound) IsAlbumFileGenerated() {}

type PinSentByEmail struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (PinSentByEmail) IsLoginByEmailResult() {}

type Project struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	Status         ProjectStatus            `json:"status"`
	StartAt        *time.Time               `json:"startAt,omitempty"`
	EndAt          *time.Time               `json:"endAt,omitempty"`
	Period         *string                  `json:"period,omitempty"`
	Contacts       *ProjectContacts         `json:"contacts"`
	Houses         *ProjectHouses           `json:"houses"`
	Visualizations *ProjectVisualizations   `json:"visualizations"`
	Files          *ProjectFiles            `json:"files"`
	Albums         *ProjectAlbums           `json:"albums"`
	PublicSite     ProjectPublicSite        `json:"publicSite"`
	Statuses       *ProjectStatusDictionary `json:"statuses"`
}

func (Project) IsAlbumProjectResult() {}

func (Project) IsProjectResult() {}

type ProjectAlbums struct {
	List  ProjectAlbumsListResult  `json:"list"`
	Total ProjectAlbumsTotalResult `json:"total"`
}

type ProjectAlbumsList struct {
	Items []*Album `json:"items"`
}

func (ProjectAlbumsList) IsProjectAlbumsListResult() {}

type ProjectAlbumsTotal struct {
	Total int `json:"total"`
}

func (ProjectAlbumsTotal) IsProjectAlbumsTotalResult() {}

type ProjectContacts struct {
	List  ProjectContactsListResult  `json:"list"`
	Total ProjectContactsTotalResult `json:"total"`
}

type ProjectContactsFilter struct {
	Type []ContactType `json:"type,omitempty"`
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

type ProjectFiles struct {
	List  ProjectFilesListResult  `json:"list"`
	Total ProjectFilesTotalResult `json:"total"`
}

type ProjectFilesList struct {
	Items []*File `json:"items"`
}

func (ProjectFilesList) IsProjectFilesListResult() {}

type ProjectFilesListFilter struct {
	Type []FileType `json:"type,omitempty"`
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
	ID []string `json:"ID,omitempty"`
}

type ProjectHousesList struct {
	Items []*House `json:"items"`
}

func (ProjectHousesList) IsProjectHousesListResult() {}

type ProjectHousesTotal struct {
	Total int `json:"total"`
}

func (ProjectHousesTotal) IsProjectHousesTotalResult() {}

type ProjectIsAlreadyNotPublic struct {
	Message string `json:"message"`
}

func (ProjectIsAlreadyNotPublic) IsMakeProjectNotPublicResult() {}

func (ProjectIsAlreadyNotPublic) IsError()                {}
func (this ProjectIsAlreadyNotPublic) GetMessage() string { return this.Message }

type ProjectIsAlreadyPublic struct {
	Message string `json:"message"`
}

func (ProjectIsAlreadyPublic) IsMakeProjectPublicResult() {}

func (ProjectIsAlreadyPublic) IsError()                {}
func (this ProjectIsAlreadyPublic) GetMessage() string { return this.Message }

type ProjectMadeNotPublic struct {
	PublicSite *PublicSite `json:"publicSite"`
}

func (ProjectMadeNotPublic) IsMakeProjectNotPublicResult() {}

type ProjectMadePublic struct {
	PublicSite *PublicSite `json:"publicSite"`
}

func (ProjectMadePublic) IsMakeProjectPublicResult() {}

type ProjectStatusChanged struct {
	Project *Project `json:"project"`
}

func (ProjectStatusChanged) IsChangeProjectStatusResult() {}

type ProjectStatusDictionary struct {
	Items []*ProjectStatusDictionaryItem `json:"items"`
}

type ProjectStatusDictionaryItem struct {
	Key   ProjectStatus `json:"key"`
	Value string        `json:"value"`
}

type ProjectVisualizationRoomIDFilter struct {
	Eq []string `json:"eq,omitempty"`
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
	RoomID *ProjectVisualizationRoomIDFilter  `json:"roomID,omitempty"`
	Status *ProjectVisualizationsStatusFilter `json:"status,omitempty"`
}

type ProjectVisualizationsStatusFilter struct {
	Eq []VisualizationStatus `json:"eq,omitempty"`
}

type ProjectVisualizationsTotal struct {
	Total int `json:"total"`
}

func (ProjectVisualizationsTotal) IsProjectVisualizationsTotalResult() {}

type PublicSite struct {
	ID       string              `json:"id"`
	Status   PublicSiteStatus    `json:"status"`
	URL      string              `json:"url"`
	Settings *PublicSiteSettings `json:"settings"`
}

func (PublicSite) IsProjectPublicSite() {}

type PublicSiteSettings struct {
	Visualizations bool `json:"visualizations"`
	Albums         bool `json:"albums"`
}

type Query struct {
}

type Room struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Square     *float64  `json:"square,omitempty"`
	Level      *int      `json:"level,omitempty"`
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

type RoomMovedToPosition struct {
	Room *Room `json:"room"`
}

func (RoomMovedToPosition) IsMoveRoomToPositionResult() {}

type RoomUpdated struct {
	Room *Room `json:"room"`
}

func (RoomUpdated) IsUpdateRoomResult() {}

type ServerError struct {
	Message string `json:"message"`
}

func (ServerError) IsAcceptInviteResult() {}

func (ServerError) IsAddContactResult() {}

func (ServerError) IsAddHouseResult() {}

func (ServerError) IsAddVisualizationsToAlbumResult() {}

func (ServerError) IsChangeAlbumPageOrientationResult() {}

func (ServerError) IsChangeAlbumPageSizeResult() {}

func (ServerError) IsChangeProjectDatesResult() {}

func (ServerError) IsChangeProjectStatusResult() {}

func (ServerError) IsConfirmLoginLinkResult() {}

func (ServerError) IsConfirmLoginPinResult() {}

func (ServerError) IsCreateAlbumResult() {}

func (ServerError) IsCreateProjectResult() {}

func (ServerError) IsDeleteAlbumResult() {}

func (ServerError) IsDeleteContactResult() {}

func (ServerError) IsDeleteVisualizationsResult() {}

func (ServerError) IsGenerateAlbumFileResult() {}

func (ServerError) IsInviteUserToWorkspaceResult() {}

func (ServerError) IsLoginByEmailResult() {}

func (ServerError) IsMakeProjectNotPublicResult() {}

func (ServerError) IsMakeProjectPublicResult() {}

func (ServerError) IsMoveRoomToPositionResult() {}

func (ServerError) IsUpdateContactResult() {}

func (ServerError) IsUpdateHouseResult() {}

func (ServerError) IsUploadFileResult() {}

func (ServerError) IsUploadVisualizationResult() {}

func (ServerError) IsUploadVisualizationsResult() {}

func (ServerError) IsAlbumResult() {}

func (ServerError) IsAlbumProjectResult() {}

func (ServerError) IsAlbumPagesResult() {}

func (ServerError) IsAlbumPageCoverResult() {}

func (ServerError) IsCoverFileResult() {}

func (ServerError) IsAlbumPageVisualizationResult() {}

func (ServerError) IsAlbumRecentFileResult() {}

func (ServerError) IsAlbumCoverResult() {}

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

func (ServerError) IsProjectAlbumsListResult() {}

func (ServerError) IsProjectAlbumsTotalResult() {}

func (ServerError) IsProjectPublicSite() {}

func (ServerError) IsWorkspaceResult() {}

func (ServerError) IsWorkspaceProjectsListResult() {}

func (ServerError) IsWorkspaceProjectsTotalResult() {}

func (ServerError) IsWorkspaceUsersListResult() {}

func (ServerError) IsWorkspaceUsersTotalResult() {}

func (ServerError) IsError()                {}
func (this ServerError) GetMessage() string { return this.Message }

func (ServerError) IsAlbumFileGenerated() {}

type SomeVisualizationsDeleted struct {
	Visualizations []*Visualization `json:"visualizations"`
}

func (SomeVisualizationsDeleted) IsDeleteVisualizationsResult() {}

type SomeVisualizationsUploaded struct {
	Visualizations []*Visualization `json:"visualizations"`
}

func (SomeVisualizationsUploaded) IsUploadVisualizationsResult() {}

type Subscription struct {
}

type Unknown struct {
	Message string `json:"message"`
}

func (Unknown) IsError()                {}
func (this Unknown) GetMessage() string { return this.Message }

func (Unknown) IsAlbumFileGenerated() {}

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
	Square *float64 `json:"square,omitempty"`
	Level  *int     `json:"level,omitempty"`
}

type UploadFileInput struct {
	ProjectID string         `json:"projectId"`
	Type      FileType       `json:"type"`
	Data      graphql.Upload `json:"data"`
}

type UserProfile struct {
	ID               string     `json:"id"`
	Email            string     `json:"email"`
	FullName         string     `json:"fullName"`
	Abbr             string     `json:"abbr"`
	Gravatar         *Gravatar  `json:"gravatar,omitempty"`
	DefaultWorkspace *Workspace `json:"defaultWorkspace"`
}

func (UserProfile) IsUserProfileResult() {}

type Visualization struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Version     int                 `json:"version"`
	Status      VisualizationStatus `json:"status"`
	CreatedAt   time.Time           `json:"createdAt"`
	ModifiedAt  time.Time           `json:"modifiedAt"`
	File        *File               `json:"file"`
	Room        *Room               `json:"room,omitempty"`
}

func (Visualization) IsAlbumPageVisualizationResult() {}

type VisualizationUploaded struct {
	Visualization *Visualization `json:"visualization"`
}

func (VisualizationUploaded) IsUploadVisualizationResult() {}

type VisualizationsAddedToAlbum struct {
	Album *Album                    `json:"album"`
	Pages []*AlbumPageVisualization `json:"pages"`
}

func (VisualizationsAddedToAlbum) IsAddVisualizationsToAlbumResult() {}

type VisualizationsDeleted struct {
	Visualizations []*Visualization `json:"visualizations"`
}

func (VisualizationsDeleted) IsDeleteVisualizationsResult() {}

type VisualizationsUploaded struct {
	Visualizations []*Visualization `json:"visualizations"`
}

func (VisualizationsUploaded) IsUploadVisualizationsResult() {}

type Workspace struct {
	ID       string                       `json:"id"`
	Name     string                       `json:"name"`
	Projects *WorkspaceProjects           `json:"projects"`
	Users    *WorkspaceUsers              `json:"users"`
	Roles    *WorkspaceUserRoleDictionary `json:"roles"`
}

func (Workspace) IsWorkspaceResult() {}

type WorkspaceProjects struct {
	List  WorkspaceProjectsListResult  `json:"list"`
	Total WorkspaceProjectsTotalResult `json:"total"`
}

type WorkspaceProjectsFilter struct {
	Status []ProjectStatus `json:"status,omitempty"`
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

type WorkspaceUserRoleDictionary struct {
	Items []*WorkspaceUserRoleDictionaryItem `json:"items"`
}

type WorkspaceUserRoleDictionaryItem struct {
	Key   WorkspaceUserRole `json:"key"`
	Value string            `json:"value"`
}

type WorkspaceUsers struct {
	List  WorkspaceUsersListResult  `json:"list"`
	Total WorkspaceUsersTotalResult `json:"total"`
}

type WorkspaceUsersFilter struct {
	Role []WorkspaceUserRole `json:"role,omitempty"`
}

type WorkspaceUsersList struct {
	Items []*WorkspaceUser `json:"items"`
}

func (WorkspaceUsersList) IsWorkspaceUsersListResult() {}

type WorkspaceUsersTotal struct {
	Total int `json:"total"`
}

func (WorkspaceUsersTotal) IsWorkspaceUsersTotalResult() {}

type AlbumFileStatus string

const (
	AlbumFileStatusNew                  AlbumFileStatus = "NEW"
	AlbumFileStatusGeneratingInProgress AlbumFileStatus = "GENERATING_IN_PROGRESS"
	AlbumFileStatusGeneratingDone       AlbumFileStatus = "GENERATING_DONE"
)

var AllAlbumFileStatus = []AlbumFileStatus{
	AlbumFileStatusNew,
	AlbumFileStatusGeneratingInProgress,
	AlbumFileStatusGeneratingDone,
}

func (e AlbumFileStatus) IsValid() bool {
	switch e {
	case AlbumFileStatusNew, AlbumFileStatusGeneratingInProgress, AlbumFileStatusGeneratingDone:
		return true
	}
	return false
}

func (e AlbumFileStatus) String() string {
	return string(e)
}

func (e *AlbumFileStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlbumFileStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlbumFileStatus", str)
	}
	return nil
}

func (e AlbumFileStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

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

type FileType string

const (
	FileTypeNone          FileType = "NONE"
	FileTypeVisualization FileType = "VISUALIZATION"
)

var AllFileType = []FileType{
	FileTypeNone,
	FileTypeVisualization,
}

func (e FileType) IsValid() bool {
	switch e {
	case FileTypeNone, FileTypeVisualization:
		return true
	}
	return false
}

func (e FileType) String() string {
	return string(e)
}

func (e *FileType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FileType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FileType", str)
	}
	return nil
}

func (e FileType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PageOrientation string

const (
	PageOrientationPortrait  PageOrientation = "PORTRAIT"
	PageOrientationLandscape PageOrientation = "LANDSCAPE"
)

var AllPageOrientation = []PageOrientation{
	PageOrientationPortrait,
	PageOrientationLandscape,
}

func (e PageOrientation) IsValid() bool {
	switch e {
	case PageOrientationPortrait, PageOrientationLandscape:
		return true
	}
	return false
}

func (e PageOrientation) String() string {
	return string(e)
}

func (e *PageOrientation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PageOrientation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PageOrientation", str)
	}
	return nil
}

func (e PageOrientation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PageSize string

const (
	PageSizeA4 PageSize = "A4"
	PageSizeA3 PageSize = "A3"
)

var AllPageSize = []PageSize{
	PageSizeA4,
	PageSizeA3,
}

func (e PageSize) IsValid() bool {
	switch e {
	case PageSizeA4, PageSizeA3:
		return true
	}
	return false
}

func (e PageSize) String() string {
	return string(e)
}

func (e *PageSize) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PageSize(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PageSize", str)
	}
	return nil
}

func (e PageSize) MarshalGQL(w io.Writer) {
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

type PublicSiteStatus string

const (
	PublicSiteStatusPublic    PublicSiteStatus = "PUBLIC"
	PublicSiteStatusNotPublic PublicSiteStatus = "NOT_PUBLIC"
)

var AllPublicSiteStatus = []PublicSiteStatus{
	PublicSiteStatusPublic,
	PublicSiteStatusNotPublic,
}

func (e PublicSiteStatus) IsValid() bool {
	switch e {
	case PublicSiteStatusPublic, PublicSiteStatusNotPublic:
		return true
	}
	return false
}

func (e PublicSiteStatus) String() string {
	return string(e)
}

func (e *PublicSiteStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PublicSiteStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PublicSiteStatus", str)
	}
	return nil
}

func (e PublicSiteStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VisualizationStatus string

const (
	VisualizationStatusUnknown  VisualizationStatus = "UNKNOWN"
	VisualizationStatusApproved VisualizationStatus = "APPROVED"
	VisualizationStatusDeleted  VisualizationStatus = "DELETED"
)

var AllVisualizationStatus = []VisualizationStatus{
	VisualizationStatusUnknown,
	VisualizationStatusApproved,
	VisualizationStatusDeleted,
}

func (e VisualizationStatus) IsValid() bool {
	switch e {
	case VisualizationStatusUnknown, VisualizationStatusApproved, VisualizationStatusDeleted:
		return true
	}
	return false
}

func (e VisualizationStatus) String() string {
	return string(e)
}

func (e *VisualizationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VisualizationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VisualizationStatus", str)
	}
	return nil
}

func (e VisualizationStatus) MarshalGQL(w io.Writer) {
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
