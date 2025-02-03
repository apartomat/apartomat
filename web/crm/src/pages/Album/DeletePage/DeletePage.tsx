import { Trash } from "grommet-icons"
import { Button } from "grommet"
import { useDeleteAlbumPage } from "pages/Album/DeletePage/api"
import React, { useEffect, useState } from "react"
import { Confirm } from "widgets/confirm"

export function DeletePage({
    albumId,
    pageNumber,
    onPageDeleted,
}: {
    albumId: string
    pageNumber
    onPageDeleted?: (pageNumber: number) => void
}) {
    const [deletePage, { loading, data, success }] = useDeleteAlbumPage(albumId)

    const [showConfirm, setShowConfirm] = useState(false)

    useEffect(() => {
        if (data?.deleteAlbumPage.__typename === "AlbumPageDeleted") {
            const n = data.deleteAlbumPage.page.number

            setShowConfirm(false)

            onPageDeleted && onPageDeleted(n)
        }
    }, [success])

    return (
        <>
            <Button plain onClick={() => setShowConfirm(true)}>
                <Trash color="icon" />
            </Button>

            {showConfirm && (
                <Confirm
                    header="Удалить страницу?"
                    text="Удаленную страницу не возможно будет восстановить."
                    confirmLabel="Удалить"
                    onClickClose={() => setShowConfirm(false)}
                    onCancel={() => setShowConfirm(false)}
                    onConfirm={() => deletePage(pageNumber)}
                    confirmBusy={loading}
                    confirmSuccess={success}
                />
            )}
        </>
    )
}
