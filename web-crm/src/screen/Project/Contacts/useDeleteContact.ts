import { FetchResult, useApolloClient } from "@apollo/client"
import { useDeleteContactMutation, DeleteContactMutation, DeleteContactMutationResult, Contact as ContactType } from "api/types.d"

export type DeleteContactFn = (id: string) => Promise<FetchResult<DeleteContactMutation>>

export function useDeleteContact(): [
    DeleteContactFn,
    DeleteContactMutationResult
] {
    const client = useApolloClient()
    const [ deleteContact, result ] = useDeleteContactMutation({ client, errorPolicy: 'all' })

    return [
        (id: string) => deleteContact({ variables: { id } }),
        result,
    ]
}

export default useDeleteContact

export { ContactType } from "api/types.d" // @todo conflicts with Contact alias

export type Contact = Pick<ContactType, "id" | "fullName" | "photo" | "details">