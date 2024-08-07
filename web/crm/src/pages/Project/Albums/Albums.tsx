import React, { useEffect, useState } from "react"

import { Box, BoxExtendedProps, Grid } from "grommet"

import { Album } from "./Album"
import { ConfirmDelete } from "./ConfirmDelete"

import { ProjectScreenAlbumsFragment as ProjectScreenAlbums } from "api/graphql"

import { useDeleteAlbum } from "./useDeleteAlbum"

export function Albums({
    albums,
    onDelete,
    ...props
}: {
    albums: ProjectScreenAlbums
    onDelete?: (id: string[]) => void
} & BoxExtendedProps) {
    const [showConfirmDeleteDialog, setShowConfirmDeleteDialog] = useState<string | undefined>()

    const [deleteAlbum, { data, loading }] = useDeleteAlbum()

    useEffect(() => {
        switch (data?.deleteAlbum.__typename) {
            case "AlbumDeleted":
                setShowConfirmDeleteDialog(undefined)
                onDelete && onDelete([data?.deleteAlbum.album.id])
        }
    }, [data])

    const handleClickDelete = (id: string) => {
        setShowConfirmDeleteDialog(id)
    }

    const handleClickCancelDelete = () => {
        setShowConfirmDeleteDialog(undefined)
    }

    const handleClickConfirmDelete = () => {
        if (showConfirmDeleteDialog) {
            deleteAlbum(showConfirmDeleteDialog)
        }
    }

    switch (albums.list.__typename) {
        case "ProjectAlbumsList":
            if (albums.list.items.length === 0) {
                return null
            }

            return (
                <Box {...props}>
                    {showConfirmDeleteDialog && (
                        <ConfirmDelete
                            disableButton={loading}
                            onEsc={handleClickCancelDelete}
                            onClickClose={handleClickCancelDelete}
                            onClickDelete={handleClickConfirmDelete}
                        />
                    )}

                    <Box overflow="auto">
                        <Grid columns="small" style={{ gridAutoFlow: "column", overflowX: "scroll" }} gap="xsmall">
                            {albums.list.items.map((item) => {
                                switch (item.__typename) {
                                    case "Album":
                                        return <Album album={item} key={item.id} onClickDelete={handleClickDelete} />
                                    default:
                                        return null
                                }
                            })}
                        </Grid>
                    </Box>
                </Box>
            )
        default:
            return null
    }
}
