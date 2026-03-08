import { MouseEvent, useEffect, useState, useCallback } from "react"
import { useParams } from "react-router-dom"
import { useVisualizations, VisualizationStatus } from "./api/useVisualizations"
import { VisualizationsScreenHouseRoomFragment, VisualizationsScreenVisualizationFragment } from "api/graphql"
import { MainLayout } from "widgets/main-layout/MainLayout"
import { AnchorLink } from "shared/ui"
import { LinkPrevious } from "grommet-icons"
import { Box, Button, Text } from "grommet"

import { Rooms, useSearchParamsRoomsFilter } from "./Rooms"
import { DeleteVisualizations } from "./DeleteVisualizations"
import { useNotifications } from "shared/context/notifications/context"
import { Upload } from "./Upload"

type Rooms = Pick<VisualizationsScreenHouseRoomFragment, "id" | "name">[]

const MAX_SIZE = 192

function VisualizationThumbnail({
    src,
    selected,
}: {
    src: string
    selected: boolean
}) {
    const [size, setSize] = useState<{ w: number; h: number } | null>(null)

    const onLoad = useCallback((e: React.SyntheticEvent<HTMLImageElement>) => {
        const img = e.currentTarget
        const nw = img.naturalWidth
        const nh = img.naturalHeight
        if (nw <= 0 || nh <= 0) return
        let w = nw
        let h = nh
        if (w > MAX_SIZE || h > MAX_SIZE) {
            const r = Math.min(MAX_SIZE / w, MAX_SIZE / h)
            w = Math.round(w * r)
            h = Math.round(h * r)
        }
        setSize({ w, h })
    }, [])

    return (
        <span
            style={{
                display: "inline-block",
                padding: "3px",
                borderRadius: "4px",
                boxShadow: selected ? "0 0 0 2px #7D4CDB" : "none",
            }}
        >
            <span
                style={{
                    display: "block",
                    overflow: "hidden",
                    width: size ? `${size.w}px` : `${MAX_SIZE}px`,
                    height: size ? `${size.h}px` : `${MAX_SIZE}px`,
                    maxWidth: MAX_SIZE,
                    maxHeight: MAX_SIZE,
                }}
            >
                <img
                    src={src}
                    alt=""
                    onLoad={onLoad}
                    style={{
                        display: "block",
                        width: size ? `${size.w}px` : "auto",
                        height: size ? `${size.h}px` : "auto",
                        maxWidth: MAX_SIZE,
                        maxHeight: MAX_SIZE,
                        objectFit: "contain",
                        verticalAlign: "top",
                        margin: 0,
                        padding: 0,
                    }}
                />
            </span>
        </span>
    )
}

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
                            onVisualizationsUploaded={({ uploadedCount, failedCount }) => {
                                let message = `Загружено визуализаций ${uploadedCount}`

                                if (uploadedCount === 1) {
                                    message = "Визуализация загружена"
                                }

                                if (failedCount > 0) {
                                    message = `Не все визуализации были загружены`
                                }

                                const severity = failedCount > 0 ? "warning" : "ok"

                                notify({ message, severity, callback: refetch })
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

                <div
                    style={{
                        display: "grid",
                        gridTemplateColumns: "repeat(auto-fill, 198px)",
                        gap: "24px",
                    }}
                    onClick={() => {
                        setSelected([])
                    }}
                >
                    {visualizations?.map(({ id, file }) => {
                        return (
                            <Box
                                key={id}
                                width="198px"
                                height="198px"
                                justify="center"
                                align="center"
                                focusIndicator={false}
                                onClick={(event: MouseEvent) => {
                                    selectVisisualization(id, event.metaKey)
                                    event.stopPropagation()
                                }}
                            >
                                <VisualizationThumbnail
                                    src={`${file.url}?h=192`}
                                    selected={selected.includes(id)}
                                />
                            </Box>
                        )
                    })}
                </div>
            </Box>
        </MainLayout>
    )
}
