import { NetworkStatus, useApolloClient } from "@apollo/client"
import { useProjectScreenQuery } from "api/graphql"

import type { Contact, ProjectContactsTotal, Forbidden, ServerError } from "api/graphql"

export function useProject(id: string) {
    const client = useApolloClient();
    const result = useProjectScreenQuery({ client, errorPolicy: "all", variables: { id }, notifyOnNetworkStatusChange: true })

    return { ...result, refetching: result.networkStatus === NetworkStatus.refetch }
}

export default useProject

export type {
    ProjectScreenHousesFragment as ProjectHouses,
    ProjectScreenHouseRoomsFragment as HouseRooms,
    ProjectScreenHouseRoomFragment as Room,
    ProjectScreenProjectFragment as Project,
    ProjectScreenHouseFragment as ProjectScreenHouse
} from "api/graphql"

export { ProjectStatus }  from "api/graphql"

export type ProjectContacts = (
  { __typename?: 'ProjectContacts' }
  & { list: (
    { __typename: 'ProjectContactsList' }
    & { items: Array<(
      { __typename?: 'Contact' }
      & Pick<Contact, 'id' | 'fullName' | 'photo' | 'details'>
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