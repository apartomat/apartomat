import { FetchResult, useApolloClient } from "@apollo/client"
import { useConfirmLoginPinMutation, ConfirmLoginPinMutation, ConfirmLoginPinMutationResult } from "api/graphql"

export type { ConfirmLoginPinMutation } from "api/graphql"

export function useConfirmLoginPin(): [
    (token: string, pin: string) => Promise<FetchResult<ConfirmLoginPinMutation>>,
    ConfirmLoginPinMutationResult,
] {
    const client = useApolloClient()
    const [confirmLogin, result] = useConfirmLoginPinMutation({ client, errorPolicy: "all" })

    return [(token: string, pin: string) => confirmLogin({ variables: { token, pin } }), result]
}

export default useConfirmLoginPin
