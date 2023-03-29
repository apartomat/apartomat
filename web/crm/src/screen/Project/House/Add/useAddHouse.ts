import { FetchResult, useApolloClient } from "@apollo/client"
import {
    useAddHouseMutation,
    AddHouseMutation,
    AddHouseInput,
    AddHouseMutationResult,
    AddHouseFragment
} from "api/graphql"

export type AddHouse = (projectId: string, contact: AddHouseInput) => Promise<FetchResult<AddHouseMutation>>

export function useAddHouse(): [
    AddHouse,
    AddHouseMutationResult
] {
    const client = useApolloClient()
    const [add, result] = useAddHouseMutation({ client, errorPolicy: 'all' })

    return [
        (projectId: string, house: AddHouseInput) => add({ variables: { projectId, house } }),
        result,
    ]
}

export default useAddHouse

export type House = AddHouseFragment
