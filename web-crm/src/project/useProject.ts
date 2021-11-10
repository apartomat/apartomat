import { useApolloClient } from "@apollo/client";
import { useProjectQuery } from "../api/types.d";

import type { ProjectFile, ProjectFilesTotal, Contact, ProjectContactsTotal, Forbidden, ServerError, MenuItem } from "../api/types.d";

export function useProject(id: number) {
    const client = useApolloClient(); 
    return useProjectQuery({client, errorPolicy: "all", variables: { id }});
}

export default useProject;

export type ProjectFiles = (
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
);

export type ProjectContacts = (
  { __typename?: 'ProjectContacts' }
  & { list: (
    { __typename: 'ProjectContactsList' }
    & { items: Array<(
      { __typename?: 'Contact' }
      & Pick<Contact, 'id' | 'fullName'>
    )> }
  ) | (
    { __typename: 'Forbidden' }
    & Pick<Forbidden, 'message'>
  ) | (
    { __typename: 'ServerError' }
    & Pick<ServerError, 'message'>
  ), total: (
    { __typename: 'ProjectContactsTotal' }
    & Pick<ProjectContactsTotal, 'total'>
  ) | { __typename: 'Forbidden' } | { __typename: 'ServerError' } }
);

export type MenuResult = (
    { __typename: 'MenuItems' }
    & { items: Array<(
      { __typename?: 'MenuItem' }
      & Pick<MenuItem, 'title' | 'url'>
    )> }
  ) | { __typename: 'ServerError' };