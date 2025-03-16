import { useState } from "react"
import { Box, BoxExtendedProps, Button } from "grommet"
import { Link } from "grommet-icons"
import { ProjectScreenProjectPageFragment, ProjectPageStatus } from "api/graphql"
import { EditForm } from "./EditForm"

export function ProjectPage({
    projectId,
    page,
    onChange,
    onClose,
    ...props
}: {
    projectId: string
    page: ProjectScreenProjectPageFragment
    onChange?: () => void
    onClose?: (changed: boolean) => void
} & BoxExtendedProps) {
    const [showForm, setShowForm] = useState(false)

    return (
        <Box {...props} justify="center">
            <Box>
                <Button
                    label="Ссылка"
                    icon={<Link color={pageIsPublic(page) ? "status-ok" : "status-unknown"} />}
                    onClick={() => setShowForm(!showForm)}
                    title="Ссылка"
                />
            </Box>

            {showForm && (
                <EditForm
                    projectId={projectId}
                    projectPage={page}
                    onEsc={() => setShowForm(false)}
                    onClickClose={(changed: boolean) => {
                        setShowForm(false)
                        onClose && onClose(changed)
                    }}
                    onChange={onChange}
                />
            )}
        </Box>
    )
}

function pageIsPublic(page: ProjectScreenProjectPageFragment): boolean {
    return page.__typename === "ProjectPage" && page.status === ProjectPageStatus.Public
}
