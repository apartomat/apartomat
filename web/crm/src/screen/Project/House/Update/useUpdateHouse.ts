import { FetchResult, useApolloClient } from "@apollo/client"
import {
    useUpdateHouseMutation,
    UpdateHouseMutation,
    UpdateHouseInput,
    UpdateHouseMutationResult,
    UpdateHouseFragment
} from "api/graphql"

export type UpdateHouse = (projectId: string, contact: UpdateHouseInput) => Promise<FetchResult<UpdateHouseMutation>>

export function useUpdateHouse(): [
    UpdateHouse,
    UpdateHouseMutationResult
] {
    const client = useApolloClient()
    const [update, result] = useUpdateHouseMutation({ client, errorPolicy: 'all' })

    return [
        (houseId: string, data: UpdateHouseInput) => update({ variables: { houseId, house: data } }),
        result,
    ]
}

export default useUpdateHouse

export type House = UpdateHouseFragment
