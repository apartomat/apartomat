import { FetchResult, useApolloClient } from "@apollo/client"

import { useDeleteAlbumMutation, DeleteAlbumMutation, DeleteAlbumMutationResult } from "api/graphql"

export function useDeleteAlbum(): [
    (id: string) => Promise<FetchResult<DeleteAlbumMutation>>,
    DeleteAlbumMutationResult
] {
    const client = useApolloClient()

    const [ change, result ] = useDeleteAlbumMutation({ client, errorPolicy: 'all' })

    return [
        (id: string) => change({ variables: { id } }),
        result,
    ]
}
