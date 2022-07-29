import { FetchResult, useApolloClient } from "@apollo/client"
import { useLoginByEmailMutation, LoginByEmailMutation, LoginByEmailMutationResult } from "api/types.d"

export type LoginByEmailFn = (email: string) => Promise<FetchResult<LoginByEmailMutation>>

export function useLoginByEmail(): [
    LoginByEmailFn,
    LoginByEmailMutationResult
] {
    const client = useApolloClient()
    const [login, result] = useLoginByEmailMutation({ client, errorPolicy: 'all' })

    return [(email: string) => login({ variables: { email } }), result]
}

export default useLoginByEmail