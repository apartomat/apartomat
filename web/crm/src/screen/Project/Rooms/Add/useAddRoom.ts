import { FetchResult, useApolloClient } from "@apollo/client"
import { useAddRoomMutation, AddRoomMutation, AddRoomInput, AddRoomMutationResult } from "api/graphql"

export type AddRoom = (houseId: string, room: AddRoomInput) => Promise<FetchResult<AddRoomMutation>>

export function useAddRoom(): [
    AddRoom,
    AddRoomMutationResult
] {
    const client = useApolloClient()
    const [add, result] = useAddRoomMutation({ client, errorPolicy: 'all' })

    return [
        (houseId: string, room: AddRoomInput) => add({ variables: { houseId, room } }),
        result,
    ]
}

export default useAddRoom

export type { ProjectScreenHouseRoomFragment } from "api/graphql"