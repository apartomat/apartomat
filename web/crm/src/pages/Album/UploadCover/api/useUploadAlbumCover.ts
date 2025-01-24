import { ApolloError, FetchResult, useApolloClient } from "@apollo/client"
import { useUploadAlbumCoverMutation, UploadAlbumCoverMutation, UploadAlbumCoverMutationResult } from "api/graphql"

export function useUploadAlbumCover(albumId: string): [
    (file: File) => Promise<FetchResult<UploadAlbumCoverMutation>>,
    // UploadAlbumCoverMutationResult,
    { loading: boolean; called: boolean; success: boolean; error?: string },
] {
    const [upload, result] = useUploadAlbumCoverMutation({
        client: useApolloClient(),
        errorPolicy: "all",
    })

    const { error: apolloError, loading, called, data } = result

    let error = undefined

    switch (data?.uploadAlbumCover.__typename) {
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
        (file) =>
            upload({
                variables: { albumId, file },
            }),
        {
            error,
            loading,
            called,
            success: data?.uploadAlbumCover.__typename === "AlbumCoverUploaded",
        },
    ]
}
