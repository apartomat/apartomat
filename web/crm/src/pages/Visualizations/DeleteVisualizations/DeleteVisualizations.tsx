import { useEffect, useState } from "react"
import { Button } from "grommet"
import { Trash } from "grommet-icons"
import { ConfirmDelete } from "./ConfirmDelete"

import useDeleteVisualizations from "./api/useDeleteVisualizations"

export function DeleteVisualizations({
    visualizations,
    onDelete,
}: {
    visualizations: string[]
    onDelete?: (n: number) => void
}) {
    const [showConfirmDialog, setShowConfirmDialog] = useState(false)

    const [deleteVisualizations, { data: deleteData, loading: deleteLoading }] = useDeleteVisualizations()

    const [deleting, setDeleting] = useState(false)

    const handleClickDelete = () => {
        if (visualizations.length === 0) {
            return
        }

        setShowConfirmDialog(true)
    }

    const handleClickConfirmDelete = async () => {
        setDeleting(true)
        await deleteVisualizations(visualizations)
        setShowConfirmDialog(false)
    }

    const handleClickCancelDelete = () => {
        setShowConfirmDialog(false)
    }

    useEffect(() => {
        if (deleting && deleteData?.deleteVisualizations.__typename === "VisualizationsDeleted") {
            setDeleting(false)
            onDelete && onDelete(deleteData?.deleteVisualizations.visualizations.length)
        }
    }, [deleteData, deleting])

    return (
        <>
            <Button
                disabled={visualizations.length === 0}
                icon={<Trash color="brand" />}
                // label="Удалить"
                onClick={handleClickDelete}
            />

            {showConfirmDialog && (
                <ConfirmDelete
                    count={visualizations.length}
                    disableButton={deleteLoading}
                    onEsc={handleClickCancelDelete}
                    onClickClose={handleClickCancelDelete}
                    onClickDelete={handleClickConfirmDelete}
                />
            )}
        </>
    )
}
