import { FetchResult, useApolloClient } from "@apollo/client"
import { useAddContactMutation, AddContactMutation, AddContactInput, AddContactMutationResult, Contact as ContactType } from "api/graphql"

export type AddContact = (projectId: string, contact: AddContactInput) => Promise<FetchResult<AddContactMutation>>

export function useAddContact(): [
    AddContact,
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

export { ContactType } from "api/graphql"  // @todo conflicts with Contact alias

export type ProjectContact = Pick<ContactType, "id" | "fullName" | "photo" | "details">