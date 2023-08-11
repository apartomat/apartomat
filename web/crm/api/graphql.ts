import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
  Upload: any;
  Url: any;
};

export type AcceptInviteResult = AlreadyInWorkspace | ExpiredToken | InvalidToken | InviteAccepted | ServerError;

export type AddContactInput = {
  details: Array<ContactDetailsInput>;
  fullName: Scalars['String'];
};

export type AddContactResult = ContactAdded | Forbidden | ServerError;

export type AddHouseInput = {
  address: Scalars['String'];
  city: Scalars['String'];
  housingComplex: Scalars['String'];
};

export type AddHouseResult = Forbidden | HouseAdded | NotFound | ServerError;

export type AddRoomInput = {
  level?: InputMaybe<Scalars['Int']>;
  name: Scalars['String'];
  square?: InputMaybe<Scalars['Float']>;
};

export type AddRoomResult = Forbidden | NotFound | RoomAdded;

export type AddVisualizationsToAlbumResult = Forbidden | ServerError | VisualizationsAddedToAlbum;

export type Album = {
  __typename?: 'Album';
  id: Scalars['String'];
  name: Scalars['String'];
  pages: AlbumPagesResult;
  project: AlbumProjectResult;
  settings: AlbumSettings;
};

export type AlbumCreated = {
  __typename?: 'AlbumCreated';
  album: Album;
};

export type AlbumDeleted = {
  __typename?: 'AlbumDeleted';
  album: Album;
};

export type AlbumPage = AlbumPageCover | AlbumPageVisualization;

export type AlbumPageCover = {
  __typename?: 'AlbumPageCover';
  position: Scalars['Int'];
};

export type AlbumPageOrientationChanged = {
  __typename?: 'AlbumPageOrientationChanged';
  album: Album;
};

export type AlbumPageSizeChanged = {
  __typename?: 'AlbumPageSizeChanged';
  album: Album;
};

export type AlbumPageVisualization = {
  __typename?: 'AlbumPageVisualization';
  position: Scalars['Int'];
  visualization: AlbumPageVisualizationResult;
};

export type AlbumPageVisualizationResult = NotFound | ServerError | Visualization;

export type AlbumPages = {
  __typename?: 'AlbumPages';
  items: Array<AlbumPage>;
};

export type AlbumPagesResult = AlbumPages | ServerError;

export type AlbumProjectResult = Forbidden | NotFound | Project | ServerError;

export type AlbumResult = Album | Forbidden | NotFound | ServerError;

export type AlbumSettings = {
  __typename?: 'AlbumSettings';
  pageOrientation: PageOrientation;
  pageSize: PageSize;
};

export type AlreadyExists = Error & {
  __typename?: 'AlreadyExists';
  message: Scalars['String'];
};

export type AlreadyInWorkspace = Error & {
  __typename?: 'AlreadyInWorkspace';
  message: Scalars['String'];
};

export type ChangeAlbumPageOrientationResult = AlbumPageOrientationChanged | Forbidden | NotFound | ServerError;

export type ChangeAlbumPageSizeResult = AlbumPageSizeChanged | Forbidden | NotFound | ServerError;

export type ChangeProjectDatesInput = {
  endAt?: InputMaybe<Scalars['Time']>;
  startAt?: InputMaybe<Scalars['Time']>;
};

export type ChangeProjectDatesResult = Forbidden | NotFound | ProjectDatesChanged | ServerError;

export type ChangeProjectStatusResult = Forbidden | NotFound | ProjectStatusChanged | ServerError;

export type ConfirmLoginLinkResult = ExpiredToken | InvalidToken | LoginConfirmed | ServerError;

export type ConfirmLoginPinResult = ExpiredToken | InvalidPin | InvalidToken | LoginConfirmed | ServerError;

export type Contact = {
  __typename?: 'Contact';
  createdAt: Scalars['Time'];
  details: Array<ContactDetails>;
  fullName: Scalars['String'];
  id: Scalars['String'];
  modifiedAt: Scalars['Time'];
  photo: Scalars['String'];
};

export type ContactAdded = {
  __typename?: 'ContactAdded';
  contact: Contact;
};

export type ContactDeleted = {
  __typename?: 'ContactDeleted';
  contact: Contact;
};

export type ContactDetails = {
  __typename?: 'ContactDetails';
  type: ContactType;
  value: Scalars['String'];
};

export type ContactDetailsInput = {
  type: ContactType;
  value: Scalars['String'];
};

export enum ContactType {
  Email = 'EMAIL',
  Instagram = 'INSTAGRAM',
  Phone = 'PHONE',
  Telegram = 'TELEGRAM',
  Unknown = 'UNKNOWN',
  Whatsapp = 'WHATSAPP'
}

export type ContactUpdated = {
  __typename?: 'ContactUpdated';
  contact: Contact;
};

export type CreateAlbumResult = AlbumCreated | Forbidden | ServerError;

export type CreateAlbumSettingsInput = {
  orientation: PageOrientation;
  pageSize: PageSize;
};

export type CreateProjectInput = {
  endAt?: InputMaybe<Scalars['Time']>;
  name: Scalars['String'];
  startAt?: InputMaybe<Scalars['Time']>;
  workspaceId: Scalars['String'];
};

export type CreateProjectResult = Forbidden | ProjectCreated | ServerError;

export type DeleteAlbumResult = AlbumDeleted | Forbidden | NotFound | ServerError;

export type DeleteContactResult = ContactDeleted | Forbidden | NotFound | ServerError;

export type DeleteRoomResult = Forbidden | NotFound | RoomDeleted;

export type DeleteVisualizationsResult = Forbidden | NotFound | ServerError | SomeVisualizationsDeleted | VisualizationsDeleted;

export type Error = {
  message: Scalars['String'];
};

export type ExpiredToken = Error & {
  __typename?: 'ExpiredToken';
  message: Scalars['String'];
};

export type File = {
  __typename?: 'File';
  id: Scalars['String'];
  mimeType: Scalars['String'];
  name: Scalars['String'];
  type: FileType;
  url: Scalars['Url'];
};

export enum FileType {
  None = 'NONE',
  Visualization = 'VISUALIZATION'
}

export type FileUploaded = {
  __typename?: 'FileUploaded';
  file: File;
};

export type Forbidden = Error & {
  __typename?: 'Forbidden';
  message: Scalars['String'];
};

export type Gravatar = {
  __typename?: 'Gravatar';
  url: Scalars['String'];
};

export type House = {
  __typename?: 'House';
  address: Scalars['String'];
  city: Scalars['String'];
  createdAt: Scalars['Time'];
  housingComplex: Scalars['String'];
  id: Scalars['String'];
  modifiedAt: Scalars['Time'];
  rooms: HouseRooms;
};

export type HouseAdded = {
  __typename?: 'HouseAdded';
  house: House;
};

export type HouseRooms = {
  __typename?: 'HouseRooms';
  list: HouseRoomsListResult;
};


export type HouseRoomsListArgs = {
  limit?: Scalars['Int'];
  offset?: Scalars['Int'];
};

export type HouseRoomsList = {
  __typename?: 'HouseRoomsList';
  items: Array<Room>;
};

export type HouseRoomsListResult = Forbidden | HouseRoomsList | ServerError;

export type HouseUpdated = {
  __typename?: 'HouseUpdated';
  house: House;
};

export type Id = {
  __typename?: 'Id';
  id: Scalars['String'];
};

export type InvalidEmail = Error & {
  __typename?: 'InvalidEmail';
  message: Scalars['String'];
};

export type InvalidPin = Error & {
  __typename?: 'InvalidPin';
  message: Scalars['String'];
};

export type InvalidToken = Error & {
  __typename?: 'InvalidToken';
  message: Scalars['String'];
};

export type InviteAccepted = {
  __typename?: 'InviteAccepted';
  token: Scalars['String'];
};

export type InviteSent = {
  __typename?: 'InviteSent';
  to: Scalars['String'];
  /**  token lifetime in seconds  */
  tokenExpiration: Scalars['Int'];
};

export type InviteUserToWorkspaceResult = AlreadyInWorkspace | Forbidden | InviteSent | NotFound | ServerError;

export type LinkSentByEmail = {
  __typename?: 'LinkSentByEmail';
  email: Scalars['String'];
};

export type LoginByEmailResult = InvalidEmail | LinkSentByEmail | PinSentByEmail | ServerError;

export type LoginConfirmed = {
  __typename?: 'LoginConfirmed';
  token: Scalars['String'];
};

export type MakeProjectNotPublicResult = Forbidden | NotFound | ProjectIsAlreadyNotPublic | ProjectMadeNotPublic | ServerError;

export type MakeProjectPublicResult = Forbidden | NotFound | ProjectIsAlreadyPublic | ProjectMadePublic | ServerError;

export type Mutation = {
  __typename?: 'Mutation';
  acceptInvite: AcceptInviteResult;
  addContact: AddContactResult;
  addHouse: AddHouseResult;
  addRoom: AddRoomResult;
  addVisualizationsToAlbum: AddVisualizationsToAlbumResult;
  changeAlbumPageOrientation: ChangeAlbumPageOrientationResult;
  changeAlbumPageSize: ChangeAlbumPageSizeResult;
  changeProjectDates: ChangeProjectDatesResult;
  changeProjectStatus: ChangeProjectStatusResult;
  confirmLoginLink: ConfirmLoginLinkResult;
  confirmLoginPin: ConfirmLoginPinResult;
  createAlbum: CreateAlbumResult;
  createProject: CreateProjectResult;
  deleteAlbum: DeleteAlbumResult;
  deleteContact: DeleteContactResult;
  deleteRoom: DeleteRoomResult;
  deleteVisualizations: DeleteVisualizationsResult;
  inviteUser: InviteUserToWorkspaceResult;
  loginByEmail: LoginByEmailResult;
  makeProjectNotPublic: MakeProjectNotPublicResult;
  makeProjectPublic: MakeProjectPublicResult;
  ping: Scalars['String'];
  updateContact: UpdateContactResult;
  updateHouse: UpdateHouseResult;
  updateRoom: UpdateRoomResult;
  uploadFile: UploadFileResult;
  uploadVisualization: UploadVisualizationResult;
  uploadVisualizations: UploadVisualizationsResult;
};


export type MutationAcceptInviteArgs = {
  token: Scalars['String'];
};


export type MutationAddContactArgs = {
  contact: AddContactInput;
  projectId: Scalars['String'];
};


export type MutationAddHouseArgs = {
  house: AddHouseInput;
  projectId: Scalars['String'];
};


export type MutationAddRoomArgs = {
  houseId: Scalars['String'];
  room: AddRoomInput;
};


export type MutationAddVisualizationsToAlbumArgs = {
  albumId: Scalars['String'];
  visualizations: Array<Scalars['String']>;
};


export type MutationChangeAlbumPageOrientationArgs = {
  albumId: Scalars['String'];
  orientation: PageOrientation;
};


export type MutationChangeAlbumPageSizeArgs = {
  albumId: Scalars['String'];
  size: PageSize;
};


export type MutationChangeProjectDatesArgs = {
  dates: ChangeProjectDatesInput;
  projectId: Scalars['String'];
};


export type MutationChangeProjectStatusArgs = {
  projectId: Scalars['String'];
  status: ProjectStatus;
};


export type MutationConfirmLoginLinkArgs = {
  token: Scalars['String'];
};


export type MutationConfirmLoginPinArgs = {
  pin: Scalars['String'];
  token: Scalars['String'];
};


export type MutationCreateAlbumArgs = {
  name: Scalars['String'];
  projectId: Scalars['String'];
  settings?: CreateAlbumSettingsInput;
};


export type MutationCreateProjectArgs = {
  input: CreateProjectInput;
};


export type MutationDeleteAlbumArgs = {
  id: Scalars['String'];
};


export type MutationDeleteContactArgs = {
  id: Scalars['String'];
};


export type MutationDeleteRoomArgs = {
  id: Scalars['String'];
};


export type MutationDeleteVisualizationsArgs = {
  id: Array<Scalars['String']>;
};


export type MutationInviteUserArgs = {
  email: Scalars['String'];
  role: WorkspaceUserRole;
  workspaceId: Scalars['String'];
};


export type MutationLoginByEmailArgs = {
  email: Scalars['String'];
  workspaceName?: Scalars['String'];
};


export type MutationMakeProjectNotPublicArgs = {
  projectId: Scalars['String'];
};


export type MutationMakeProjectPublicArgs = {
  projectId: Scalars['String'];
};


export type MutationUpdateContactArgs = {
  contactId: Scalars['String'];
  data: UpdateContactInput;
};


export type MutationUpdateHouseArgs = {
  data: UpdateHouseInput;
  houseId: Scalars['String'];
};


export type MutationUpdateRoomArgs = {
  data: UpdateRoomInput;
  roomId: Scalars['String'];
};


