import { FetchResult, useApolloClient } from "@apollo/client"
import { useAddVisualizationsToAlbumMutation, AddVisualizationsToAlbumMutation, AddVisualizationsToAlbumMutationResult } from "api/graphql"

export function useAddVisualizationsToAlbum(albumId: string): [
    (visualizations: string[]) => Promise<FetchResult<AddVisualizationsToAlbumMutation>>,
    AddVisualizationsToAlbumMutationResult
] {
    const [ add, result ] = useAddVisualizationsToAlbumMutation({ client: useApolloClient(), errorPolicy: "all" })

    return [
        (visualizations: string[]) => add({ variables: { albumId, visualizations } }),
        result,
    ]
}
