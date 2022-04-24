import { useApolloClient } from "@apollo/client"
import { useWorkspaceQuery } from "../api/types.d"

export type { WorkspaceUsersResult, WorkspaceProjectsListResult, WorkspaceProject } from "../api/types.d"

export function useWorkspace({ id, timezone }: { id: string, timezone: string }) {
    const client = useApolloClient(); 
    return useWorkspaceQuery({client, errorPolicy: "all", variables: { id, timezone }})
}

export default useWorkspace