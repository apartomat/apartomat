import { NetworkStatus, useApolloClient } from "@apollo/client"
import { useProjectQuery } from "api/types.d"

import type { Contact, ProjectContactsTotal, Forbidden, ServerError } from "api/types.d"

export function useProject(id: string) {
    const client = useApolloClient();
    const result = useProjectQuery({ client, errorPolicy: "all", variables: { id }, notifyOnNetworkStatusChange: true })

    return {...result, refetching: result.networkStatus === NetworkStatus.refetch }
}

export default useProject

export type {
    ProjectScreenHousesFragment as ProjectHouses,
    ProjectScreenHouseRoomsFragment as HouseRooms,
    ProjectScreenHouseRoomFragment as Room,
    ProjectScreenProjectFragment as Project,
    ProjectScreenHouseFragment as ProjectScreenHouse
} from "api/types.d"

export { ProjectStatus }  from "api/types.d"

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