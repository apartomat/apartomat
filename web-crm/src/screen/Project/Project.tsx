import React, { useEffect, useState } from "react"
import { useParams } from "react-router-dom"

import { Main, Box, Header, Heading, Text, Layer, Button } from "grommet"
import { StatusGood } from "grommet-icons"

import AnchorLink from "common/AnchorLink"
import UserAvatar from "./UserAvatar/UserAvatar"

import { useAuthContext } from "common/context/auth/useAuthContext"

import { useProject, Project  as ProjectType } from "./useProject"
import { ProjectFileType } from "./useUploadProjectFile"
import { ProjectEnums } from "api/types"

import ChangeStatus from "./ChangeStatus/ChangeStatus"
import Contacts from "./Contacts/Contacts"
import Loading from "./Loading/Loading"
import AddSomething from "./AddSomething/AddSomething"
import ProjectDates from "./Dates/Dates"
import House from "./House/House"
import Rooms from "./Rooms/Rooms"
import Visualizations from "./Visualizations/Visualizations"
import UploadFiles from "./UploadFiles/UploadFiles"


interface RouteParams {
    id: string
};

export function Project () {
    let { id } = useParams<RouteParams>()

    const [ error, setError ] = useState<string | undefined>(undefined)

    const { user } = useAuthContext()

    const { data, loading, error: fetchError, refetch, refetching } = useProject(id)

    const [ notification, setNotification ] = useState("")
    const [ showNotification, setShowNotification ] = useState(false)

    const notify = ({ message, timeout = 250, duration = 1500 }: { message: string, timeout?: number, duration?: number }) => {
        setNotification(message)
        
        setTimeout(() => {
            setShowNotification(true)

            setTimeout(() => {
                setShowNotification(false)
            }, duration)
        }, timeout)
    }

    const [ project, setProject ] = useState<ProjectType | undefined>(undefined)

    const [ projectEnums, setProjectEnums ] = useState<ProjectEnums | undefined>(undefined)

    useEffect(() => {
        console.log({ loading, data, fetchError })
    }, [ loading, data, fetchError ])

    useEffect(() => {
        if (fetchError) {
            setError("Ошибка сервера")
            console.error({...fetchError})
        }
    }, [ fetchError ])

    useEffect(() => {
        const screen = data?.screen.projectScreen

        if (screen?.project) {
            switch (screen.project.__typename) {
                case "Project":
                    setProject(screen.project)
                    setProjectEnums(screen.enums)
                    break
                case "NotFound":
                    setError("Проект не найден")
                    break
                case "Forbidden":
                    setError("Доступ запрещен")
                    break
            }
        }        
    }, [ data ])

    const [showUploadFiles, setShowUploadFiles] = useState(false);

    if (loading && !refetching) {
        return (
            <Main pad="large">
                <Box direction="row" gap="small" align="center">
                    <Loading message="Загрузка..."/>
                    <Text>Загрузка...</Text>
                </Box>
            </Main>
        );
    }

    if (error) {
        return (
        <Main pad="large">
            <Heading level={2}>Ошибка</Heading>
            <Box>
                <Text>{error}</Text>
            </Box>
        </Main>
        )
    }

    if (project) {
        return (
            <Main pad={{vertical: "medium", horizontal: "large"}}>

            {refetching ?
                <Layer position="top" margin="medium" plain animate={false}>
                    <Box direction="row" gap="small">
                        <Loading message="Загрузка..."/>
                        <Text>Загрузка...</Text>
                    </Box>
                </Layer> : null}

            {showNotification ? <Layer
                position="top"
                modal={false}
                responsive={false}
                margin={{ vertical: "small", horizontal: "small"}}
            >
                <Box
                    align="center"
                    direction="row"
                    gap="xsmall"
                    justify="between"
                    elevation="small"
                    background="status-ok"
                    round="medium"
                    pad={{ vertical: "xsmall", horizontal: "small"}}
                >
                    <StatusGood/>
                    <Text>{notification}</Text>
                </Box>
            </Layer> : null}

            <Header background="white" margin={{vertical: "medium"}}>
                <Box>
                    <Text size="xlarge" weight="bold" color="brand">
                        <AnchorLink to="/">apartomat</AnchorLink>
                    </Text>
                </Box>
                <Box><UserAvatar user={user} className="header-user" /></Box>
            </Header>

            <Box>
                <Box direction="row" justify="between" margin={{vertical: "medium"}}>
                    <Box direction="row" justify="center">
                        <Heading level={2} margin="none">{project.title}</Heading>
                        <ChangeStatus
                            margin={{ horizontal: "medium"}}
                            projectId={project.id}
                            status={project.status}
                            values={projectEnums?.status}
                            onProjectStatusChanged={({ status }) => {
                                setProject({ ...project, status })
                            }}
                        />
                    </Box>
                    <AddSomething showUploadFiles={setShowUploadFiles}/>
                </Box>

                <Box direction="row" justify="between" wrap>
                    <Box width={{min: "35%"}}>
                        <Box margin="none">
                            <Heading level={4} margin={{ bottom: "xsmall"}}>Сроки проекта</Heading>
                            <ProjectDates
                                projectId={project.id}
                                startAt={project.startAt}
                                endAt={project.endAt}
                                onChange={({ startAt, endAt }) => {
                                    notify({ message: "Даты изменены" })
                                    setProject({ ...project, startAt, endAt })
                                }}
                            />
                        </Box>
                        <Box margin={{top: "small"}}>
                            <Heading level={4} margin={{ bottom: "xsmall" }}>Заказчик</Heading>
                            <Contacts
                                projectId={project.id}
                                contacts={project.contacts}
                                notify={notify}
                                onAdd={() => notify({ message: "Контакт добавлен"})}
                                onDelete={() => notify({ message: "Контакт удален"})}
                                onUpdate={() => notify({ message: "Контакт сохранен"})}
                            />
                        </Box>

                    </Box>
                    <Box width={{min: "35%"}}>
                        <Box margin="none">
                            <Heading level={4} margin={{ bottom: "xsmall" }}>Адрес</Heading>
                            <House
                                projectId={project.id}
                                houses={project.houses}
                                onAdd={() => refetch()}
                                onUpdate={() => refetch()}
                            />
                        </Box>
                        <Box margin={{top: "small"}}>
                            <Heading level={4} margin={{ bottom: "xsmall"}}>Комнаты</Heading>
                            <Rooms houses={project.houses}/>
                        </Box>
                    </Box>
                </Box>

                {project.files.list.__typename === "ProjectFilesList" && project.files.list.items.length > 0 &&
                    <Box margin={{vertical: "large"}}>
                        <Box direction="row" justify="between">
                            <Heading level={3}>Визуализации</Heading>
                            <Box justify="center">
                                <Button color="brand" label="Загрузить" onClick={() => setShowUploadFiles(true)} />
                            </Box>
                        </Box>
                        <Visualizations files={project.files}/>
                    </Box>
                }

                {showUploadFiles ?
                    <UploadFiles
                        projectId={project.id}
                        type={ProjectFileType.Visualization}
                        setShow={setShowUploadFiles}
                        onUploadComplete={({message}) => notify({ message })}
                    /> : null}
                </Box>
            </Main>
        );
    } else {
        return <></>
    }
}

export default Project;