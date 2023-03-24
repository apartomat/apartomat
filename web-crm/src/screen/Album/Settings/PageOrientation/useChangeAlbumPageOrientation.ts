import { FetchResult, useApolloClient } from "@apollo/client"
import { ChangeAlbumPageOrientationMutation, useChangeAlbumPageOrientationMutation, PageSize, ChangeAlbumPageOrientationMutationResult, PageOrientation } from "api/types.d"

export { PageOrientation } from "api/types.d"

export function useChangeAlbumPageOrientation(albumId: string): [
    (pageSize: PageOrientation) => Promise<FetchResult<ChangeAlbumPageOrientationMutation>>,
    ChangeAlbumPageOrientationMutationResult
] {
    const [ change, result ] = useChangeAlbumPageOrientationMutation({ client: useApolloClient(), errorPolicy: "all" })

    return [
        (orientation: PageOrientation) => change({ variables: { albumId, orientation } }),
        result,
    ]
}
