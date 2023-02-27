import React, { useEffect, useState } from "react"
import { useParams } from "react-router-dom"

import { useAuthContext } from "common/context/auth/useAuthContext"
import { useWorkspace, WorkspaceScreenFragment } from "./useWorkspace"

import { Box, Button, Header, Heading, Text, Main, Paragraph } from "grommet"
import UserAvatar from "./UserAvatar"
import CreateProject from "./CreateProject/CreateProject"
import Loading from "./Loading/Loading"
import Projects from "./Projects/Projects"
import Users from "./Users/Users"
import Archive from "./Archive/Archive"

interface RouteParams {
    id: string
}

export default function Workspace () {
    let { id } = useParams<RouteParams>()
    
    const { user } = useAuthContext()

    const [ error, setError ] = useState<string | undefined>(undefined)

    const [ screen, setScreen ] = useState<WorkspaceScreenFragment | undefined>(undefined)

    const { data, loading, error: fetchError } = useWorkspace({ id, timezone: Intl.DateTimeFormat().resolvedOptions().timeZone })

    useEffect(() => {
        setError(fetchError ? "Ошибка сервера" : undefined)
    }, [ fetchError ])

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
    }, [ data ])

    const [ showCreateProjectLayer, setShowCreateProjectLayer ] = useState(false)
    
    if (loading) {
        return (
            <Main pad="large">
                <Box direction="row" gap="small" align="center">
                    <Loading message="Загрузка..."/>
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
            <Header background="white" margin={{ top:"large", horizontal:"large", bottom:"medium" }}>
                <Box>
                    <Text size="xlarge" weight="bold" color="brand">apartomat</Text>
                </Box>
                <Box><UserAvatar user={user}/></Box>
            </Header>

            <Box margin={{ horizontal: "large" }}>
                <Box margin={{bottom: "medium"}}>
                    <Box direction="row" margin={{vertical: "medium"}} justify="between">
                        <Heading level={2} margin="none">{screen.name}</Heading>
                        <Box>
                            <Button color="brand" label="Новый проект" onClick={() => setShowCreateProjectLayer(true)} />
                        </Box>
                    </Box>

                    <Projects projects={projects}/>
                </Box>

                <Archive projects={projects}/>

                <Box margin={{vertical: "medium"}}>
                    <Users users={users}/>
                </Box>
            </Box>

            {showCreateProjectLayer && <CreateProject workspaceId={screen.id} setShow={setShowCreateProjectLayer} />}
        </Main>
    )
}