export type MutationUploadFileArgs = {
  input: UploadFileInput;
};


export type MutationUploadVisualizationArgs = {
  file: Scalars['Upload'];
  projectId: Scalars['String'];
  roomId?: InputMaybe<Scalars['String']>;
};


export type MutationUploadVisualizationsArgs = {
  files: Array<Scalars['Upload']>;
  projectId: Scalars['String'];
  roomId?: InputMaybe<Scalars['String']>;
};

export type NotFound = Error & {
  __typename?: 'NotFound';
  message: Scalars['String'];
};

export enum PageOrientation {
  Landscape = 'LANDSCAPE',
  Portrait = 'PORTRAIT'
}

export enum PageSize {
  A3 = 'A3',
  A4 = 'A4'
}

export type PinSentByEmail = {
  __typename?: 'PinSentByEmail';
  email: Scalars['String'];
  token: Scalars['String'];
};

export type Project = {
  __typename?: 'Project';
  albums: ProjectAlbums;
  contacts: ProjectContacts;
  endAt?: Maybe<Scalars['Time']>;
  files: ProjectFiles;
  houses: ProjectHouses;
  id: Scalars['String'];
  name: Scalars['String'];
  period?: Maybe<Scalars['String']>;
  publicSite: ProjectPublicSite;
  startAt?: Maybe<Scalars['Time']>;
  status: ProjectStatus;
  statuses: ProjectStatusDictionary;
  visualizations: ProjectVisualizations;
};


export type ProjectPeriodArgs = {
  timezone?: InputMaybe<Scalars['String']>;
};

export type ProjectAlbums = {
  __typename?: 'ProjectAlbums';
  list: ProjectAlbumsListResult;
  total: ProjectAlbumsTotalResult;
};


export type ProjectAlbumsListArgs = {
  limit?: Scalars['Int'];
  offset?: Scalars['Int'];
};

export type ProjectAlbumsList = {
  __typename?: 'ProjectAlbumsList';
  items: Array<Album>;
};

export type ProjectAlbumsListResult = Forbidden | ProjectAlbumsList | ServerError;

export type ProjectAlbumsTotal = {
  __typename?: 'ProjectAlbumsTotal';
  total: Scalars['Int'];
};

export type ProjectAlbumsTotalResult = Forbidden | ProjectAlbumsTotal | ServerError;

export type ProjectContacts = {
  __typename?: 'ProjectContacts';
  list: ProjectContactsListResult;
  total: ProjectContactsTotalResult;
};


export type ProjectContactsListArgs = {
  filter?: ProjectContactsFilter;
  limit?: Scalars['Int'];
  offset?: Scalars['Int'];
};


export type ProjectContactsTotalArgs = {
  filter?: ProjectContactsFilter;
};

export type ProjectContactsFilter = {
  type?: InputMaybe<Array<ContactType>>;
};

export type ProjectContactsList = {
  __typename?: 'ProjectContactsList';
  items: Array<Contact>;
};

export type ProjectContactsListResult = Forbidden | ProjectContactsList | ServerError;

export type ProjectContactsTotal = {
  __typename?: 'ProjectContactsTotal';
  total: Scalars['Int'];
};

export type ProjectContactsTotalResult = Forbidden | ProjectContactsTotal | ServerError;

export type ProjectCreated = {
  __typename?: 'ProjectCreated';
  project: Project;
};

export type ProjectDatesChanged = {
  __typename?: 'ProjectDatesChanged';
  project: Project;
};

export type ProjectFiles = {
  __typename?: 'ProjectFiles';
  list: ProjectFilesListResult;
  total: ProjectFilesTotalResult;
};


export type ProjectFilesListArgs = {
  filter?: ProjectFilesListFilter;
  limit?: Scalars['Int'];
  offset?: Scalars['Int'];
};


export type ProjectFilesTotalArgs = {
  filter?: ProjectFilesListFilter;
};

export type ProjectFilesList = {
  __typename?: 'ProjectFilesList';
  items: Array<File>;
};

export type ProjectFilesListFilter = {
  type?: InputMaybe<Array<FileType>>;
};

export type ProjectFilesListResult = Forbidden | ProjectFilesList | ServerError;

export type ProjectFilesTotal = {
  __typename?: 'ProjectFilesTotal';
  total: Scalars['Int'];
};

export type ProjectFilesTotalResult = Forbidden | ProjectFilesTotal | ServerError;

export type ProjectHouses = {
  __typename?: 'ProjectHouses';
  list: ProjectHousesListResult;
  total: ProjectHousesTotalResult;
};


export type ProjectHousesListArgs = {
  filter?: ProjectHousesFilter;
  limit?: Scalars['Int'];
  offset?: Scalars['Int'];
};


export type ProjectHousesTotalArgs = {
  filter?: ProjectHousesFilter;
};

export type ProjectHousesFilter = {
  ID?: InputMaybe<Array<Scalars['String']>>;
};

export type ProjectHousesList = {
  __typename?: 'ProjectHousesList';
  items: Array<House>;
};

export type ProjectHousesListResult = Forbidden | ProjectHousesList | ServerError;

export type ProjectHousesTotal = {
  __typename?: 'ProjectHousesTotal';
  total: Scalars['Int'];
};

export type ProjectHousesTotalResult = Forbidden | ProjectHousesTotal | ServerError;

export type ProjectIsAlreadyNotPublic = Error & {
  __typename?: 'ProjectIsAlreadyNotPublic';
  message: Scalars['String'];
};

export type ProjectIsAlreadyPublic = Error & {
  __typename?: 'ProjectIsAlreadyPublic';
  message: Scalars['String'];
};

export type ProjectMadeNotPublic = {
  __typename?: 'ProjectMadeNotPublic';
  publicSite: PublicSite;
};

export type ProjectMadePublic = {
  __typename?: 'ProjectMadePublic';
  publicSite: PublicSite;
};

export type ProjectPublicSite = NotFound | PublicSite | ServerError;

export type ProjectResult = Forbidden | NotFound | Project | ServerError;

export enum ProjectStatus {
  Canceled = 'CANCELED',
  Done = 'DONE',
  InProgress = 'IN_PROGRESS',
  New = 'NEW'
}

export type ProjectStatusChanged = {
  __typename?: 'ProjectStatusChanged';
  project: Project;
};

export type ProjectStatusDictionary = {
  __typename?: 'ProjectStatusDictionary';
  items: Array<ProjectStatusDictionaryItem>;
};

export type ProjectStatusDictionaryItem = {
  __typename?: 'ProjectStatusDictionaryItem';
  key: ProjectStatus;
  value: Scalars['String'];
};

export type ProjectVisualizationRoomIdFilter = {
  eq?: InputMaybe<Array<Scalars['String']>>;
};

export type ProjectVisualizations = {
  __typename?: 'ProjectVisualizations';
  list: ProjectVisualizationsListResult;
  total: ProjectVisualizationsTotalResult;
};


export type ProjectVisualizationsListArgs = {
  filter?: ProjectVisualizationsListFilter;
  limit?: Scalars['Int'];
  offset?: Scalars['Int'];
};


export type ProjectVisualizationsTotalArgs = {
  filter?: ProjectVisualizationsListFilter;
};

export type ProjectVisualizationsList = {
  __typename?: 'ProjectVisualizationsList';
  items: Array<Visualization>;
};

export type ProjectVisualizationsListFilter = {
  roomID?: InputMaybe<ProjectVisualizationRoomIdFilter>;
  status?: InputMaybe<ProjectVisualizationsStatusFilter>;
};

export type ProjectVisualizationsListResult = Forbidden | ProjectVisualizationsList | ServerError;

export type ProjectVisualizationsStatusFilter = {
  eq?: InputMaybe<Array<VisualizationStatus>>;
};

export type ProjectVisualizationsTotal = {
  __typename?: 'ProjectVisualizationsTotal';
  total: Scalars['Int'];
};

export type ProjectVisualizationsTotalResult = Forbidden | ProjectVisualizationsTotal | ServerError;

export type PublicSite = {
  __typename?: 'PublicSite';
  id: Scalars['String'];
  settings: PublicSiteSettings;
  status: PublicSiteStatus;
  url: Scalars['String'];
};

export type PublicSiteSettings = {
  __typename?: 'PublicSiteSettings';
  albums: Scalars['Boolean'];
  visualizations: Scalars['Boolean'];
};

export enum PublicSiteStatus {
  NotPublic = 'NOT_PUBLIC',
  Public = 'PUBLIC'
}

export type Query = {
  __typename?: 'Query';
  album: AlbumResult;
  profile: UserProfileResult;
  project: ProjectResult;
  version: Scalars['String'];
  workspace: WorkspaceResult;
};


export type QueryAlbumArgs = {
  id: Scalars['String'];
};


export type QueryProjectArgs = {
  id: Scalars['String'];
};


export type QueryWorkspaceArgs = {
  id: Scalars['String'];
};

export type Room = {
  __typename?: 'Room';
  createdAt: Scalars['Time'];
  id: Scalars['String'];
  level?: Maybe<Scalars['Int']>;
  modifiedAt: Scalars['Time'];
  name: Scalars['String'];
  square?: Maybe<Scalars['Float']>;
};

export type RoomAdded = {
  __typename?: 'RoomAdded';
  room: Room;
};

export type RoomDeleted = {
  __typename?: 'RoomDeleted';
  room: Room;
};

export type RoomUpdated = {
  __typename?: 'RoomUpdated';
  room: Room;
};

export type ServerError = Error & {
  __typename?: 'ServerError';
  message: Scalars['String'];
};

export type SomeVisualizationsDeleted = {
  __typename?: 'SomeVisualizationsDeleted';
  visualizations: Array<Visualization>;
};

export type SomeVisualizationsUploaded = {
  __typename?: 'SomeVisualizationsUploaded';
  visualizations: Array<Visualization>;
};

export type UpdateContactInput = {
  details: Array<ContactDetailsInput>;
  fullName: Scalars['String'];
};

export type UpdateContactResult = ContactUpdated | Forbidden | NotFound | ServerError;

export type UpdateHouseInput = {
  address: Scalars['String'];
  city: Scalars['String'];
  housingComplex: Scalars['String'];
};

export type UpdateHouseResult = Forbidden | HouseUpdated | NotFound | ServerError;

export type UpdateRoomInput = {
  level?: InputMaybe<Scalars['Int']>;
  name: Scalars['String'];
  square?: InputMaybe<Scalars['Float']>;
};

export type UpdateRoomResult = Forbidden | NotFound | RoomUpdated;

export type UploadFileInput = {
  data: Scalars['Upload'];
  projectId: Scalars['String'];
  type: FileType;
};

export type UploadFileResult = AlreadyExists | FileUploaded | Forbidden | ServerError;

export type UploadVisualizationResult = Forbidden | ServerError | VisualizationUploaded;

export type UploadVisualizationsResult = Forbidden | ServerError | SomeVisualizationsUploaded | VisualizationsUploaded;

export type UserProfile = {
  __typename?: 'UserProfile';
  abbr: Scalars['String'];
  defaultWorkspace: Workspace;
  email: Scalars['String'];
  fullName: Scalars['String'];
  gravatar?: Maybe<Gravatar>;
  id: Scalars['String'];
};

export type UserProfileResult = Forbidden | ServerError | UserProfile;

export type Visualization = {
  __typename?: 'Visualization';
  createdAt: Scalars['Time'];
  description: Scalars['String'];
  file: File;
  id: Scalars['String'];
  modifiedAt: Scalars['Time'];
  name: Scalars['String'];
  room?: Maybe<Room>;
  status: VisualizationStatus;
  version: Scalars['Int'];
};

export enum VisualizationStatus {
  Approved = 'APPROVED',
  Deleted = 'DELETED',
  Unknown = 'UNKNOWN'
}

export type VisualizationUploaded = {
  __typename?: 'VisualizationUploaded';
  visualization: Visualization;
};

export type VisualizationsAddedToAlbum = {
  __typename?: 'VisualizationsAddedToAlbum';
  pages: Array<AlbumPageVisualization>;
};

export type VisualizationsDeleted = {
  __typename?: 'VisualizationsDeleted';
  visualizations: Array<Visualization>;
};

export type VisualizationsUploaded = {
  __typename?: 'VisualizationsUploaded';
  visualizations: Array<Visualization>;
};

export type Workspace = {
  __typename?: 'Workspace';
  id: Scalars['String'];
  name: Scalars['String'];
  projects: WorkspaceProjects;
  roles: WorkspaceUserRoleDictionary;
  users: WorkspaceUsers;
};

export type WorkspaceProjects = {
  __typename?: 'WorkspaceProjects';
  list: WorkspaceProjectsListResult;
  total: WorkspaceProjectsTotalResult;
};


