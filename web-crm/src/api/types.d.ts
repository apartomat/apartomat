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

export type AlreadyExists = Error & {
  __typename?: 'AlreadyExists';
  message: Scalars['String'];
};

export type CheckEmail = {
  __typename?: 'CheckEmail';
  email: Scalars['String'];
};

export type ConfirmLoginResult = ExpiredToken | InvalidToken | LoginConfirmed | ServerError;

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

export type CreateProjectInput = {
  endAt?: Maybe<Scalars['Time']>;
  startAt?: Maybe<Scalars['Time']>;
  title: Scalars['String'];
  workspaceId: Scalars['Int'];
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

export type Id = {
  __typename?: 'Id';
  id: Scalars['Int'];
};

export type InvalidEmail = Error & {
  __typename?: 'InvalidEmail';
  message: Scalars['String'];
};

export type InvalidToken = Error & {
  __typename?: 'InvalidToken';
  message: Scalars['String'];
};

export type LoginByEmailResult = CheckEmail | InvalidEmail | ServerError;

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
  confirmLogin: ConfirmLoginResult;
  createProject: CreateProjectResult;
  deleteContact: DeleteContactResult;
  loginByEmail: LoginByEmailResult;
  ping: Scalars['String'];
  uploadProjectFile: UploadProjectFileResult;
};


export type MutationAddContactArgs = {
  contact: AddContactInput;
  projectId: Scalars['String'];
};


export type MutationConfirmLoginArgs = {
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


export type MutationUploadProjectFileArgs = {
  input: UploadProjectFileInput;
};

export type NotFound = Error & {
  __typename?: 'NotFound';
  message: Scalars['String'];
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
  id: Scalars['Int'];
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

export type ProjectFile = {
  __typename?: 'ProjectFile';
  id: Scalars['Int'];
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
  menu: MenuResult;
  project: ProjectResult;
};

export enum ProjectStatus {
  Canceled = 'CANCELED',
  Done = 'DONE',
  InProgress = 'IN_PROGRESS',
  New = 'NEW'
}

export type Query = {
  __typename?: 'Query';
  profile: UserProfileResult;
  screen: ScreenQuery;
  shoppinglist: ShoppinglistQuery;
  version: Scalars['String'];
  workspace: WorkspaceResult;
};


export type QueryWorkspaceArgs = {
  id: Scalars['Int'];
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
  projectId: Scalars['Int'];
};


export type ScreenQueryProjectArgs = {
  id: Scalars['Int'];
};


export type ScreenQuerySpecArgs = {
  projectId: Scalars['Int'];
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

export type UploadProjectFileInput = {
  file: Scalars['Upload'];
  projectId: Scalars['Int'];
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
  id: Scalars['Int'];
};

export type UserProfileResult = Forbidden | ServerError | UserProfile;

export type Workspace = {
  __typename?: 'Workspace';
  id: Scalars['Int'];
  name: Scalars['String'];
  projects: WorkspaceProjects;
  users: WorkspaceUsersResult;
};

export type WorkspaceProject = {
  __typename?: 'WorkspaceProject';
  id: Scalars['Int'];
  name: Scalars['String'];
  period?: Maybe<Scalars['String']>;
  status: ProjectStatus;
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
  id: Scalars['Int'];
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
  id: Scalars['Int'];
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

export type ConfirmLoginMutationVariables = Exact<{
  token: Scalars['String'];
}>;


export type ConfirmLoginMutation = { __typename?: 'Mutation', confirmLogin: { __typename: 'ExpiredToken', message: string } | { __typename: 'InvalidToken', message: string } | { __typename: 'LoginConfirmed', token: string } | { __typename: 'ServerError', message: string } };

export type CreateProjectMutationVariables = Exact<{
  input: CreateProjectInput;
}>;


export type CreateProjectMutation = { __typename?: 'Mutation', createProject: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectCreated', project: { __typename?: 'Project', id: number, title: string, startAt?: any | null | undefined, endAt?: any | null | undefined } } | { __typename: 'ServerError', message: string } };

export type DeleteContactMutationVariables = Exact<{
  id: Scalars['String'];
}>;


export type DeleteContactMutation = { __typename?: 'Mutation', deleteContact: { __typename: 'ContactDeleted', contact: { __typename?: 'Contact', id: string } } | { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } };

export type LoginByEmailMutationVariables = Exact<{
  email: Scalars['String'];
}>;


export type LoginByEmailMutation = { __typename?: 'Mutation', loginByEmail: { __typename: 'CheckEmail', email: string } | { __typename: 'InvalidEmail', message: string } | { __typename: 'ServerError', message: string } };

export type ProfileQueryVariables = Exact<{ [key: string]: never; }>;


export type ProfileQuery = { __typename?: 'Query', profile: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'UserProfile', id: number, email: string, gravatar?: { __typename?: 'Gravatar', url: string } | null | undefined, defaultWorkspace: { __typename?: 'Workspace', id: number, name: string } } };

export type ProjectScreenHouseRoomsFragment = { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null | undefined, design: boolean }> } | { __typename?: 'ServerError', message: string } };

export type ProjectScreenHousesFragment = { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any, rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null | undefined, design: boolean }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } };

