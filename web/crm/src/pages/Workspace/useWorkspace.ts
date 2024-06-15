import { NetworkStatus, useApolloClient } from "@apollo/client"
import { useWorkspaceScreenQuery } from "api/graphql"

export type {
    WorkspaceScreenFragment,
    WorkspaceScreenUsersFragment,
    WorkspaceScreenArchiveProjectsFragment,
    WorkspaceScreenCurrentProjectsFragment,
    WorkspaceScreenProjectFragment,
} from "api/graphql"

export function useWorkspace({ id, timezone }: { id: string; timezone: string }) {
    const client = useApolloClient()
    const result = useWorkspaceScreenQuery({ client, errorPolicy: "all", variables: { id, timezone } })

    return { ...result, refetching: result.networkStatus === NetworkStatus.refetch }
}

export default useWorkspace
