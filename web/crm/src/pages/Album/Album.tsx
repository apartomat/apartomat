import React, { MouseEventHandler, useEffect, useRef, useState } from "react"
import { useParams, useNavigate } from "react-router-dom"
import { InView } from "react-intersection-observer"

import { Box, BoxExtendedProps, Button, Grid, Heading, Header, Image, Text, Drop, Main } from "grommet"
import { Add, Close } from "grommet-icons"

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
import AddVisualizations from "pages/Album/AddVisualizations/AddVisualizations"
import { GenerateFile as GenerateAlbumFile } from "pages/Album/GenerateFile"

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

    const [ scale, setScale ] = useState(0.5)

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

    const navigate = useNavigate()

    const [inView, setInView] = useState(0)

    return (
        <Main overflow="scroll" style={{ position: "fixed", inset: 0 }}>
            <Header
                style={{
                    position: "fixed",
                    top: 0,
                    left: 0,
                    right: 0,
                    zIndex: 2,
                }}
                background={{
                    color: "white",
                    opacity: "strong",
                }}
                pad={{ top: "medium", bottom: "small", horizontal: "large" }}
                justify="between"
            >
                <Grid columns={{ count: 3, size: "auto" }} gap="small" width="100%">
                    <Box align="start" justify="center">
                        <Text size="xlarge" weight="bold">
                            {data?.album.__typename === "Album" && <>{data?.album.name}</>}
                        </Text>
                    </Box>
                    <Box align="center" justify="center">
                        {pages.length > 0 && (
                            <Text weight="bold" color="text-xweak">
                                {inView + 1} / {pages.length}
                            </Text>
                        )}
                    </Box>
                    <Box align="end" justify="center">
                        <Button
                            icon={<Close />}
                            onClick={() => {
                                project && navigate(`/p/${project.id}`)
                            }}
                        />
                    </Box>
                </Grid>
            </Header>

            <Box
                style={{
                    position: "fixed",
                    top: "84px",
                    right: "60px",
                    zIndex: 1
                }}
                background="background-back"
                pad="medium"
                round="xsmall"
                margin={{ top: "large" }}
            >
                {settings && (
                    <Box gap="small" align="end">
                        <Heading level={5} margin={{ top: "none" }}>
                            Настройки для печати
                        </Heading>
                        <PageSize albumId={id} size={settings.pageSize} onAlbumPageSizeChanged={() => refetch()} />
                        <PageOrientation
                            albumId={id}
                            orientation={settings.pageOrientation}
                            onAlbumPageOrientationChanged={() => refetch()}
                        />
                    </Box>
                )}
                {data?.album.__typename === "Album" && (
                    <Box gap="small" align="end" margin={{ top: "medium" }}>
                        <GenerateAlbumFile album={data.album} onAlbumFileGenerated={() => refetch()} />
                    </Box>
                )}
            </Box>

            <AddVariants
                style={{
                    position: "fixed",
                    bottom: 0,
                    left: 0,
                    right: 0,
                    zIndex: 1
                }}
                direction="row"
                justify="center"
                pad="small"
                onClickAddVisualizations={() => setShowAddVisualizations(true)}
            />

            {pages.length > 0 && (
                <Box align="center" pad={{ top: "84px", bottom: "68px" }}>
                    <Grid>
                        {pages.map((p, key) => {
                            return (
                                <InView
                                    key={key}
                                    onChange={(inView) => {
                                        if (inView) {
                                            setInView(key)
                                        }
                                    }}
                                    threshold={1.0}
                                >
                                    <Box
                                        background="background-contrast"
                                        // round="small"
                                        margin={{ vertical: "xxsmall" }}
                                        width={orientationWidth(settings?.pageOrientation, scale)}
                                        height={orientationHeight(settings?.pageOrientation, scale)}
                                    >
                                        {(() => {
                                            switch (p.__typename) {
                                                case "AlbumPageVisualization":
                                                    switch (p.visualization.__typename) {
                                                        case "Visualization":
                                                            return (
                                                                <Box
                                                                    style={{transform: `scale(${scale})`, transformOrigin: "left top"}}
                                                                >
                                                                    {p.svg.__typename === "Svg" &&
                                                                        <div dangerouslySetInnerHTML={svg(p.svg.svg)}/>
                                                                    }
                                                                </Box>
                                                            )
                                                        default:
                                                            return <></>
                                                    }
                                                case "AlbumPageCover":
                                                    switch (p.cover.__typename) {
                                                        case "CoverUploaded":
                                                            return (
                                                                <Box
                                                                    style={{transform: `scale(${scale})`, transformOrigin: "left top"}}
                                                                >
                                                                    {p.svg.__typename === "Svg" &&
                                                                        <div dangerouslySetInnerHTML={svg(p.svg.svg)}/>
                                                                    }
                                                                </Box>
                                                            )
                                                        default:
                                                            return <></>
                                                    }
                                                default:
                                                    return <></>
                                            }
                                        })()}
                                    </Box>
                                </InView>
                            )
                        })}
                    </Grid>
                </Box>
            )}

            {showAddVisualizations && (
                <AddVisualizations
                    albumId={id}
                    visualizations={visualizations}
                    rooms={rooms}
                    alreadyAdded={ids(pages)}
                    onVisualizationsAdded={() => {
                        setShowAddVisualizations(false)
                        refetch()
                    }}
                    onEsc={() => setShowAddVisualizations(false)}
                    onClickOutside={() => setShowAddVisualizations(false)}
                    onClickClose={() => setShowAddVisualizations(false)}
                />
            )}
        </Main>
    )
}

