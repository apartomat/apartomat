import React, { useState, DragEvent, ChangeEvent, useEffect, useCallback } from "react"

import { Layer, Grid, Form, Box, Heading, Button, Image, Text, LayerExtendedProps, BoxExtendedProps, Stack, GridExtendedProps, FormField, Select  } from "grommet"
import { Image as ImageIcon, FormClose } from "grommet-icons"

import { Room, ProjectHouses } from "../useProject"
import { useUploadVisualizations } from "./useUploadVisualizations"
import { useRef } from "react"
import { useImperativeHandle } from "react"
import { forwardRef } from "react"

export default function UploadVisualizations({
    projectId,
    onClickClose,
    onUploadComplete,
    houses,
    ...layerProps
}: {
    projectId: string,
    houses: ProjectHouses,
    onClickClose?: () => void,
    onUploadComplete?: ({ files }: { files: File[] }) => void
} & LayerExtendedProps) {
    const [ files, setFiles ] = useState<File[]>([])

    const [ room, setRoom ] = useState<{ key: string, value: string } | undefined>(undefined)

    const [ errorMessage, setErrorMessage ] = useState<string | undefined>(undefined)

    const [ upload, { loading, error, data } ] = useUploadVisualizations()

    useEffect(() => {
        const complete = data?.uploadVisualizations.__typename === "VisualizationsUploaded"

        if (complete) {
            onUploadComplete && onUploadComplete({ files })
        }
    }, [ data, files, onUploadComplete ])

    useEffect(() => {
        const complete = data?.uploadVisualizations.__typename === "SomeVisualizationsUploaded"

        if (complete) {
            onUploadComplete && onUploadComplete({ files })
        }
    }, [ data, files, onUploadComplete ])

    useEffect(() => {
        const error = data?.uploadVisualizations.__typename === "Forbidden"

        if (error) {
            setErrorMessage("Доступ запрещен")
        }
    }, [ data ])

    useEffect(() => {
        setErrorMessage(error ? "Ошибка сервера" : undefined)
    }, [ error ])

    const [ dragCounter, setDragCounter ] = useState(0)

    const handleSubmit = (event: React.FormEvent) => {
        event.preventDefault()

        if (files && !loading) {
            upload({ projectId, files, roomId: room?.value })
        }
    }

    const addFiles = (list: File[]) => {
        const nextFiles = [...files]

        list.forEach((item) => {
            const exists = nextFiles.filter((file) => item.name === file.name && item.size === file.size).length > 0

            if (!exists) {
                nextFiles.push(item)
            }
        })

        setFiles(nextFiles)
    }

    const removeFile = (i: number) => {
        const nextFiles: File[] = []

        files.forEach((file, index) => {
            if (i !== index) {
                nextFiles.push(file)
            }
        })

        setFiles(nextFiles)
    }

    const inputFile = useRef<HTMLInputElement>()

    const roomscb = useCallback(() => {
        return rooms(houses)
    }, [ houses ])

    return (
        <Layer
            {...layerProps}
        >
            <Box
                width="large"
                onDragEnter={() => {
                    setDragCounter(dragCounter+1)
                }}
                onDragLeave={() => {
                    setDragCounter(dragCounter-1)
                }}
            >

                <UploadFiles
                    ref={inputFile}
                    onAdd={(files) => {
                        addFiles(files)
                        setDragCounter(0)
                    }}
                    pad="medium"
                    gap="medium"
                    border={{color: dragCounter ? "focus" : "background", style: "dashed", size: "medium"}}
                    round="small"
                >                    
                    <Box direction="row" justify="between" align="center">
                        <Heading level={3} margin="none">Загрузить визуализации</Heading>
                        <Button icon={ <FormClose/> } onClick={onClickClose}/>
                    </Box>

                    {errorMessage &&
                        <Box
                            pad="small"
                            round="small"
                            direction="row"
                            gap="small"
                            align="center"
                            background={{ color: "status-critical", opacity: "weak"}}
                        >
                            <Box border={{ color: "status-critical", size: "small"}} round="large">
                                <FormClose color="status-critical" size="medium"/>
                            </Box>
                            <Text weight="bold" size="medium">{errorMessage}</Text>
                        </Box>
                    }

                    <Form>
                        <FormField
                            contentProps={{ border: false }}
                            label="Файлы">
                                {files.length === 0 &&
                                    <Box
                                        align="center"
                                        justify="center"
                                        round="small"
                                        background="light-1"
                                        height="xsmall"
                                        direction="column"
                                        gap="small"
                                    >
                                            <ImageIcon size="medium"/>
                                            <Box><Text size="medium">Для загрузки перетащите файлы сюда или выбирите файл</Text></Box>                                    
                                    </Box>
                                }
                                {files.length > 0 &&
                                    <Files
                                        files={files}
                                        onClickRemove={removeFile}
                                    />
                                }
                            </FormField>

                        {roomscb() &&
                            <FormField
                                label="Комната"
                                width="medium"
                            >
                                <Select
                                    labelKey="label"
                                    valueKey="value"
                                    value={room}
                                    options={roomscb().map(room => {
                                        return { label: room.name, value: room.id }
                                    })}
                                    onChange={({ value }) => {
                                        setRoom(value)
                                    }}
                                />
                            </FormField>
                        }

                        <Box direction="row" justify="between" margin={{ top: "large" }}>
                            <Button
                                onClick={handleSubmit}
                                primary
                                label={loading ? "Загрузка..." : "Загрузить" }
                                disabled={files.length === 0 || loading}
                            />
                            <Button
                                label="Обзор"
                                onClick={() => {
                                    inputFile.current?.click()
                                }}
                            />                            
                        </Box>
                    </Form>
                </UploadFiles>
            </Box>
        </Layer>
    )
}

