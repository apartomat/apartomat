import { useApolloClient } from "@apollo/client"
import { useAlbumScreenQuery, VisualizationStatus } from "api/graphql"
import { useEffect, useState } from "react"

export type {
    AlbumScreenProjectFragment,
    AlbumScreenVisualizationFragment,
    AlbumScreenAlbumPageCoverFragment,
    AlbumScreenAlbumPageVisualizationFragment,
    AlbumScreenSettingsFragment,
    AlbumScreenHouseRoomFragment,
} from "api/graphql"

import {
    AlbumScreenAlbumFragment,
    AlbumScreenProjectFragment,
    AlbumScreenAlbumPageCoverFragment,
    AlbumScreenAlbumPageVisualizationFragment,
} from "api/graphql"

export { PageSize, PageOrientation } from "api/graphql"

export function useAlbum({ id }: { id: string }) {
    const { data, loading, refetch } = useAlbumScreenQuery({
        client: useApolloClient(),
        errorPolicy: "all",
        variables: {
            id,
            filter: { status: { eq: [VisualizationStatus.Approved, VisualizationStatus.Unknown] } },
        },
    })

    const [album, setAlbum] = useState<AlbumScreenAlbumFragment | undefined>()

    const [project, setProject] = useState<AlbumScreenProjectFragment | undefined>()

    const [pages, setPages] = useState<
        (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]
    >([])

    useEffect(() => {
        switch (data?.album.__typename) {
            case "Album": {
                setAlbum(data?.album)

                if (data?.album?.project?.__typename === "Project") {
                    setProject(data.album.project)
                }

                if (data.album.pages.__typename === "AlbumPages") {
                    setPages(data.album.pages.items)
                }
                break
            }
        }
    }, [data])

    return { data, loading, refetch, extracted: { album, project, pages } }
}

export default useAlbum