export type WorkspaceProjectsListArgs = {
  filter?: WorkspaceProjectsFilter;
  limit?: Scalars['Int'];
};


export type WorkspaceProjectsTotalArgs = {
  filter?: WorkspaceProjectsFilter;
};

export type WorkspaceProjectsFilter = {
  status?: InputMaybe<Array<ProjectStatus>>;
};

export type WorkspaceProjectsList = {
  __typename?: 'WorkspaceProjectsList';
  items: Array<Project>;
};

export type WorkspaceProjectsListResult = Forbidden | ServerError | WorkspaceProjectsList;

export type WorkspaceProjectsTotal = {
  __typename?: 'WorkspaceProjectsTotal';
  total: Scalars['Int'];
};

export type WorkspaceProjectsTotalResult = Forbidden | ServerError | WorkspaceProjectsTotal;

export type WorkspaceResult = Forbidden | NotFound | ServerError | Workspace;

export type WorkspaceUser = {
  __typename?: 'WorkspaceUser';
  id: Scalars['String'];
  profile: UserProfile;
  role: WorkspaceUserRole;
  workspace: Id;
};

export enum WorkspaceUserRole {
  Admin = 'ADMIN',
  User = 'USER'
}

export type WorkspaceUserRoleDictionary = {
  __typename?: 'WorkspaceUserRoleDictionary';
  items: Array<WorkspaceUserRoleDictionaryItem>;
};

export type WorkspaceUserRoleDictionaryItem = {
  __typename?: 'WorkspaceUserRoleDictionaryItem';
  key: WorkspaceUserRole;
  value: Scalars['String'];
};

export type WorkspaceUsers = {
  __typename?: 'WorkspaceUsers';
  list: WorkspaceUsersListResult;
  total: WorkspaceUsersTotalResult;
};


export type WorkspaceUsersListArgs = {
  filter?: WorkspaceUsersFilter;
  limit?: Scalars['Int'];
};


export type WorkspaceUsersTotalArgs = {
  filter?: WorkspaceUsersFilter;
};

export type WorkspaceUsersFilter = {
  role?: InputMaybe<Array<WorkspaceUserRole>>;
};

export type WorkspaceUsersList = {
  __typename?: 'WorkspaceUsersList';
  items: Array<WorkspaceUser>;
};

export type WorkspaceUsersListResult = Forbidden | ServerError | WorkspaceUsersList;

export type WorkspaceUsersTotal = {
  __typename?: 'WorkspaceUsersTotal';
  total: Scalars['Int'];
};

export type WorkspaceUsersTotalResult = Forbidden | ServerError | WorkspaceUsersTotal;

export type ProfileQueryVariables = Exact<{ [key: string]: never; }>;


export type ProfileQuery = { __typename?: 'Query', profile: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'UserProfile', id: string, email: string, gravatar?: { __typename?: 'Gravatar', url: string } | null, defaultWorkspace: { __typename?: 'Workspace', id: string, name: string } } };

export type AcceptInviteMutationVariables = Exact<{
  token: Scalars['String'];
}>;


export type AcceptInviteMutation = { __typename?: 'Mutation', acceptInvite: { __typename: 'AlreadyInWorkspace', message: string } | { __typename: 'ExpiredToken', message: string } | { __typename: 'InvalidToken', message: string } | { __typename: 'InviteAccepted', token: string } | { __typename: 'ServerError', message: string } };

export type ChangeAlbumPageOrientationMutationVariables = Exact<{
  albumId: Scalars['String'];
  orientation: PageOrientation;
}>;


