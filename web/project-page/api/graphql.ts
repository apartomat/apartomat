import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
  Url: { input: any; output: any; }
};

export type Album = {
  __typename?: 'Album';
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  url: Scalars['String']['output'];
};

export type Error = {
  message: Scalars['String']['output'];
};

export type Forbidden = Error & {
  __typename?: 'Forbidden';
  message: Scalars['String']['output'];
};

export type House = {
  __typename?: 'House';
  address: Scalars['String']['output'];
  city: Scalars['String']['output'];
  housingComplex: Scalars['String']['output'];
  id: Scalars['String']['output'];
};

export type NotFound = Error & {
  __typename?: 'NotFound';
  message: Scalars['String']['output'];
};

export type ProjectPage = {
  __typename?: 'ProjectPage';
  album: ProjectPageAlbumResult;
  description: Scalars['String']['output'];
  house: ProjectPageHouseResult;
  id: Scalars['String']['output'];
  title: Scalars['String']['output'];
  visualizations: ProjectPageVisualizations;
};

export type ProjectPageAlbumResult = Album | Forbidden | NotFound | ServerError;

export type ProjectPageHouseResult = Forbidden | House | NotFound | ServerError;

export type ProjectPageResult = Forbidden | NotFound | ProjectPage | ServerError;

export type ProjectPageVisualizations = {
  __typename?: 'ProjectPageVisualizations';
  list: ProjectPageVisualizationsListResult;
  total: ProjectPageVisualizationsTotalResult;
};


export type ProjectPageVisualizationsListArgs = {
  filter?: ProjectPageVisualizationsFilter;
  limit?: Scalars['Int']['input'];
  offset?: Scalars['Int']['input'];
};


export type ProjectPageVisualizationsTotalArgs = {
  filter?: ProjectPageVisualizationsFilter;
};

export type ProjectPageVisualizationsFilter = {
  roomId?: InputMaybe<StringFilter>;
};

export type ProjectPageVisualizationsListResult = Forbidden | ServerError | VisualizationsList;

export type ProjectPageVisualizationsTotalResult = Forbidden | ServerError | VisualizationsTotal;

export type Query = {
  __typename?: 'Query';
  projectPage: ProjectPageResult;
  version: Scalars['String']['output'];
};


export type QueryProjectPageArgs = {
  id: Scalars['String']['input'];
};

export type Room = {
  __typename?: 'Room';
  id: Scalars['String']['output'];
  level?: Maybe<Scalars['Int']['output']>;
  name: Scalars['String']['output'];
  square?: Maybe<Scalars['Float']['output']>;
};

export type ServerError = Error & {
  __typename?: 'ServerError';
  message: Scalars['String']['output'];
};

export type StringFilter = {
  eq?: InputMaybe<Array<Scalars['String']['input']>>;
};

export type Unknown = Error & {
  __typename?: 'Unknown';
  message: Scalars['String']['output'];
};

export type Visualization = {
  __typename?: 'Visualization';
  description: Scalars['String']['output'];
  file: VisualizationFileResult;
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  room: VisualizationRoomResult;
};

export type VisualizationFile = {
  __typename?: 'VisualizationFile';
  id: Scalars['String']['output'];
  mimeType: Scalars['String']['output'];
  url: Scalars['Url']['output'];
};

export type VisualizationFileResult = Forbidden | NotFound | ServerError | VisualizationFile;

export type VisualizationRoomResult = Forbidden | NotFound | Room | ServerError;

export type VisualizationsList = {
  __typename?: 'VisualizationsList';
  items: Array<Visualization>;
};

export type VisualizationsTotal = {
  __typename?: 'VisualizationsTotal';
  total: Scalars['Int']['output'];
};

