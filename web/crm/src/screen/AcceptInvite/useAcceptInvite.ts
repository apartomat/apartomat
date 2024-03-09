import { FetchResult, useApolloClient } from "@apollo/client"
import { useAcceptInviteMutation, AcceptInviteMutation, AcceptInviteMutationResult } from "api/graphql"

export function useAcceptInvite(): [
    (email: string) => Promise<FetchResult<AcceptInviteMutation>>,
    AcceptInviteMutationResult,
] {
    const client = useApolloClient()

    const [accept, result] = useAcceptInviteMutation({ client, errorPolicy: "all" })

    return [(token: string) => accept({ variables: { token } }), result]
}

export default useAcceptInvite
