import { FetchResult, useApolloClient } from "@apollo/client"
import { useCreateAlbumMutation, CreateAlbumMutation, CreateAlbumMutationResult } from "api/graphql"

export type CreateAlbum = (projectId: string, name: string) => Promise<FetchResult<CreateAlbumMutation>>

export function useCreatePrintAlbum(): [
    CreateAlbum,
    CreateAlbumMutationResult
] {
    const client = useApolloClient()
    const [ create, result ] = useCreateAlbumMutation({ client, errorPolicy: 'all' })

    return [
        (projectId: string, name: string) => create({ variables: { projectId, name } }),
        result,
    ]
}

export default useCreatePrintAlbum

export type { ProjectScreenAlbumFragment } from "api/graphql"