export type ProjectPageScreenQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type ProjectPageScreenQuery = { __typename?: 'Query', projectPage: { __typename: 'Forbidden', message: string } | { __typename: 'NotFound', message: string } | { __typename: 'ProjectPage', id: string, title: string, description: string, house: { __typename?: 'Forbidden', message: string } | { __typename?: 'House', city: string, address: string, housingComplex: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string }, visualizations: { __typename?: 'ProjectPageVisualizations', list: { __typename: 'Forbidden' } | { __typename: 'ServerError' } | { __typename: 'VisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, file: { __typename?: 'Forbidden', message: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'VisualizationFile', url: any } }> }, total: { __typename?: 'Forbidden' } | { __typename?: 'ServerError' } | { __typename?: 'VisualizationsTotal', total: number } }, album: { __typename?: 'Album', id: string, name: string, url: string } | { __typename?: 'Forbidden', message: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } } | { __typename: 'ServerError', message: string } };

export type ProjectPageScreenProjectFragment = { __typename?: 'ProjectPage', id: string, title: string, description: string, house: { __typename?: 'Forbidden', message: string } | { __typename?: 'House', city: string, address: string, housingComplex: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string }, visualizations: { __typename?: 'ProjectPageVisualizations', list: { __typename: 'Forbidden' } | { __typename: 'ServerError' } | { __typename: 'VisualizationsList', items: Array<{ __typename?: 'Visualization', id: string, file: { __typename?: 'Forbidden', message: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } | { __typename?: 'VisualizationFile', url: any } }> }, total: { __typename?: 'Forbidden' } | { __typename?: 'ServerError' } | { __typename?: 'VisualizationsTotal', total: number } }, album: { __typename?: 'Album', id: string, name: string, url: string } | { __typename?: 'Forbidden', message: string } | { __typename?: 'NotFound', message: string } | { __typename?: 'ServerError', message: string } };

export type ProjectPageScreenAlbumFragment = { __typename?: 'Album', id: string, name: string, url: string };

export type ProjectPageScreenHouseFragment = { __typename?: 'House', city: string, address: string, housingComplex: string };

export const ProjectPageScreenHouseFragmentDoc = gql`
    fragment ProjectPageScreenHouse on House {
  city
  address
  housingComplex
}
    `;
export const ProjectPageScreenAlbumFragmentDoc = gql`
    fragment ProjectPageScreenAlbum on Album {
  id
  name
  url
}
    `;
export const ProjectPageScreenProjectFragmentDoc = gql`
    fragment ProjectPageScreenProject on ProjectPage {
  id
  title
  description
  house {
    ... on House {
      ...ProjectPageScreenHouse
    }
    ... on Error {
      message
    }
  }
  visualizations {
    list(limit: 100, offset: 0) {
      __typename
      ... on VisualizationsList {
        items {
          id
          file {
            ... on VisualizationFile {
              url
            }
            ... on Error {
              message
            }
          }
        }
      }
    }
    total {
      ... on VisualizationsTotal {
        total
      }
    }
  }
  album {
    ... on Album {
      ...ProjectPageScreenAlbum
    }
    ... on Error {
      message
    }
  }
}
    ${ProjectPageScreenHouseFragmentDoc}
${ProjectPageScreenAlbumFragmentDoc}`;
export const ProjectPageScreenDocument = gql`
    query projectPageScreen($id: String!) {
  projectPage(id: $id) {
    __typename
    ... on ProjectPage {
      ...ProjectPageScreenProject
    }
    ... on Error {
      message
    }
  }
}
    ${ProjectPageScreenProjectFragmentDoc}`;

/**
 * __useProjectPageScreenQuery__
 *
 * To run a query within a React component, call `useProjectPageScreenQuery` and pass it any options that fit your needs.
 * When your component renders, `useProjectPageScreenQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useProjectPageScreenQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useProjectPageScreenQuery(baseOptions: Apollo.QueryHookOptions<ProjectPageScreenQuery, ProjectPageScreenQueryVariables> & ({ variables: ProjectPageScreenQueryVariables; skip?: boolean; } | { skip: boolean; }) ) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ProjectPageScreenQuery, ProjectPageScreenQueryVariables>(ProjectPageScreenDocument, options);
      }
export function useProjectPageScreenLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ProjectPageScreenQuery, ProjectPageScreenQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ProjectPageScreenQuery, ProjectPageScreenQueryVariables>(ProjectPageScreenDocument, options);
        }
export function useProjectPageScreenSuspenseQuery(baseOptions?: Apollo.SkipToken | Apollo.SuspenseQueryHookOptions<ProjectPageScreenQuery, ProjectPageScreenQueryVariables>) {
          const options = baseOptions === Apollo.skipToken ? baseOptions : {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ProjectPageScreenQuery, ProjectPageScreenQueryVariables>(ProjectPageScreenDocument, options);
        }
export type ProjectPageScreenQueryHookResult = ReturnType<typeof useProjectPageScreenQuery>;
export type ProjectPageScreenLazyQueryHookResult = ReturnType<typeof useProjectPageScreenLazyQuery>;
export type ProjectPageScreenSuspenseQueryHookResult = ReturnType<typeof useProjectPageScreenSuspenseQuery>;
export type ProjectPageScreenQueryResult = Apollo.QueryResult<ProjectPageScreenQuery, ProjectPageScreenQueryVariables>;