function AddVisualizationsCircleButton({
                                           onClick,
                                           ...boxProps
                                       }: {
    onClick?: MouseEventHandler | undefined
} & BoxExtendedProps) {
    return (
        <Box {...boxProps} round="full" overflow="hidden" background="brand">
            <Button icon={<Add/>} hoverIndicator onClick={onClick}/>
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

function orientationWidth(orientation: PageOrientationEnum = PageOrientationEnum.Landscape, scale: Number = 1.0): string {
    const landscapeWidth = 420, portraitWidth = 297

    switch (orientation) {
        case PageOrientationEnum.Landscape:
            return `${landscapeWidth  * scale}mm`
        case PageOrientationEnum.Portrait:
            return `${portraitWidth  * scale}mm`
    }
}

function orientationHeight(orientation: PageOrientationEnum = PageOrientationEnum.Landscape, scale: Number = 1.0): string {
    const landscapeHeight = 297, portraitHeight = 420

    switch (orientation) {
        case PageOrientationEnum.Landscape:
            return `${landscapeHeight  * scale}mm`
        case PageOrientationEnum.Portrait:
            return `${portraitHeight  * scale}mm`
    }
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
function ids(pages: (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]): string[] {
    const vis = pages.filter(
        (p) => p.__typename === "AlbumPageVisualization" && p.visualization.__typename === "Visualization"
    ) as { visualization: { __typename?: "Visualization"; id: string; file: { __typename?: "File"; url: any } } }[]

    return vis.map((v) => v.visualization.id)
}

function AddVariants({
    onClickAddVisualizations,
    ...boxProps
}: { onClickAddVisualizations?: () => void } & {boxProps: BoxExtendedProps}) {
    const [open, setOpen] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box {...boxProps}>
            <Box ref={targetRef} border={{ color: "background-front", size: "medium" }} round="large">
                <Button label="Добавить..." icon={<Add />} primary onClick={() => setOpen(true)} />
            </Box>

            {open && targetRef.current && (
                <Drop
                    elevation="none"
                    target={targetRef.current}
                    onClickOutside={() => setOpen(false)}
                    onEsc={() => setOpen(false)}
                    align={{ bottom: "bottom" }}
                    round="large"
                >
                    <Box gap="small" border={{ color: "background-front", size: "medium" }} direction="row">
                        <Button primary label="Обложку" />
                        <Button
                            primary
                            label="Визуализации"
                            color="accent-3"
                            onClick={() => {
                                setOpen(false)
                                onClickAddVisualizations && onClickAddVisualizations()
                            }}
                        />
                        <Button primary label="Ссылку на сайт" color="status-ok" />
                    </Box>
                </Drop>
            )}
        </Box>
    )
}

function svg(html: string) {
    return {__html: html}
}