export type ProjectQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type ProjectQuery = { __typename?: 'Query', screen: { __typename?: 'ScreenQuery', projectScreen: { __typename?: 'ProjectScreen', project: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'Project', id: number, title: string, startAt?: any | null | undefined, endAt?: any | null | undefined, contacts: { __typename?: 'ProjectContacts', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectContactsList', items: Array<{ __typename?: 'Contact', id: string, fullName: string, photo: string, details: Array<{ __typename?: 'ContactDetails', type: ContactType, value: string }> }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectContactsTotal', total: number } | { __typename: 'ServerError' } }, houses: { __typename?: 'ProjectHouses', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectHousesList', items: Array<{ __typename?: 'House', id: string, city: string, address: string, housingComplex: string, createdAt: any, modifiedAt: any, rooms: { __typename?: 'HouseRooms', list: { __typename?: 'Forbidden', message: string } | { __typename?: 'HouseRoomsList', items: Array<{ __typename?: 'Room', id: string, name: string, square?: number | null | undefined, design: boolean }> } | { __typename?: 'ServerError', message: string } } }> } | { __typename: 'ServerError', message: string } }, files: { __typename?: 'ProjectFiles', list: { __typename: 'Forbidden', message: string } | { __typename: 'ProjectFilesList', items: Array<{ __typename?: 'ProjectFile', id: number, name: string, url: any, type: ProjectFileType }> } | { __typename: 'ServerError', message: string }, total: { __typename: 'Forbidden' } | { __typename: 'ProjectFilesTotal', total: number } | { __typename: 'ServerError' } } } | { __typename: 'ServerError', message: string }, menu: { __typename: 'MenuItems', items: Array<{ __typename?: 'MenuItem', title: string, url: string }> } | { __typename: 'ServerError' } } } };

export type SpecScreenQueryVariables = Exact<{
  projectId: Scalars['Int'];
}>;


export type SpecScreenQuery = { __typename?: 'Query', screen: { __typename?: 'ScreenQuery', screen: { __typename?: 'SpecScreen', project: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'Project', id: number, title: string } | { __typename: 'ServerError', message: string } } } };

export type UploadProjectFileMutationVariables = Exact<{
  input: UploadProjectFileInput;
}>;


export type UploadProjectFileMutation = { __typename?: 'Mutation', uploadProjectFile: { __typename: 'AlreadyExists', message: string } | { __typename: 'Forbidden', message: string } | { __typename: 'ProjectFileUploaded', file: { __typename?: 'ProjectFile', id: number, url: any } } | { __typename: 'ServerError', message: string } };

export type WorkspaceQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type WorkspaceQuery = { __typename?: 'Query', workspace: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'Workspace', id: number, name: string, users: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceUsers', items: Array<{ __typename?: 'WorkspaceUser', id: number, role: WorkspaceUserRole, workspace: { __typename?: 'Id', id: number }, profile: { __typename?: 'WorkspaceUserProfile', id: number, email: string, fullName: string, abbr: string, gravatar?: { __typename?: 'Gravatar', url: string } | null | undefined } }> }, projects: { __typename: 'WorkspaceProjects', current: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceProjectsList', items: Array<{ __typename?: 'WorkspaceProject', id: number, name: string, status: ProjectStatus, period?: string | null | undefined }> }, done: { __typename: 'Forbidden', message: string } | { __typename: 'ServerError', message: string } | { __typename: 'WorkspaceProjectsList', items: Array<{ __typename?: 'WorkspaceProject', id: number, name: string, status: ProjectStatus, period?: string | null | undefined }> } } } };

export const ProjectScreenHouseRoomsFragmentDoc = gql`
    fragment ProjectScreenHouseRooms on HouseRooms {
  list {
    ... on HouseRoomsList {
      items {
        id
        name
        square
        design
      }
    }
    ... on Error {
      message
    }
  }
}
    `;
export const ProjectScreenHousesFragmentDoc = gql`
    fragment ProjectScreenHouses on ProjectHouses {
  list(filter: {}, limit: 1, offset: 0) {
    __typename
    ... on ProjectHousesList {
      items {
        id
        city
        address
        housingComplex
        createdAt
        modifiedAt
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
    ${ProjectScreenHouseRoomsFragmentDoc}`;
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
export const ConfirmLoginDocument = gql`
    mutation confirmLogin($token: String!) {
  confirmLogin(token: $token) {
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
export type ConfirmLoginMutationFn = Apollo.MutationFunction<ConfirmLoginMutation, ConfirmLoginMutationVariables>;

/**
 * __useConfirmLoginMutation__
 *
 * To run a mutation, you first call `useConfirmLoginMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useConfirmLoginMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [confirmLoginMutation, { data, loading, error }] = useConfirmLoginMutation({
 *   variables: {
 *      token: // value for 'token'
 *   },
 * });
 */
export function useConfirmLoginMutation(baseOptions?: Apollo.MutationHookOptions<ConfirmLoginMutation, ConfirmLoginMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ConfirmLoginMutation, ConfirmLoginMutationVariables>(ConfirmLoginDocument, options);
      }
export type ConfirmLoginMutationHookResult = ReturnType<typeof useConfirmLoginMutation>;
export type ConfirmLoginMutationResult = Apollo.MutationResult<ConfirmLoginMutation>;
export type ConfirmLoginMutationOptions = Apollo.BaseMutationOptions<ConfirmLoginMutation, ConfirmLoginMutationVariables>;
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
    ... on CheckEmail {
      email
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
    query project($id: Int!) {
  screen {
    projectScreen: project(id: $id) {
      project {
        __typename
        ... on Project {
          id
          title
          startAt
          endAt
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
    }
  }
}
    ${ProjectScreenHousesFragmentDoc}`;

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
    query specScreen($projectId: Int!) {
  screen {
    screen: spec(projectId: $projectId) {
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
    query workspace($id: Int!) {
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
              period
            }
          }
          ... on Error {
            message
          }
        }
        done: list(filter: {status: [DONE]}, limit: 10) {
          __typename
          ... on WorkspaceProjectsList {
            items {
              id
              name
              status
              period
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