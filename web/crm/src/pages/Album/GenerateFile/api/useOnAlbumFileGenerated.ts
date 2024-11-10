import { useOnAlbumFileGeneratedSubscription } from "api/graphql"
import { useApolloClient } from "@apollo/client"

export function useOnAlbumFileGenerated({ albumId }: { albumId: string }) {
    return useOnAlbumFileGeneratedSubscription({
        client: useApolloClient(),
        variables: {
            albumId,
        },
        shouldResubscribe: true,
    })
}