import React, { useEffect, useState } from "react"
import { useParams } from "react-router-dom"

import { useAuthContext } from "context/auth/context"
import { useWorkspace, WorkspaceScreenFragment } from "./useWorkspace"

import { Box, Button, Header, Heading, Text, Main, Paragraph } from "grommet"
import UserAvatar from "./UserAvatar"
import CreateProject from "./CreateProject/CreateProject"
import Projects from "./Projects/Projects"
import Users from "./Users/Users"
import Archive from "./Archive/Archive"
import Notification from "./Notification/Notification"
import { Spinner } from "shared/ui/Spinner"

export function Workspace() {
    const { id } = useParams<"id">() as { id: string }

    const { user } = useAuthContext()

    const [error, setError] = useState<string | undefined>(undefined)

    const [screen, setScreen] = useState<WorkspaceScreenFragment | undefined>(undefined)

    const {
        data,
        loading,
        error: fetchError,
    } = useWorkspace({ id, timezone: Intl.DateTimeFormat().resolvedOptions().timeZone })

    useEffect(() => {
        setError(fetchError ? "Ошибка сервера" : undefined)
    }, [fetchError])

    const [notification, setNotification] = useState<string | undefined>(undefined)

    const notify = ({
        message,
        callback,
        timeout = 250,
        duration = 1500,
    }: {
        message: string
        callback?: () => void
        timeout?: number
        duration?: number
    }) => {
        setTimeout(() => {
            setNotification(message)

            setTimeout(() => {
                setNotification(undefined)
                callback && callback()
            }, duration)
        }, timeout)
    }

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

    const [showCreateProjectLayer, setShowCreateProjectLayer] = useState(false)

    if (loading) {
        return (
            <Main pad="large">
                <Box direction="row" gap="small" align="center">
                    <Spinner message="Загрузка..." />
                    <Text>Загрузка...</Text>
                </Box>
            </Main>
        )
    }

    if (error) {
        return (
            <Main pad="large">
                <Heading>Ошибка</Heading>
                <Paragraph>{error}</Paragraph>
            </Main>
        )
    }

    if (!screen) {
        return null
    }

    const { projects, users } = screen

    return (
        <Main>
            {notification && <Notification message={notification} />}

            <Header margin={{ top: "large", horizontal: "large", bottom: "medium" }}>
                <Box>
                    <Text size="xlarge" weight="bold" color="brand">
                        apartomat
                    </Text>
                </Box>
                <Box>
                    <UserAvatar user={user} />
                </Box>
            </Header>

            <Box margin={{ horizontal: "large" }}>
                <Box margin={{ bottom: "medium" }}>
                    <Box direction="row" margin={{ vertical: "medium" }} justify="between">
                        <Heading level={2} margin="none">
                            {screen.name}
                        </Heading>
                        <Box justify="center">
                            <Button
                                color="brand"
                                label="Новый проект"
                                onClick={() => setShowCreateProjectLayer(true)}
                            />
                        </Box>
                    </Box>

                    <Projects projects={projects} />
                </Box>

                <Archive projects={projects} />

                <Users
                    workspaceId={screen.id}
                    users={users}
                    margin={{ vertical: "medium" }}
                    roles={screen.roles}
                    onUserInviteSent={(email) => {
                        notify({ message: `Приглашение отправлено на ${email}` })
                    }}
                />
            </Box>

            {showCreateProjectLayer && <CreateProject workspaceId={screen.id} setShow={setShowCreateProjectLayer} />}
        </Main>
    )
}
