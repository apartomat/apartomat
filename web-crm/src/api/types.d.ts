import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions =  {}
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

export type AlreadyExists = Error & {
  __typename?: 'AlreadyExists';
  message: Scalars['String'];
};

export type ChangeProjectDatesInput = {
  endAt?: Maybe<Scalars['Time']>;
  startAt?: Maybe<Scalars['Time']>;
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

export type CreateProjectInput = {
  endAt?: Maybe<Scalars['Time']>;
  startAt?: Maybe<Scalars['Time']>;
  title: Scalars['String'];
  workspaceId: Scalars['String'];
};

export type CreateProjectResult = Forbidden | ProjectCreated | ServerError;

export type DeleteContactResult = ContactDeleted | Forbidden | NotFound | ServerError;

export type Error = {
  message: Scalars['String'];
};

export type ExpiredToken = Error & {
  __typename?: 'ExpiredToken';
  message: Scalars['String'];
};

export type FilesScreen = {
  __typename?: 'FilesScreen';
  menu: MenuResult;
  project: ProjectResult;
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

export type LinkSentByEmail = {
  __typename?: 'LinkSentByEmail';
  email: Scalars['String'];
};

export type LoginByEmailResult = InvalidEmail | LinkSentByEmail | PinSentByEmail | ServerError;

export type LoginConfirmed = {
  __typename?: 'LoginConfirmed';
  token: Scalars['String'];
};

export type MenuItem = {
  __typename?: 'MenuItem';
  title: Scalars['String'];
  url: Scalars['String'];
};

export type MenuItems = {
  __typename?: 'MenuItems';
  items: Array<MenuItem>;
};

export type MenuResult = MenuItems | ServerError;

export type Mutation = {
  __typename?: 'Mutation';
  addContact: AddContactResult;
  addHouse: AddHouseResult;
  changeProjectDates: ChangeProjectDatesResult;
  changeProjectStatus: ChangeProjectStatusResult;
  confirmLoginLink: ConfirmLoginLinkResult;
  confirmLoginPin: ConfirmLoginPinResult;
  createProject: CreateProjectResult;
  deleteContact: DeleteContactResult;
  loginByEmail: LoginByEmailResult;
  ping: Scalars['String'];
  updateContact: UpdateContactResult;
  updateHouse: UpdateHouseResult;
  uploadProjectFile: UploadProjectFileResult;
};


export type MutationAddContactArgs = {
  contact: AddContactInput;
  projectId: Scalars['String'];
};


export type MutationAddHouseArgs = {
  house: AddHouseInput;
  projectId: Scalars['String'];
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


export type MutationCreateProjectArgs = {
  input: CreateProjectInput;
};


export type MutationDeleteContactArgs = {
  id: Scalars['String'];
};


export type MutationLoginByEmailArgs = {
  email: Scalars['String'];
  workspaceName?: Scalars['String'];
};


export type MutationUpdateContactArgs = {
  contactId: Scalars['String'];
  data: UpdateContactInput;
};


export type MutationUpdateHouseArgs = {
  data: UpdateHouseInput;
  houseId: Scalars['String'];
};


export type MutationUploadProjectFileArgs = {
  input: UploadProjectFileInput;
};

export type NotFound = Error & {
  __typename?: 'NotFound';
  message: Scalars['String'];
};

export type PinSentByEmail = {
  __typename?: 'PinSentByEmail';
  email: Scalars['String'];
  token: Scalars['String'];
};

export type Product = {
  __typename?: 'Product';
  description: Scalars['String'];
  image: Scalars['String'];
  name: Scalars['String'];
};

export type Project = {
  __typename?: 'Project';
  contacts: ProjectContacts;
  endAt?: Maybe<Scalars['Time']>;
  files: ProjectFiles;
  houses: ProjectHouses;
  id: Scalars['String'];
  startAt?: Maybe<Scalars['Time']>;
  status: ProjectStatus;
  title: Scalars['String'];
};

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
  type?: Maybe<Array<ContactType>>;
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

export type ProjectEnums = {
  __typename?: 'ProjectEnums';
  status: ProjectStatusEnum;
};

export type ProjectFile = {
  __typename?: 'ProjectFile';
  id: Scalars['String'];
  mimeType: Scalars['String'];
  name: Scalars['String'];
  type: ProjectFileType;
  url: Scalars['Url'];
};

export enum ProjectFileType {
  None = 'NONE',
  Visualization = 'VISUALIZATION'
}

export type ProjectFileUploaded = {
  __typename?: 'ProjectFileUploaded';
  file: ProjectFile;
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

export type ProjectFilesList = {
  __typename?: 'ProjectFilesList';
  items: Array<ProjectFile>;
};

export type ProjectFilesListFilter = {
  type?: Maybe<Array<ProjectFileType>>;
};

export type ProjectFilesListResult = Forbidden | ProjectFilesList | ServerError;

export type ProjectFilesResult = Forbidden | ProjectFiles | ServerError;

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
  ID?: Maybe<Array<Scalars['String']>>;
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

export type ProjectResult = Forbidden | NotFound | Project | ServerError;

export type ProjectScreen = {
  __typename?: 'ProjectScreen';
  enums: ProjectEnums;
  menu: MenuResult;
  project: ProjectResult;
};

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

export type ProjectStatusEnum = {
  __typename?: 'ProjectStatusEnum';
  items: Array<ProjectStatusEnumItem>;
};

export type ProjectStatusEnumItem = {
  __typename?: 'ProjectStatusEnumItem';
  key: ProjectStatus;
  value: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  profile: UserProfileResult;
  screen: ScreenQuery;
  shoppinglist: ShoppinglistQuery;
  version: Scalars['String'];
  workspace: WorkspaceResult;
};


export type QueryWorkspaceArgs = {
  id: Scalars['String'];
};

export type Room = {
  __typename?: 'Room';
  createdAt: Scalars['Time'];
  design: Scalars['Boolean'];
  id: Scalars['String'];
  modifiedAt: Scalars['Time'];
  name: Scalars['String'];
  square?: Maybe<Scalars['Float']>;
};

export type ScreenQuery = {
  __typename?: 'ScreenQuery';
  files: FilesScreen;
  project: ProjectScreen;
  spec: SpecScreen;
  version: Scalars['String'];
};


export type ScreenQueryFilesArgs = {
  projectId: Scalars['String'];
};


export type ScreenQueryProjectArgs = {
  id: Scalars['String'];
};


export type ScreenQuerySpecArgs = {
  projectId: Scalars['String'];
};

export type ServerError = Error & {
  __typename?: 'ServerError';
  message: Scalars['String'];
};

export type ShoppinglistQuery = {
  __typename?: 'ShoppinglistQuery';
  productOnPage?: Maybe<Product>;
};


export type ShoppinglistQueryProductOnPageArgs = {
  url: Scalars['String'];
};

export type SpecScreen = {
  __typename?: 'SpecScreen';
  menu: MenuResult;
  project: ProjectResult;
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

export type UploadProjectFileInput = {
  file: Scalars['Upload'];
  projectId: Scalars['String'];
  type: ProjectFileType;
};

export type UploadProjectFileResult = AlreadyExists | Forbidden | ProjectFileUploaded | ServerError;

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

export type Workspace = {
  __typename?: 'Workspace';
  id: Scalars['String'];
  name: Scalars['String'];
  projects: WorkspaceProjects;
  users: WorkspaceUsersResult;
};

export type WorkspaceProject = {
  __typename?: 'WorkspaceProject';
  endAt?: Maybe<Scalars['Time']>;
  id: Scalars['String'];
  name: Scalars['String'];
  period?: Maybe<Scalars['String']>;
  startAt?: Maybe<Scalars['Time']>;
  status: ProjectStatus;
};


export type WorkspaceProjectPeriodArgs = {
  timezone?: Maybe<Scalars['String']>;
};

export type WorkspaceProjects = {
  __typename?: 'WorkspaceProjects';
  list: WorkspaceProjectsListResult;
  total: WorkspaceProjectsTotalResult;
  workspace?: Maybe<Id>;
};


export type WorkspaceProjectsListArgs = {
  filter?: WorkspaceProjectsFilter;
  limit?: Scalars['Int'];
};


export type WorkspaceProjectsTotalArgs = {
  filter?: WorkspaceProjectsFilter;
};

export type WorkspaceProjectsFilter = {
  status?: Maybe<Array<ProjectStatus>>;
};

export type WorkspaceProjectsList = {
  __typename?: 'WorkspaceProjectsList';
  items: Array<WorkspaceProject>;
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
  profile: WorkspaceUserProfile;
  role: WorkspaceUserRole;
  workspace: Id;
};

export type WorkspaceUserProfile = {
  __typename?: 'WorkspaceUserProfile';
  abbr: Scalars['String'];
  email: Scalars['String'];
  fullName: Scalars['String'];
  gravatar?: Maybe<Gravatar>;
  id: Scalars['String'];
};

export enum WorkspaceUserRole {
  Admin = 'ADMIN',
  User = 'USER'
}

export type WorkspaceUsers = {
  __typename?: 'WorkspaceUsers';
  items: Array<WorkspaceUser>;
};

export type WorkspaceUsersResult = Forbidden | ServerError | WorkspaceUsers;

export type AddContactMutationVariables = Exact<{
  projectId: Scalars['String'];
  contact: AddContactInput;
}>;


export type AddContactMutation = { __typename?: 'Mutation', addContact: { __typename: 'ContactAdded', contact: { __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> } } | { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } };

export type AddHouseFragment = { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any };

export type AddHouseMutationVariables = Exact<{
  projectId: Scalars['String'];
  house: AddHouseInput;
}>;


export type AddHouseMutation = { __typename?: 'Mutation', addHouse: { __typename: 'Forbidden', message: string } | { __typename: 'HouseAdded', house: { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any } } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type ChangeProjectDatesMutationVariables = Exact<{
  projectId: Scalars['String'];
  dates: ChangeProjectDatesInput;
}>;


export type ChangeProjectDatesMutation = { __typename?: 'Mutation', changeProjectDates: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ProjectDatesChanged', project: { __typename?: 'Project', startAt?: any | null | undefined, endAt?: any | null | undefined } } | { __typename: 'ServerError', message: string } };

export type ChangeProjectStatusMutationVariables = Exact<{
  projectId: Scalars['String'];
  status: ProjectStatus;
}>;


export type ChangeProjectStatusMutation = { __typename?: 'Mutation', changeProjectStatus: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ProjectStatusChanged', project: { __typename?: 'Project', status: ProjectStatus } } | { __typename: 'ServerError', message: string } };

export type ConfirmLoginLinkMutationVariables = Exact<{
  token: Scalars['String'];
}>;


export type ConfirmLoginLinkMutation = { __typename?: 'Mutation', confirmLoginLink: { __typename: 'ExpiredToken', message: string } | { __typename: 'InvalidToken', message: string } | { __typename: 'LoginConfirmed', token: string } | { __typename: 'ServerError', message: string } };

export type ConfirmLoginPinMutationVariables = Exact<{
  token: Scalars['String'];
  pin: Scalars['String'];
}>;


export type ConfirmLoginPinMutation = { __typename?: 'Mutation', confirmLoginPin: { __typename: 'ExpiredToken', message: string } | { __typename: 'InvalidPin', message: string } | { __typename: 'InvalidToken', message: string } | { __typename: 'LoginConfirmed', token: string } | { __typename: 'ServerError', message: string } };

export type CreateProjectMutationVariables = Exact<{
  input: CreateProjectInput;
}>;


export type CreateProjectMutation = { __typename?: 'Mutation', createProject: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectCreated', project: { __typename?: 'Project', id: string, title: string, startAt?: any | null | undefined, endAt?: any | null | undefined } } | { __typename: 'ServerError', message: string } };

export type DeleteContactMutationVariables = Exact<{
  id: Scalars['String'];
}>;


export type DeleteContactMutation = { __typename?: 'Mutation', deleteContact: { __typename: 'ContactDeleted', contact: { __typename?: 'Contact', id: string } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type LoginByEmailMutationVariables = Exact<{
  email: Scalars['String'];
}>;


export type LoginByEmailMutation = { __typename?: 'Mutation', loginByEmail: { __typename: 'InvalidEmail', message: string } | { __typename: 'LinkSentByEmail', email: string } | { __typename: 'PinSentByEmail', email: string, token: string } | { __typename: 'ServerError', message: string } };

export type ProfileQueryVariables = Exact<{ [key: string]: never; }>;


export type ProfileQuery = { __typename?: 'Query', profile: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'UserProfile', id: string, email: string, gravatar?: { __typename?: 'Gravatar', url: string } | null | undefined, defaultWorkspace: { __typename?: 'Workspace', id: string, name: string } } };

export type ProjectQueryVariables = Exact<{
  id: Scalars['String'];
}>;


export type ProjectQuery = { __typename?: 'Query', screen: { __typename?: 'ScreenQuery', projectScreen: { __typename?: 'ProjectScreen', project: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'Project', id: string, title: string, startAt?: any | null | undefined, endAt?: any | null | undefined, status: ProjectStatus, contacts: { __typename?: 'ProjectContacts', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectContactsList', items: Array<{ __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectContactsTotal', total: number } | { __typename: 'ServerError' } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any, rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null | undefined, design: boolean }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } }, files: { __typename?: 'ProjectFiles', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectFilesList', items: Array<{ __typename?: 'ProjectFile', id: string, name: string, url: any, type: ProjectFileType }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectFilesTotal', total: number } | { __typename: 'ServerError' } } } | { __typename: 'ServerError', message: string }, menu: { __typename: 'MenuItems', items: Array<{ __typename?: 'MenuItem', title: string, url: string }> } | { __typename: 'ServerError' }, enums: { __typename?: 'ProjectEnums', status: { __typename?: 'ProjectStatusEnum', items: Array<{ __typename?: 'ProjectStatusEnumItem', key: ProjectStatus, value: string }> } } } } };

export type ProjectScreenProjectFragment = { __typename?: 'Project', id: string, title: string, startAt?: any | null | undefined, endAt?: any | null | undefined, status: ProjectStatus, contacts: { __typename?: 'ProjectContacts', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectContactsList', items: Array<{ __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectContactsTotal', total: number } | { __typename: 'ServerError' } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any, rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null | undefined, design: boolean }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } }, files: { __typename?: 'ProjectFiles', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectFilesList', items: Array<{ __typename?: 'ProjectFile', id: string, name: string, url: any, type: ProjectFileType }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectFilesTotal', total: number } | { __typename: 'ServerError' } } };

export type ProjectScreenHousesFragment = { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any, rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null | undefined, design: boolean }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } };

