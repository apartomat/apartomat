import { useApolloClient } from "@apollo/client"
import { useAlbumScreenQuery, VisualizationStatus } from "api/graphql"

export type {
    AlbumScreenProjectFragment,
    AlbumScreenVisualizationFragment,
    AlbumScreenAlbumPageCoverFragment,
    AlbumScreenAlbumPageVisualizationFragment,
    AlbumScreenSettingsFragment,
    AlbumScreenHouseRoomFragment,
} from "api/graphql"

export { PageSize, PageOrientation } from "api/graphql"

export function useAlbum({ id }: { id: string }) {
    return useAlbumScreenQuery({
        client: useApolloClient(),
        errorPolicy: "all",
        variables: {
            id,
            filter: { status: { eq: [VisualizationStatus.Approved, VisualizationStatus.Unknown] } },
        },
    })
}

export default useAlbum
