import React, { useEffect, useState } from "react"
import { useParams } from "react-router-dom"

import { useWorkspace, WorkspaceScreenFragment } from "./useWorkspace"

import { Box, Button } from "grommet"
import CreateProject from "./CreateProject/CreateProject"
import Projects from "./Projects/Projects"
import Users from "./Users/Users"
import Archive from "./Archive/Archive"
import { MainLayout } from "widgets/main-layout/MainLayout"
import { useNotifications } from "shared/context/notifications/context"

export function Workspace() {
    const { id } = useParams<"id">() as { id: string }

    const [error, setError] = useState<string | undefined>(undefined)

    const [screen, setScreen] = useState<WorkspaceScreenFragment | undefined>(undefined)

    const {
        data,
        loading,
        error: fetchError,
        refetch,
    } = useWorkspace({ id, timezone: Intl.DateTimeFormat().resolvedOptions().timeZone })

    useEffect(() => {
        setError(fetchError ? "Ошибка сервера" : undefined)
    }, [fetchError])

    const { notify } = useNotifications()

    useEffect(() => {
        if (data) {
            switch (data.workspace.__typename) {
                case "Workspace":
                    setScreen(data.workspace)
                    break
                case "NotFound":
                    setError("Проект не найден")
                    break
                case "Forbidden":
                    setError("Доступ запрещен")
                    break
                default:
                    setError("Неизвестная ошибка")
                    break
            }
        }
    }, [data])

    const [showCreateProject, setShowCreateProject] = useState(false)

    return (
        <MainLayout
            loading={loading}
            error={error}
            header={screen?.name}
            headerMenu={
                <Box justify="center">
                    <Button color="brand" label="Новый проект" onClick={() => setShowCreateProject(true)} />
                </Box>
            }
        >
            {screen && (
                <>
                    <Box margin={{ bottom: "medium" }}>
                        <Projects projects={screen.projects} />
                    </Box>

                    <Archive projects={screen.projects} />

                    <Users
                        workspaceId={screen.id}
                        users={screen.users}
                        margin={{ vertical: "medium" }}
                        roles={screen.roles}
                        onUserInviteSent={(email) => {
                            notify({ message: `Приглашение отправлено на ${email}` })
                        }}
                    />

                    {showCreateProject && (
                        <CreateProject
                            workspaceId={screen.id}
                            onClickClose={() => setShowCreateProject(false)}
                            onCreate={async () => {
                                setShowCreateProject(false)
                                notify({
                                    message: "Проект создан",
                                })
                                await refetch()
                            }}
                        />
                    )}
                </>
            )}
        </MainLayout>
    )
}
