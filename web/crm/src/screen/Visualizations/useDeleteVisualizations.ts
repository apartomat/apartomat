import { FetchResult, useApolloClient } from "@apollo/client"
import {
    useDeleteVisualizationsMutation,
    DeleteVisualizationsMutationResult,
    DeleteVisualizationsMutation,
} from "api/graphql"

export type DeleteContactFn = (id: string[]) => Promise<FetchResult<DeleteVisualizationsMutation>>

export function useDeleteVisualizations(): [DeleteContactFn, DeleteVisualizationsMutationResult] {
    const client = useApolloClient()
    const [deleteVisualizations, result] = useDeleteVisualizationsMutation({ client, errorPolicy: "all" })

    return [(id: string[]) => deleteVisualizations({ variables: { id } }), result]
}

export default useDeleteVisualizations
