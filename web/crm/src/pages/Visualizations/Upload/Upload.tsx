import { Button } from "grommet"
import { useState } from "react"
import { UploadVisualizations } from "features/upload-visualizations"

export function Upload({
    projectId,
    rooms,
    roomId,
    onVisualizationsUploaded,
}: {
    projectId: string
    rooms: { id: string; name: string }[]
    roomId?: string
    onVisualizationsUploaded?: ({ files }: { files: File[] }) => void
}) {
    const [showUploadForm, setShowUploadForm] = useState(false)

    return (
        <>
            <Button label="Загрузить" onClick={() => setShowUploadForm(true)} />

            {showUploadForm && (
                <UploadVisualizations
                    projectId={projectId}
                    rooms={rooms}
                    roomId={roomId}
                    onClickClose={() => setShowUploadForm(false)}
                    onClickOutside={() => setShowUploadForm(false)}
                    onEsc={() => setShowUploadForm(false)}
                    onVisualizationsUploaded={(files) => {
                        setShowUploadForm(false)
                        onVisualizationsUploaded && onVisualizationsUploaded(files)
                    }}
                />
            )}
        </>
    )
}
