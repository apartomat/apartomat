import { FetchResult, useApolloClient } from "@apollo/client"
import { useInviteUserMutation, InviteUserMutation, InviteUserMutationResult, WorkspaceUserRole } from "api/graphql"

export function useInviteUser(workspaceId: string): [
    (email: string, role: WorkspaceUserRole) => Promise<FetchResult<InviteUserMutation>>,
    InviteUserMutationResult
] {
    const client = useApolloClient()

    const [ invite, result ] = useInviteUserMutation({ client, errorPolicy: "all" })

    return [
        (email: string, role: WorkspaceUserRole) => invite({ variables: { workspaceId, email, role } }),
        result,
    ]
}

export default useInviteUser
