import { MouseEvent, useEffect, useState } from "react"

import { Box, Button, Grid, Heading, Image, Layer, LayerExtendedProps, Text } from "grommet"
import { FormClose } from "grommet-icons"

import { AlbumScreenHouseRoomFragment, AlbumScreenVisualizationFragment } from "api/graphql"
import { RoomsFilter } from "features/rooms-filter/RoomsFilter"
import { useAddVisualizationsToAlbum } from "pages/Album/AddVisualizations/useAddVisualizationsToAlbum"

export default function AddVisualizations({
    albumId,
    visualizations,
    rooms,
    alreadyAdded,
    onClickClose,
    onVisualizationsAdded,
    ...layerProps
}: {
    albumId: string
    visualizations: AlbumScreenVisualizationFragment[]
    rooms: AlbumScreenHouseRoomFragment[]
    alreadyAdded: string[]
    onVisualizationsAdded?: () => void
    onClickClose?: () => void
} & LayerExtendedProps) {
    const [errorMessage] = useState<string | undefined>(undefined)

    const [selected, setSelected] = useState<string[]>([])

    const select = (id: string) => {
        if (selected.includes(id)) {
            setSelected(selected.filter((s) => s !== id))
        } else {
            setSelected([...selected, id])
        }
    }

    const [selectedRooms, setSelectedRooms] = useState<string[]>([])

    const [addVisualizations, { data }] = useAddVisualizationsToAlbum(albumId)

    useEffect(() => {
        if (data?.addVisualizationsToAlbum.__typename === "VisualizationsAddedToAlbum") {
            onVisualizationsAdded && onVisualizationsAdded()
        }
    }, [data])

    return (
        <Layer {...layerProps}>
            <Box pad="medium" gap="medium" width="large">
                <Box direction="row" justify="between" align="center">
                    <Heading level={2} margin="none">
                        Визуализации
                    </Heading>
                    <Button icon={<FormClose />} onClick={onClickClose} />
                </Box>

                {errorMessage && (
                    <Box
                        pad="small"
                        round="small"
                        direction="row"
                        gap="small"
                        align="center"
                        background={{ color: "status-critical", opacity: "weak" }}
                    >
                        <Box border={{ color: "status-critical", size: "small" }} round="large">
                            <FormClose color="status-critical" size="medium" />
                        </Box>
                        <Text weight="bold" size="medium">
                            {errorMessage}
                        </Text>
                    </Box>
                )}

                {rooms && (
                    <RoomsFilter
                        rooms={rooms}
                        margin={{ bottom: "medium" }}
                        onSelectRooms={(id: string[]) => setSelectedRooms(id)}
                    />
                )}

                <Box overflow="auto" pad={{ vertical: "medium", horizontal: "small" }}>
                    {visualizations && (
                        <Grid columns="xsmall" rows="xsmall" gap="large">
                            {visualizations
                                .filter((vis) => {
                                    if (selectedRooms.length === 0) {
                                        return true
                                    } else if (vis.room?.id) {
                                        return selectedRooms.includes(vis.room.id)
                                    }

                                    return false
                                })
                                .map((vis, key) => {
                                    return (
                                        <Box
                                            key={key}
                                            width={{ max: "xsmall" }}
                                            height={{ max: "xsmall" }}
                                            justify="center"
                                            align="center"
                                        >
                                            <Button
                                                badge={
                                                    selected.includes(vis.id)
                                                        ? {
                                                              value: selected.indexOf(vis.id) + 1,
                                                              max: 999,
                                                          }
                                                        : undefined
                                                }
                                            >
                                                <Box
                                                    round={selected.includes(vis.id) ? "xsmall" : "none"}
                                                    pad="xxsmall"
                                                    style={{
                                                        boxShadow: selected.includes(vis.id)
                                                            ? "0 0 0px 2px #7D4CDB"
                                                            : "none",
                                                    }}
                                                    focusIndicator={false}
                                                    onClick={(event: MouseEvent) => {
                                                        select(vis.id)
                                                        event.stopPropagation()
                                                    }}
                                                >
                                                    <Image
                                                        fit="contain"
                                                        fill
                                                        src={`${vis.file.url}?h=96`}
                                                        opacity={alreadyAdded.includes(vis.id) ? "0.5" : undefined}
                                                    />
                                                </Box>
                                            </Button>
                                        </Box>
                                    )
                                })}
                        </Grid>
                    )}
                </Box>

                <Box direction="row" height="xxsmall" align="center">
                    <Button
                        type="button"
                        primary
                        label="Добавить"
                        disabled={selected.length === 0}
                        onClick={() => {
                            addVisualizations(selected)
                            setSelected([])
                        }}
                    />
                </Box>
            </Box>
        </Layer>
    )
}
