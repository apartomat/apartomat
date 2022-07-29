import { FetchResult, useApolloClient } from "@apollo/client"
import { useConfirmLoginPinMutation, ConfirmLoginPinMutation, ConfirmLoginPinMutationResult } from "api/types.d"

export type { ConfirmLoginPinMutation } from "api/types.d"

export function useConfirmLoginPin(): [
    (token: string, pin: string) => Promise<FetchResult<ConfirmLoginPinMutation>>,
    ConfirmLoginPinMutationResult
] {
    const client = useApolloClient()
    const [ confirmLogin, result ] = useConfirmLoginPinMutation({ client, errorPolicy: 'all' })

    return [(token: string, pin: string) => confirmLogin({ variables: { token, pin } }), result];
}

export default useConfirmLoginPin;