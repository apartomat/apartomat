import {useEffect, useState} from "react"
import {DocumentPdf} from "grommet-icons"
import {Button, ButtonExtendedProps} from "grommet"

import {AlbumScreenAlbumFragment} from "api/graphql";

import { useOnAlbumFileGenerated } from "./useOnAlbumFileGenerated"
import { useGenerateAlbumFile } from "./useGenerateAlbumFile"

export function GenerateFile({
    album,
    onAlbumFileGenerated,
}: {
    album: AlbumScreenAlbumFragment,
    onAlbumFileGenerated?: (version: number) => void,
}) {
    const [ version, setVersion ] = useState<number>(album.version)

    const [ fileVersion, setFileVersion ] = useState<number | undefined>()

    const [ generateAlbumFile, { data: generateAlbumFileResultData, loading }] = useGenerateAlbumFile(album.id)

    useEffect(() => {
        if (album.file?.__typename === "AlbumFile") {
            setVersion(album.version)
            setFileVersion(album.file.version)
        }
    }, [ album ])

    useEffect(() => {
        if (generateAlbumFileResultData?.generateAlbumFile.__typename === "AlbumFileGenerationStarted") {
            setWaitingFile(true)
        }
    }, [ generateAlbumFileResultData ]);

    const [ waitingFile, setWaitingFile ] = useState(false)

    if (waitingFile) {
        return (
            <ButtonInProgress
                primary
                color="brand"
                label="Сгенерировать"
                icon={<DocumentPdf/>}
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
                icon={<DocumentPdf/>}
                download={"test"}
                target="_blank"
                href={getFileUrl(album)}
                as="a"
            />
        )
    }

    return (
        <Button
            primary
            color="brand"
            label="Сгенерировать"
            icon={<DocumentPdf/>}
            onClick={() => generateAlbumFile()}
        />
    )
}

export default GenerateFile

function ButtonInProgress({
    albumId,
    onAlbumFileGenerated,
    ...buttonProps
}: {
    albumId: string,
    onAlbumFileGenerated?: (version: number) => void
} & ButtonExtendedProps ) {
    const { data, loading} = useOnAlbumFileGenerated({ albumId})

    const [ success, setSuccess ] = useState(false)

    useEffect(() => {
        if (data?.albumFileGenerated.__typename === "AlbumFile") {
            setSuccess(true)
            setTimeout(() => {
                if (data?.albumFileGenerated.__typename === "AlbumFile") {
                    onAlbumFileGenerated && onAlbumFileGenerated(data.albumFileGenerated.version)
                }
            }, 1000)
        }
    }, [ data ])

    return (
        <Button {...buttonProps} busy={loading} success={success}/>
    )
}


function getFileUrl(album: AlbumScreenAlbumFragment): string {
    console.log(album)
    if (album?.file?.__typename === "AlbumFile" && album.file.file) {
        return album.file.file.url
    }

    return ""
}
