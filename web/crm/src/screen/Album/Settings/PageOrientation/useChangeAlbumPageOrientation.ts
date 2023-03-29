import { FetchResult, useApolloClient } from "@apollo/client"
import { ChangeAlbumPageOrientationMutation, useChangeAlbumPageOrientationMutation, PageSize, ChangeAlbumPageOrientationMutationResult, PageOrientation } from "api/graphql"

export { PageOrientation } from "api/graphql"

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