export type ChangeAlbumPageOrientationMutation = { __typename?: 'Mutation', changeAlbumPageOrientation: { __typename: 'AlbumPageOrientationChanged', album: { __typename?: 'Album', settings: { __typename?: 'AlbumSettings', pageOrientation: PageOrientation } } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type ChangeAlbumPageSizeMutationVariables = Exact<{
  albumId: Scalars['String'];
  size: PageSize;
}>;


export type ChangeAlbumPageSizeMutation = { __typename?: 'Mutation', changeAlbumPageSize: { __typename: 'AlbumPageSizeChanged', album: { __typename?: 'Album', settings: { __typename?: 'AlbumSettings', pageSize: PageSize } } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type AddVisualizationsToAlbumMutationVariables = Exact<{
  albumId: Scalars['String'];
  visualizations: Array<Scalars['String']> | Scalars['String'];
}>;


export type AddVisualizationsToAlbumMutation = { __typename?: 'Mutation', addVisualizationsToAlbum: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'VisualizationsAddedToAlbum', pages: Array<{ __typename?: 'AlbumPageVisualization', position: number, visualization: { __typename?: 'NotFound' } | { __typename?: 'ServerError' } | { __typename?: 'Visualization', id: string, file: { __typename?: 'File', url: any } } }> } };

export type AlbumScreenQueryVariables = Exact<{
  id: Scalars['String'];
  filter: ProjectVisualizationsListFilter;
}>;


export type AlbumScreenQuery = { __typename?: 'Query', album: { __typename: 'Album', id: string, name: string, project: { __typename?: 'Forbidden', message: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'Project', id: string, name: string, visualizations: { __typename?: 'ProjectVisualizations', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'ProjectVisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null }> } | { __typename?: 'ServerError', message: string } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } } } | { __typename?: 'ServerError', message: string }, pages: { __typename?: 'AlbumPages', items: Array<{ __typename?: 'AlbumPageCover', position: number } | { __typename?: 'AlbumPageVisualization', position: number, visualization: { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'Visualization', id: string, file: { __typename?: 'File', url: any } } }> } | { __typename?: 'ServerError', message: string }, settings: { __typename?: 'AlbumSettings', pageSize: PageSize, pageOrientation: PageOrientation } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type AlbumScreenAlbumFragment = { __typename?: 'Album', id: string, name: string, project: { __typename?: 'Forbidden', message: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'Project', id: string, name: string, visualizations: { __typename?: 'ProjectVisualizations', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'ProjectVisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null }> } | { __typename?: 'ServerError', message: string } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } } } | { __typename?: 'ServerError', message: string }, pages: { __typename?: 'AlbumPages', items: Array<{ __typename?: 'AlbumPageCover', position: number } | { __typename?: 'AlbumPageVisualization', position: number, visualization: { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'Visualization', id: string, file: { __typename?: 'File', url: any } } }> } | { __typename?: 'ServerError', message: string }, settings: { __typename?: 'AlbumSettings', pageSize: PageSize, pageOrientation: PageOrientation } };

export type AlbumScreenSettingsFragment = { __typename?: 'AlbumSettings', pageSize: PageSize, pageOrientation: PageOrientation };

export type AlbumScreenProjectFragment = { __typename?: 'Project', id: string, name: string, visualizations: { __typename?: 'ProjectVisualizations', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'ProjectVisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null }> } | { __typename?: 'ServerError', message: string } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } } };

export type AlbumScreenVisualizationFragment = { __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null };

export type AlbumScreenHouseRoomFragment = { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null };

export type AlbumScreenHousesFragment = { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } };

export type AlbumScreenHouseRoomsFragment = { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } };

export type AlbumScreenAlbumPageCoverFragment = { __typename?: 'AlbumPageCover', position: number };

export type AlbumScreenAlbumPageVisualizationFragment = { __typename?: 'AlbumPageVisualization', position: number, visualization: { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'Visualization', id: string, file: { __typename?: 'File', url: any } } };

export type ConfirmLoginLinkMutationVariables = Exact<{
  token: Scalars['String'];
}>;


export type ConfirmLoginLinkMutation = { __typename?: 'Mutation', confirmLoginLink: { __typename: 'ExpiredToken', message: string } | { __typename: 'InvalidToken', message: string } | { __typename: 'LoginConfirmed', token: string } | { __typename: 'ServerError', message: string } };

export type ConfirmLoginPinMutationVariables = Exact<{
  token: Scalars['String'];
  pin: Scalars['String'];
}>;


export type ConfirmLoginPinMutation = { __typename?: 'Mutation', confirmLoginPin: { __typename: 'ExpiredToken', message: string } | { __typename: 'InvalidPin', message: string } | { __typename: 'InvalidToken', message: string } | { __typename: 'LoginConfirmed', token: string } | { __typename: 'ServerError', message: string } };

export type LoginByEmailMutationVariables = Exact<{
  email: Scalars['String'];
}>;


export type LoginByEmailMutation = { __typename?: 'Mutation', loginByEmail: { __typename: 'InvalidEmail', message: string } | { __typename: 'LinkSentByEmail', email: string } | { __typename: 'PinSentByEmail', email: string, token: string } | { __typename: 'ServerError', message: string } };

export type DeleteAlbumMutationVariables = Exact<{
  id: Scalars['String'];
}>;


export type DeleteAlbumMutation = { __typename?: 'Mutation', deleteAlbum: { __typename: 'AlbumDeleted', album: { __typename?: 'Album', id: string } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type ChangeProjectStatusMutationVariables = Exact<{
  projectId: Scalars['String'];
  status: ProjectStatus;
}>;


export type ChangeProjectStatusMutation = { __typename?: 'Mutation', changeProjectStatus: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ProjectStatusChanged', project: { __typename?: 'Project', status: ProjectStatus } } | { __typename: 'ServerError', message: string } };

export type AddContactMutationVariables = Exact<{
  projectId: Scalars['String'];
  contact: AddContactInput;
}>;


export type AddContactMutation = { __typename?: 'Mutation', addContact: { __typename: 'ContactAdded', contact: { __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> } } | { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } };

export type DeleteContactMutationVariables = Exact<{
  id: Scalars['String'];
}>;


export type DeleteContactMutation = { __typename?: 'Mutation', deleteContact: { __typename: 'ContactDeleted', contact: { __typename?: 'Contact', id: string } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type UpdateContactMutationVariables = Exact<{
  contactId: Scalars['String'];
  data: UpdateContactInput;
}>;


export type UpdateContactMutation = { __typename?: 'Mutation', updateContact: { __typename: 'ContactUpdated', contact: { __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type CreateAlbumMutationVariables = Exact<{
  projectId: Scalars['String'];
  name: Scalars['String'];
}>;


export type CreateAlbumMutation = { __typename?: 'Mutation', createAlbum: { __typename: 'AlbumCreated', album: { __typename?: 'Album', id: string, name: string } } | { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } };

export type ChangeProjectDatesMutationVariables = Exact<{
  projectId: Scalars['String'];
  dates: ChangeProjectDatesInput;
}>;


export type ChangeProjectDatesMutation = { __typename?: 'Mutation', changeProjectDates: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ProjectDatesChanged', project: { __typename?: 'Project', startAt?: any | null, endAt?: any | null } } | { __typename: 'ServerError', message: string } };

export type AddHouseFragment = { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any };

export type AddHouseMutationVariables = Exact<{
  projectId: Scalars['String'];
  house: AddHouseInput;
}>;


export type AddHouseMutation = { __typename?: 'Mutation', addHouse: { __typename: 'Forbidden', message: string } | { __typename: 'HouseAdded', house: { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any } } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type UpdateHouseFragment = { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any };

export type UpdateHouseMutationVariables = Exact<{
  houseId: Scalars['String'];
  house: UpdateHouseInput;
}>;


export type UpdateHouseMutation = { __typename?: 'Mutation', updateHouse: { __typename: 'Forbidden', message: string } | { __typename: 'HouseUpdated', house: { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any } } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type MakeProjectPublicMutationVariables = Exact<{
  projectId: Scalars['String'];
}>;


export type MakeProjectPublicMutation = { __typename?: 'Mutation', makeProjectPublic: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ProjectIsAlreadyPublic' } | { __typename: 'ProjectMadePublic', publicSite: { __typename?: 'PublicSite', status: PublicSiteStatus } } | { __typename: 'ServerError', message: string } };

export type MakeProjectNotPublicMutationVariables = Exact<{
  projectId: Scalars['String'];
}>;


export type MakeProjectNotPublicMutation = { __typename?: 'Mutation', makeProjectNotPublic: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ProjectIsAlreadyNotPublic' } | { __typename: 'ProjectMadeNotPublic', publicSite: { __typename?: 'PublicSite', status: PublicSiteStatus } } | { __typename: 'ServerError', message: string } };

export type AddRoomMutationVariables = Exact<{
  houseId: Scalars['String'];
  room: AddRoomInput;
}>;


export type AddRoomMutation = { __typename?: 'Mutation', addRoom: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'RoomAdded', room: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } } };

export type DeleteRoomMutationVariables = Exact<{
  id: Scalars['String'];
}>;


export type DeleteRoomMutation = { __typename?: 'Mutation', deleteRoom: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'RoomDeleted', room: { __typename?: 'Room', id: string } } };

export type UpdateRoomMutationVariables = Exact<{
  roomId: Scalars['String'];
  data: UpdateRoomInput;
}>;


export type UpdateRoomMutation = { __typename?: 'Mutation', updateRoom: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'RoomUpdated', room: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } } };

export type UploadVisualizationsMutationVariables = Exact<{
  projectId: Scalars['String'];
  files: Array<Scalars['Upload']> | Scalars['Upload'];
  roomId?: InputMaybe<Scalars['String']>;
}>;


export type UploadVisualizationsMutation = { __typename?: 'Mutation', uploadVisualizations: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'SomeVisualizationsUploaded', visualizations: Array<{ __typename?: 'Visualization', id: string, file: { __typename?: 'File', id: string, url: any } }> } | { __typename: 'VisualizationsUploaded', visualizations: Array<{ __typename?: 'Visualization', id: string, file: { __typename?: 'File', id: string, url: any } }> } };

export type ProjectScreenQueryVariables = Exact<{
  id: Scalars['String'];
}>;


export type ProjectScreenQuery = { __typename?: 'Query', project: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'Project', id: string, name: string, startAt?: any | null, endAt?: any | null, status: ProjectStatus, statuses: { __typename?: 'ProjectStatusDictionary', items: Array<{ __typename?: 'ProjectStatusDictionaryItem', key: ProjectStatus, value: string }> }, contacts: { __typename?: 'ProjectContacts', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectContactsList', items: Array<{ __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectContactsTotal', total: number } | { __typename: 'ServerError' } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any, rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } }, visualizations: { __typename?: 'ProjectVisualizations', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectVisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectVisualizationsTotal', total: number } | { __typename: 'ServerError' } }, albums: { __typename?: 'ProjectAlbums', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectAlbumsList', items: Array<{ __typename: 'Album', id: string, name: string }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectAlbumsTotal', total: number } | { __typename: 'ServerError' } }, publicSite: { __typename: 'NotFound' } | { __typename: 'PublicSite', id: string, status: PublicSiteStatus, url: string, settings: { __typename?: 'PublicSiteSettings', visualizations: boolean, albums: boolean } } | { __typename: 'ServerError' } } | { __typename: 'ServerError', message: string } };

export type ProjectScreenProjectFragment = { __typename?: 'Project', id: string, name: string, startAt?: any | null, endAt?: any | null, status: ProjectStatus, statuses: { __typename?: 'ProjectStatusDictionary', items: Array<{ __typename?: 'ProjectStatusDictionaryItem', key: ProjectStatus, value: string }> }, contacts: { __typename?: 'ProjectContacts', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectContactsList', items: Array<{ __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectContactsTotal', total: number } | { __typename: 'ServerError' } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any, rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } }, visualizations: { __typename?: 'ProjectVisualizations', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectVisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectVisualizationsTotal', total: number } | { __typename: 'ServerError' } }, albums: { __typename?: 'ProjectAlbums', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectAlbumsList', items: Array<{ __typename: 'Album', id: string, name: string }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectAlbumsTotal', total: number } | { __typename: 'ServerError' } }, publicSite: { __typename: 'NotFound' } | { __typename: 'PublicSite', id: string, status: PublicSiteStatus, url: string, settings: { __typename?: 'PublicSiteSettings', visualizations: boolean, albums: boolean } } | { __typename: 'ServerError' } };

export type ProjectScreenVisualizationsFragment = { __typename?: 'ProjectVisualizations', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectVisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectVisualizationsTotal', total: number } | { __typename: 'ServerError' } };

export type ProjectScreenVisualizationFragment = { __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null };

export type ProjectScreenHousesFragment = { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any, rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } };

export type ProjectScreenHouseFragment = { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any };

export type ProjectScreenHouseRoomsFragment = { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } };

export type ProjectScreenHouseRoomFragment = { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null };

export type ProjectScreenAlbumsFragment = { __typename?: 'ProjectAlbums', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectAlbumsList', items: Array<{ __typename: 'Album', id: string, name: string }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectAlbumsTotal', total: number } | { __typename: 'ServerError' } };

export type ProjectScreenAlbumFragment = { __typename?: 'Album', id: string, name: string };

type ProjectScreenPublicSite_NotFound_Fragment = { __typename: 'NotFound' };

type ProjectScreenPublicSite_PublicSite_Fragment = { __typename: 'PublicSite', id: string, status: PublicSiteStatus, url: string, settings: { __typename?: 'PublicSiteSettings', visualizations: boolean, albums: boolean } };

type ProjectScreenPublicSite_ServerError_Fragment = { __typename: 'ServerError' };

export type ProjectScreenPublicSiteFragment = ProjectScreenPublicSite_NotFound_Fragment | ProjectScreenPublicSite_PublicSite_Fragment | ProjectScreenPublicSite_ServerError_Fragment;

export type DeleteVisualizationsMutationVariables = Exact<{
  id: Array<Scalars['String']> | Scalars['String'];
}>;


export type DeleteVisualizationsMutation = { __typename?: 'Mutation', deleteVisualizations: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'SomeVisualizationsDeleted', visualizations: Array<{ __typename?: 'Visualization', id: string }> } | { __typename: 'VisualizationsDeleted', visualizations: Array<{ __typename?: 'Visualization', id: string }> } };

export type VisualizationsScreenQueryVariables = Exact<{
  id: Scalars['String'];
  filter: ProjectVisualizationsListFilter;
}>;


export type VisualizationsScreenQuery = { __typename?: 'Query', project: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'Project', id: string, name: string, visualizations: { __typename?: 'ProjectVisualizations', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'ProjectVisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null }> } | { __typename?: 'ServerError', message: string } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } } } | { __typename: 'ServerError', message: string } };

export type VisualizationsScreenVisualizationFragment = { __typename?: 'Visualization', id: string, name: string, description: string, version: number, file: { __typename?: 'File', id: string, name: string, url: any, type: FileType, mimeType: string }, room?: { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null } | null };

export type VisualizationsScreenHouseRoomFragment = { __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null };

export type VisualizationsScreenHousesFragment = { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } };

export type VisualizationsScreenHouseRoomsFragment = { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null, level?: number | null }> } | { __typename?: 'ServerError', message: string } };

export type CreateProjectMutationVariables = Exact<{
  input: CreateProjectInput;
}>;


export type CreateProjectMutation = { __typename?: 'Mutation', createProject: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectCreated', project: { __typename?: 'Project', id: string, name: string, startAt?: any | null, endAt?: any | null } } | { __typename: 'ServerError', message: string } };

export type InviteUserMutationVariables = Exact<{
  workspaceId: Scalars['String'];
  email: Scalars['String'];
  role: WorkspaceUserRole;
}>;


export type InviteUserMutation = { __typename?: 'Mutation', inviteUser: { __typename?: 'AlreadyInWorkspace', message: string } | { __typename?: 'Forbidden', message: string } | { __typename?: 'InviteSent', to: string, tokenExpiration: number } | { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } };

export type WorkspaceScreenQueryVariables = Exact<{
  id: Scalars['String'];
  timezone?: InputMaybe<Scalars['String']>;
}>;


export type WorkspaceScreenQuery = { __typename?: 'Query', workspace: { __typename?: 'Forbidden', message: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'Workspace', id: string, name: string, users: { __typename?: 'WorkspaceUsers', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'WorkspaceUsersList', items: Array<{ __typename?: 'WorkspaceUser', id: string, role: WorkspaceUserRole, workspace: { __typename?: 'Id', id: string }, profile: { __typename?: 'UserProfile', id: string, email: string, fullName: string, abbr: string, gravatar?: { __typename?: 'Gravatar', url: string } | null } }> } }, projects: { __typename?: 'WorkspaceProjects', current: { __typename?: 'Forbidden', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'WorkspaceProjectsList', items: Array<{ __typename?: 'Project', id: string, name: string, status: ProjectStatus, startAt?: any | null, endAt?: any | null, period?: string | null }> }, done: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceProjectsList', items: Array<{ __typename?: 'Project', id: string, name: string, status: ProjectStatus, startAt?: any | null, endAt?: any | null, period?: string | null }> } }, roles: { __typename?: 'WorkspaceUserRoleDictionary', items: Array<{ __typename?: 'WorkspaceUserRoleDictionaryItem', key: WorkspaceUserRole, value: string }> } } };

export type WorkspaceScreenFragment = { __typename?: 'Workspace', id: string, name: string, users: { __typename?: 'WorkspaceUsers', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'WorkspaceUsersList', items: Array<{ __typename?: 'WorkspaceUser', id: string, role: WorkspaceUserRole, workspace: { __typename?: 'Id', id: string }, profile: { __typename?: 'UserProfile', id: string, email: string, fullName: string, abbr: string, gravatar?: { __typename?: 'Gravatar', url: string } | null } }> } }, projects: { __typename?: 'WorkspaceProjects', current: { __typename?: 'Forbidden', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'WorkspaceProjectsList', items: Array<{ __typename?: 'Project', id: string, name: string, status: ProjectStatus, startAt?: any | null, endAt?: any | null, period?: string | null }> }, done: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceProjectsList', items: Array<{ __typename?: 'Project', id: string, name: string, status: ProjectStatus, startAt?: any | null, endAt?: any | null, period?: string | null }> } }, roles: { __typename?: 'WorkspaceUserRoleDictionary', items: Array<{ __typename?: 'WorkspaceUserRoleDictionaryItem', key: WorkspaceUserRole, value: string }> } };

export type WorkspaceScreenUsersFragment = { __typename?: 'WorkspaceUsers', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'WorkspaceUsersList', items: Array<{ __typename?: 'WorkspaceUser', id: string, role: WorkspaceUserRole, workspace: { __typename?: 'Id', id: string }, profile: { __typename?: 'UserProfile', id: string, email: string, fullName: string, abbr: string, gravatar?: { __typename?: 'Gravatar', url: string } | null } }> } };

export type WorkspaceScreenUserFragment = { __typename?: 'WorkspaceUser', id: string, role: WorkspaceUserRole, workspace: { __typename?: 'Id', id: string }, profile: { __typename?: 'UserProfile', id: string, email: string, fullName: string, abbr: string, gravatar?: { __typename?: 'Gravatar', url: string } | null } };

export type WorkspaceScreenCurrentProjectsFragment = { __typename?: 'WorkspaceProjects', current: { __typename?: 'Forbidden', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'WorkspaceProjectsList', items: Array<{ __typename?: 'Project', id: string, name: string, status: ProjectStatus, startAt?: any | null, endAt?: any | null, period?: string | null }> } };

export type WorkspaceScreenArchiveProjectsFragment = { __typename?: 'WorkspaceProjects', done: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceProjectsList', items: Array<{ __typename?: 'Project', id: string, name: string, status: ProjectStatus, startAt?: any | null, endAt?: any | null, period?: string | null }> } };

export type WorkspaceScreenProjectFragment = { __typename?: 'Project', id: string, name: string, status: ProjectStatus, startAt?: any | null, endAt?: any | null, period?: string | null };

export const AlbumScreenHouseRoomFragmentDoc = gql`
    fragment AlbumScreenHouseRoom on Room {
  id
  name
  square
  level
}
    `;
export const AlbumScreenVisualizationFragmentDoc = gql`
    fragment AlbumScreenVisualization on Visualization {
  id
  name
  description
  version
  file {
    id
    name
    url
    type
    mimeType
  }
  room {
    ...AlbumScreenHouseRoom
  }
}
    ${AlbumScreenHouseRoomFragmentDoc}`;
export const AlbumScreenHouseRoomsFragmentDoc = gql`
    fragment AlbumScreenHouseRooms on HouseRooms {
  list {
    ... on HouseRoomsList {
      items {
        ...AlbumScreenHouseRoom
      }
    }
    ... on Error {
      message
    }
  }
}
    ${AlbumScreenHouseRoomFragmentDoc}`;
export const AlbumScreenHousesFragmentDoc = gql`
    fragment AlbumScreenHouses on ProjectHouses {
  list(filter: {}, limit: 1, offset: 0) {
    __typename
    ... on ProjectHousesList {
      items {
        rooms {
          ...AlbumScreenHouseRooms
        }
      }
    }
    ... on Error {
      message
    }
  }
}
    ${AlbumScreenHouseRoomsFragmentDoc}`;
export const AlbumScreenProjectFragmentDoc = gql`
    fragment AlbumScreenProject on Project {
  id
  name
  visualizations {
    list(filter: $filter, limit: 100, offset: 0) {
      ... on ProjectVisualizationsList {
        items {
          ...AlbumScreenVisualization
        }
      }
      ... on Error {
        message
      }
    }
  }
  houses {
    ...AlbumScreenHouses
  }
}
    ${AlbumScreenVisualizationFragmentDoc}
${AlbumScreenHousesFragmentDoc}`;
export const AlbumScreenAlbumPageCoverFragmentDoc = gql`
    fragment AlbumScreenAlbumPageCover on AlbumPageCover {
  position
}
    `;
export const AlbumScreenAlbumPageVisualizationFragmentDoc = gql`
    fragment AlbumScreenAlbumPageVisualization on AlbumPageVisualization {
  position
  visualization {
    ... on Visualization {
      id
      file {
        url
      }
    }
    ... on Error {
      message
    }
  }
}
    `;
export const AlbumScreenSettingsFragmentDoc = gql`
    fragment AlbumScreenSettings on AlbumSettings {
  pageSize
  pageOrientation
}
    `;
export const AlbumScreenAlbumFragmentDoc = gql`
    fragment AlbumScreenAlbum on Album {
  id
  name
  project {
    ... on Project {
      ...AlbumScreenProject
    }
    ... on Error {
      message
    }
  }
  pages {
    ... on AlbumPages {
      items {
        ... on AlbumPageCover {
          ...AlbumScreenAlbumPageCover
        }
        ... on AlbumPageVisualization {
          ...AlbumScreenAlbumPageVisualization
        }
      }
    }
    ... on Error {
      message
    }
  }
  settings {
    ...AlbumScreenSettings
  }
}
    ${AlbumScreenProjectFragmentDoc}
${AlbumScreenAlbumPageCoverFragmentDoc}
${AlbumScreenAlbumPageVisualizationFragmentDoc}
${AlbumScreenSettingsFragmentDoc}`;
export const AddHouseFragmentDoc = gql`
    fragment AddHouse on House {
  id
  city
  address
  housingComplex
  createdAt
  modifiedAt
}
    `;
export const UpdateHouseFragmentDoc = gql`
    fragment UpdateHouse on House {
  id
  city
  address
  housingComplex
  createdAt
  modifiedAt
}
    `;
export const ProjectScreenHouseFragmentDoc = gql`
    fragment ProjectScreenHouse on House {
  id
  city
  address
  housingComplex
  createdAt
  modifiedAt
}
    `;
export const ProjectScreenHouseRoomFragmentDoc = gql`
    fragment ProjectScreenHouseRoom on Room {
  id
  name
  square
  level
}
    `;
export const ProjectScreenHouseRoomsFragmentDoc = gql`
    fragment ProjectScreenHouseRooms on HouseRooms {
  list {
    ... on HouseRoomsList {
      items {
        ...ProjectScreenHouseRoom
      }
    }
    ... on Error {
      message
    }
  }
}
    ${ProjectScreenHouseRoomFragmentDoc}`;
export const ProjectScreenHousesFragmentDoc = gql`
    fragment ProjectScreenHouses on ProjectHouses {
  list(filter: {}, limit: 1, offset: 0) {
    __typename
    ... on ProjectHousesList {
      items {
        ...ProjectScreenHouse
        rooms {
          ...ProjectScreenHouseRooms
        }
      }
    }
    ... on Error {
      message
    }
  }
}
    ${ProjectScreenHouseFragmentDoc}
${ProjectScreenHouseRoomsFragmentDoc}`;
export const ProjectScreenVisualizationFragmentDoc = gql`
    fragment ProjectScreenVisualization on Visualization {
  id
  name
  description
  version
  file {
    id
    name
    url
    type
    mimeType
  }
  room {
    ...ProjectScreenHouseRoom
  }
}
    ${ProjectScreenHouseRoomFragmentDoc}`;
export const ProjectScreenVisualizationsFragmentDoc = gql`
    fragment ProjectScreenVisualizations on ProjectVisualizations {
  list(filter: {status: {eq: [UNKNOWN, APPROVED]}}, limit: 20, offset: 0) {
    __typename
    ... on ProjectVisualizationsList {
      items {
        ...ProjectScreenVisualization
      }
    }
    ... on Error {
      message
    }
  }
  total(filter: {}) {
    __typename
    ... on ProjectVisualizationsTotal {
      total
    }
  }
}
    ${ProjectScreenVisualizationFragmentDoc}`;
export const ProjectScreenAlbumFragmentDoc = gql`
    fragment ProjectScreenAlbum on Album {
  id
  name
}
    `;
export const ProjectScreenAlbumsFragmentDoc = gql`
    fragment ProjectScreenAlbums on ProjectAlbums {
  list(limit: 20, offset: 0) {
    __typename
    ... on ProjectAlbumsList {
      items {
        __typename
        ... on Album {
          ...ProjectScreenAlbum
        }
      }
    }
    ... on Error {
      message
    }
  }
  total {
    __typename
    ... on ProjectAlbumsTotal {
      total
    }
  }
}
    ${ProjectScreenAlbumFragmentDoc}`;
export const ProjectScreenPublicSiteFragmentDoc = gql`
    fragment ProjectScreenPublicSite on ProjectPublicSite {
  __typename
  ... on PublicSite {
    id
    status
    url
    settings {
      visualizations
      albums
    }
  }
}
    `;
export const ProjectScreenProjectFragmentDoc = gql`
    fragment ProjectScreenProject on Project {
  id
  name
  startAt
  endAt
  status
  statuses {
    items {
      key
      value
    }
  }
  contacts {
    list(filter: {}, limit: 10, offset: 0) {
      __typename
      ... on ProjectContactsList {
        items {
          id
          fullName
          photo
          details {
            type
            value
          }
        }
      }
      ... on Error {
        message
      }
    }
    total {
      __typename
      ... on ProjectContactsTotal {
        total
      }
    }
  }
  houses {
    ...ProjectScreenHouses
  }
  visualizations {
    ...ProjectScreenVisualizations
  }
  albums {
    ...ProjectScreenAlbums
  }
  publicSite {
    ...ProjectScreenPublicSite
  }
}
    ${ProjectScreenHousesFragmentDoc}
${ProjectScreenVisualizationsFragmentDoc}
${ProjectScreenAlbumsFragmentDoc}
${ProjectScreenPublicSiteFragmentDoc}`;
export const VisualizationsScreenHouseRoomFragmentDoc = gql`
    fragment VisualizationsScreenHouseRoom on Room {
  id
  name
  square
  level
}
    `;
export const VisualizationsScreenVisualizationFragmentDoc = gql`
    fragment VisualizationsScreenVisualization on Visualization {
  id
  name
  description
  version
  file {
    id
    name
    url
    type
    mimeType
  }
  room {
    ...VisualizationsScreenHouseRoom
  }
}
    ${VisualizationsScreenHouseRoomFragmentDoc}`;
export const VisualizationsScreenHouseRoomsFragmentDoc = gql`
    fragment VisualizationsScreenHouseRooms on HouseRooms {
  list {
    ... on HouseRoomsList {
      items {
        ...VisualizationsScreenHouseRoom
      }
    }
    ... on Error {
      message
    }
  }
}
    ${VisualizationsScreenHouseRoomFragmentDoc}`;
export const VisualizationsScreenHousesFragmentDoc = gql`
    fragment VisualizationsScreenHouses on ProjectHouses {
  list(filter: {}, limit: 1, offset: 0) {
    __typename
    ... on ProjectHousesList {
      items {
        rooms {
          ...VisualizationsScreenHouseRooms
        }
      }
    }
    ... on Error {
      message
    }
  }
}
    ${VisualizationsScreenHouseRoomsFragmentDoc}`;
export const WorkspaceScreenUserFragmentDoc = gql`
    fragment WorkspaceScreenUser on WorkspaceUser {
  id
  workspace {
    id
  }
  role
  profile {
    id
    email
    fullName
    abbr
    gravatar {
      url
    }
  }
}
    `;
export const WorkspaceScreenUsersFragmentDoc = gql`
    fragment WorkspaceScreenUsers on WorkspaceUsers {
  list {
    ... on WorkspaceUsersList {
      items {
        ...WorkspaceScreenUser
      }
    }
    ... on Error {
      message
    }
  }
}
    ${WorkspaceScreenUserFragmentDoc}`;
export const WorkspaceScreenProjectFragmentDoc = gql`
    fragment WorkspaceScreenProject on Project {
  id
  name
  status
  startAt
  endAt
  period(timezone: $timezone)
}
    `;
export const WorkspaceScreenCurrentProjectsFragmentDoc = gql`
    fragment WorkspaceScreenCurrentProjects on WorkspaceProjects {
  current: list(filter: {status: [NEW, IN_PROGRESS]}, limit: 10) {
    ... on WorkspaceProjectsList {
      items {
        ...WorkspaceScreenProject
      }
    }
    ... on Error {
      message
    }
  }
}
    ${WorkspaceScreenProjectFragmentDoc}`;
export const WorkspaceScreenArchiveProjectsFragmentDoc = gql`
    fragment WorkspaceScreenArchiveProjects on WorkspaceProjects {
  done: list(filter: {status: [DONE, CANCELED]}, limit: 10) {
    __typename
    ... on WorkspaceProjectsList {
      items {
        ...WorkspaceScreenProject
      }
    }
    ... on Error {
      message
    }
  }
}
    ${WorkspaceScreenProjectFragmentDoc}`;
export const WorkspaceScreenFragmentDoc = gql`
    fragment WorkspaceScreen on Workspace {
  id
  name
  users {
    ...WorkspaceScreenUsers
  }
  projects {
    ...WorkspaceScreenCurrentProjects
    ...WorkspaceScreenArchiveProjects
  }
  roles {
    items {
      key
      value
    }
  }
}
    ${WorkspaceScreenUsersFragmentDoc}
${WorkspaceScreenCurrentProjectsFragmentDoc}
${WorkspaceScreenArchiveProjectsFragmentDoc}`;
export const ProfileDocument = gql`
    query profile {
  profile {
    __typename
    ... on UserProfile {
      id
      email
      gravatar {
        url
      }
      defaultWorkspace {
        id
        name
      }
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;

/**
 * __useProfileQuery__
 *
 * To run a query within a React component, call `useProfileQuery` and pass it any options that fit your needs.
 * When your component renders, `useProfileQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useProfileQuery({
 *   variables: {
 *   },
 * });
 */
export function useProfileQuery(baseOptions?: Apollo.QueryHookOptions<ProfileQuery, ProfileQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ProfileQuery, ProfileQueryVariables>(ProfileDocument, options);
      }
export function useProfileLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ProfileQuery, ProfileQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ProfileQuery, ProfileQueryVariables>(ProfileDocument, options);
        }
export type ProfileQueryHookResult = ReturnType<typeof useProfileQuery>;
export type ProfileLazyQueryHookResult = ReturnType<typeof useProfileLazyQuery>;
export type ProfileQueryResult = Apollo.QueryResult<ProfileQuery, ProfileQueryVariables>;
export const AcceptInviteDocument = gql`
    mutation acceptInvite($token: String!) {
  acceptInvite(token: $token) {
    __typename
    ... on InviteAccepted {
      token
    }
    ... on Error {
      message
    }
  }
}
    `;
export type AcceptInviteMutationFn = Apollo.MutationFunction<AcceptInviteMutation, AcceptInviteMutationVariables>;

/**
 * __useAcceptInviteMutation__
 *
 * To run a mutation, you first call `useAcceptInviteMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAcceptInviteMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [acceptInviteMutation, { data, loading, error }] = useAcceptInviteMutation({
 *   variables: {
 *      token: // value for 'token'
 *   },
 * });
 */
export function useAcceptInviteMutation(baseOptions?: Apollo.MutationHookOptions<AcceptInviteMutation, AcceptInviteMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AcceptInviteMutation, AcceptInviteMutationVariables>(AcceptInviteDocument, options);
      }
export type AcceptInviteMutationHookResult = ReturnType<typeof useAcceptInviteMutation>;
export type AcceptInviteMutationResult = Apollo.MutationResult<AcceptInviteMutation>;
export type AcceptInviteMutationOptions = Apollo.BaseMutationOptions<AcceptInviteMutation, AcceptInviteMutationVariables>;
export const ChangeAlbumPageOrientationDocument = gql`
    mutation changeAlbumPageOrientation($albumId: String!, $orientation: PageOrientation!) {
  changeAlbumPageOrientation(albumId: $albumId, orientation: $orientation) {
    __typename
    ... on AlbumPageOrientationChanged {
      album {
        settings {
          pageOrientation
        }
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type ChangeAlbumPageOrientationMutationFn = Apollo.MutationFunction<ChangeAlbumPageOrientationMutation, ChangeAlbumPageOrientationMutationVariables>;

/**
 * __useChangeAlbumPageOrientationMutation__
 *
 * To run a mutation, you first call `useChangeAlbumPageOrientationMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useChangeAlbumPageOrientationMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [changeAlbumPageOrientationMutation, { data, loading, error }] = useChangeAlbumPageOrientationMutation({
 *   variables: {
 *      albumId: // value for 'albumId'
 *      orientation: // value for 'orientation'
 *   },
 * });
 */
export function useChangeAlbumPageOrientationMutation(baseOptions?: Apollo.MutationHookOptions<ChangeAlbumPageOrientationMutation, ChangeAlbumPageOrientationMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ChangeAlbumPageOrientationMutation, ChangeAlbumPageOrientationMutationVariables>(ChangeAlbumPageOrientationDocument, options);
      }
export type ChangeAlbumPageOrientationMutationHookResult = ReturnType<typeof useChangeAlbumPageOrientationMutation>;
export type ChangeAlbumPageOrientationMutationResult = Apollo.MutationResult<ChangeAlbumPageOrientationMutation>;
export type ChangeAlbumPageOrientationMutationOptions = Apollo.BaseMutationOptions<ChangeAlbumPageOrientationMutation, ChangeAlbumPageOrientationMutationVariables>;
export const ChangeAlbumPageSizeDocument = gql`
    mutation changeAlbumPageSize($albumId: String!, $size: PageSize!) {
  changeAlbumPageSize(albumId: $albumId, size: $size) {
    __typename
    ... on AlbumPageSizeChanged {
      album {
        settings {
          pageSize
        }
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type ChangeAlbumPageSizeMutationFn = Apollo.MutationFunction<ChangeAlbumPageSizeMutation, ChangeAlbumPageSizeMutationVariables>;

/**
 * __useChangeAlbumPageSizeMutation__
 *
 * To run a mutation, you first call `useChangeAlbumPageSizeMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useChangeAlbumPageSizeMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [changeAlbumPageSizeMutation, { data, loading, error }] = useChangeAlbumPageSizeMutation({
 *   variables: {
 *      albumId: // value for 'albumId'
 *      size: // value for 'size'
 *   },
 * });
 */
export function useChangeAlbumPageSizeMutation(baseOptions?: Apollo.MutationHookOptions<ChangeAlbumPageSizeMutation, ChangeAlbumPageSizeMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ChangeAlbumPageSizeMutation, ChangeAlbumPageSizeMutationVariables>(ChangeAlbumPageSizeDocument, options);
      }
export type ChangeAlbumPageSizeMutationHookResult = ReturnType<typeof useChangeAlbumPageSizeMutation>;
export type ChangeAlbumPageSizeMutationResult = Apollo.MutationResult<ChangeAlbumPageSizeMutation>;
export type ChangeAlbumPageSizeMutationOptions = Apollo.BaseMutationOptions<ChangeAlbumPageSizeMutation, ChangeAlbumPageSizeMutationVariables>;
export const AddVisualizationsToAlbumDocument = gql`
    mutation addVisualizationsToAlbum($albumId: String!, $visualizations: [String!]!) {
  addVisualizationsToAlbum(albumId: $albumId, visualizations: $visualizations) {
    __typename
    ... on VisualizationsAddedToAlbum {
      pages {
        position
        visualization {
          ... on Visualization {
            id
            file {
              url
            }
          }
        }
      }
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type AddVisualizationsToAlbumMutationFn = Apollo.MutationFunction<AddVisualizationsToAlbumMutation, AddVisualizationsToAlbumMutationVariables>;

/**
 * __useAddVisualizationsToAlbumMutation__
 *
 * To run a mutation, you first call `useAddVisualizationsToAlbumMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddVisualizationsToAlbumMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addVisualizationsToAlbumMutation, { data, loading, error }] = useAddVisualizationsToAlbumMutation({
 *   variables: {
 *      albumId: // value for 'albumId'
 *      visualizations: // value for 'visualizations'
 *   },
 * });
 */
export function useAddVisualizationsToAlbumMutation(baseOptions?: Apollo.MutationHookOptions<AddVisualizationsToAlbumMutation, AddVisualizationsToAlbumMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddVisualizationsToAlbumMutation, AddVisualizationsToAlbumMutationVariables>(AddVisualizationsToAlbumDocument, options);
      }
export type AddVisualizationsToAlbumMutationHookResult = ReturnType<typeof useAddVisualizationsToAlbumMutation>;
export type AddVisualizationsToAlbumMutationResult = Apollo.MutationResult<AddVisualizationsToAlbumMutation>;
export type AddVisualizationsToAlbumMutationOptions = Apollo.BaseMutationOptions<AddVisualizationsToAlbumMutation, AddVisualizationsToAlbumMutationVariables>;
export const AlbumScreenDocument = gql`
    query albumScreen($id: String!, $filter: ProjectVisualizationsListFilter!) {
  album: album(id: $id) {
    __typename
    ... on Album {
      ...AlbumScreenAlbum
    }
    ... on Error {
      message
    }
  }
}
    ${AlbumScreenAlbumFragmentDoc}`;

/**
 * __useAlbumScreenQuery__
 *
 * To run a query within a React component, call `useAlbumScreenQuery` and pass it any options that fit your needs.
 * When your component renders, `useAlbumScreenQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAlbumScreenQuery({
 *   variables: {
 *      id: // value for 'id'
 *      filter: // value for 'filter'
 *   },
 * });
 */
export function useAlbumScreenQuery(baseOptions: Apollo.QueryHookOptions<AlbumScreenQuery, AlbumScreenQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<AlbumScreenQuery, AlbumScreenQueryVariables>(AlbumScreenDocument, options);
      }
export function useAlbumScreenLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<AlbumScreenQuery, AlbumScreenQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<AlbumScreenQuery, AlbumScreenQueryVariables>(AlbumScreenDocument, options);
        }
export type AlbumScreenQueryHookResult = ReturnType<typeof useAlbumScreenQuery>;
export type AlbumScreenLazyQueryHookResult = ReturnType<typeof useAlbumScreenLazyQuery>;
export type AlbumScreenQueryResult = Apollo.QueryResult<AlbumScreenQuery, AlbumScreenQueryVariables>;
export const ConfirmLoginLinkDocument = gql`
    mutation confirmLoginLink($token: String!) {
  confirmLoginLink(token: $token) {
    __typename
    ... on LoginConfirmed {
      token
    }
    ... on InvalidToken {
      message
    }
    ... on ExpiredToken {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type ConfirmLoginLinkMutationFn = Apollo.MutationFunction<ConfirmLoginLinkMutation, ConfirmLoginLinkMutationVariables>;

/**
 * __useConfirmLoginLinkMutation__
 *
 * To run a mutation, you first call `useConfirmLoginLinkMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useConfirmLoginLinkMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [confirmLoginLinkMutation, { data, loading, error }] = useConfirmLoginLinkMutation({
 *   variables: {
 *      token: // value for 'token'
 *   },
 * });
 */
export function useConfirmLoginLinkMutation(baseOptions?: Apollo.MutationHookOptions<ConfirmLoginLinkMutation, ConfirmLoginLinkMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ConfirmLoginLinkMutation, ConfirmLoginLinkMutationVariables>(ConfirmLoginLinkDocument, options);
      }
export type ConfirmLoginLinkMutationHookResult = ReturnType<typeof useConfirmLoginLinkMutation>;
export type ConfirmLoginLinkMutationResult = Apollo.MutationResult<ConfirmLoginLinkMutation>;
export type ConfirmLoginLinkMutationOptions = Apollo.BaseMutationOptions<ConfirmLoginLinkMutation, ConfirmLoginLinkMutationVariables>;
export const ConfirmLoginPinDocument = gql`
    mutation confirmLoginPin($token: String!, $pin: String!) {
  confirmLoginPin(token: $token, pin: $pin) {
    __typename
    ... on LoginConfirmed {
      token
    }
    ... on InvalidPin {
      message
    }
    ... on InvalidToken {
      message
    }
    ... on ExpiredToken {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type ConfirmLoginPinMutationFn = Apollo.MutationFunction<ConfirmLoginPinMutation, ConfirmLoginPinMutationVariables>;

/**
 * __useConfirmLoginPinMutation__
 *
 * To run a mutation, you first call `useConfirmLoginPinMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useConfirmLoginPinMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [confirmLoginPinMutation, { data, loading, error }] = useConfirmLoginPinMutation({
 *   variables: {
 *      token: // value for 'token'
 *      pin: // value for 'pin'
 *   },
 * });
 */
export function useConfirmLoginPinMutation(baseOptions?: Apollo.MutationHookOptions<ConfirmLoginPinMutation, ConfirmLoginPinMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ConfirmLoginPinMutation, ConfirmLoginPinMutationVariables>(ConfirmLoginPinDocument, options);
      }
export type ConfirmLoginPinMutationHookResult = ReturnType<typeof useConfirmLoginPinMutation>;
export type ConfirmLoginPinMutationResult = Apollo.MutationResult<ConfirmLoginPinMutation>;
export type ConfirmLoginPinMutationOptions = Apollo.BaseMutationOptions<ConfirmLoginPinMutation, ConfirmLoginPinMutationVariables>;
export const LoginByEmailDocument = gql`
    mutation loginByEmail($email: String!) {
  loginByEmail(email: $email) {
    __typename
    ... on LinkSentByEmail {
      email
    }
    ... on PinSentByEmail {
      email
      token
    }
    ... on InvalidEmail {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type LoginByEmailMutationFn = Apollo.MutationFunction<LoginByEmailMutation, LoginByEmailMutationVariables>;

/**
 * __useLoginByEmailMutation__
 *
 * To run a mutation, you first call `useLoginByEmailMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useLoginByEmailMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [loginByEmailMutation, { data, loading, error }] = useLoginByEmailMutation({
 *   variables: {
 *      email: // value for 'email'
 *   },
 * });
 */
export function useLoginByEmailMutation(baseOptions?: Apollo.MutationHookOptions<LoginByEmailMutation, LoginByEmailMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<LoginByEmailMutation, LoginByEmailMutationVariables>(LoginByEmailDocument, options);
      }
export type LoginByEmailMutationHookResult = ReturnType<typeof useLoginByEmailMutation>;
export type LoginByEmailMutationResult = Apollo.MutationResult<LoginByEmailMutation>;
export type LoginByEmailMutationOptions = Apollo.BaseMutationOptions<LoginByEmailMutation, LoginByEmailMutationVariables>;
export const DeleteAlbumDocument = gql`
    mutation deleteAlbum($id: String!) {
  deleteAlbum(id: $id) {
    __typename
    ... on AlbumDeleted {
      album {
        id
      }
    }
    ... on Error {
      __typename
      message
    }
  }
}
    `;
export type DeleteAlbumMutationFn = Apollo.MutationFunction<DeleteAlbumMutation, DeleteAlbumMutationVariables>;

/**
 * __useDeleteAlbumMutation__
 *
 * To run a mutation, you first call `useDeleteAlbumMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteAlbumMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteAlbumMutation, { data, loading, error }] = useDeleteAlbumMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteAlbumMutation(baseOptions?: Apollo.MutationHookOptions<DeleteAlbumMutation, DeleteAlbumMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteAlbumMutation, DeleteAlbumMutationVariables>(DeleteAlbumDocument, options);
      }
export type DeleteAlbumMutationHookResult = ReturnType<typeof useDeleteAlbumMutation>;
export type DeleteAlbumMutationResult = Apollo.MutationResult<DeleteAlbumMutation>;
export type DeleteAlbumMutationOptions = Apollo.BaseMutationOptions<DeleteAlbumMutation, DeleteAlbumMutationVariables>;
export const ChangeProjectStatusDocument = gql`
    mutation changeProjectStatus($projectId: String!, $status: ProjectStatus!) {
  changeProjectStatus(projectId: $projectId, status: $status) {
    __typename
    ... on ProjectStatusChanged {
      project {
        status
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type ChangeProjectStatusMutationFn = Apollo.MutationFunction<ChangeProjectStatusMutation, ChangeProjectStatusMutationVariables>;

/**
 * __useChangeProjectStatusMutation__
 *
 * To run a mutation, you first call `useChangeProjectStatusMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useChangeProjectStatusMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [changeProjectStatusMutation, { data, loading, error }] = useChangeProjectStatusMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      status: // value for 'status'
 *   },
 * });
 */
export function useChangeProjectStatusMutation(baseOptions?: Apollo.MutationHookOptions<ChangeProjectStatusMutation, ChangeProjectStatusMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ChangeProjectStatusMutation, ChangeProjectStatusMutationVariables>(ChangeProjectStatusDocument, options);
      }
export type ChangeProjectStatusMutationHookResult = ReturnType<typeof useChangeProjectStatusMutation>;
export type ChangeProjectStatusMutationResult = Apollo.MutationResult<ChangeProjectStatusMutation>;
export type ChangeProjectStatusMutationOptions = Apollo.BaseMutationOptions<ChangeProjectStatusMutation, ChangeProjectStatusMutationVariables>;
export const AddContactDocument = gql`
    mutation addContact($projectId: String!, $contact: AddContactInput!) {
  addContact(projectId: $projectId, contact: $contact) {
    __typename
    ... on ContactAdded {
      contact {
        id
        fullName
        photo
        details {
          type
          value
        }
      }
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type AddContactMutationFn = Apollo.MutationFunction<AddContactMutation, AddContactMutationVariables>;

/**
 * __useAddContactMutation__
 *
 * To run a mutation, you first call `useAddContactMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddContactMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addContactMutation, { data, loading, error }] = useAddContactMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      contact: // value for 'contact'
 *   },
 * });
 */
export function useAddContactMutation(baseOptions?: Apollo.MutationHookOptions<AddContactMutation, AddContactMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddContactMutation, AddContactMutationVariables>(AddContactDocument, options);
      }
export type AddContactMutationHookResult = ReturnType<typeof useAddContactMutation>;
export type AddContactMutationResult = Apollo.MutationResult<AddContactMutation>;
export type AddContactMutationOptions = Apollo.BaseMutationOptions<AddContactMutation, AddContactMutationVariables>;
export const DeleteContactDocument = gql`
    mutation deleteContact($id: String!) {
  deleteContact(id: $id) {
    __typename
    ... on ContactDeleted {
      contact {
        id
      }
    }
    ... on Forbidden {
      message
    }
    ... on NotFound {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type DeleteContactMutationFn = Apollo.MutationFunction<DeleteContactMutation, DeleteContactMutationVariables>;

/**
 * __useDeleteContactMutation__
 *
 * To run a mutation, you first call `useDeleteContactMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteContactMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteContactMutation, { data, loading, error }] = useDeleteContactMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteContactMutation(baseOptions?: Apollo.MutationHookOptions<DeleteContactMutation, DeleteContactMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteContactMutation, DeleteContactMutationVariables>(DeleteContactDocument, options);
      }
