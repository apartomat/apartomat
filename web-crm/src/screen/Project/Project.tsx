import React, { useEffect, useState, useContext } from "react"
import { useParams } from "react-router-dom"

import { Main, Box, Grid, Header, Heading, Text, Layer, Button, ResponsiveContext } from "grommet"
import { StatusGood } from "grommet-icons"

import AnchorLink from "common/AnchorLink"
import UserAvatar from "./UserAvatar/UserAvatar"

import { useAuthContext } from "common/context/auth/useAuthContext"

import { useProject, Project  as ProjectType } from "./useProject"
import { ProjectEnums } from "api/types"

import ChangeStatus from "./ChangeStatus/ChangeStatus"
import Contacts from "./Contacts/Contacts"
import Loading from "./Loading/Loading"
import AddSomething from "./AddSomething/AddSomething"
import ProjectDates from "./Dates/Dates"
import House from "./House/House"
import Rooms from "./Rooms/Rooms"
import Visualizations from "./Visualizations/Visualizations"
import UploadVisualizations from "./UploadVisualizations/UploadVisualizations"


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

    const respSize = useContext(ResponsiveContext);

    const notify = ({
        message,
        callback,
        timeout = 250,
        duration = 1500
    }: {
        message: string,
        callback?: () => void,
        timeout?: number,
        duration?: number
    }) => {
        setNotification(message)

        callback && callback()
        
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
        if (fetchError) {
            setError("Ошибка сервера")
        }
    }, [ fetchError ])

    useEffect(() => {
        if (data?.project) {
            switch (data.project.__typename) {
                case "Project":
                    setProject(data.project)
                    break
                case "NotFound":
                    setError("Проект не найден")
                    break
                case "Forbidden":
                    setError("Доступ запрещен")
                    break
            }
        }

        if (data?.enums) {
            setProjectEnums(data.enums.project)
        }
    }, [ data ])

    const [showUploadVisualizations, setShowUploadVisualizations] = useState(false);

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
                        <Heading level={2} margin="none">{project.name}</Heading>
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
                    <AddSomething
                        onClickAddVisualizations={() => {
                            setShowUploadVisualizations(true)
                        }}
                    />
                </Box>

                <Grid columns={{count: respSize === "small" ? 1 : 2, size: "auto"}} gap="small" responsive>
                    <Box>
                        <Heading level={4}>Сроки проекта</Heading>
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

                    <Box>
                        <Heading level={4}>Адрес</Heading>
                        <House
                            projectId={project.id}
                            houses={project.houses}
                            onAdd={() => refetch()}
                            onUpdate={() => refetch()}
                        />
                    </Box>

                    <Box>
                        <Heading level={4}>Заказчик</Heading>
                        <Contacts
                            projectId={project.id}
                            contacts={project.contacts}
                            notify={notify}
                            onAdd={() => notify({ message: "Контакт добавлен", callback: refetch })}
                            onDelete={() => notify({ message: "Контакт удален" })}
                            onUpdate={() => notify({ message: "Контакт сохранен" })}
                        />
                    </Box>

                    <Box>
                        <Heading level={4}>Комнаты</Heading>
                        <Rooms
                            houses={project.houses}
                            onAddRoom={() => notify({ message: "Комната добавлена", callback: refetch })}
                            onDeleteRoom={() => notify({ message: "Комната удалена", callback: refetch })}
                            onUpdateRoom={() => notify({ message: "Комната сохранена", callback: refetch })}
                        />
                    </Box>
                </Grid>

                {project.visualizations.list.__typename === "ProjectVisualizationsList" && project.visualizations.list.items.length > 0 &&
                    <Box margin={{vertical: "large"}}>
                        <Box direction="row" justify="between">
                            <Heading level={3}><AnchorLink to={`/p/${id}/vis`}>Визуализации</AnchorLink></Heading>
                            <Box justify="center">
                                <Button color="brand" label="Загрузить" onClick={() => setShowUploadVisualizations(true)} />
                            </Box>
                        </Box>
                        <Visualizations visualizations={project.visualizations}/>
                    </Box>
                }

                {showUploadVisualizations &&
                    <UploadVisualizations
                        projectId={project.id}
                        houses={project.houses}
                        onUploadComplete={({ files }: { files: File[] }) => {
                            setShowUploadVisualizations(false)
                            notify({ message: files?.length === 1 ? "Файл загружен" : `Загружено файлов ${files?.length}` })
                            refetch()
                        }}
                        onClickOutside={() => {
                            setShowUploadVisualizations(false)
                        }}
                        onClickClose={() => {
                            setShowUploadVisualizations(false)
                        }}
                    />}
                </Box>
            </Main>
        );
    } else {
        return <></>
    }
}

export default Project;