import { useApolloClient } from "@apollo/client"
import { useAlbumScreenQuery, VisualizationStatus } from "api/types.d"

import type { ProjectVisualizationsListFilter } from "api/types.d"

export type {
    AlbumScreenProjectFragment,
    AlbumScreenVisualizationFragment,
    AlbumScreenAlbumPageCoverFragment,
    AlbumScreenAlbumPageVisualizationFragment,
    AlbumScreenSettingsFragment
} from "api/types.d"

export function useAlbum({
    id
}: {
    id: string
}) {
    return useAlbumScreenQuery({
        client: useApolloClient(),
        errorPolicy: "all",
        variables: {
            id,
            filter: { status: { eq: [ VisualizationStatus.Approved, VisualizationStatus.Unknown ] } }
        }
    })
}

export default useAlbum