export type DeleteContactMutationHookResult = ReturnType<typeof useDeleteContactMutation>;
export type DeleteContactMutationResult = Apollo.MutationResult<DeleteContactMutation>;
export type DeleteContactMutationOptions = Apollo.BaseMutationOptions<DeleteContactMutation, DeleteContactMutationVariables>;
export const UpdateContactDocument = gql`
    mutation updateContact($contactId: String!, $data: UpdateContactInput!) {
  updateContact(contactId: $contactId, data: $data) {
    __typename
    ... on ContactUpdated {
      contact {
        id
        fullName
        photo
        details {
          type
          value
        }
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type UpdateContactMutationFn = Apollo.MutationFunction<UpdateContactMutation, UpdateContactMutationVariables>;

/**
 * __useUpdateContactMutation__
 *
 * To run a mutation, you first call `useUpdateContactMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateContactMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateContactMutation, { data, loading, error }] = useUpdateContactMutation({
 *   variables: {
 *      contactId: // value for 'contactId'
 *      data: // value for 'data'
 *   },
 * });
 */
export function useUpdateContactMutation(baseOptions?: Apollo.MutationHookOptions<UpdateContactMutation, UpdateContactMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateContactMutation, UpdateContactMutationVariables>(UpdateContactDocument, options);
      }
