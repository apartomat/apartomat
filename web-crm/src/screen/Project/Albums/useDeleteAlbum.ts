import { FetchResult, useApolloClient } from "@apollo/client"

import { useDeleteAlbumMutation, DeleteAlbumMutation, ProjectStatus, DeleteAlbumMutationResult } from "api/types.d"

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
