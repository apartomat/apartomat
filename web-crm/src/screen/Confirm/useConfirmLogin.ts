import { FetchResult, useApolloClient } from "@apollo/client";
import { useConfirmLoginLinkMutation, ConfirmLoginLinkMutation, ConfirmLoginLinkMutationResult } from "api/types.d";

export function useConfirmLogin(): [
    (email: string) => Promise<FetchResult<ConfirmLoginLinkMutation>>,
    ConfirmLoginLinkMutationResult
] {
    const client = useApolloClient();
    const [confirmLogin, result ] = useConfirmLoginLinkMutation({ client, errorPolicy: 'all' });

    return [(token: string) => confirmLogin({ variables: { token } }), result];
}

export default useConfirmLogin;