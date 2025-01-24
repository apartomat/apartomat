import React, { forwardRef, useEffect, useImperativeHandle, useRef, useState } from "react"
import { Box, BoxExtendedProps, Button, LayerExtendedProps, Text, GridExtendedProps, Grid, Stack, Image } from "grommet"
import { Modal, Header as ModalHeader } from "widgets/modal"
import { Image as ImageIcon, FormClose } from "grommet-icons/icons"
import { useUploadAlbumCover } from "./api"

export function UploadCover({
    albumId,
    onClickClose,
    onAlbumCoverUploaded,
    ...props
}: {
    albumId: string
    onClickClose?: () => void
    onAlbumCoverUploaded?: () => void
} & LayerExtendedProps) {
    const [files, addFiles, removeFile] = useFiles()

    const [upload, { loading, error, success }] = useUploadAlbumCover(albumId)

    useEffect(() => {
        if (success && onAlbumCoverUploaded) {
            onAlbumCoverUploaded()
        }
    }, [success, onAlbumCoverUploaded])

    const inputFile = useRef<HTMLInputElement>()

    const handleClickUpload = async (event: React.FormEvent) => {
        event.preventDefault()

        const [file] = files

        if (file) {
            await upload(file)
        }
    }

    const handleClickBrowse = () => {
        inputFile.current?.click()
    }

    return (
        <Modal header="Добавить обложку" onClickClose={onClickClose} error={error} {...props}>
            <UploadFiles ref={inputFile} onAdd={addFiles}>
                {files.length === 0 && (
                    <Box
                        align="center"
                        justify="center"
                        round="small"
                        background="light-1"
                        direction="column"
                        gap="small"
                        pad="medium"
                        style={{ width: "50vh", aspectRatio: "1.41/1" }}
                        flex={{ shrink: 0 }}
                    >
                        <ImageIcon size="medium" />
                        <Box align="center">
                            <Text size="medium">Для загрузки перетащите файл сюда или выбирите файл</Text>
                        </Box>
                    </Box>
                )}

                <Box>{files.length > 0 && <Files files={files} onClickRemove={removeFile} />}</Box>
            </UploadFiles>

            <Box direction="row" justify="between" margin={{ top: "medium" }}>
                <Button
                    onClick={handleClickUpload}
                    primary
                    busy={loading}
                    success={success}
                    label={"Загрузить обложку"}
                    disabled={files.length === 0}
                />
                <Button label="Обзор" onClick={handleClickBrowse} />
            </Box>
        </Modal>
    )
}

const UploadFiles = forwardRef(
    (
        {
            onAdd,
            children,
            ...boxProps
        }: {
            onAdd?: (files: File[]) => void
            children: React.ReactNode
        } & BoxExtendedProps,
        ref: React.Ref<unknown>
    ) => {
        const fileInput = useRef<HTMLInputElement>(null)

        useImperativeHandle(ref, () => ({
            click: () => {
                fileInput.current?.click()
            },
        }))

        const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
            event.preventDefault()

            if (event.dataTransfer.files) {
                const nextFiles: File[] = []

                for (const h in event.dataTransfer.files) {
                    const item = event.dataTransfer.files[h]

                    if (item instanceof File) {
                        nextFiles.push(item)
                    }
                }

                if (nextFiles.length > 0) {
                    onAdd && onAdd(nextFiles)
                    setDragCounter(0)
                }
            }
        }

        const handleSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
            event.preventDefault()

            const nextFiles: File[] = []

            if (event.target.files) {
                for (const h in event.target.files) {
                    const item = event.target?.files[h]

                    if (item instanceof File) {
                        nextFiles.push(item)
                    }
                }
            }

            event.target.files = new DataTransfer().files

            if (nextFiles.length > 0) {
                onAdd && onAdd(nextFiles)
            }
        }

        const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
            event.preventDefault()
        }

        const [dragCounter, setDragCounter] = useState(0)

        return (
            <Box
                onDragEnter={() => {
                    setDragCounter(dragCounter + 1)
                }}
                onDragLeave={() => {
                    setDragCounter(dragCounter - 1)
                }}
            >
                <Box
                    onDrop={handleDrop}
                    onDragOver={handleDragOver}
                    border={{ color: dragCounter ? "brand" : "background", style: "dashed", size: "small" }}
                    round="xsmall"
                >
                    <Box {...boxProps}>
                        {React.Children.map(children, (child) => {
                            return child
                        })}
                    </Box>

                    <input type="file" hidden ref={fileInput} onChange={handleSelect} />
                </Box>
            </Box>
        )
    }
)

UploadFiles.displayName = "UploadFiles"

function Files({
    files,
    onClickRemove,
    ...gridProps
}: {
    files: File[]
    onClickRemove?: (i: number) => void
} & GridExtendedProps) {
    return (
        <Grid {...gridProps} gap="xsmall">
            {files?.map((file, index) => {
                return (
                    <FileForUpload
                        key={index}
                        file={file}
                        onClickRemove={() => onClickRemove && onClickRemove(index)}
                    />
                )
            })}
        </Grid>
    )
}

function FileForUpload({
    file,
    onClickRemove,
}: {
    file: File
    onClickRemove?: () => void
} & BoxExtendedProps) {
    const [hover, setHover] = useState(false)

    return (
        <Box flex={{ shrink: 0 }} onMouseEnter={() => setHover(true)} onMouseLeave={() => setHover(false)}>
            <Stack anchor="top-right">
                <Box style={{ width: "50vh", aspectRatio: "1.41/1" }} flex={{ shrink: 0 }} background="light-2">
                    <Image fit="contain" src={URL.createObjectURL(file)} />
                </Box>
                {hover && (
                    <Button
                        plain
                        icon={<FormClose color="light-1" />}
                        onClick={() => onClickRemove && onClickRemove()}
                    />
                )}
            </Stack>
        </Box>
    )
}

function useFiles() {
    const [files, setFiles] = useState<File[]>([])

    const removeFile = (i: number) => {
        const nextFiles: File[] = []

        files.forEach((file, index) => {
            if (i !== index) {
                nextFiles.push(file)
            }
        })

        setFiles(nextFiles)
    }

    const addFiles = (list: File[]) => {
        const [file] = list

        if (file) {
            setFiles([file])
        }
    }

    return [files, addFiles, removeFile]
}
