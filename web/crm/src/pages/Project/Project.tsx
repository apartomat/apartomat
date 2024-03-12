import React, { useEffect, useState, useContext, useRef } from "react"
import { useParams, useNavigate } from "react-router-dom"

import {
    Main,
    Box,
    Grid,
    Header,
    Heading,
    Text,
    Layer,
    Button,
    ResponsiveContext,
    BoxExtendedProps,
    Drop,
} from "grommet"
import { Add, StatusGood } from "grommet-icons"

import { AnchorLink } from "shared/ui/AnchorLink"
import UserAvatar from "./UserAvatar/UserAvatar"

import { useAuthContext } from "context/auth/context"

import { useProject, Project as ProjectType } from "./useProject"
import { ProjectStatusDictionary } from "api/graphql"

import ChangeStatus from "./ChangeStatus/ChangeStatus"
import Contacts from "./Contacts/Contacts"
import AddSomething from "./AddSomething/AddSomething"
import ProjectDates from "./Dates/Dates"
import House from "./House/House"
import Rooms from "./Rooms/Rooms"
import Visualizations from "./Visualizations/Visualizations"
import UploadVisualizations from "./UploadVisualizations/UploadVisualizations"
import CreateAlbumOnClick from "./CreateAlbum/CreateAlbum"
import Albums from "./Albums/Albums"
import PublicSite from "./PublicSite/PublicSite"
import { Spinner } from "shared/ui/Spinner"

