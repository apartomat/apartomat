import { MouseEventHandler, useEffect, useState } from "react"
import { useParams, useNavigate } from "react-router-dom"

import { Box, BoxExtendedProps, Button, Grid, Heading, Header, Image, Text } from "grommet"
import { Add, Close, Sort } from "grommet-icons"

import useAlbum, {
    AlbumScreenVisualizationFragment,
    AlbumScreenProjectFragment,
    AlbumScreenAlbumPageCoverFragment,
    AlbumScreenAlbumPageVisualizationFragment,
    AlbumScreenSettingsFragment,
    PageOrientation as PageOrientationEnum,
    AlbumScreenHouseRoomFragment,
} from "./useAlbum"

import { PageSize, PageOrientation } from "./Settings/"
import AddVisualizations from "screen/Album/AddVisualizations/AddVisualizations"
import GenerateAlbumFile from "screen/Album/GenerateFile/GenerateFile"
import { Pages } from "./Pages"

export function Album() {
    const { id } = useParams<"id">() as { id: string }

    const [project, setProject] = useState<AlbumScreenProjectFragment | undefined>()

    const [visualizations, setVisualizations] = useState<AlbumScreenVisualizationFragment[]>([])

    const [rooms, setRooms] = useState<AlbumScreenHouseRoomFragment[]>([])

    const [pages, setPages] = useState<
        (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]
    >([])

    const [settings, setSettings] = useState<AlbumScreenSettingsFragment | undefined>()

    const { data, loading, refetch } = useAlbum({ id })

    const [showAddVisualizations, setShowAddVisualizations] = useState(false)

    useEffect(() => {
        if (data?.album?.__typename === "Album") {
            if (data?.album?.project?.__typename === "Project") {
                setProject(data.album.project)

                if (data.album.project.visualizations.list.__typename === "ProjectVisualizationsList") {
                    setVisualizations(data.album.project.visualizations.list.items)
                }

                if (
                    data.album.project.houses.__typename === "ProjectHouses" &&
                    data.album.project.houses.list.__typename === "ProjectHousesList"
                ) {
                    const list = data.album.project.houses.list.items[0].rooms.list

                    if (list.__typename === "HouseRoomsList") {
                        setRooms(list.items)
                    }
                }
            }

            if (data.album.pages.__typename === "AlbumPages") {
                setPages(data.album.pages.items)
            }

            setSettings(data.album.settings)
        }
    }, [data])

    const [currentPage, setCurrentPage] = useState<number>(0)

    const [redirectTo, setRedirectTo] = useState<string | undefined>(undefined)

    const navigate = useNavigate()

    if (redirectTo) {
        navigate(redirectTo)
    }

    return (
        <Grid
            fill
            columns={["small", "flex", "small"]}
            rows={["auto", "flex", "auto"]}
            areas={[
                { name: "header", start: [0, 0], end: [2, 0] },
                { name: "left", start: [0, 1], end: [0, 1] },
                { name: "main", start: [1, 1], end: [1, 1] },
                { name: "right", start: [2, 1], end: [2, 1] },
                { name: "footer", start: [0, 2], end: [2, 2] },
            ]}
            width={{ width: "100vw" }}
            height={{ height: "100vh" }}
            pad="medium"
            responsive
        >
            <Header gridArea="header">
                <Box>
                    <Text size="xlarge" weight="bold">
                        {data?.album.__typename === "Album" && <>{data?.album.name}</>}
                    </Text>
                </Box>
                <Box>
                    <Button
                        icon={<Close />}
                        onClick={() => {
                            project && setRedirectTo(`/p/${project.id}`)
                        }}
                    />
                </Box>
            </Header>

            <Box gridArea="right">
                {settings && (
                    <Box gap="small" align="end" margin={{ top: "large" }}>
                        <Heading level={5}>Настройки для печати</Heading>
                        <PageSize albumId={id} size={settings.pageSize} onAlbumPageSizeChanged={() => refetch()} />
                        <PageOrientation
                            albumId={id}
                            orientation={settings.pageOrientation}
                            onAlbumPageOrientationChanged={() => refetch()}
                        />
                    </Box>
                )}
                {data?.album.__typename === "Album" && (
                    <Box gap="small" align="end" margin={{ top: "large" }}>
                        <GenerateAlbumFile album={data.album} onAlbumFileGenerated={() => refetch()} />
                    </Box>
                )}
            </Box>

            <Box gridArea="left"></Box>

            {pages.length === 0 && !loading && (
                <Box gridArea="main" align="center" justify="center">
                    <Box margin={{ bottom: "medium" }}>
                        <Text size="small" color="text-xweak" textAlign="center">
                            В альбом можно добавить обложку, визуализации и другие материалы
                        </Text>
                    </Box>
                    <Button
                        icon={<Add />}
                        label="Добавить..."
                        primary
                        onClick={() => {
                            setShowAddVisualizations(true)
                        }}
                    />
                </Box>
            )}

            {pages.length > 0 && (
                <Box gridArea="main" align="center" justify="center">
                    {pages
                        .filter((_, i) => i === currentPage)
                        .map((p, key) => {
                            return (
                                <Box
                                    key={key}
                                    style={{ aspectRatio: orientationToAspectRation(settings?.pageOrientation) }}
                                    pad="medium"
                                    background="background-contrast"
                                    round="xsmall"
                                >
                                    {(() => {
                                        switch (p.__typename) {
                                            case "AlbumPageVisualization":
                                                switch (p.visualization.__typename) {
                                                    case "Visualization":
                                                        return (
                                                            <Image
                                                                key={key}
                                                                fit="contain"
                                                                src={p.visualization.file.url}
                                                            />
                                                        )
                                                    default:
                                                        return <></>
                                                }
                                            default:
                                                return <></>
                                        }
                                    })()}
                                </Box>
                            )
                        })}
                </Box>
            )}

            <Box gridArea="footer">
                {pages.length > 0 && (
                    <Box direction="row" align="center" justify="center">
                        <Box align="center" margin={{ right: "medium" }}>
                            <Box round="full" overflow="hidden" background="light-2">
                                <Button icon={<Sort />} hoverIndicator />
                            </Box>
                        </Box>

                        {pages.length > 0 && (
                            <Pages
                                pages={pages}
                                current={currentPage}
                                onClickPage={(n) => setCurrentPage(n)}
                                width={{ max: "large" }}
                            />
                        )}

                        <AddVisualizationsCircleButton
                            margin={{ left: "medium" }}
                            onClick={() => {
                                setShowAddVisualizations(true)
                            }}
                        />
                    </Box>
                )}
            </Box>

            {showAddVisualizations && (
                <AddVisualizations
                    albumId={id}
                    visualizations={visualizations}
                    rooms={rooms}
                    inAlbum={ids(pages)}
                    onVisualizationsAdded={() => {
                        setShowAddVisualizations(false)
                        refetch()
                    }}
                    onEsc={() => setShowAddVisualizations(false)}
                    onClickOutside={() => setShowAddVisualizations(false)}
                    onClickClose={() => setShowAddVisualizations(false)}
                />
            )}
        </Grid>
    )
}

export default Album

function AddVisualizationsCircleButton({
    onClick,
    ...boxProps
}: {
    onClick?: MouseEventHandler | undefined
} & BoxExtendedProps) {
    return (
        <Box {...boxProps} round="full" overflow="hidden" background="brand">
            <Button icon={<Add />} hoverIndicator onClick={onClick} />
        </Box>
    )
}

function orientationToAspectRation(orientation?: PageOrientationEnum): string {
    const land = "1.41/1",
        port = "1/1.41"

    switch (orientation) {
        case PageOrientationEnum.Landscape:
            return land
        case PageOrientationEnum.Portrait:
            return port
    }

    return port
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
function ids(pages: (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]): string[] {
    const vis = pages.filter(
        (p) => p.__typename === "AlbumPageVisualization" && p.visualization.__typename === "Visualization"
    ) as { visualization: { __typename?: "Visualization"; id: string; file: { __typename?: "File"; url: any } } }[]

    return vis.map((v) => v.visualization.id)
}
