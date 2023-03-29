import { FetchResult, useApolloClient } from "@apollo/client"
import { useDeleteRoomMutation, DeleteRoomMutation, DeleteRoomMutationResult } from "api/graphql"

export type DeleteRoom = (id: string) => Promise<FetchResult<DeleteRoomMutation>>

export function useDeleteRoom(): [
    DeleteRoom,
    DeleteRoomMutationResult
] {
    const client = useApolloClient()
    const [ deleteRoom, result ] = useDeleteRoomMutation({ client, errorPolicy: 'all' })

    return [
        (id: string) => deleteRoom({ variables: { id } }),
        result,
    ]
}

export default useDeleteRoom