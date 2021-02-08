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
};

export type LoginByEmailResult = CheckEmail | InvalidEmail | ServerError;

export type InvalidEmail = Error & {
  __typename?: 'InvalidEmail';
  message: Scalars['String'];
};

export type ConfirmLoginResult = LoginConfirmed | InvalidToken | ExpiredToken | ServerError;

export type WorkspaceUsersResult = WorkspaceUsers | Forbidden | ServerError;

export type InvalidToken = Error & {
  __typename?: 'InvalidToken';
  message: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  version: Scalars['String'];
  profile: UserProfileResult;
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

export type CheckEmail = {
  __typename?: 'CheckEmail';
  email: Scalars['String'];
};

export type UserProfileResult = UserProfile | Forbidden | ServerError;

export type WorkspaceUsers = {
  __typename?: 'WorkspaceUsers';
  items: Array<WorkspaceUser>;
};

export type WorkspaceUser = {
  __typename?: 'WorkspaceUser';
  id: Scalars['Int'];
  role: WorkspaceUserRole;
  profile: WorkspaceUserProfile;
};

/**  Common types  */
export type Error = {
  message: Scalars['String'];
};

export type ServerError = Error & {
  __typename?: 'ServerError';
  message: Scalars['String'];
};

export type Forbidden = Error & {
  __typename?: 'Forbidden';
  message: Scalars['String'];
};

export type Workspace = {
  __typename?: 'Workspace';
  id: Scalars['Int'];
  name: Scalars['String'];
  users: WorkspaceUsersResult;
};

export type Id = {
  __typename?: 'Id';
  id: Scalars['Int'];
};

export type ShoppinglistQuery = {
  __typename?: 'ShoppinglistQuery';
  productOnPage?: Maybe<Product>;
};


export type ShoppinglistQueryProductOnPageArgs = {
  url: Scalars['String'];
};

export enum WorkspaceUserRole {
  Admin = 'ADMIN',
  User = 'USER'
}

export type UserProfile = {
  __typename?: 'UserProfile';
  id: Scalars['Int'];
  email: Scalars['String'];
  gravatar?: Maybe<Gravatar>;
  defaultWorkspace: Workspace;
};

export type ExpiredToken = Error & {
  __typename?: 'ExpiredToken';
  message: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  loginByEmail: LoginByEmailResult;
  confirmLogin: ConfirmLoginResult;
};


export type MutationLoginByEmailArgs = {
  email: Scalars['String'];
  workspaceName?: Scalars['String'];
};


export type MutationConfirmLoginArgs = {
  token: Scalars['String'];
};

export type WorkspaceResult = Workspace | NotFound | Forbidden | ServerError;

export type WorkspaceUserProfile = {
  __typename?: 'WorkspaceUserProfile';
  id: Scalars['Int'];
  email: Scalars['String'];
  gravatar: Gravatar;
};

export type LoginConfirmed = {
  __typename?: 'LoginConfirmed';
  token: Scalars['String'];
};

export type NotFound = Error & {
  __typename?: 'NotFound';
  message: Scalars['String'];
};

export type Product = {
  __typename?: 'Product';
  name: Scalars['String'];
  description: Scalars['String'];
  image: Scalars['String'];
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
        & { profile: (
          { __typename?: 'WorkspaceUserProfile' }
          & Pick<WorkspaceUserProfile, 'id' | 'email'>
          & { gravatar: (
            { __typename?: 'Gravatar' }
            & Pick<Gravatar, 'url'>
          ) }
        ) }
      )> }
    ) | (
      { __typename: 'Forbidden' }
      & Pick<Forbidden, 'message'>
    ) | (
      { __typename: 'ServerError' }
      & Pick<ServerError, 'message'>
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
            profile {
              id
              email
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