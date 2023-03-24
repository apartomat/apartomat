import { FetchResult, useApolloClient } from "@apollo/client"
import { ChangeAlbumPageSizeMutation, useChangeAlbumPageSizeMutation, PageSize, ChangeAlbumPageSizeMutationResult } from "api/types.d"

export { PageSize } from "api/types.d"

export function useChangeAlbumPageSize(albumId: string): [
    (pageSize: PageSize) => Promise<FetchResult<ChangeAlbumPageSizeMutation>>,
    ChangeAlbumPageSizeMutationResult
] {
    const [ change, result ] = useChangeAlbumPageSizeMutation({ client: useApolloClient(), errorPolicy: "all" })

    return [
        (size: PageSize) => change({ variables: { albumId, size } }),
        result,
    ]
}
