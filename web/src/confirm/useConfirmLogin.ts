import { FetchResult, useApolloClient } from "@apollo/client";
import { useConfirmLoginMutation, ConfirmLoginMutation, ConfirmLoginMutationResult } from "../api/types.d";

export function useConfirmLogin(): [
    (email: string) => Promise<FetchResult<ConfirmLoginMutation>>,
    ConfirmLoginMutationResult
] {
    const client = useApolloClient();
    const [confirmLogin, result ] = useConfirmLoginMutation({ client, errorPolicy: 'all' });

    return [(token: string) => confirmLogin({ variables: { token } }), result];
}

export default useConfirmLogin;