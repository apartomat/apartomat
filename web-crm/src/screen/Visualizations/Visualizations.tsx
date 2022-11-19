import React, { useEffect, useState } from "react"
import { useParams } from "react-router-dom"

import { useAuthContext } from "common/context/auth/useAuthContext"

import { useVisualizations, VisualizationStatus } from "./useVisualizations"
import { useDeleteVisualizations } from "./useDeleteVisualizations"
import { VisualizationsScreenVisualizationFragment, VisualizationsScreenHouseRoomFragment } from "api/types"

import { Box, Button, CheckBox, Header, Heading, Grid, Image, Layer, Main, Text } from "grommet"
import Loading from "./Loading/Loading"
import AnchorLink from "common/AnchorLink"
import UserAvatar from "./UserAvatar/UserAvatar"
import ConfirmDelete from "./ConfirmDelete/ConfirmDelete"
import { Trash } from "grommet-icons"
import Notification from "./Notification/Notifications"

export default function Visualizations() {
    const { id } = useParams<{ id: string }>()

    const { user } = useAuthContext()


    const [ error, setError ] = useState<string | undefined>(undefined)

    const [ notification, setNotification ] = useState<string | undefined>(undefined)

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
        setTimeout(() => {
            setNotification(message)

            setTimeout(() => {
                setNotification(undefined)

                callback && callback()
            }, duration)
        }, timeout)
    }


    const { data, loading, refetch, first } = useVisualizations(
        id,
        { status: { eq: [VisualizationStatus.Approved, VisualizationStatus.Unknown] },
    })

    const [ project, setProject ] = useState<{ id: string, name: string }>()

    const [ visualizations, setVisualizations ] = useState<VisualizationsScreenVisualizationFragment[]>()

    const [ rooms, setRooms ] = useState<VisualizationsScreenHouseRoomFragment[]>()

    useEffect(() => {
        if (data?.project) {
            switch (data.project.__typename) {
                case "Project":
                    setProject({ id: data.project.id, name: data.project.name })

                    if (data?.project?.visualizations?.list?.__typename === "ProjectVisualizationsList") {
                        setVisualizations(data.project.visualizations.list.items)
                    }
        
                    if (data &&
                        data.project.houses.list.__typename === "ProjectHousesList" &&
                        data.project.houses.list.items.length !== 0
                        ) {
                            const h = data.project.houses.list.items[0]
            
                            if (h.rooms.list.__typename === "HouseRoomsList") {
                                setRooms(h.rooms.list.items)
                            }
                    }
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

    const [ roomsFilter, setRoomsfilter ] = useState<string[]>([])

    useEffect(() => {
        setSelected([])
        refetch({ filter: {
            roomID: { eq: roomsFilter },
            status: { eq: [ VisualizationStatus.Approved, VisualizationStatus.Unknown ] }
        }})
    }, [ roomsFilter, refetch ])

    
    const [ selected, setSelected ] = useState<string[]>([])

    const selectVis = (id: string, add: boolean) => {
        if (add && selected.includes(id)) {
            setSelected(selected.filter(s => s !== id))
        } else if (add) {
            setSelected([...selected, id])
        } else {
            setSelected([id])
        }
    }


    const [ showConfirmDialog, setShowConfirmDialog ] = useState(false)

    const handleClickDelete = () => {
        if (selected.length === 0) {
            return
        }

        setShowConfirmDialog(true)
    }

    const handleClickConfirmDelete = () => {
        setDeleting(true)
        deleteVisualizations(selected)
        setShowConfirmDialog(false)
    }

    const handleClickCancelDelete = () => {
        setShowConfirmDialog(false)
    }


    const [ deleteVisualizations, { data: deleteData, loading: deleteLoading } ] = useDeleteVisualizations()

    const [ deleting, setDeleting ] = useState(false)

    useEffect(() => {
        if (deleting && deleteData?.deleteVisualizations.__typename === "VisualizationsDeleted") {
            setSelected([])
            setDeleting(false)

            notify({
                message: `Удалено ${deleteData?.deleteVisualizations.visualizations.length} визуализаций`,
                callback: refetch,
            })
        }
    }, [ deleteData, deleting, refetch ])


    if (loading && first) {
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

    return (
        <Main pad={{vertical: "medium", horizontal: "large"}}>

            {loading && !first &&
                <Layer position="top" margin="medium" plain animate={false}>
                    <Box direction="row" gap="small">
                        <Loading message="Загрузка..."/>
                        <Text>Загрузка...</Text>
                    </Box>
                </Layer>
            }

            {notification &&
                <Notification message={notification}/>
            }

            <Header background="white" margin={{vertical: "medium"}}>
                <Box>
                    <Text size="xlarge" weight="bold" color="brand">
                        <AnchorLink to="/">apartomat</AnchorLink>
                    </Text>
                </Box>
                <Box><UserAvatar user={user} className="header-user" /></Box>
            </Header>

            <Box direction="row" justify="between" margin={{ vertical: "medium" }}>
                <Heading level={2} margin="none"><AnchorLink to={`/p/${project?.id}/`} color="black">{project?.name}</AnchorLink></Heading>
                <Box direction="row" gap="small" justify="center" align="center">
                    <Button
                        disabled={selected.length === 0}
                        icon={<Trash color="brand"/>}
                        label="Удалить"
                        onClick={handleClickDelete}
                    />
                </Box>
            </Box>

            <Box margin={{bottom: "medium"}}>
                <Box direction="row" justify="between">
                {rooms &&
                    <Box direction="row" wrap margin={{ vertical: "medium" }} gap="small">
                        {rooms.map((room) => {
                            return (
                                <CheckBox
                                    key={room.id}
                                    onChange={({ target: { checked }}) => {
                                        if (checked) {
                                            setRoomsfilter([...roomsFilter, room.id])
                                        } else {
                                            setRoomsfilter(roomsFilter.filter(item => item !== room.id))
                                        }
                                    }}
                                    checked={roomsFilter.includes(room.id)}
                                >
                                    {({ checked }: { checked: boolean }) => (
                                        <Box
                                            pad={{horizontal: "small", vertical: "xsmall"}}
                                            background={checked ? "brand" : "light-1"}
                                            round="medium"
                                        >{room.name}</Box>
                                    )}
                                </CheckBox>
                            )
                        })}
                    </Box>
                }
                    {selected.length > 1 &&
                        <Box direction="row" gap="small" align="center">
                            <Text size="small">Выбрано {selected.length}</Text>
                            <Button label="Отменить" size="small" color="dark-5" onClick={() => setSelected([])}/>
                        </Box>
                    }
                </Box>

                <Grid
                    columns="small"
                    rows="small"
                    gap="medium"
                    onClick={(event) => {
                        setSelected([])
                    }}
                >
                    {visualizations?.map(({ id, file }) => {
                        return (
                            <Box
                                key={id}
                                width={{ max: "small" }}
                                height={{ max: "small" }}
                                justify="center"
                                align="center"
                            >
                                <Box
                                    round="xsmall"
                                    pad="xxsmall"
                                    onClick={(event: MouseEvent) => {
                                        selectVis(id, event.metaKey)
                                        event.stopPropagation()
                                    }}
                                    focusIndicator={false}
                                    style={{boxShadow: selected.includes(id) ? "0 0 0px 2px #7D4CDB": "none" }}
                                >
                                    <Image
                                        fit="contain"
                                        src={`${file.url}?h=192`}
                                    />
                                </Box>
                            </Box>

                        )
                    })}
                </Grid>
            </Box>

            {showConfirmDialog &&
                <ConfirmDelete
                    count={selected.length}
                    disableButton={deleteLoading}
                    onEsc={handleClickCancelDelete}
                    onClickClose={handleClickCancelDelete}
                    onClickCancel={handleClickCancelDelete}
                    onClickDelete={handleClickConfirmDelete}
                />
            }
        </Main>
    )
}