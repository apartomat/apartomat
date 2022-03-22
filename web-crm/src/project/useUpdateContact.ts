import { FetchResult, useApolloClient } from "@apollo/client"
import { useUpdateContactMutation, UpdateContactMutation, UpdateContactInput, UpdateContactMutationResult, Contact as ContactType } from "../api/types.d"

export type UpdateContact = (contactId: string, data: UpdateContactInput) => Promise<FetchResult<UpdateContactMutation>>


export function useUpdateContact(): [
    UpdateContact,
    UpdateContactMutationResult
] {
    const client = useApolloClient()
    const [update, result] = useUpdateContactMutation({ client, errorPolicy: 'all' })

    return [
        (contactId: string, data: UpdateContactInput) => update({ variables: { contactId, data } }),
        result,
    ]
}
