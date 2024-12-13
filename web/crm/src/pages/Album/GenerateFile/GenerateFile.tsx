import { useEffect, useState } from "react"
import { filesize } from "filesize"

import { DocumentPdf } from "grommet-icons"
import { Box, Button, ButtonExtendedProps, Text } from "grommet"

import { AlbumFileStatus, AlbumScreenAlbumFragment } from "api/graphql"
import { useGenerateAlbumFile, useOnAlbumFileGenerated } from "./api/"

export function GenerateFile({
    album,
    onAlbumFileGenerated,
}: {
    album: AlbumScreenAlbumFragment
    onAlbumFileGenerated?: (version: number) => void
}) {
    const [version, setVersion] = useState<number>(album.version)

    const [fileVersion, setFileVersion] = useState<number | undefined>()

    const [waitingFile, setWaitingFile] = useState(false)

    const [generateAlbumFile, { data: generateAlbumFileResultData }] = useGenerateAlbumFile(album.id)

    useEffect(() => {
        if (album.file?.__typename === "AlbumFile") {
            switch (album.file.status) {
                case AlbumFileStatus.GeneratingDone:
                    setVersion(album.version)
                    setFileVersion(album.file.version)
                    setWaitingFile(false)
                    break
                default:
                    setWaitingFile(true)
            }
        }
    }, [album])

    useEffect(() => {
        if (generateAlbumFileResultData?.generateAlbumFile.__typename === "AlbumFileGenerationStarted") {
            setWaitingFile(true)
        }
    }, [generateAlbumFileResultData])

    if (waitingFile) {
        return (
            <ButtonInProgress
                primary
                color="brand"
                label="Сгенерировать"
                icon={<DocumentPdf />}
                albumId={album.id}
                onAlbumFileGenerated={(version: number) => {
                    setVersion(version)
                    setFileVersion(version)
                    setWaitingFile(false)
                    onAlbumFileGenerated && onAlbumFileGenerated(version)
                }}
            />
        )
    }

    if (version === fileVersion) {
        return (
            <Button
                primary
                color="brand"
                label="Скачать"
                icon={<DocumentPdf />}
                download={true}
                target="_blank"
                href={getFileUrl(album)}
                as="a"
                tip={{
                    content: (
                        <Box
                            pad={{ vertical: "xxsmall", horizontal: "small" }}
                            background="light-1"
                            round="medium"
                            margin="xsmall"
                            elevation="xsmall"
                            alignSelf="start"
                        >
                            <Text size="small">Файл {getFileSize(album)}</Text>
                        </Box>
                    ),
                    plain: true,
                    dropProps: { align: { top: "bottom" } },
                }}
            />
        )
    }

    return (
        <Button
            primary
            color="brand"
            label="Сгенерировать"
            icon={<DocumentPdf />}
            onClick={() => generateAlbumFile()}
        />
    )
}

function ButtonInProgress({
    albumId,
    onAlbumFileGenerated,
    ...buttonProps
}: {
    albumId: string
    onAlbumFileGenerated?: (version: number) => void
} & ButtonExtendedProps) {
    const { data, loading } = useOnAlbumFileGenerated({ albumId })

    const [success, setSuccess] = useState(false)

    useEffect(() => {
        if (data?.albumFileGenerated.__typename === "AlbumFile") {
            setSuccess(true)
            setTimeout(() => {
                if (data?.albumFileGenerated.__typename === "AlbumFile") {
                    onAlbumFileGenerated && onAlbumFileGenerated(data.albumFileGenerated.version)
                }
            }, 1000)
        }
    }, [data])

    return <Button {...buttonProps} busy={loading} success={success} />
}

function getFileUrl(album: AlbumScreenAlbumFragment): string {
    if (album?.file?.__typename === "AlbumFile" && album.file.file) {
        return album.file.file.url
    }

    return ""
}

function getFileSize(album: AlbumScreenAlbumFragment): number {
    if (album?.file?.__typename === "AlbumFile" && album.file.file?.size) {
        return filesize(album.file.file.size)
    }

    return 0
}
