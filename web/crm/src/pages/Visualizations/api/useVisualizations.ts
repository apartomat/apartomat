import { NetworkStatus, QueryResult, useApolloClient } from "@apollo/client"
import {
    ProjectVisualizationsListFilter,
    useVisualizationsScreenQuery,
    VisualizationsScreenQuery,
    VisualizationsScreenQueryVariables,
} from "api/graphql"

export { VisualizationStatus } from "api/graphql"
export type { VisualizationsScreenHouseRoomFragment } from "api/graphql"

export function useVisualizations(
    id: string,
    filter: ProjectVisualizationsListFilter
): QueryResult<VisualizationsScreenQuery, VisualizationsScreenQueryVariables> & { refetching: boolean } {
    const client = useApolloClient()

    const result = useVisualizationsScreenQuery({
        client,
        errorPolicy: "all",
        variables: { id, filter },
        notifyOnNetworkStatusChange: true,
    })

    return { ...result, refetching: result.networkStatus === NetworkStatus.refetch }
}
