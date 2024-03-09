import { FetchResult, useApolloClient } from "@apollo/client"
import { useGenerateAlbumFileMutation, GenerateAlbumFileMutation, GenerateAlbumFileMutationResult } from "api/graphql"

export function useGenerateAlbumFile(
    albumId: string
): [() => Promise<FetchResult<GenerateAlbumFileMutation>>, GenerateAlbumFileMutationResult] {
    const client = useApolloClient()

    const [gen, result] = useGenerateAlbumFileMutation({ client, errorPolicy: "all" })

    return [() => gen({ variables: { albumId } }), result]
}

export default useGenerateAlbumFile
