import React, { useState, useRef } from "react"

import { Box, Button, Drop, Text, BoxExtendedProps } from "grommet"
import { Next } from "grommet-icons"
import CreateAlbumOnClick from "screen/Project/CreateAlbum/CreateAlbum";

export default function AddSomething({
     projectId,
     onAlbumCreated,
     onClickAddVisualizations,
     ...boxProps
}: {
    projectId: string,
    onAlbumCreated?: (id: string) => void
    onClickAddVisualizations: () => void
} & BoxExtendedProps) {
    const [show, setShow] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box { ...boxProps} justify="center">
            <Box ref={targetRef}>
                <Button label="Добавить" icon={<Next/>} reverse onClick={() => setShow(true)} />
            </Box>

            {show && targetRef.current && (
                <Drop
                    elevation="small"
                    round="small"
                    align={{top: "bottom", right: "right"}}
                    margin={{top: "xsmall"}}
                    target={targetRef.current}
                    onClickOutside={() => setShow(false)}
                    onEsc={() => setShow(false)}
                >
                    <Button
                        plain
                        children={({ hover }: {hover: boolean}) => {
                            return <Box pad="small" background={hover ? 'light-1': ''}><Text>Визуализации</Text></Box>
                        }}
                        onClick={() => {
                            setShow(false)
                            onClickAddVisualizations()
                        }}
                    />
                    <Button
                        plain
                        children={({ hover }: {hover: boolean}) => {
                            return (
                                <CreateAlbumOnClick
                                    projectId={projectId}
                                    onAlbumCreated={onAlbumCreated}
                                >
                                    <Box pad="small" background={hover ? 'light-1': ''}><Text>Альбом</Text></Box>
                                </CreateAlbumOnClick>
                            )
                        }}
                    />
                </Drop>
            )}
        </Box>
    )
}
