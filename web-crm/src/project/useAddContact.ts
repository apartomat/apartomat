import { FetchResult, useApolloClient } from "@apollo/client"
import { useAddContactMutation, AddContactMutation, AddContactInput, AddContactMutationResult } from "../api/types.d"

export type AddContactFn = (email: string, contact: AddContactInput) => Promise<FetchResult<AddContactMutation>>

export function useAddContact(): [
    AddContactFn,
    AddContactMutationResult
] {
    const client = useApolloClient()
    const [add, result] = useAddContactMutation({ client, errorPolicy: 'all' })

    return [
        (projectId: string, contact: AddContactInput) => add({ variables: { projectId, contact } }),
        result,
    ]
}

export default useAddContact

export { ContactType } from "../api/types.d"