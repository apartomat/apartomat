import { useApolloClient } from "@apollo/client"
import { useWorkspaceScreenQuery } from "api/types.d"

export type {
    WorkspaceScreenFragment,
    WorkspaceScreenUsersFragment,
    WorkspaceScreenArchiveProjectsFragment,
    WorkspaceScreenCurrentProjectsFragment,
    WorkspaceScreenProjectFragment,
} from "api/types.d"

export function useWorkspace({ id, timezone }: { id: string, timezone: string }) {
    const client = useApolloClient();
    return useWorkspaceScreenQuery({ client, errorPolicy: "all", variables: { id, timezone }})
}

export default useWorkspace