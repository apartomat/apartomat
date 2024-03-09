import React, { useState } from "react"

import { Box, BoxExtendedProps, Button } from "grommet"
import { Print, Trash } from "grommet-icons"

import AnchorLink from "common/AnchorLink"

export function Album({
    id,
    onClickDelete,
    ...props
}: {
    id: string
    onClickDelete?: (id: string) => void
} & BoxExtendedProps) {
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
            onMouseEnter={() => setShowDeleteButton(true)}
            onMouseLeave={() => setShowDeleteButton(false)}
            style={{ position: "relative" }}
            {...props}
        >
            {showDeleteButton && (
                <Box pad="small" style={{ position: "absolute", right: 0 }}>
                    <Button plain onClick={handleClickDeleteButton}>
                        <Trash />
                    </Button>
                </Box>
            )}

            <AnchorLink to={`/album/${id}`}>
                <Box background="light-2" direction="column" justify="center" align="center" height="small">
                    <Print size="large" color="dark-6" />
                </Box>
            </AnchorLink>
        </Box>
    )
}