export type ProjectScreenHouseFragment = { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any };

export type ProjectScreenHouseRoomsFragment = { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null | undefined, design: boolean }> } | { __typename?: 'ServerError', message: string } };

export type ProjectScreenHouseRoomFragment = { __typename?: 'Room', id: string, name: string, square?: number | null | undefined, design: boolean };

export type SpecScreenQueryVariables = Exact<{
  projectId: Scalars['String'];
}>;


export type SpecScreenQuery = { __typename?: 'Query', screen: { __typename?: 'ScreenQuery', specScreen: { __typename?: 'SpecScreen', project: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'Project', id: string, title: string } | { __typename: 'ServerError', message: string } } } };

export type UpdateContactMutationVariables = Exact<{
  contactId: Scalars['String'];
  data: UpdateContactInput;
}>;


export type UpdateContactMutation = { __typename?: 'Mutation', updateContact: { __typename: 'ContactUpdated', contact: { __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type UpdateHouseFragment = { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any };

export type UpdateHouseMutationVariables = Exact<{
  houseId: Scalars['String'];
  house: UpdateHouseInput;
}>;


export type UpdateHouseMutation = { __typename?: 'Mutation', updateHouse: { __typename: 'Forbidden', message: string } | { __typename: 'HouseUpdated', house: { __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any } } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type UploadProjectFileMutationVariables = Exact<{
  input: UploadProjectFileInput;
}>;


export type UploadProjectFileMutation = { __typename?: 'Mutation', uploadProjectFile: { __typename: 'AlreadyExists', message: string } | { __typename: 'Forbidden', message: string } | { __typename: 'ProjectFileUploaded', file: { __typename?: 'ProjectFile', id: string, url: any } } | { __typename: 'ServerError', message: string } };

export type WorkspaceQueryVariables = Exact<{
  id: Scalars['String'];
  timezone?: Maybe<Scalars['String']>;
}>;


export type WorkspaceQuery = { __typename?: 'Query', workspace: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'Workspace', id: string, name: string, users: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceUsers', items: Array<{ __typename?: 'WorkspaceUser', id: string, role: WorkspaceUserRole, workspace: { __typename?: 'Id', id: string }, profile: { __typename?: 'WorkspaceUserProfile', id: string, email: string, fullName: string, abbr: string, gravatar?: { __typename?: 'Gravatar', url: string } | null | undefined } }> }, projects: { __typename: 'WorkspaceProjects', current: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceProjectsList', items: Array<{ __typename?: 'WorkspaceProject', id: string, name: string, status: ProjectStatus, startAt?: any | null | undefined, endAt?: any | null | undefined, period?: string | null | undefined }> }, done: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceProjectsList', items: Array<{ __typename?: 'WorkspaceProject', id: string, name: string, status: ProjectStatus, startAt?: any | null | undefined, endAt?: any | null | undefined, period?: string | null | undefined }> } } } };

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
  design
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
export const ProjectScreenProjectFragmentDoc = gql`
    fragment ProjectScreenProject on Project {
  id
  title
  startAt
  endAt
  status
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
  files {
    list(filter: {type: [VISUALIZATION]}, limit: 10, offset: 0) {
      __typename
      ... on ProjectFilesList {
        items {
          id
          name
          url
          type
        }
      }
      ... on Error {
        message
      }
    }
    total {
      __typename
      ... on ProjectFilesTotal {
        total
      }
    }
  }
}
    ${ProjectScreenHousesFragmentDoc}`;
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
export const CreateProjectDocument = gql`
    mutation createProject($input: CreateProjectInput!) {
  createProject(input: $input) {
    __typename
    ... on ProjectCreated {
      project {
        id
        title
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
export const ProjectDocument = gql`
    query project($id: String!) {
  screen {
    projectScreen: project(id: $id) {
      project {
        __typename
        ... on Project {
          ...ProjectScreenProject
        }
        ... on Error {
          message
        }
      }
      menu {
        __typename
        ... on MenuItems {
          items {
            title
            url
          }
        }
      }
      enums {
        status {
          items {
            key
            value
          }
        }
      }
    }
  }
}
    ${ProjectScreenProjectFragmentDoc}`;

/**
 * __useProjectQuery__
 *
 * To run a query within a React component, call `useProjectQuery` and pass it any options that fit your needs.
 * When your component renders, `useProjectQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useProjectQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useProjectQuery(baseOptions: Apollo.QueryHookOptions<ProjectQuery, ProjectQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ProjectQuery, ProjectQueryVariables>(ProjectDocument, options);
      }
export function useProjectLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ProjectQuery, ProjectQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ProjectQuery, ProjectQueryVariables>(ProjectDocument, options);
        }
export type ProjectQueryHookResult = ReturnType<typeof useProjectQuery>;
export type ProjectLazyQueryHookResult = ReturnType<typeof useProjectLazyQuery>;
export type ProjectQueryResult = Apollo.QueryResult<ProjectQuery, ProjectQueryVariables>;
export const SpecScreenDocument = gql`
    query specScreen($projectId: String!) {
  screen {
    specScreen: spec(projectId: $projectId) {
      project {
        __typename
        ... on Project {
          id
          title
        }
        ... on Error {
          message
        }
      }
    }
  }
}
    `;

/**
 * __useSpecScreenQuery__
 *
 * To run a query within a React component, call `useSpecScreenQuery` and pass it any options that fit your needs.
 * When your component renders, `useSpecScreenQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useSpecScreenQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useSpecScreenQuery(baseOptions: Apollo.QueryHookOptions<SpecScreenQuery, SpecScreenQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<SpecScreenQuery, SpecScreenQueryVariables>(SpecScreenDocument, options);
      }
export function useSpecScreenLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<SpecScreenQuery, SpecScreenQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<SpecScreenQuery, SpecScreenQueryVariables>(SpecScreenDocument, options);
        }
export type SpecScreenQueryHookResult = ReturnType<typeof useSpecScreenQuery>;
export type SpecScreenLazyQueryHookResult = ReturnType<typeof useSpecScreenLazyQuery>;
export type SpecScreenQueryResult = Apollo.QueryResult<SpecScreenQuery, SpecScreenQueryVariables>;
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
export const UploadProjectFileDocument = gql`
    mutation uploadProjectFile($input: UploadProjectFileInput!) {
  uploadProjectFile(input: $input) {
    __typename
    ... on ProjectFileUploaded {
      file {
        id
        url
      }
    }
    ... on Error {
      message
    }
  }
}
    `;
export type UploadProjectFileMutationFn = Apollo.MutationFunction<UploadProjectFileMutation, UploadProjectFileMutationVariables>;

/**
 * __useUploadProjectFileMutation__
 *
 * To run a mutation, you first call `useUploadProjectFileMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUploadProjectFileMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [uploadProjectFileMutation, { data, loading, error }] = useUploadProjectFileMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUploadProjectFileMutation(baseOptions?: Apollo.MutationHookOptions<UploadProjectFileMutation, UploadProjectFileMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UploadProjectFileMutation, UploadProjectFileMutationVariables>(UploadProjectFileDocument, options);
      }
export type UploadProjectFileMutationHookResult = ReturnType<typeof useUploadProjectFileMutation>;
export type UploadProjectFileMutationResult = Apollo.MutationResult<UploadProjectFileMutation>;
export type UploadProjectFileMutationOptions = Apollo.BaseMutationOptions<UploadProjectFileMutation, UploadProjectFileMutationVariables>;
export const WorkspaceDocument = gql`
    query workspace($id: String!, $timezone: String) {
  workspace(id: $id) {
    __typename
    ... on Workspace {
      id
      name
      users {
        __typename
        ... on WorkspaceUsers {
          items {
            id
            role
            workspace {
              id
            }
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
        }
        ... on Error {
          message
        }
      }
      projects {
        __typename
        current: list(filter: {status: [NEW, IN_PROGRESS]}, limit: 10) {
          __typename
          ... on WorkspaceProjectsList {
            items {
              id
              name
              status
              startAt
              endAt
              period(timezone: $timezone)
            }
          }
          ... on Error {
            message
          }
        }
        done: list(filter: {status: [DONE, CANCELED]}, limit: 10) {
          __typename
          ... on WorkspaceProjectsList {
            items {
              id
              name
              status
              startAt
              endAt
              period(timezone: $timezone)
            }
          }
          ... on Error {
            message
          }
        }
      }
    }
    ... on Error {
      message
    }
  }
}
    `;

/**
 * __useWorkspaceQuery__
 *
 * To run a query within a React component, call `useWorkspaceQuery` and pass it any options that fit your needs.
 * When your component renders, `useWorkspaceQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useWorkspaceQuery({
 *   variables: {
 *      id: // value for 'id'
 *      timezone: // value for 'timezone'
 *   },
 * });
 */
export function useWorkspaceQuery(baseOptions: Apollo.QueryHookOptions<WorkspaceQuery, WorkspaceQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<WorkspaceQuery, WorkspaceQueryVariables>(WorkspaceDocument, options);
      }
export function useWorkspaceLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<WorkspaceQuery, WorkspaceQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<WorkspaceQuery, WorkspaceQueryVariables>(WorkspaceDocument, options);
        }
export type WorkspaceQueryHookResult = ReturnType<typeof useWorkspaceQuery>;
export type WorkspaceLazyQueryHookResult = ReturnType<typeof useWorkspaceLazyQuery>;
export type WorkspaceQueryResult = Apollo.QueryResult<WorkspaceQuery, WorkspaceQueryVariables>;