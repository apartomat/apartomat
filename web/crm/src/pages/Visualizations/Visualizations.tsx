import { MouseEvent, useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { useVisualizations, VisualizationStatus } from "./api/useVisualizations"
import { VisualizationsScreenHouseRoomFragment, VisualizationsScreenVisualizationFragment } from "api/graphql"
import { MainLayout } from "widgets/main-layout/MainLayout"
import { AnchorLink } from "shared/ui"
import { LinkPrevious } from "grommet-icons"
import { Box, Button, Grid, Image, Text } from "grommet"

import { Rooms, useSearchParamsRoomsFilter } from "./Rooms"
import { DeleteVisualizations } from "./DeleteVisualizations"
import { useNotifications } from "shared/context/notifications/context"
import { Upload } from "./Upload"

type Rooms = Pick<VisualizationsScreenHouseRoomFragment, "id" | "name">[]

export function Visualizations() {
    const { id } = useParams<"id">() as { id: string }

    const { notify: notify } = useNotifications()

    const [roomsFilter] = useSearchParamsRoomsFilter()

    const { data, error, loading, refetch, refetching } = useVisualizations(id, {
        status: { eq: [VisualizationStatus.Approved, VisualizationStatus.Unknown] },
    })

    const [errorMessage, setErrorMessage] = useState<string | undefined>()

    const [project, setProject] = useState<{ id: string; name: string }>()

    const [visualizations, setVisualizations] = useState<VisualizationsScreenVisualizationFragment[]>()

    const [rooms, setRooms] = useState<VisualizationsScreenHouseRoomFragment[]>()

    const [selected, setSelected] = useState<string[]>([])

    const selectVisisualization = (id: string, add: boolean) => {
        if (add && selected.includes(id)) {
            setSelected(selected.filter((s) => s !== id))
        } else if (add) {
            setSelected([...selected, id])
        } else {
            setSelected([id])
        }
    }

    useEffect(() => {
        const res = data?.project

        switch (res?.__typename) {
            case "Project":
                setProject({ id: res.id, name: res.name })

                if (res.visualizations?.list?.__typename === "ProjectVisualizationsList") {
                    setVisualizations(res.visualizations.list.items)
                }

                if (res.houses.list.__typename === "ProjectHousesList" && res.houses.list.items.length > 0) {
                    const h = res.houses.list.items[0]

                    if (h.rooms.list.__typename === "HouseRoomsList") {
                        setRooms(h.rooms.list.items)
                    }
                }

                break
            case "Forbidden":
                setErrorMessage("Доступ запрещен")
                break
            case "ServerError":
                setErrorMessage("Ошибка сервера")
                break
            case "NotFound":
                setErrorMessage("Проект не найден")
                break
        }
    }, [data])

    useEffect(() => {
        refetch({
            filter: {
                status: { eq: [VisualizationStatus.Approved, VisualizationStatus.Unknown] },
                roomID: roomsFilter.length > 0 ? { eq: roomsFilter } : undefined,
            },
        })

        setSelected([])
    }, [JSON.stringify(roomsFilter)])

    return (
        <MainLayout
            loading={loading}
            error={errorMessage}
            header={
                <Box direction="row" gap="small">
                    <AnchorLink to={`/p/${project?.id}`} color="black" style={{ left: "-50px" }}>
                        <LinkPrevious />
                    </AnchorLink>
                    <AnchorLink to={`/p/${project?.id}`} color="black">
                        {project?.name}
                    </AnchorLink>
                </Box>
            }
            headerMenu={
                <Box direction="row" gap="small" justify="center" align="center">
                    <DeleteVisualizations
                        visualizations={selected}
                        onDelete={(n: number) => {
                            notify({
                                message: `Удалено ${n} визуализаций`,
                                callback: refetch,
                            })
                        }}
                    />

                    {project && (
                        <Upload
                            projectId={project.id}
                            rooms={rooms as Rooms}
                            roomId={roomsFilter[0]}
                            onVisualizationsUploaded={({ files }) => {
                                notify({
                                    message:
                                        files?.length === 1
                                            ? "Визуализация загружена"
                                            : `Загружено визуализаций ${files?.length}`,
                                    callback: refetch,
                                })
                            }}
                        />
                    )}
                </Box>
            }
        >
            <Box>
                <Box direction="row" justify="between">
                    <Rooms rooms={rooms} pad={{ vertical: "medium" }} />

                    {selected.length > 1 && (
                        <Box direction="row" gap="small" align="center">
                            <Box>
                                <Text size="small">Выбрано&nbsp;{selected.length}</Text>
                            </Box>
                            <Button label="Отменить" size="small" color="dark-5" onClick={() => setSelected([])} />
                        </Box>
                    )}
                </Box>

                <Grid
                    columns="small"
                    rows="small"
                    gap={{ row: "large", column: "medium" }}
                    onClick={() => {
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
                                <Image
                                    onClick={(event: MouseEvent) => {
                                        selectVisisualization(id, event.metaKey)
                                        event.stopPropagation()
                                    }}
                                    fit="contain"
                                    src={`${file.url}?h=192`}
                                    style={{
                                        padding: "3px",
                                        borderRadius: "4px",
                                        boxShadow: selected.includes(id) ? "0 0 0px 2px #7D4CDB" : "none",
                                    }}
                                />
                            </Box>
                        )
                    })}
                </Grid>
            </Box>
        </MainLayout>
    )
}
