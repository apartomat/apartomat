import { FetchResult, useApolloClient } from "@apollo/client"
import { useUpdateRoomMutation, UpdateRoomMutation, UpdateRoomInput, UpdateRoomMutationResult } from "api/graphql"

export type UpdateRoom = (contactId: string, data: UpdateRoomInput) => Promise<FetchResult<UpdateRoomMutation>>

export function useUpdateRoom(): [
    UpdateRoom,
    UpdateRoomMutationResult
] {
    const client = useApolloClient()
    const [update, result] = useUpdateRoomMutation({ client, errorPolicy: "all" })

    return [
        (roomId: string, data: UpdateRoomInput) => update({ variables: { roomId, data } }),
        result,
    ]
}

export type { ProjectScreenHouseRoomFragment } from "api/graphql"