export type UpdateContactMutationHookResult = ReturnType<typeof useUpdateContactMutation>;
export type UpdateContactMutationResult = Apollo.MutationResult<UpdateContactMutation>;
export type UpdateContactMutationOptions = Apollo.BaseMutationOptions<UpdateContactMutation, UpdateContactMutationVariables>;
export const CreateAlbumDocument = gql`
    mutation createAlbum($projectId: String!, $name: String!) {
  createAlbum(projectId: $projectId, name: $name) {
    __typename
    ... on AlbumCreated {
      album {
        ...ProjectScreenAlbum
      }
    }
    ... on Error {
      __typename
      message
    }
  }
}
    ${ProjectScreenAlbumFragmentDoc}`;
export type CreateAlbumMutationFn = Apollo.MutationFunction<CreateAlbumMutation, CreateAlbumMutationVariables>;

/**
 * __useCreateAlbumMutation__
 *
 * To run a mutation, you first call `useCreateAlbumMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateAlbumMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createAlbumMutation, { data, loading, error }] = useCreateAlbumMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      name: // value for 'name'
 *   },
 * });
 */
export function useCreateAlbumMutation(baseOptions?: Apollo.MutationHookOptions<CreateAlbumMutation, CreateAlbumMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateAlbumMutation, CreateAlbumMutationVariables>(CreateAlbumDocument, options);
      }
