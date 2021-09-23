import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
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

export type ShoppinglistQuery = {
  __typename?: 'ShoppinglistQuery';
  productOnPage?: Maybe<Product>;
};


export type ShoppinglistQueryProductOnPageArgs = {
  url: Scalars['String'];
};

export type UserProfileResult = UserProfile | Forbidden | ServerError;

export type Query = {
  __typename?: 'Query';
  version: Scalars['String'];
  profile: UserProfileResult;
  screen: ScreenQuery;
  shoppinglist: ShoppinglistQuery;
  workspace: WorkspaceResult;
};


export type QueryWorkspaceArgs = {
  id: Scalars['Int'];
};

export type Gravatar = {
  __typename?: 'Gravatar';
  url: Scalars['String'];
};

export type Id = {
  __typename?: 'Id';
  id: Scalars['Int'];
};

export type ConfirmLoginResult = LoginConfirmed | InvalidToken | ExpiredToken | ServerError;

export type LoginConfirmed = {
  __typename?: 'LoginConfirmed';
  token: Scalars['String'];
};

export type ProjectFilesList = {
  __typename?: 'ProjectFilesList';
  items: Array<ProjectFile>;
};

export type ScreenQuery = {
  __typename?: 'ScreenQuery';
  version: Scalars['String'];
  files: FilesScreen;
  project: ProjectScreen;
  spec: SpecScreen;
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

export type MenuItem = {
  __typename?: 'MenuItem';
  title: Scalars['String'];
  url: Scalars['String'];
};

export enum WorkspaceUserRole {
  Admin = 'ADMIN',
  User = 'USER'
}

export type ProjectFilesListFilter = {
  type?: Maybe<Array<ProjectFileType>>;
};

/**  Common types  */
export type Error = {
  message: Scalars['String'];
};

export type Forbidden = Error & {
  __typename?: 'Forbidden';
  message: Scalars['String'];
};

export type WorkspaceUser = {
  __typename?: 'WorkspaceUser';
  id: Scalars['Int'];
  workspace: Id;
  role: WorkspaceUserRole;
  profile: WorkspaceUserProfile;
};

export type WorkspaceProjectsList = {
  __typename?: 'WorkspaceProjectsList';
  items: Array<WorkspaceProject>;
};

export enum ProjectFileType {
  None = 'NONE',
  Image = 'IMAGE'
}

export type ProjectFilesResult = ProjectFiles | Forbidden | ServerError;

export type FilesScreen = {
  __typename?: 'FilesScreen';
  project: ProjectResult;
  menu: MenuResult;
};

export type ProjectScreen = {
  __typename?: 'ProjectScreen';
  project: ProjectResult;
  menu: MenuResult;
};

export type WorkspaceUsersResult = WorkspaceUsers | Forbidden | ServerError;

export type InvalidToken = Error & {
  __typename?: 'InvalidToken';
  message: Scalars['String'];
};

export type ProjectFile = {
  __typename?: 'ProjectFile';
  id: Scalars['Int'];
  name: Scalars['String'];
  url: Scalars['Url'];
  type: ProjectFileType;
  mimeType: Scalars['String'];
};

export type UploadProjectFileInput = {
  projectId: Scalars['Int'];
  file: Scalars['Upload'];
  type?: Maybe<ProjectFileType>;
};

export type MenuResult = MenuItems | ServerError;

export type CreateProjectResult = Project | Forbidden | ServerError;

export type Product = {
  __typename?: 'Product';
  name: Scalars['String'];
  description: Scalars['String'];
  image: Scalars['String'];
};

export type InvalidEmail = Error & {
  __typename?: 'InvalidEmail';
  message: Scalars['String'];
};

export type WorkspaceProjectsListResult = WorkspaceProjectsList | Forbidden | ServerError;

export type WorkspaceProject = {
  __typename?: 'WorkspaceProject';
  id: Scalars['Int'];
  name: Scalars['String'];
  period?: Maybe<Scalars['String']>;
};

export type LoginByEmailResult = CheckEmail | InvalidEmail | ServerError;

export type NotFound = Error & {
  __typename?: 'NotFound';
  message: Scalars['String'];
};

export type WorkspaceProjects = {
  __typename?: 'WorkspaceProjects';
  workspace?: Maybe<Id>;
  list: WorkspaceProjectsListResult;
  total: WorkspaceProjectsTotalResult;
};

export type WorkspaceProjectsTotalResult = WorkspaceProjectsTotal | Forbidden | ServerError;

export type Project = {
  __typename?: 'Project';
  id: Scalars['Int'];
  title: Scalars['String'];
  startAt?: Maybe<Scalars['Time']>;
  endAt?: Maybe<Scalars['Time']>;
  files: ProjectFiles;
};

export type AlreadyExists = Error & {
  __typename?: 'AlreadyExists';
  message: Scalars['String'];
};

export type ProjectFilesListResult = ProjectFilesList | Forbidden | ServerError;

export type UploadProjectFileResult = ProjectFile | Forbidden | AlreadyExists | ServerError;

export type WorkspaceUsers = {
  __typename?: 'WorkspaceUsers';
  items: Array<WorkspaceUser>;
};

export type WorkspaceUserProfile = {
  __typename?: 'WorkspaceUserProfile';
  id: Scalars['Int'];
  email: Scalars['String'];
  fullName: Scalars['String'];
  abbr: Scalars['String'];
  gravatar?: Maybe<Gravatar>;
};

export type UserProfile = {
  __typename?: 'UserProfile';
  id: Scalars['Int'];
  email: Scalars['String'];
  fullName: Scalars['String'];
  abbr: Scalars['String'];
  gravatar?: Maybe<Gravatar>;
  defaultWorkspace: Workspace;
};

export type ProjectResult = Project | NotFound | Forbidden | ServerError;

export type Mutation = {
  __typename?: 'Mutation';
  loginByEmail: LoginByEmailResult;
  confirmLogin: ConfirmLoginResult;
  createProject: CreateProjectResult;
  uploadProjectFile: UploadProjectFileResult;
};


export type MutationLoginByEmailArgs = {
  email: Scalars['String'];
  workspaceName?: Scalars['String'];
};


export type MutationConfirmLoginArgs = {
  token: Scalars['String'];
};


export type MutationCreateProjectArgs = {
  input: CreateProjectInput;
};


export type MutationUploadProjectFileArgs = {
  file: UploadProjectFileInput;
};

export type CheckEmail = {
  __typename?: 'CheckEmail';
  email: Scalars['String'];
};

export type WorkspaceProjectsTotal = {
  __typename?: 'WorkspaceProjectsTotal';
  total: Scalars['Int'];
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

export type ServerError = Error & {
  __typename?: 'ServerError';
  message: Scalars['String'];
};


export type Workspace = {
  __typename?: 'Workspace';
  id: Scalars['Int'];
  name: Scalars['String'];
  users: WorkspaceUsersResult;
  projects: WorkspaceProjects;
};

export type ProjectFilesTotal = {
  __typename?: 'ProjectFilesTotal';
  total: Scalars['Int'];
};

export type CreateProjectInput = {
  workspaceId: Scalars['Int'];
  title: Scalars['String'];
};


export type ProjectFilesTotalResult = ProjectFilesTotal | Forbidden | ServerError;


export type MenuItems = {
  __typename?: 'MenuItems';
  items: Array<MenuItem>;
};

export type SpecScreen = {
  __typename?: 'SpecScreen';
  project: ProjectResult;
  menu: MenuResult;
};

export type WorkspaceResult = Workspace | NotFound | Forbidden | ServerError;

export type ExpiredToken = Error & {
  __typename?: 'ExpiredToken';
  message: Scalars['String'];
};

export type ConfirmLoginMutationVariables = Exact<{
  token: Scalars['String'];
}>;


export type ConfirmLoginMutation = (
  { __typename?: 'Mutation' }
  & { confirmLogin: (
    { __typename: 'LoginConfirmed' }
    & Pick<LoginConfirmed, 'token'>
  ) | (
    { __typename: 'InvalidToken' }
    & Pick<InvalidToken, 'message'>
  ) | (
    { __typename: 'ExpiredToken' }
    & Pick<ExpiredToken, 'message'>
  ) | (
    { __typename: 'ServerError' }
    & Pick<ServerError, 'message'>
  ) }
);

export type CreateProjectMutationVariables = Exact<{
  input: CreateProjectInput;
}>;


export type CreateProjectMutation = (
  { __typename?: 'Mutation' }
  & { createProject: (
    { __typename: 'Project' }
    & Pick<Project, 'id' | 'title'>
  ) | (
    { __typename: 'Forbidden' }
    & Pick<Forbidden, 'message'>
  ) | (
    { __typename: 'ServerError' }
    & Pick<ServerError, 'message'>
  ) }
);

export type LoginByEmailMutationVariables = Exact<{
  email: Scalars['String'];
}>;


export type LoginByEmailMutation = (
  { __typename?: 'Mutation' }
  & { loginByEmail: (
    { __typename: 'CheckEmail' }
    & Pick<CheckEmail, 'email'>
  ) | (
    { __typename: 'InvalidEmail' }
    & Pick<InvalidEmail, 'message'>
  ) | (
    { __typename: 'ServerError' }
    & Pick<ServerError, 'message'>
  ) }
);

export type ProfileQueryVariables = Exact<{ [key: string]: never; }>;


export type ProfileQuery = (
  { __typename?: 'Query' }
  & { profile: (
    { __typename: 'UserProfile' }
    & Pick<UserProfile, 'id' | 'email'>
    & { gravatar?: Maybe<(
      { __typename?: 'Gravatar' }
      & Pick<Gravatar, 'url'>
    )>, defaultWorkspace: (
      { __typename?: 'Workspace' }
      & Pick<Workspace, 'id' | 'name'>
    ) }
  ) | (
    { __typename: 'Forbidden' }
    & Pick<Forbidden, 'message'>
  ) | (
    { __typename: 'ServerError' }
    & Pick<ServerError, 'message'>
  ) }
);

export type ProjectQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type ProjectQuery = (
  { __typename?: 'Query' }
  & { screen: (
    { __typename?: 'ScreenQuery' }
    & { projectScreen: (
      { __typename?: 'ProjectScreen' }
      & { project: (
        { __typename: 'Project' }
        & Pick<Project, 'id' | 'title'>
        & { files: (
          { __typename?: 'ProjectFiles' }
          & { list: (
            { __typename: 'ProjectFilesList' }
            & { items: Array<(
              { __typename?: 'ProjectFile' }
              & Pick<ProjectFile, 'id' | 'name' | 'url' | 'type'>
            )> }
          ) | (
            { __typename: 'Forbidden' }
            & Pick<Forbidden, 'message'>
          ) | (
            { __typename: 'ServerError' }
            & Pick<ServerError, 'message'>
          ), total: (
            { __typename: 'ProjectFilesTotal' }
            & Pick<ProjectFilesTotal, 'total'>
          ) | { __typename: 'Forbidden' } | { __typename: 'ServerError' } }
        ) }
      ) | (
        { __typename: 'NotFound' }
        & Pick<NotFound, 'message'>
      ) | (
        { __typename: 'Forbidden' }
        & Pick<Forbidden, 'message'>
      ) | (
        { __typename: 'ServerError' }
        & Pick<ServerError, 'message'>
      ), menu: (
        { __typename: 'MenuItems' }
        & { items: Array<(
          { __typename?: 'MenuItem' }
          & Pick<MenuItem, 'title' | 'url'>
        )> }
      ) | { __typename: 'ServerError' } }
    ) }
  ) }
);

export type SpecScreenQueryVariables = Exact<{
  projectId: Scalars['Int'];
}>;


export type SpecScreenQuery = (
  { __typename?: 'Query' }
  & { screen: (
    { __typename?: 'ScreenQuery' }
    & { screen: (
      { __typename?: 'SpecScreen' }
      & { project: (
        { __typename: 'Project' }
        & Pick<Project, 'id' | 'title'>
      ) | (
        { __typename: 'NotFound' }
        & Pick<NotFound, 'message'>
      ) | (
        { __typename: 'Forbidden' }
        & Pick<Forbidden, 'message'>
      ) | (
        { __typename: 'ServerError' }
        & Pick<ServerError, 'message'>
      ) }
    ) }
  ) }
);

export type UploadProjectFileMutationVariables = Exact<{
  file: UploadProjectFileInput;
}>;


export type UploadProjectFileMutation = (
  { __typename?: 'Mutation' }
  & { uploadProjectFile: (
    { __typename: 'ProjectFile' }
    & Pick<ProjectFile, 'id' | 'url'>
  ) | { __typename: 'Forbidden' } | { __typename: 'AlreadyExists' } | (
    { __typename: 'ServerError' }
    & Pick<ServerError, 'message'>
  ) }
);

export type WorkspaceQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type WorkspaceQuery = (
  { __typename?: 'Query' }
  & { workspace: (
    { __typename: 'Workspace' }
    & Pick<Workspace, 'id' | 'name'>
    & { users: (
      { __typename: 'WorkspaceUsers' }
      & { items: Array<(
        { __typename?: 'WorkspaceUser' }
        & Pick<WorkspaceUser, 'id' | 'role'>
        & { workspace: (
          { __typename?: 'Id' }
          & Pick<Id, 'id'>
        ), profile: (
          { __typename?: 'WorkspaceUserProfile' }
          & Pick<WorkspaceUserProfile, 'id' | 'email' | 'fullName' | 'abbr'>
          & { gravatar?: Maybe<(
            { __typename?: 'Gravatar' }
            & Pick<Gravatar, 'url'>
          )> }
        ) }
      )> }
    ) | (
      { __typename: 'Forbidden' }
      & Pick<Forbidden, 'message'>
    ) | (
      { __typename: 'ServerError' }
      & Pick<ServerError, 'message'>
    ), projects: (
      { __typename: 'WorkspaceProjects' }
      & { list: (
        { __typename: 'WorkspaceProjectsList' }
        & { items: Array<(
          { __typename?: 'WorkspaceProject' }
          & Pick<WorkspaceProject, 'id' | 'name' | 'period'>
        )> }
      ) | (
        { __typename: 'Forbidden' }
        & Pick<Forbidden, 'message'>
      ) | (
        { __typename: 'ServerError' }
        & Pick<ServerError, 'message'>
      ) }
    ) }
  ) | (
    { __typename: 'NotFound' }
    & Pick<NotFound, 'message'>
  ) | (
    { __typename: 'Forbidden' }
    & Pick<Forbidden, 'message'>
  ) | (
    { __typename: 'ServerError' }
    & Pick<ServerError, 'message'>
  ) }
);


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
        return Apollo.useMutation<ConfirmLoginMutation, ConfirmLoginMutationVariables>(ConfirmLoginDocument, baseOptions);
      }
export type ConfirmLoginMutationHookResult = ReturnType<typeof useConfirmLoginMutation>;
export type ConfirmLoginMutationResult = Apollo.MutationResult<ConfirmLoginMutation>;
export type ConfirmLoginMutationOptions = Apollo.BaseMutationOptions<ConfirmLoginMutation, ConfirmLoginMutationVariables>;
export const CreateProjectDocument = gql`
    mutation createProject($input: CreateProjectInput!) {
  createProject(input: $input) {
    __typename
    ... on ServerError {
      message
    }
    ... on Forbidden {
      message
    }
    ... on Project {
      id
      title
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
        return Apollo.useMutation<CreateProjectMutation, CreateProjectMutationVariables>(CreateProjectDocument, baseOptions);
      }
export type CreateProjectMutationHookResult = ReturnType<typeof useCreateProjectMutation>;
export type CreateProjectMutationResult = Apollo.MutationResult<CreateProjectMutation>;
export type CreateProjectMutationOptions = Apollo.BaseMutationOptions<CreateProjectMutation, CreateProjectMutationVariables>;
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
        return Apollo.useMutation<LoginByEmailMutation, LoginByEmailMutationVariables>(LoginByEmailDocument, baseOptions);
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
        return Apollo.useQuery<ProfileQuery, ProfileQueryVariables>(ProfileDocument, baseOptions);
      }
export function useProfileLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ProfileQuery, ProfileQueryVariables>) {
          return Apollo.useLazyQuery<ProfileQuery, ProfileQueryVariables>(ProfileDocument, baseOptions);
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
          files {
            list(filter: {type: [IMAGE]}, limit: 10, offset: 0) {
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
    `;

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
export function useProjectQuery(baseOptions?: Apollo.QueryHookOptions<ProjectQuery, ProjectQueryVariables>) {
        return Apollo.useQuery<ProjectQuery, ProjectQueryVariables>(ProjectDocument, baseOptions);
      }
export function useProjectLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ProjectQuery, ProjectQueryVariables>) {
          return Apollo.useLazyQuery<ProjectQuery, ProjectQueryVariables>(ProjectDocument, baseOptions);
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
export function useSpecScreenQuery(baseOptions?: Apollo.QueryHookOptions<SpecScreenQuery, SpecScreenQueryVariables>) {
        return Apollo.useQuery<SpecScreenQuery, SpecScreenQueryVariables>(SpecScreenDocument, baseOptions);
      }
export function useSpecScreenLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<SpecScreenQuery, SpecScreenQueryVariables>) {
          return Apollo.useLazyQuery<SpecScreenQuery, SpecScreenQueryVariables>(SpecScreenDocument, baseOptions);
        }
