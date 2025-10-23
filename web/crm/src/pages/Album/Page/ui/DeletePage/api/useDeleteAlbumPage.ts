import { FetchResult, useApolloClient } from "@apollo/client"
import { useDeleteAlbumPageMutation, DeleteAlbumPageMutation, DeleteAlbumPageMutationResult } from "api/graphql"

export function useDeleteAlbumPage(
    albumId: string
): [
    (pageNumber: number) => Promise<FetchResult<DeleteAlbumPageMutation>>,
    { data?: DeleteAlbumPageMutation | null; loading: boolean; called: boolean; success: boolean; error?: string },
] {
    const client = useApolloClient()

    const [deletePage, result] = useDeleteAlbumPageMutation({ client, errorPolicy: "all" })

    const { error: apolloError, loading, called, data } = result

    let error = undefined

    switch (data?.deleteAlbumPage.__typename) {
        case "Forbidden":
            error = "Доступ запрещен"
            break
        case "ServerError":
            error = "Ошибка сервера"
            break
    }

    if (apolloError) {
        error = `Ошибка: ${apolloError.message}`
    }

    return [
        (pageNumber: number) => deletePage({ variables: { albumId, pageNumber } }),
        {
            data,
            loading,
            called,
            success: data?.deleteAlbumPage.__typename === "AlbumPageDeleted",
            error,
        },
    ]
}
