import { useApolloClient } from "@apollo/client";
import { useWorkspaceQuery,  } from "../api/types.d";

export type { WorkspaceUsersResult, WorkspaceProjectsListResult } from "../api/types.d";

export function useWorkspace(id: number) {
    const client = useApolloClient(); 
    return useWorkspaceQuery({client, errorPolicy: "all", variables: { id }});
}

export default useWorkspace;