import React, { useState, useRef } from "react"

import { Box, Button, Drop, Text, BoxExtendedProps } from "grommet"
import { Next } from "grommet-icons"
import CreateAlbumOnClick from "pages/Project/CreateAlbum/CreateAlbum"

export default function AddSomething({
    projectId,
    onAlbumCreated,
    onClickAddVisualizations,
    ...boxProps
}: {
    projectId: string
    onAlbumCreated?: (id: string) => void
    onClickAddVisualizations: () => void
} & BoxExtendedProps) {
    const [show, setShow] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box {...boxProps} justify="center">
            <Box ref={targetRef}>
                <Button label="Добавить" icon={<Next />} reverse onClick={() => setShow(true)} />
            </Box>

            {show && targetRef.current && (
                <Drop
                    elevation="small"
                    round="small"
                    align={{ top: "bottom", right: "right" }}
                    margin={{ top: "xsmall" }}
                    target={targetRef.current}
                    onClickOutside={() => setShow(false)}
                    onEsc={() => setShow(false)}
                >
                    <Button
                        plain
                        hoverIndicator
                        onClick={() => {
                            setShow(false)
                            onClickAddVisualizations()
                        }}
                    >
                        <Box pad="small">
                            <Text>Визуализации</Text>
                        </Box>
                    </Button>
                    <Button plain hoverIndicator>
                        <CreateAlbumOnClick projectId={projectId} onAlbumCreated={onAlbumCreated}>
                            <Box pad="small">
                                <Text>Альбом</Text>
                            </Box>
                        </CreateAlbumOnClick>
                    </Button>
                </Drop>
            )}
        </Box>
    )
}