export type CreateAlbumMutationHookResult = ReturnType<typeof useCreateAlbumMutation>;
export type CreateAlbumMutationResult = Apollo.MutationResult<CreateAlbumMutation>;
export type CreateAlbumMutationOptions = Apollo.BaseMutationOptions<CreateAlbumMutation, CreateAlbumMutationVariables>;
export const ChangeProjectDatesDocument = gql`
    mutation changeProjectDates($projectId: String!, $dates: ChangeProjectDatesInput!) {
  changeProjectDates(projectId: $projectId, dates: $dates) {
    __typename
    ... on ProjectDatesChanged {
      project {
        startAt
        endAt
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type ChangeProjectDatesMutationFn = Apollo.MutationFunction<ChangeProjectDatesMutation, ChangeProjectDatesMutationVariables>;

/**
 * __useChangeProjectDatesMutation__
 *
 * To run a mutation, you first call `useChangeProjectDatesMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useChangeProjectDatesMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [changeProjectDatesMutation, { data, loading, error }] = useChangeProjectDatesMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      dates: // value for 'dates'
 *   },
 * });
 */
export function useChangeProjectDatesMutation(baseOptions?: Apollo.MutationHookOptions<ChangeProjectDatesMutation, ChangeProjectDatesMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ChangeProjectDatesMutation, ChangeProjectDatesMutationVariables>(ChangeProjectDatesDocument, options);
      }
export type ChangeProjectDatesMutationHookResult = ReturnType<typeof useChangeProjectDatesMutation>;
export type ChangeProjectDatesMutationResult = Apollo.MutationResult<ChangeProjectDatesMutation>;
export type ChangeProjectDatesMutationOptions = Apollo.BaseMutationOptions<ChangeProjectDatesMutation, ChangeProjectDatesMutationVariables>;
export const AddHouseDocument = gql`
    mutation addHouse($projectId: String!, $house: AddHouseInput!) {
  addHouse(projectId: $projectId, house: $house) {
    __typename
    ... on HouseAdded {
      house {
        ...AddHouse
      }
    }
    ... on Error {
      __typename
      message
    }
  }
}
    ${AddHouseFragmentDoc}`;
export type AddHouseMutationFn = Apollo.MutationFunction<AddHouseMutation, AddHouseMutationVariables>;

/**
 * __useAddHouseMutation__
 *
 * To run a mutation, you first call `useAddHouseMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddHouseMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addHouseMutation, { data, loading, error }] = useAddHouseMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      house: // value for 'house'
 *   },
 * });
 */
export function useAddHouseMutation(baseOptions?: Apollo.MutationHookOptions<AddHouseMutation, AddHouseMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddHouseMutation, AddHouseMutationVariables>(AddHouseDocument, options);
      }
export type AddHouseMutationHookResult = ReturnType<typeof useAddHouseMutation>;
export type AddHouseMutationResult = Apollo.MutationResult<AddHouseMutation>;
export type AddHouseMutationOptions = Apollo.BaseMutationOptions<AddHouseMutation, AddHouseMutationVariables>;
export const UpdateHouseDocument = gql`
    mutation updateHouse($houseId: String!, $house: UpdateHouseInput!) {
  updateHouse(houseId: $houseId, data: $house) {
    __typename
    ... on HouseUpdated {
      house {
        ...UpdateHouse
      }
    }
    ... on Error {
      __typename
      message
    }
  }
}
    ${UpdateHouseFragmentDoc}`;
export type UpdateHouseMutationFn = Apollo.MutationFunction<UpdateHouseMutation, UpdateHouseMutationVariables>;

/**
 * __useUpdateHouseMutation__
 *
 * To run a mutation, you first call `useUpdateHouseMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateHouseMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateHouseMutation, { data, loading, error }] = useUpdateHouseMutation({
 *   variables: {
 *      houseId: // value for 'houseId'
 *      house: // value for 'house'
 *   },
 * });
 */
export function useUpdateHouseMutation(baseOptions?: Apollo.MutationHookOptions<UpdateHouseMutation, UpdateHouseMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateHouseMutation, UpdateHouseMutationVariables>(UpdateHouseDocument, options);
      }
export type UpdateHouseMutationHookResult = ReturnType<typeof useUpdateHouseMutation>;
export type UpdateHouseMutationResult = Apollo.MutationResult<UpdateHouseMutation>;
export type UpdateHouseMutationOptions = Apollo.BaseMutationOptions<UpdateHouseMutation, UpdateHouseMutationVariables>;
export const MakeProjectPublicDocument = gql`
    mutation makeProjectPublic($projectId: String!) {
  makeProjectPublic(projectId: $projectId) {
    __typename
    ... on ProjectMadePublic {
      publicSite {
        status
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type MakeProjectPublicMutationFn = Apollo.MutationFunction<MakeProjectPublicMutation, MakeProjectPublicMutationVariables>;

/**
 * __useMakeProjectPublicMutation__
 *
 * To run a mutation, you first call `useMakeProjectPublicMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useMakeProjectPublicMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [makeProjectPublicMutation, { data, loading, error }] = useMakeProjectPublicMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useMakeProjectPublicMutation(baseOptions?: Apollo.MutationHookOptions<MakeProjectPublicMutation, MakeProjectPublicMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<MakeProjectPublicMutation, MakeProjectPublicMutationVariables>(MakeProjectPublicDocument, options);
      }
export type MakeProjectPublicMutationHookResult = ReturnType<typeof useMakeProjectPublicMutation>;
export type MakeProjectPublicMutationResult = Apollo.MutationResult<MakeProjectPublicMutation>;
export type MakeProjectPublicMutationOptions = Apollo.BaseMutationOptions<MakeProjectPublicMutation, MakeProjectPublicMutationVariables>;
export const MakeProjectNotPublicDocument = gql`
    mutation makeProjectNotPublic($projectId: String!) {
  makeProjectNotPublic(projectId: $projectId) {
    __typename
    ... on ProjectMadeNotPublic {
      publicSite {
        status
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type MakeProjectNotPublicMutationFn = Apollo.MutationFunction<MakeProjectNotPublicMutation, MakeProjectNotPublicMutationVariables>;

/**
 * __useMakeProjectNotPublicMutation__
 *
 * To run a mutation, you first call `useMakeProjectNotPublicMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useMakeProjectNotPublicMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [makeProjectNotPublicMutation, { data, loading, error }] = useMakeProjectNotPublicMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useMakeProjectNotPublicMutation(baseOptions?: Apollo.MutationHookOptions<MakeProjectNotPublicMutation, MakeProjectNotPublicMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<MakeProjectNotPublicMutation, MakeProjectNotPublicMutationVariables>(MakeProjectNotPublicDocument, options);
      }
export type MakeProjectNotPublicMutationHookResult = ReturnType<typeof useMakeProjectNotPublicMutation>;
export type MakeProjectNotPublicMutationResult = Apollo.MutationResult<MakeProjectNotPublicMutation>;
export type MakeProjectNotPublicMutationOptions = Apollo.BaseMutationOptions<MakeProjectNotPublicMutation, MakeProjectNotPublicMutationVariables>;
export const AddRoomDocument = gql`
    mutation addRoom($houseId: String!, $room: AddRoomInput!) {
  addRoom(houseId: $houseId, room: $room) {
    __typename
    ... on RoomAdded {
      room {
        ...ProjectScreenHouseRoom
      }
    }
    ... on Error {
      message
    }
  }
}
    ${ProjectScreenHouseRoomFragmentDoc}`;
export type AddRoomMutationFn = Apollo.MutationFunction<AddRoomMutation, AddRoomMutationVariables>;

/**
 * __useAddRoomMutation__
 *
 * To run a mutation, you first call `useAddRoomMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddRoomMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addRoomMutation, { data, loading, error }] = useAddRoomMutation({
 *   variables: {
 *      houseId: // value for 'houseId'
 *      room: // value for 'room'
 *   },
 * });
 */
export function useAddRoomMutation(baseOptions?: Apollo.MutationHookOptions<AddRoomMutation, AddRoomMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddRoomMutation, AddRoomMutationVariables>(AddRoomDocument, options);
      }
export type AddRoomMutationHookResult = ReturnType<typeof useAddRoomMutation>;
export type AddRoomMutationResult = Apollo.MutationResult<AddRoomMutation>;
export type AddRoomMutationOptions = Apollo.BaseMutationOptions<AddRoomMutation, AddRoomMutationVariables>;
export const DeleteRoomDocument = gql`
    mutation deleteRoom($id: String!) {
  deleteRoom(id: $id) {
    __typename
    ... on RoomDeleted {
      room {
        id
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
  }
}
    `;