const UploadFiles = forwardRef(({
    onAdd,
    children,
    ...boxProps
}: {
    onAdd?: (files: File[]) => void
    children: React.ReactNode
} & BoxExtendedProps, ref: React.Ref<unknown>) => {
    const fileInput = useRef<HTMLInputElement>(null)

    useImperativeHandle(ref, () => ({
        click: () => {
            fileInput.current?.click()
        }
    }))

    const handleDrop = (event: DragEvent<HTMLDivElement>) => {
        event.preventDefault()

        if (event.dataTransfer.files) {
            let nextFiles: File[] = []
            
            for (const h in event.dataTransfer.files) {
                const item = event.dataTransfer.files[h]

                if (item instanceof File) {
                    nextFiles.push(item)
                }
            }

            if (nextFiles.length > 0 ) {
                onAdd && onAdd(nextFiles)
            }
        }
    }

    const handleSelect = (event: ChangeEvent<HTMLInputElement>) => {
        event.preventDefault()

        let nextFiles: File[] = []

        if (event.target.files) {
            for (const h in event.target.files) {
                const item = event.target.files[h]
    
                if (item instanceof File) {
                    nextFiles.push(item)
                }
            }
        }

        event.target.files = (new DataTransfer()).files

        if (nextFiles.length > 0) {
            onAdd && onAdd(nextFiles)
        }
    }

    const handleDragOver = (event: DragEvent<HTMLDivElement>) => {
        event.preventDefault()
    }

    return (
        <Box
            onDrop={handleDrop}
            onDragOver={handleDragOver}
        >
            <Box {...boxProps}>
                {React.Children.map(children, ( child ) => {
                    return child
                })}
            </Box>

            <input
                type="file"
                hidden
                multiple
                ref={fileInput}
                onChange={handleSelect}
            />
        </Box>
    )
})

function Files({
    files,
    onClickRemove,
    ...gridProps
}: {
    files: File[],
    onClickRemove?: (i: number) => void
} & GridExtendedProps) {
    return (
        <Grid {...gridProps} columns="xsmall" rows="xsmall" gap="xsmall">
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
    ...boxProps
}: {
    file: File,
    onClickRemove?: () => void
} & BoxExtendedProps) {
    const [ hover, setHover ] = useState(false)

    return (
        <Box width="xsmall" height="xsmall" flex={{"shrink":0}} onMouseEnter={() => setHover(true)} onMouseLeave={() => setHover(false)}>
            <Stack anchor="top-right">
                <Box
                    width="xsmall"
                    height="xsmall"
                    flex={{"shrink":0}}
                    background="light-2"
                >
                    <Image
                        fit="cover"
                        src={ URL.createObjectURL(file) }
                    />
                </Box>
                {hover && <Button plain icon={<FormClose color="light-1"/>} onClick={() => onClickRemove && onClickRemove()}/>}
            </Stack>
        </Box>
    )
}

function rooms(houses: ProjectHouses ): Room[] {
    const house = firstHouse(houses)

    if (house) {
        switch (house.rooms.list.__typename) {
            case "HouseRoomsList":
                return house.rooms.list.items
            default:
                return []
        }
    }

    return []
}

function firstHouse(houses: ProjectHouses) {
    switch (houses.list.__typename) {
        case "ProjectHousesList":
            return houses.list.items[0]
        default:
            return undefined
    }
}