import React, { useState } from "react"

import { Box, BoxExtendedProps, Button, Image } from "grommet"
import { Print, Trash } from "grommet-icons"

import { AnchorLink } from "shared/ui/AnchorLink"

import { ProjectScreenAlbum } from "pages/Project/useProject"

interface AlbumProps extends BoxExtendedProps {
    album: ProjectScreenAlbum
    onClickDelete?: (id: string) => void
}

export function Album({ album: { id, cover }, onClickDelete, ...props }: AlbumProps) {
    const [showDeleteButton, setShowDeleteButton] = useState(false)

    const handleClickDeleteButton = (event: React.MouseEvent) => {
        event.preventDefault()

        if (onClickDelete) {
            onClickDelete(id)
        }
    }

    return (
        <Box
            height="small"
            width="small"
            flex={{ shrink: 0 }}
            background="light-2"
            onMouseEnter={() => setShowDeleteButton(true)}
            onMouseLeave={() => setShowDeleteButton(false)}
            style={{ position: "relative" }}
            {...props}
        >
            {showDeleteButton && (
                <Box
                    pad="xsmall"
                    style={{ position: "absolute", right: 0 }}
                    background="background-back"
                    round="xxsmall"
                    margin="xxsmall"
                >
                    <Button plain onClick={handleClickDeleteButton}>
                        <Trash color="control" />
                    </Button>
                </Box>
            )}

            <AnchorLink to={`/album/${id}`}>
                {cover.__typename === "File" ? (
                    <Box background="light-2" height="small" width="small">
                        <Image fit="cover" src={cover.url} />
                    </Box>
                ) : (
                    <Box background="light-2" height="small" width="small" align="center" justify="center">
                        <Print size="large" color="dark-6" />
                    </Box>
                )}
            </AnchorLink>
        </Box>
    )
}