export function Project() {
    const { id } = useParams<"id">() as { id: string }

    const [error, setError] = useState<string | undefined>(undefined)

    const { user } = useAuthContext()

    const { data, loading, error: fetchError, refetch, refetching } = useProject(id)

    const [notification, setNotification] = useState("")

    const [showNotification, setShowNotification] = useState(false)

    const respSize = useContext(ResponsiveContext)

    const navigate = useNavigate()

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
        setNotification(message)

        callback && callback()

        setTimeout(() => {
            setShowNotification(true)

            setTimeout(() => {
                setShowNotification(false)
            }, duration)
        }, timeout)
    }

    const [project, setProject] = useState<ProjectType | undefined>(undefined)

    const [projectStatusDictionary, setProjectStatusDictionary] = useState<ProjectStatusDictionary | undefined>(
        undefined
    )

    useEffect(() => {
        if (fetchError) {
            setError("Ошибка сервера")
        }
    }, [fetchError])

    useEffect(() => {
        if (data?.project) {
            switch (data.project.__typename) {
                case "Project":
                    setProject(data.project)
                    setProjectStatusDictionary(data.project.statuses)
                    break
                case "NotFound":
                    setError("Проект не найден")
                    break
                case "Forbidden":
                    setError("Доступ запрещен")
                    break
            }
        }
    }, [data])

    const [showUploadVisualizations, setShowUploadVisualizations] = useState(false)

    const [redirectTo, setRedirectTo] = useState<string | undefined>(undefined)

    if (redirectTo) {
        navigate(redirectTo)
    }

    if (loading && !refetching) {
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
                <Heading level={2}>Ошибка</Heading>
                <Box>
                    <Text>{error}</Text>
                </Box>
            </Main>
        )
    }

    if (!project) {
        return <></>
    }

    return (
        <Main pad={{ vertical: "medium", horizontal: "large" }}>
            {refetching && (
                <Layer position="top" margin="medium" plain animate={false}>
                    <Box direction="row" gap="small">
                        <Spinner message="Загрузка..." />
                        <Text>Загрузка...</Text>
                    </Box>
                </Layer>
            )}

            {showNotification ? (
                <Layer
                    position="top"
                    modal={false}
                    responsive={false}
                    margin={{ vertical: "small", horizontal: "small" }}
                >
                    <Box
                        align="center"
                        direction="row"
                        gap="xsmall"
                        justify="between"
                        elevation="small"
                        background="status-ok"
                        round="medium"
                        pad={{ vertical: "xsmall", horizontal: "small" }}
                    >
                        <StatusGood />
                        <Text>{notification}</Text>
                    </Box>
                </Layer>
            ) : null}

            <Header margin={{ vertical: "medium" }}>
                <Box>
                    <Text size="xlarge" weight="bold" color="brand">
                        <AnchorLink to="/">apartomat</AnchorLink>
                    </Text>
                </Box>
                <Box>
                    <UserAvatar user={user} className="header-user" />
                </Box>
            </Header>

            <Box>
                <Box direction="row" justify="between" margin={{ vertical: "medium" }}>
                    <Box direction="row" justify="center">
                        <Heading level={2} margin="none">
                            {project.name}
                        </Heading>
                        <ChangeStatus
                            margin={{ horizontal: "medium" }}
                            projectId={project.id}
                            status={project.status}
                            values={projectStatusDictionary}
                            onProjectStatusChanged={({ status }) => {
                                if (project) {
                                    setProject({ ...project, status })
                                }
                            }}
                        />
                    </Box>
                    <Box direction="row" gap="small">
                        <PublicSite
                            projectId={id}
                            site={project.publicSite}
                            onClose={(changed: boolean) => {
                                if (changed) {
                                    refetch()
                                }
                            }}
                        />
                        <AddSomething
                            projectId={id}
                            onAlbumCreated={(id) => {
                                setRedirectTo(`/album/${id}`)
                            }}
                            onClickAddVisualizations={() => {
                                setShowUploadVisualizations(true)
                            }}
                        />
                    </Box>
                </Box>

                <Grid
                    columns={{ count: respSize === "small" ? 1 : 2, size: "auto" }}
                    gap="small"
                    responsive
                    margin={{ bottom: "large" }}
                >
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
                        <Contacts projectId={project.id} contacts={project.contacts} notify={notify} />
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

                {project.visualizations.list.__typename === "ProjectVisualizationsList" &&
                    project.visualizations.list.items.length > 0 && (
                        <Box margin={{ bottom: "large" }}>
                            <Box direction="row" justify="between">
                                <Heading level={3}>
                                    <AnchorLink to={`/vis/${id}`}>Визуализации</AnchorLink>
                                </Heading>
                                <Box justify="center">
                                    <Button
                                        color="brand"
                                        label="Загрузить"
                                        onClick={() => setShowUploadVisualizations(true)}
                                    />
                                </Box>
                            </Box>
                            <Visualizations visualizations={project.visualizations} />
                        </Box>
                    )}

                {project.albums.list.__typename === "ProjectAlbumsList" && project.albums.list.items.length > 0 && (
                    <Box margin={{ bottom: "large" }}>
                        <Box direction="row" justify="between">
                            <Heading level={3}>Альбомы</Heading>
                            <Box justify="center">
                                {/* <Button color="brand" label="Загрузить" onClick={() => setShowUploadVisualizations(true)} /> */}
                            </Box>
                        </Box>
                        <Albums
                            albums={project.albums}
                            onDelete={(albums) => {
                                notify({
                                    message:
                                        albums?.length === 1 ? "Альбом удален" : `Удалено альбомов ${albums.length}`,
                                })
                                refetch()
                            }}
                        />
                    </Box>
                )}

                {showUploadVisualizations && (
                    <UploadVisualizations
                        projectId={project.id}
                        houses={project.houses}
                        onUploadComplete={({ files }: { files: File[] }) => {
                            setShowUploadVisualizations(false)
                            notify({
                                message: files?.length === 1 ? "Файл загружен" : `Загружено файлов ${files?.length}`,
                            })
                            refetch()
                        }}
                        onClickOutside={() => {
                            setShowUploadVisualizations(false)
                        }}
                        onClickClose={() => {
                            setShowUploadVisualizations(false)
                        }}
                    />
                )}
            </Box>

            <AddSomething2
                projectId={id}
                style={{
                    position: "fixed",
                    bottom: 0,
                    left: 0,
                    right: 0,
                }}
                direction="row"
                justify="center"
                pad="small"
                onClickAddVisualizations={() => {
                    setShowUploadVisualizations(true)
                }}
                onAlbumCreated={(id: string) => {
                    setRedirectTo(`/album/${id}`)
                }}
            />
        </Main>
    )
}

function AddSomething2({
    projectId,
    onClickAddVisualizations,
    onAlbumCreated,
    ...boxProps
}: {
    projectId: string
    onClickAddVisualizations?: () => void
    onAlbumCreated?: (id: string) => void
} & BoxExtendedProps) {
    const [show, setShow] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box {...boxProps}>
            <Box ref={targetRef} border={{ color: "background-front", size: "medium" }} round="large">
                <Button label="Добавить..." icon={<Add />} primary onClick={() => setShow(true)} />
            </Box>

            {show && targetRef.current && (
                <Drop
                    elevation="none"
                    target={targetRef.current}
                    onClickOutside={() => setShow(false)}
                    onEsc={() => setShow(false)}
                    align={{ bottom: "bottom" }}
                    round="large"
                >
                    <Box gap="small" border={{ color: "background-front", size: "medium" }} direction="row">
                        <Button primary label="Визуализации" color="accent-3" onClick={onClickAddVisualizations} />
                        <CreateAlbumOnClick projectId={projectId} onAlbumCreated={onAlbumCreated}>
                            <Button primary label="Альбом" />
                        </CreateAlbumOnClick>
                        {/* <AnchorLink to={`/p/${projectId}/album`} weight="normal"><Button primary label="Альбом"/></AnchorLink> */}
                    </Box>
                </Drop>
            )}
        </Box>
    )
}