export type DeleteRoomMutationFn = Apollo.MutationFunction<DeleteRoomMutation, DeleteRoomMutationVariables>;

/**
 * __useDeleteRoomMutation__
 *
 * To run a mutation, you first call `useDeleteRoomMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteRoomMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteRoomMutation, { data, loading, error }] = useDeleteRoomMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteRoomMutation(baseOptions?: Apollo.MutationHookOptions<DeleteRoomMutation, DeleteRoomMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteRoomMutation, DeleteRoomMutationVariables>(DeleteRoomDocument, options);
      }
export type DeleteRoomMutationHookResult = ReturnType<typeof useDeleteRoomMutation>;
export type DeleteRoomMutationResult = Apollo.MutationResult<DeleteRoomMutation>;
export type DeleteRoomMutationOptions = Apollo.BaseMutationOptions<DeleteRoomMutation, DeleteRoomMutationVariables>;
export const UpdateRoomDocument = gql`
    mutation updateRoom($roomId: String!, $data: UpdateRoomInput!) {
  updateRoom(roomId: $roomId, data: $data) {
    __typename
    ... on RoomUpdated {
      room {
        ...ProjectScreenHouseRoom
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
  }
}
    ${ProjectScreenHouseRoomFragmentDoc}`;
export type UpdateRoomMutationFn = Apollo.MutationFunction<UpdateRoomMutation, UpdateRoomMutationVariables>;

/**
 * __useUpdateRoomMutation__
 *
 * To run a mutation, you first call `useUpdateRoomMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateRoomMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateRoomMutation, { data, loading, error }] = useUpdateRoomMutation({
 *   variables: {
 *      roomId: // value for 'roomId'
 *      data: // value for 'data'
 *   },
 * });
 */
export function useUpdateRoomMutation(baseOptions?: Apollo.MutationHookOptions<UpdateRoomMutation, UpdateRoomMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateRoomMutation, UpdateRoomMutationVariables>(UpdateRoomDocument, options);
      }
export type UpdateRoomMutationHookResult = ReturnType<typeof useUpdateRoomMutation>;
export type UpdateRoomMutationResult = Apollo.MutationResult<UpdateRoomMutation>;
export type UpdateRoomMutationOptions = Apollo.BaseMutationOptions<UpdateRoomMutation, UpdateRoomMutationVariables>;
export const UploadVisualizationsDocument = gql`
    mutation uploadVisualizations($projectId: String!, $files: [Upload!]!, $roomId: String) {
  uploadVisualizations(projectId: $projectId, files: $files, roomId: $roomId) {
    __typename
    ... on VisualizationsUploaded {
      visualizations {
        id
        file {
          id
          url
        }
      }
    }
    ... on SomeVisualizationsUploaded {
      visualizations {
        id
        file {
          id
          url
        }
      }
    }
    ... on Error {
      message
    }
  }
}
    `;
export type UploadVisualizationsMutationFn = Apollo.MutationFunction<UploadVisualizationsMutation, UploadVisualizationsMutationVariables>;

/**
 * __useUploadVisualizationsMutation__
 *
 * To run a mutation, you first call `useUploadVisualizationsMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUploadVisualizationsMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [uploadVisualizationsMutation, { data, loading, error }] = useUploadVisualizationsMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      files: // value for 'files'
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useUploadVisualizationsMutation(baseOptions?: Apollo.MutationHookOptions<UploadVisualizationsMutation, UploadVisualizationsMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UploadVisualizationsMutation, UploadVisualizationsMutationVariables>(UploadVisualizationsDocument, options);
      }
export type UploadVisualizationsMutationHookResult = ReturnType<typeof useUploadVisualizationsMutation>;
export type UploadVisualizationsMutationResult = Apollo.MutationResult<UploadVisualizationsMutation>;
export type UploadVisualizationsMutationOptions = Apollo.BaseMutationOptions<UploadVisualizationsMutation, UploadVisualizationsMutationVariables>;
export const ProjectScreenDocument = gql`
    query projectScreen($id: String!) {
  project(id: $id) {
    __typename
    ... on Project {
      ...ProjectScreenProject
    }
    ... on Error {
      message
    }
  }
}
    ${ProjectScreenProjectFragmentDoc}`;

/**
 * __useProjectScreenQuery__
 *
 * To run a query within a React component, call `useProjectScreenQuery` and pass it any options that fit your needs.
 * When your component renders, `useProjectScreenQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useProjectScreenQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useProjectScreenQuery(baseOptions: Apollo.QueryHookOptions<ProjectScreenQuery, ProjectScreenQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ProjectScreenQuery, ProjectScreenQueryVariables>(ProjectScreenDocument, options);
      }
export function useProjectScreenLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ProjectScreenQuery, ProjectScreenQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ProjectScreenQuery, ProjectScreenQueryVariables>(ProjectScreenDocument, options);
        }
export type ProjectScreenQueryHookResult = ReturnType<typeof useProjectScreenQuery>;
export type ProjectScreenLazyQueryHookResult = ReturnType<typeof useProjectScreenLazyQuery>;
export type ProjectScreenQueryResult = Apollo.QueryResult<ProjectScreenQuery, ProjectScreenQueryVariables>;
export const DeleteVisualizationsDocument = gql`
    mutation deleteVisualizations($id: [String!]!) {
  deleteVisualizations(id: $id) {
    __typename
    ... on VisualizationsDeleted {
      visualizations {
        id
      }
    }
    ... on SomeVisualizationsDeleted {
      visualizations {
        id
      }
    }
    ... on NotFound {
      message
    }
    ... on Forbidden {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type DeleteVisualizationsMutationFn = Apollo.MutationFunction<DeleteVisualizationsMutation, DeleteVisualizationsMutationVariables>;

/**
 * __useDeleteVisualizationsMutation__
 *
 * To run a mutation, you first call `useDeleteVisualizationsMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteVisualizationsMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteVisualizationsMutation, { data, loading, error }] = useDeleteVisualizationsMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteVisualizationsMutation(baseOptions?: Apollo.MutationHookOptions<DeleteVisualizationsMutation, DeleteVisualizationsMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteVisualizationsMutation, DeleteVisualizationsMutationVariables>(DeleteVisualizationsDocument, options);
      }
export type DeleteVisualizationsMutationHookResult = ReturnType<typeof useDeleteVisualizationsMutation>;
export type DeleteVisualizationsMutationResult = Apollo.MutationResult<DeleteVisualizationsMutation>;
export type DeleteVisualizationsMutationOptions = Apollo.BaseMutationOptions<DeleteVisualizationsMutation, DeleteVisualizationsMutationVariables>;
export const VisualizationsScreenDocument = gql`
    query visualizationsScreen($id: String!, $filter: ProjectVisualizationsListFilter!) {
  project(id: $id) {
    __typename
    ... on Project {
      id
      name
      visualizations {
        list(filter: $filter, limit: 100, offset: 0) {
          ... on ProjectVisualizationsList {
            items {
              ...VisualizationsScreenVisualization
            }
          }
          ... on Error {
            message
          }
        }
      }
      houses {
        ...VisualizationsScreenHouses
      }
    }
    ... on Error {
      message
    }
  }
}
    ${VisualizationsScreenVisualizationFragmentDoc}
${VisualizationsScreenHousesFragmentDoc}`;

/**
 * __useVisualizationsScreenQuery__
 *
 * To run a query within a React component, call `useVisualizationsScreenQuery` and pass it any options that fit your needs.
 * When your component renders, `useVisualizationsScreenQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useVisualizationsScreenQuery({
 *   variables: {
 *      id: // value for 'id'
 *      filter: // value for 'filter'
 *   },
 * });
 */
export function useVisualizationsScreenQuery(baseOptions: Apollo.QueryHookOptions<VisualizationsScreenQuery, VisualizationsScreenQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<VisualizationsScreenQuery, VisualizationsScreenQueryVariables>(VisualizationsScreenDocument, options);
      }
export function useVisualizationsScreenLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<VisualizationsScreenQuery, VisualizationsScreenQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<VisualizationsScreenQuery, VisualizationsScreenQueryVariables>(VisualizationsScreenDocument, options);
        }
export type VisualizationsScreenQueryHookResult = ReturnType<typeof useVisualizationsScreenQuery>;
export type VisualizationsScreenLazyQueryHookResult = ReturnType<typeof useVisualizationsScreenLazyQuery>;
export type VisualizationsScreenQueryResult = Apollo.QueryResult<VisualizationsScreenQuery, VisualizationsScreenQueryVariables>;
export const CreateProjectDocument = gql`
    mutation createProject($input: CreateProjectInput!) {
  createProject(input: $input) {
    __typename
    ... on ProjectCreated {
      project {
        id
        name
        startAt
        endAt
      }
    }
    ... on ServerError {
      message
    }
    ... on Forbidden {
      message
    }
  }
}
    `;
export type CreateProjectMutationFn = Apollo.MutationFunction<CreateProjectMutation, CreateProjectMutationVariables>;

/**
 * __useCreateProjectMutation__
 *
 * To run a mutation, you first call `useCreateProjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateProjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createProjectMutation, { data, loading, error }] = useCreateProjectMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateProjectMutation(baseOptions?: Apollo.MutationHookOptions<CreateProjectMutation, CreateProjectMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateProjectMutation, CreateProjectMutationVariables>(CreateProjectDocument, options);
      }
export type CreateProjectMutationHookResult = ReturnType<typeof useCreateProjectMutation>;
export type CreateProjectMutationResult = Apollo.MutationResult<CreateProjectMutation>;
export type CreateProjectMutationOptions = Apollo.BaseMutationOptions<CreateProjectMutation, CreateProjectMutationVariables>;
export const InviteUserDocument = gql`
    mutation inviteUser($workspaceId: String!, $email: String!, $role: WorkspaceUserRole!) {
  inviteUser(workspaceId: $workspaceId, email: $email, role: $role) {
    ... on InviteSent {
      to
      tokenExpiration
    }
    ... on AlreadyInWorkspace {
      message
    }
    ... on Forbidden {
      message
    }
    ... on NotFound {
      message
    }
    ... on ServerError {
      message
    }
  }
}
    `;
export type InviteUserMutationFn = Apollo.MutationFunction<InviteUserMutation, InviteUserMutationVariables>;

/**
 * __useInviteUserMutation__
 *
 * To run a mutation, you first call `useInviteUserMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useInviteUserMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [inviteUserMutation, { data, loading, error }] = useInviteUserMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      email: // value for 'email'
 *      role: // value for 'role'
 *   },
 * });
 */
export function useInviteUserMutation(baseOptions?: Apollo.MutationHookOptions<InviteUserMutation, InviteUserMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<InviteUserMutation, InviteUserMutationVariables>(InviteUserDocument, options);
      }
export type InviteUserMutationHookResult = ReturnType<typeof useInviteUserMutation>;
export type InviteUserMutationResult = Apollo.MutationResult<InviteUserMutation>;
export type InviteUserMutationOptions = Apollo.BaseMutationOptions<InviteUserMutation, InviteUserMutationVariables>;
export const WorkspaceScreenDocument = gql`
    query workspaceScreen($id: String!, $timezone: String) {
  workspace(id: $id) {
    ... on Workspace {
      ...WorkspaceScreen
    }
    ... on Error {
      message
    }
  }
}
    ${WorkspaceScreenFragmentDoc}`;

/**
 * __useWorkspaceScreenQuery__
 *
 * To run a query within a React component, call `useWorkspaceScreenQuery` and pass it any options that fit your needs.
 * When your component renders, `useWorkspaceScreenQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useWorkspaceScreenQuery({
 *   variables: {
 *      id: // value for 'id'
 *      timezone: // value for 'timezone'
 *   },
 * });
 */
export function useWorkspaceScreenQuery(baseOptions: Apollo.QueryHookOptions<WorkspaceScreenQuery, WorkspaceScreenQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<WorkspaceScreenQuery, WorkspaceScreenQueryVariables>(WorkspaceScreenDocument, options);
      }
export function useWorkspaceScreenLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<WorkspaceScreenQuery, WorkspaceScreenQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<WorkspaceScreenQuery, WorkspaceScreenQueryVariables>(WorkspaceScreenDocument, options);
        }
export type WorkspaceScreenQueryHookResult = ReturnType<typeof useWorkspaceScreenQuery>;
export type WorkspaceScreenLazyQueryHookResult = ReturnType<typeof useWorkspaceScreenLazyQuery>;
export type WorkspaceScreenQueryResult = Apollo.QueryResult<WorkspaceScreenQuery, WorkspaceScreenQueryVariables>;