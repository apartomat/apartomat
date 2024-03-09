import { FetchResult, useApolloClient } from "@apollo/client"
import {
    useMoveRoomToPositionMutation,
    MoveRoomToPositionMutation,
    MoveRoomToPositionMutationResult,
} from "api/graphql"

export function useMoveRoomToPosition(): [
    (roomId: string, position: number) => Promise<FetchResult<MoveRoomToPositionMutation>>,
    MoveRoomToPositionMutationResult,
] {
    const client = useApolloClient()
    const [moveRoom, result] = useMoveRoomToPositionMutation({ client, errorPolicy: "all" })

    return [(roomId: string, position: number) => moveRoom({ variables: { roomId, position } }), result]
}

export default useMoveRoomToPosition
