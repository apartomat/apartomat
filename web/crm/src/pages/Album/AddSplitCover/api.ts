import { useMutation } from "@apollo/client"
import { gql } from "@apollo/client"

const ADD_SPLIT_COVER_TO_ALBUM = gql`
    mutation AddSplitCoverToAlbum($albumId: String!, $input: AddSplitCoverToAlbumInput!) {
        addSplitCoverToAlbum(albumId: $albumId, input: $input) {
            ... on SplitCoverAdded {
                cover {
                    title
                    subtitle
                    image {
                        ... on File {
                            id
                            url
                        }
                        ... on Error {
                            message
                        }
                    }
                    qrCodeSrc
                    city
                    year
                    variant
                }
            }
            ... on Forbidden {
                message
            }
            ... on NotFound {
                message
            }
            ... on ServerError {
                message
            }
        }
    }
`

export function useAddSplitCoverToAlbum(albumId: string) {
    const [addSplitCover, { loading, error, data }] = useMutation(ADD_SPLIT_COVER_TO_ALBUM)

    const success = data?.addSplitCoverToAlbum?.__typename === "SplitCoverAdded"

    const callAddSplitCover = async (input: {
        title: string
        subtitle?: string
        imgFileId: string
        withQr: boolean
        city?: string
        year?: number
    }) => {
        return addSplitCover({
            variables: {
                albumId,
                input,
            },
        })
    }

    return [
        callAddSplitCover,
        {
            loading,
            error: error?.message,
            success,
        },
    ] as const
}