export type SpecScreenQueryHookResult = ReturnType<typeof useSpecScreenQuery>;
export type SpecScreenLazyQueryHookResult = ReturnType<typeof useSpecScreenLazyQuery>;
export type SpecScreenQueryResult = Apollo.QueryResult<SpecScreenQuery, SpecScreenQueryVariables>;
export const UploadProjectFileDocument = gql`
    mutation uploadProjectFile($file: UploadProjectFileInput!) {
  uploadProjectFile(file: $file) {
    __typename
    ... on ProjectFile {
      id
      url
    }
    ... on ServerError {
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
 *      file: // value for 'file'
 *   },
 * });
 */
export function useUploadProjectFileMutation(baseOptions?: Apollo.MutationHookOptions<UploadProjectFileMutation, UploadProjectFileMutationVariables>) {
        return Apollo.useMutation<UploadProjectFileMutation, UploadProjectFileMutationVariables>(UploadProjectFileDocument, baseOptions);
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
        list {
          __typename
          ... on WorkspaceProjectsList {
            items {
              id
              name
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
export function useWorkspaceQuery(baseOptions?: Apollo.QueryHookOptions<WorkspaceQuery, WorkspaceQueryVariables>) {
        return Apollo.useQuery<WorkspaceQuery, WorkspaceQueryVariables>(WorkspaceDocument, baseOptions);
      }
export function useWorkspaceLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<WorkspaceQuery, WorkspaceQueryVariables>) {
          return Apollo.useLazyQuery<WorkspaceQuery, WorkspaceQueryVariables>(WorkspaceDocument, baseOptions);
        }
export type WorkspaceQueryHookResult = ReturnType<typeof useWorkspaceQuery>;
export type WorkspaceLazyQueryHookResult = ReturnType<typeof useWorkspaceLazyQuery>;
export type WorkspaceQueryResult = Apollo.QueryResult<WorkspaceQuery, WorkspaceQueryVariables>;